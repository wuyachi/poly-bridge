package zilliqamonitor

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/chainsdk"
	"poly-bridge/conf"
	"time"
)

type ZilliqaMonitor struct {
	monitorConfig *conf.HealthMonitorConfig
	sdks          map[string]*chainsdk.ZilliqaSdk
	nodeHeight    map[string]uint64
	nodeStatus    map[string]string
}

func NewZilliqaHealthMonitor(monitorConfig *conf.HealthMonitorConfig) *ZilliqaMonitor {
	zilliqaMonitor := &ZilliqaMonitor{}
	zilliqaMonitor.monitorConfig = monitorConfig
	sdks := make(map[string]*chainsdk.ZilliqaSdk, 0)
	for _, node := range monitorConfig.ChainNodes.Nodes {
		sdks[node.Url] = chainsdk.NewZilliqaSdk(node.Url)
	}
	zilliqaMonitor.sdks = sdks
	zilliqaMonitor.nodeHeight = make(map[string]uint64, len(sdks))
	zilliqaMonitor.nodeStatus = make(map[string]string, len(sdks))
	return zilliqaMonitor
}

func (z *ZilliqaMonitor) GetChainName() string {
	return z.monitorConfig.ChainName
}

func (z *ZilliqaMonitor) RelayerBalanceMonitor() ([]*basedef.RelayerAccountStatus, error) {
	return nil, nil
}

func (z *ZilliqaMonitor) NodeMonitor() ([]basedef.NodeStatus, error) {
	nodeStatuses := make([]basedef.NodeStatus, 0)
	for url, sdk := range z.sdks {
		status := basedef.NodeStatus{
			ChainId:   z.monitorConfig.ChainId,
			ChainName: z.monitorConfig.ChainName,
			Url:       url,
			Status:    make([]string, 0),
			Time:      time.Now().Unix(),
		}
		height, err := z.GetCurrentHeight(sdk)
		if err == nil {
			status.Height = height
			z.nodeHeight[url] = height
			err = z.CheckAbiCall(sdk)
		}
		if err != nil {
			z.nodeStatus[url] = err.Error()
		} else {
			z.nodeStatus[url] = basedef.StatusOk
		}
		status.Status = append(status.Status, z.nodeStatus[url])
		nodeStatuses = append(nodeStatuses, status)
	}
	data, _ := json.Marshal(nodeStatuses)
	_, err := cacheRedis.Redis.Set(cacheRedis.NodeStatusPrefix+z.monitorConfig.ChainName, data, time.Hour*24)
	if err != nil {
		logs.Error("set %s node status error: %s", z.GetChainName(), err)
	}
	return nodeStatuses, err
}

func (z *ZilliqaMonitor) GetCurrentHeight(sdk *chainsdk.ZilliqaSdk) (uint64, error) {
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil || height == 0 || height == math.MaxUint64 {
		err := fmt.Errorf("get current block height err: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", z.GetChainName(), sdk.GetUrl(), err))
		return 0, err
	}
	logs.Info("%s node: %s, latest height: %d", z.GetChainName(), sdk.GetUrl(), height)
	return height, nil
}

func (z *ZilliqaMonitor) CheckAbiCall(sdk *chainsdk.ZilliqaSdk) error {
	_, err := sdk.GetBlock(z.nodeHeight[sdk.GetUrl()] - 1)
	if err != nil {
		err := fmt.Errorf("call GetBlock error: %s", err)
		logs.Error(fmt.Sprintf("%s node: %s, %s ", z.GetChainName(), sdk.GetUrl(), err))
		return err
	}
	return nil
}
