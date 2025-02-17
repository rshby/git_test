package director

import (
	"git_test/builder/builder"
	"git_test/builder/house"
)

type Director struct {
	builder builder.HouseBuilder
}

func NewDirector(b ...builder.HouseBuilder) *Director {
	d := &Director{}
	if len(b) > 0 {
		d.SetBuilder(b[0])
	}

	return d
}

func (d *Director) SetBuilder(b builder.HouseBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() house.House {
	// create walls
	d.builder.BuildWalls()
	d.builder.BuidGarage()
	return d.builder.GetHouse()
}
