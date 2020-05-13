package distance

import "math"

func DistanceBetween2Points(la1, lo1, la2, lo2 float64) float64 {
	R := 6371e3
	dLat := (la2 - la1) * (math.Pi / 180)
	dLon := (lo2 - lo1) * (math.Pi / 180)
	la1ToRad := la1 * (math.Pi / 180)
	la2ToRad := la2 * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(la1ToRad)*math.Cos(la2ToRad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c
	return d
}
