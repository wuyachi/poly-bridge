package neomonitor

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"time"
)

type NeoMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.NeoSdk
	nodeHeight    map[string]uint64
}

func NewNeoHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *NeoMonitor {
	neoMonitor := &NeoMonitor{}
	neoMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.NeoSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdk := chainsdk.NewNeoSdk(node.Url)
		if sdk.GetClient() == nil {
			if _, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+node.Url, fmt.Sprintf("initial sdk error:sdk.client is nil"), time.Hour*24); err != nil {
				logs.Error("set %s node[%s] status error: %s", monitorConfig.ChainName, node.Url, err)
			}
			logs.Error("%s node: %s, initial sdk error:sdk.client is nil", monitorConfig.ChainName, node.Url)
			continue
		}
		sdks[node.Url] = sdk
	}
	neoMonitor.sdks = sdks
	neoMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	return neoMonitor
}

func (n *NeoMonitor) GetChainName() string {
	return n.monitorConfig.ChainName
}

func (n *NeoMonitor) GetChainId() uint64 {
	return n.monitorConfig.ChainId
}

func (n *NeoMonitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	balanceSuccessMap := make(map[string]float64, 0)
	balanceFailedMap := make(map[string]string, 0)
	var precision float64 = 1
	for _, sdk := range n.sdks {
		for _, address := range n.monitorConfig.RelayerAccount.Address {
			if _, ok := balanceSuccessMap[address]; ok {
				continue
			}
			txBuilder := &tx.TransactionBuilder{
				EndPoint: sdk.GetUrl(),
				Client:   sdk.GetClient(),
			}
			walletHelper := wallet.NewWalletHelper(txBuilder, nil)
			_, gasBalance, err := walletHelper.GetBalance(address)

			if err == nil {
				balanceSuccessMap[address] = gasBalance
				delete(balanceFailedMap, address)
			} else {
				balanceFailedMap[address] = err.Error()
			}
		}
	}
	relayerStatus := make([]*basedef.RelayerAccountStatus, 0)
	for address, balance := range balanceSuccessMap {
		status := basedef.RelayerAccountStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Address:   address,
			Balance:   balance / precision,
			Threshold: n.monitorConfig.RelayerAccount.Threshold / precision,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	for address, err := range balanceFailedMap {
		status := basedef.RelayerAccountStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Address:   address,
			Balance:   0,
			Threshold: n.monitorConfig.RelayerAccount.Threshold / precision,
			Status:    err,
			Time:      time.Now().Unix(),
		}
		relayerStatus = append(relayerStatus, &status)
	}
	return relayerStatus, nil
}

func (n *NeoMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range n.sdks {
		status := basedef.NodeStatus{
			ChainId:   n.monitorConfig.ChainId,
			ChainName: n.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := n.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			n.nodeHeight[url] = height
			err = n.CheckAbiCall(sdk)
		}
		err = n.CheckAbiCall(sdk)
		if err != nil {
			status.Status = append(status.Status, err.Error())
		}
		nodeStatuses = append(nodeStatuses, status)
	}
	//data, _ := json.Marshal(nodeStatuses)
	//_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+n.monitorConfig.ChainName, data, time.Hour*24)
	//if err != nil {
	//	logs.Error("set %s node status error: %s", n.GetChainName(), err)
	//}
	return nodeStatuses, nil
}

func (n *NeoMonitor) GetCurrentHeight(sdk *chainsdk.NeoSdk) (uint64, error) {
	height, err := sdk.GetBlockCount()
	if err != nil || height == 0 || height == math.MaxUint64 {
		e := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", n.GetChainName(), sdk.GetUrl(), e))
		return 0, e
	}
	logs.Info("%s node: %s, latest height: %d", n.GetChainName(), sdk.GetUrl(), height)
	return height, nil
}

func (n *NeoMonitor) CheckAbiCall(sdk *chainsdk.NeoSdk) error {
	_, err := sdk.GetBlockByIndex(n.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		e := fmt.Errorf("call GetBlockByIndex error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", n.GetChainName(), sdk.GetUrl(), e))
		return err
	}
	return nil
}
