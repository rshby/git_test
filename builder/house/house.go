package house

type House struct {
	Wall   bool `json:"wall"`
	Door   bool `json:"door"`
	Window bool `json:"window"`
	Roof   bool `json:"roof"`
	Garage bool `json:"garage"`
}
