package orm

import (
	"sync"

	"github.com/fyf2173/ysdk-go/util"
	"github.com/fyf2173/ysdk-go/xdb"
	"gorm.io/gorm"
)

var instance *gorm.DB
var once sync.Once

func InitConn() {
	once.Do(func() {
		db, err := xdb.NewGorm(util.ViperMustGetNode("env", ""), util.ViperMustGetNode("db", xdb.DbConfig{}))
		if err != nil {
			panic(err)
		}
		instance = db
	})
	return
}

func DB() *gorm.DB {
	return instance
}
