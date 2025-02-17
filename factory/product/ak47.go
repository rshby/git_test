package product

type AK47 struct {
	Gun
}

// NewAK47 create new instance AK47 type of Gun. it implements interface IGun
func NewAK47() IGun {
	// create gun
	gun := &AK47{}
	gun.SetName("AK47 gun")
	gun.SetPower(4)

	return gun
}
