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
	Mode    string       `json:"mode" yaml:"mode"`
	Mysql   xdb.DbConfig `json:"mysql"`
	MiniApp MiniApp      `json:"mini_app"`
	Kepler  KplConfig    `json:"kepler"`
}

type MiniApp struct {
	Appid  string `json:"appid" yaml:"appid"`
	Secret string `json:"secret" yaml:"secret"`
}

type KplConfig struct {
	App     KplApp     `json:"app" yaml:"app"`
	Account KplAccount `json:"account" yaml:"account"`
	JCQ     KplJCQ     `json:"jcq" yaml:"jcq"`
}

type KplAccount struct {
	Pin        string `json:"pin" yaml:"pin"`
	ChannelId  int64  `json:"channel_id" yaml:"channel_id"`
	CustomerId int64  `json:"customer_id" yaml:"customer_id"`
	OpName     string `json:"op_name" yaml:"op_name"`
}

type KplApp struct {
	Key    string `json:"key" yaml:"key"`
	Secret string `json:"secret" yaml:"secret"`
}

type KplJCQ struct {
	AccessKey    string `json:"access_key" yaml:"access_key"`
	AccessSecret string `json:"access_secret" yaml:"access_secret"`
	TenantId     int64  `json:"tenant_id" yaml:"tenant_id"`
	GroupId      int64  `json:"group_id" yaml:"group_id"`
}

type PartnerMerchant struct {
	ServiceMerchant
	StandaloneMerchant
}

// ServiceMerchant 服务商户配置
type ServiceMerchant struct {
	SpAppId       string
	SpMchId       string
	SpSerialNum   string
	SpPrivateKey  string
	SpCertificate string
	SpApiV3Key    string
}

// StandaloneMerchant 独立商户配置
type StandaloneMerchant struct {
	SubAppId   string
	SubAppName string
	SubMchId   string
	SubKey     string
}

func InitConfigData() {
	CfgData.Mode = viper.GetString("mode")
	if err := utils.ViperGetNode("db", &CfgData.Mysql); err != nil {
		panic(err)
	}
	if err := utils.ViperGetNode("mini_app", &CfgData.MiniApp); err != nil {
		panic(err)
	}
	if err := utils.ViperGetNode("kepler", &CfgData.Kepler); err != nil {
		panic(err)
	}

	cfgDataB, _ := json.Marshal(CfgData)
	log.Println("load config ", string(cfgDataB))
}
