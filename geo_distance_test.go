package gobag

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoDistance(t *testing.T) {
	chicagoLat := 41.881832
	chicagoLong := -87.623177
	madisonLat := 43.073051
	madisonLong := -89.401230

	gp := NewGeoPoint(chicagoLat, chicagoLong)

	result := gp.GreatCircleDistance(NewGeoPoint(madisonLat, madisonLong))
	assert.Equal(t, float64(197209.38331706668), result)

	result = gp.Distance(NewGeoPoint(madisonLat, madisonLong))
	assert.Equal(t, float64(196980), math.Round(result))
}
