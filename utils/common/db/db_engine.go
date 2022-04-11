package db

import (
	"log"
	"os"
	"user-context/rhombic/acl/adapters/pl"
	"user-context/utils/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	// dbDSN    = "host=127.0.0.1,user=postgres,password=realibox2021_postgres,dbname=order,port=54321,sslmode=disable,TimeZone=Asia/Shanghai"
// 	dbDSN    = "host=127.0.0.1,user=postgres,password=driver,dbname=,port=54321,sslmode=disable,TimeZone=Asia/Shanghai"
// 	dbDriver = "postgres"
// )

var (
	db       *gorm.DB
	dbDriver string
)

// ConnectDB return orm.DB
func ConnectDB() (*gorm.DB, error) {
	var err error
	config := common.FileConfig
	if err != nil {
		return nil, err
	}

	if driver := os.Getenv("DRIVER"); driver == "" {
		dbDriver = config.DB.Driver
	} else {
		dbDriver = driver
	}

	if db == nil {
		if dbDriver == "postgres" {
			db, err = gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
		} else if dbDriver == "sqlite" {
			//db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		}

		log.Println("Connect to db successfully.")
		if err != nil {
			return nil, err
		}
	}

	return db, err
}

// InitTables init tables
func InitTables(db *gorm.DB) {
	account := pl.Account{}
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
func DisconnectDB() error {
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
	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}
	return db
}
