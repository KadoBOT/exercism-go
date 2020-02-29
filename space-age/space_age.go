package space

type Planet string

var planets = map[Planet]float64 {
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Earth": 1,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

const earthSeconds = 31557600

func (p Planet) toYears(s float64) float64 {
	return (s/planets[p])/earthSeconds
}

func Age(s float64, p Planet) float64 {
	return p.toYears(s)
}
