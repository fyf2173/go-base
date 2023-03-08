package config

import (
	"encoding/json"
	"go-base/internal/pkg/utils"
	"log"

	"github.com/fyf2173/ysdk-go/xdb"
	"github.com/spf13/viper"
)

var CfgData Config

type Config struct {
	Mode  string       `json:"mode" yaml:"mode"`
	Mysql xdb.DbConfig `json:"mysql"`
}

func InitConfigData() {
	CfgData.Mode = viper.GetString("mode")
	if err := utils.ViperGetNode("db", &CfgData.Mysql); err != nil {
		panic(err)
	}

	cfgDataB, _ := json.Marshal(CfgData)
	log.Println("load config ", string(cfgDataB))
}
