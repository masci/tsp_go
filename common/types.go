package common

import (
	"errors"
	"fmt"
	"math"
)

type City struct {
	x, y float64
}

func (c *City) Distance(other *City) (float64, error) {
	if other.x < 0 || other.y < 0 {
		return 0.0, errors.New("Distance only allows positive coordinates")
	}
	sx := math.Pow(other.x-c.x, 2)
	sy := math.Pow(other.y-c.y, 2)
	return math.Sqrt(sx + sy), nil
}

func NewCity(x float64, y float64) *City {
	c := City{x, y}
	return &c
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
	t._length = 0.0 // length needs to be recomputed
}

// append a new City to this Tour
func (t *Tour) Append(c City) {
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
				d, _ := c.Distance(&(t.cities[i-1]))
				t._length += d
			} else {
				d, _ := c.Distance(&(t.cities[len(t.cities)-1]))
				t._length += d
			}
		}
	}
	return t._length
}

// does a Tour contain a City?
func (t *Tour) Contains(city City) bool {
	for _, c := range t.cities {
		if c == city {
			return true
		}
	}
	return false
}

// reverse a segment of the Tour
func (t *Tour) ReverseSegmentIfBetter(i int, j int) {
	if i >= len(t.cities) || j >= len(t.cities) {
		return
	}

	t._length = 0.0 // length needs to be recomputed

	var a City
	if i == 0 {
		a = t.cities[len(t.cities)-1]

	} else {
		a = t.cities[i-1]
	}
	b := t.cities[i]
	c := t.cities[j]
	d := t.cities[(j+1)%len(t.cities)]

	ab, _ := a.Distance(&b)
	cd, _ := c.Distance(&d)
	ac, _ := a.Distance(&c)
	bd, _ := b.Distance(&d)

	PlotTour(*t, "foo")

	if (ab + cd) > (ac + bd) {
		before := t.cities[:i]
		rev := Reverse(t.cities[i : j+1])
		after := t.cities[j+1:]

		t.cities = before
		t.cities = append(t.cities, rev...)
		t.cities = append(t.cities, after...)
	}
}

func (t Tour) String() string {
	return fmt.Sprintf("Cities:%v, Length:%f", t.cities, t._length)
}

type Edge struct {
	A, B City
}

func (e Edge) Length() float64 {
	res, _ := e.A.Distance(&e.B)
	return res
}

type Edges []Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Edges) Less(i, j int) bool {
	return e[i].Length() < e[j].Length()
}

func (e Edges) String() string {
	res := ""
	for _, edge := range e {
		res += fmt.Sprintf("Edge[A:%v, B:%v] ", edge.A, edge.B)
	}
	res += "\n"
	return res
}
