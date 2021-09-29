package models

type Status struct {
	Status struct {
		Water int
		Wind  int
	}
}

func (s Status) WaterStatus() string {
	if s.Status.Water >= 6 && s.Status.Water <= 8 {
		return "Siaga"
	} else if s.Status.Water > 8 {
		return "Bahaya"
	}
	return "Aman"
}

func (s Status) WindStatus() string {
	if s.Status.Wind >= 7 && s.Status.Wind <= 15 {
		return "Siaga"
	} else if s.Status.Wind > 15 {
		return "Bahaya"
	}
	return "Aman"
}
