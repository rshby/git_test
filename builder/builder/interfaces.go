package builder

import "git_test/builder/house"

type HouseBuilder interface {
	BuildWalls()
	BuildDoor()
	BuildWindows()
	BuildRoof()
	BuidGarage()
	GetHouse() house.House
}
