package builder

import "git_test/builder/house"

type iglooBuilder struct {
	House *house.House
}

// NewIglooBuilder creates new instance of iglooBuilder. it implements of interface HouseBuilder
func NewIglooBuilder() HouseBuilder {
	return &iglooBuilder{House: &house.House{}}
}

func (i *iglooBuilder) BuildWalls() {
	if i.House != nil {
		i.House.Wall = true
	}
}

func (i *iglooBuilder) BuildDoor() {
	//TODO implement me
	panic("implement me")
}

func (i *iglooBuilder) BuildWindows() {
	//TODO implement me
	panic("implement me")
}

func (i *iglooBuilder) BuildRoof() {
	//TODO implement me
	panic("implement me")
}

func (i *iglooBuilder) BuidGarage() {
	if i.House != nil {
		i.House.Garage = false
	}
}

func (i *iglooBuilder) GetHouse() house.House {
	return *i.House
}
