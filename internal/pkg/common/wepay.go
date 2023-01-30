package common

import (
	"context"

	"github.com/fyf2173/ysdk-go/wepay"
)

var WXPay wepay.PartnerMerchantClient

func InitWXPayClient(conf wepay.PartnerMerchantConfig) {
	WXPay = *wepay.NewPartnerMerchantClient(context.Background(), conf)
}
