package fuel

import "math"

type Module struct {
  Mass float64
}

func (m Module) Fuel() float64 {
  return fuelForMass(m.Mass)
}

func (m Module) TotalFuel() float64 {
  fuel := m.Fuel()
  var sum float64
  for fuel > 0 {
    sum += fuel
    fuel = fuelForMass(fuel)
  }

  return sum
}

func fuelForMass(mass float64) float64 {
  fuel := math.Floor(mass / 3)
  return fuel - 2
}