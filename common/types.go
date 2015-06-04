package common

import (
	"errors"
	"math"
)

type City struct {
	x, y float64
}

func (c *City) distance(other *City) (float64, error) {
	if other.x < 0 || other.y < 0 {
		return 0.0, errors.New("Distance only allows positive coordinates")
	}
	sx := math.Pow(other.x-c.x, 2)
	sy := math.Pow(other.y-c.y, 2)
	return math.Sqrt(sx + sy), nil
}

type Tour struct {
	cities  []City
	_length float64
}

// Tour constructor
func NewTour(cities []City) *Tour {
	t := Tour{cities: cities, _length: 0.0}
	return &t
}

// swap two cities from a tour
func (t *Tour) Swap(i int, j int) {
	tmp := t.cities[i]
	t.cities[i] = t.cities[j]
	t.cities[j] = tmp
}

// append a new City to this Tour
func (t *Tour) append(c City) {
	t.cities = append(t.cities, c)
	t._length = 0.0 // length needs to be recomputed
}

func (t *Tour) Cities() []City {
	return t.cities
}

// replace the array of cities
func (t *Tour) setCities(cities []City) {
	t.cities = cities
	t._length = 0.0
}

// compute Tour length
func (t *Tour) Length() float64 {
	if t._length == 0.0 {
		for i, c := range t.cities {
			if i > 0 {
				d, _ := c.distance(&(t.cities[i-1]))
				t._length += d
			} else {
				d, _ := c.distance(&(t.cities[len(t.cities)-1]))
				t._length += d
			}
		}
	}
	return t._length
}

// does a Tour contain a City?
func (t *Tour) contains(city City) bool {
	for _, c := range t.cities {
		if c == city {
			return true
		}
	}
	return false
}
