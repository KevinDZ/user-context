package db

import (
	"log"
	"os"
	"user-context/diamond/acl/adapters/pl/dao"
	"user-context/utils/common/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitTables init tables
func InitTables(db *gorm.DB) {
	account := dao.Account{}
	tables := map[string]interface{}{
		account.TableName(): account,
	}
	for k, v := range tables {
		if !db.Migrator().HasTable(k) {
			err := db.Migrator().CreateTable(v)
			if err != nil {
				log.Fatalf("init table %s error: %v", k, err)
			}
		}
	}
}

// DisconnectDB ...
func DisconnectDB(db *gorm.DB) error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		db = nil
		return err
	}
	return nil
}

// NewDBEngine create db engine
func NewDBEngine() *gorm.DB {
	var err error
	var db *gorm.DB
	var dbDriver string
	config := config.FileConfig

	if driver := os.Getenv("DRIVER"); driver == "" {
		dbDriver = config.DB.Driver
	} else {
		dbDriver = driver
	}

	if dbDriver == "postgres" {
		db, err = gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	} else if dbDriver == "sqlite" {
		//db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}

	log.Println("Connect to db successfully.")
	return db
}
