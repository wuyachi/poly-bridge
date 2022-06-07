package test

import (
	"github.com/polynetwork/poly-bridge/basedef"
	"github.com/polynetwork/poly-bridge/conf"
	"github.com/polynetwork/poly-bridge/crosschainlisten"
	"github.com/polynetwork/poly-bridge/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestHandleSingleBlock(t *testing.T) {
	c := conf.NewConfig("./../../../conf/config_testnet.json")
	assert.NotNil(t, c)

	var (
		chainId            = basedef.BSC_CROSSCHAIN_ID
		blockHeight uint64 = 8022375
	)

	dbcfg := c.DBConfig
	Logger := logger.Default
	Logger = Logger.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(c.DBConfig.User+":"+dbcfg.Password+"@tcp("+dbcfg.URL+")/"+
		dbcfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	assert.NoError(t, err)

	cfg := c.GetChainListenConfig(chainId)
	assert.NotNil(t, cfg)

	handler := crosschainlisten.NewChainHandle(cfg)

	wpTxs, srcTxs, polyTxs, dstTxs, err := handler.HandleNewBlock(blockHeight)
	assert.NoError(t, err)

	err = onlyUpdateEvents(db, wpTxs, srcTxs, polyTxs, dstTxs)
	assert.NoError(t, err)
}

func onlyUpdateEvents(
	db *gorm.DB,
	wptxs []*models.WrapperTransaction,
	srcTxs []*models.SrcTransaction,
	polyTxs []*models.PolyTransaction,
	dstTxs []*models.DstTransaction,
) error {

	if wptxs != nil && len(wptxs) > 0 {
		res := db.Save(wptxs)
		if res.Error != nil {
			return res.Error
		}
	}
	if srcTxs != nil && len(srcTxs) > 0 {
		res := db.Save(srcTxs)
		if res.Error != nil {
			return res.Error
		}
	}
	if polyTxs != nil && len(polyTxs) > 0 {
		res := db.Save(polyTxs)
		if res.Error != nil {
			return res.Error
		}
	}
	if dstTxs != nil && len(dstTxs) > 0 {
		res := db.Save(dstTxs)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
