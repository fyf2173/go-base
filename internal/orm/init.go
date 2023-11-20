package orm

import (
	"sync"

	"github.com/fyf2173/ysdk-go/util"
	"github.com/fyf2173/ysdk-go/xdb"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB
var once sync.Once

func InitConn() {
	once.Do(func() {
		if viper.GetString("data_driver") == "sqlite3" {
			db, err := gorm.Open(sqlite.Open(viper.GetString("sqliteDsn")), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			instance = db
		}
		if viper.GetString("data_driver") == "mysql" {
			db, err := xdb.NewGorm(viper.GetString("env"), util.ViperMustGetNode("db", xdb.DbConfig{}))
			if err != nil {
				panic(err)
			}
			instance = db
		}
	})
	return
}

func DB() *gorm.DB {
	return instance
}
