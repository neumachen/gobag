package gobag

import (
	"math"
)

// NewGeoPoint ...
func NewGeoPoint(latitude, longitude float64) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// GeoPoint ...
type GeoPoint struct {
	Latitude  float64
	Longitude float64
}

// GreatCircleDistance ...
func (g *GeoPoint) GreatCircleDistance(gp *GeoPoint) float64 {
	dLat := (gp.Latitude - g.Latitude) * (math.Pi / 180.0)
	dLon := (gp.Longitude - g.Longitude) * (math.Pi / 180.0)

	lat1 := g.Latitude * (math.Pi / 180.0)
	lat2 := gp.Latitude * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return 6378100 * c
}

// Distance ...
func (g *GeoPoint) Distance(gp *GeoPoint) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * g.Latitude / 180)
	radlat2 := float64(PI * gp.Latitude / 180)

	theta := float64(g.Longitude - gp.Longitude)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	return dist * 1.609344 * 1000
}
