package main

import (
	"git_test/factory/factory"
	"git_test/factory/product"
	"github.com/sirupsen/logrus"
)

func main() {
	// get gun
	ak47Gun, err := factory.GetGun("ak47")
	if err != nil {
		logrus.Error(err)
	}

	musketGun, err := factory.GetGun("musket")
	if err != nil {
		logrus.Error(err)
	}

	PrintDetail(ak47Gun)
	PrintDetail(musketGun)

}

func PrintDetail(request product.IGun) {
	logrus.Infof("get gun %s with power %d", request.GetName(), request.GetPower())
}
