package main

import (
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
	"poly-bridge/models"
)

func migrateTable(src, dst *gorm.DB, table string, model interface{}) {
	logs.Info("Migrating table %s", table)
	err := src.Find(model).Error
	checkError(err, "Loading table")
	err = dst.Save(model).Error
	checkError(err, "Saving table")
	countTables(table, table, src, dst)
}

func updateTable(src, dst *gorm.DB, table string, model interface{}) {
	logs.Info("Migrating table %s", table)
	err := src.Find(model).Error
	checkError(err, "Loading table")
	err = dst.Updates(model).Error
	checkError(err, "updating table")
	countTables(table, table, src, dst)
}

func countTables(tableA, tableB string, src, dst *gorm.DB) {
	var a, b int64
	err := src.Table(tableA).Count(&a).Error
	checkError(err, "count table error")
	err = src.Table(tableA).Count(&b).Error
	checkError(err, "count table error")
	logs.Info("===============  Compare table size %s %d:%d %s ============", tableA, a, b, tableB)
}

func migrateBridgeBasicTables(bri, db *gorm.DB) {
	migrateTable(bri, db, "token_basics", &[]*models.TokenBasic{})
	migrateTable(bri, db, "price_markets", &[]*models.PriceMarket{})
	migrateTable(bri, db, "chains", &[]*models.Chain{})
	migrateTable(bri, db, "chain_fees", &[]*models.ChainFee{})
	migrateTable(bri, db, "nft_profiles", &[]*models.NFTProfile{})
	migrateTable(bri, db, "tokens", &[]*models.Token{})
	migrateTable(bri, db, "token_maps", &[]*models.TokenMap{})
}

func syncBridgeBasicTables(bri, db *gorm.DB) {
	updateTable(bri, db, "token_basics", &[]*models.TokenBasic{})
	updateTable(bri, db, "price_markets", &[]*models.PriceMarket{})
	updateTable(bri, db, "chains", &[]*models.Chain{})
	updateTable(bri, db, "chain_fees", &[]*models.ChainFee{})
	updateTable(bri, db, "nft_profiles", &[]*models.NFTProfile{})
	updateTable(bri, db, "tokens", &[]*models.Token{})
	updateTable(bri, db, "token_maps", &[]*models.TokenMap{})
}
