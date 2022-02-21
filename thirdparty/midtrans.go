package thirdparty

import (
	"Back-End-Ecommers/configs"

	"github.com/veritrans/go-midtrans"
)

func GetMidtransCoreGateway() midtrans.CoreGateway {
	midclient := midtrans.NewClient()
	midclient.ServerKey = configs.Conf.MidtransServerKey
	midclient.ClientKey = configs.Conf.MidtransClientKey
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	return coreGateway
}
