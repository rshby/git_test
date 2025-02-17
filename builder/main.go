package main

import (
	"encoding/json"
	"git_test/builder/builder"
	"git_test/builder/director"
	"github.com/sirupsen/logrus"
)

func main() {
	normalHouseBuilder := builder.NewHouseBuilder()
	normalDirector := director.NewDirector(normalHouseBuilder)

	normalHouse := normalDirector.BuildHouse()
	if marshal, err := json.Marshal(&normalHouse); err == nil {
		logrus.Infof("normal house : %s", string(marshal))
	}

	iglooHouseBuilder := builder.NewIglooBuilder()
	iglooDirector := director.NewDirector(iglooHouseBuilder)

	iglooHouse := iglooDirector.BuildHouse()
	if marshal, err := json.Marshal(&iglooHouse); err == nil {
		logrus.Infof("igloo house : %s", string(marshal))
	}
}
