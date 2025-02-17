package factory

import (
	"fmt"
	"git_test/factory/product"
)

// GetGun gets new gun by given type
func GetGun(gunType string) (product.IGun, error) {
	switch gunType {
	case "ak47":
		return product.NewAK47(), nil
	case "musket":
		return product.NewMusket(), nil
	}

	return nil, fmt.Errorf("invalid gun type")
}
