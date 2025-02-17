package product

type Musket struct {
	Gun
}

// NewMusket creates new instance Musket type of Gun. it implements interface IGun
func NewMusket() IGun {
	// create gun
	gun := &Musket{}
	gun.SetName("Musket gun")
	gun.SetPower(1)

	return gun
}
