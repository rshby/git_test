package builder

import "git_test/builder/house"

type houseBuilder struct {
	House *house.House
}

// NewHouseBuilder creates new instanfe of houseBuilder. it implements interface HouseBuilder
func NewHouseBuilder() HouseBuilder {
	h := &house.House{}
	return &houseBuilder{h}
}

func (h *houseBuilder) BuildWalls() {
	if h.House != nil {
		h.House.Wall = true
	}
}

func (h *houseBuilder) BuildDoor() {
	//TODO implement me
	panic("implement me")
}

func (h *houseBuilder) BuildWindows() {
	//TODO implement me
	panic("implement me")
}

func (h *houseBuilder) BuildRoof() {
	//TODO implement me
	panic("implement me")
}

func (h *houseBuilder) BuidGarage() {
	if h.House != nil {
		h.House.Garage = true
	}
}

func (h *houseBuilder) GetHouse() house.House {
	return *h.House
}
