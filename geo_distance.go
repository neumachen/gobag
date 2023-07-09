package gobag

import (
	"math"
)

// NewGeoPoint creates a new GeoPoint with the given latitude and longitude
// values. It returns a pointer to the created GeoPoint object.
func NewGeoPoint(latitude, longitude float64) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// GeoPoint represents a geographic point with latitude and longitude
// coordinates.
type GeoPoint struct {
	// Latitude represents the angular distance of the point north or south of the equator.
	Latitude float64
	// Longitude represents the angular distance of the point east or west of the prime meridian.
	Longitude float64
}

// GreatCircleDistance calculates the great circle distance between two GeoPoint objects.
// It takes another GeoPoint as input and returns the distance in meters.
//
// The great circle distance is the shortest distance along the surface of a sphere
// between two points. This function utilizes the haversine formula to calculate the
// distance, taking into account the curvature of the Earth. It considers the latitude
// and longitude values of the two GeoPoints to determine the great circle distance.
//
// The returned distance represents the actual path that would be traveled on the
// Earth's surface between the two GeoPoints, assuming a perfect sphere. It provides
// a more accurate measure of distance than a straight-line distance, especially for
// longer distances.
//
// Example:
//
//	point1 := GeoPoint{Latitude: 52.5200, Longitude: 13.4050} // Berlin, Germany
//	point2 := GeoPoint{Latitude: 48.8566, Longitude: 2.3522}  // Paris, France
//	distance := point1.GreatCircleDistance(&point2)
//	fmt.Printf("The great circle distance between Berlin and Paris is: %.2f meters\n", distance)
func (g *GeoPoint) GreatCircleDistance(gp *GeoPoint) float64 {
	// The calculation of the difference in latitude is necessary to determine
	// the angular separation between the two GeoPoints along the north-south
	// direction. By converting the difference from degrees to radians, it
	// ensures consistency with the mathematical operations performed later in
	// the calculation.
	deltaLat := (gp.Latitude - g.Latitude) * (math.Pi / 180.0)
	// Calculate the difference in longitude between the two GeoPoints and
	// convert it to radians. Similarly, the calculation of the difference in
	// longitude is essential to determine the angular separation between the
	// two GeoPoints along the east-west direction. Converting the difference
	// from degrees to radians ensures consistency with subsequent
	// calculations.
	deltaLon := (gp.Longitude - g.Longitude) * (math.Pi / 180.0)

	// Calculate the difference in longitude between the two GeoPoints and
	// convert it to radians. Converting the latitude values of both GeoPoints
	// from degrees to radians is necessary to ensure consistency when using
	// trigonometric functions later in the calculation. These converted values
	// are used in the calculation of the haversine of the longitude
	// difference.
	lat1Rad := g.Latitude * (math.Pi / 180.0)
	lat2Rad := gp.Latitude * (math.Pi / 180.0)

	// Calculate the haversine of half the latitude difference. Convert the
	// latitude of the second GeoPoint to radians. The haversine of half the
	// latitude difference is calculated using the math.Sin function. The
	// haversine formula is commonly employed to calculate distances on a
	// sphere. This calculation represents the squared value of the sine of
	// half the latitude difference.
	haversineLat := math.Sin(deltaLat/2) * math.Sin(deltaLat/2)

	// Calculate the haversine of half the longitude difference, multiplied by
	// the cosine of the first latitude. Similar to the previous line, the
	// haversine of half the longitude difference is calculated. However, in
	// this case, the result is multiplied by the cosine of the latitude of the
	// first GeoPoint (g) and the cosine of the latitude of the second GeoPoint
	// (gp). This additional multiplication accounts for the convergence of the
	// meridians toward the poles and ensures an accurate representation of the
	// distance on a sphere.
	haversineLon := math.Sin(deltaLon/2) * math.Sin(deltaLon/2) * math.Cos(lat1Rad) * math.Cos(lat2Rad)

	// The sum of the haversine values for latitude and longitude differences
	// is calculated. This sum represents an intermediate step in the
	// determination of the central angle between the two GeoPoints.
	haversineSum := haversineLat + haversineLon

	// The central angle between the two GeoPoints is calculated using the
	// haversine formula. By applying the inverse tangent function (math.Atan2)
	// to the square root of the haversine sum divided by the square root of 1
	// - haversineSum, the central angle is determined. Multiplying the result
	// by 2 accounts for the symmetrical nature of the great circle.
	centralAngle := 2 * math.Atan2(math.Sqrt(haversineSum), math.Sqrt(1-haversineSum))

	// The definition of the Earth's radius in meters is necessary to convert
	// the central angle into an actual distance. The value 6378100 represents
	// an approximate average radius of the Earth and is commonly used for
	// calculations involving distances on Earth.
	earthRadius := 6378100

	// Finally, the great circle distance between the two GeoPoints is
	// determined by multiplying the central angle (expressed in radians) by
	// the Earth's radius. The resulting value is returned as a float64
	// representing the distance in meters.
	return float64(earthRadius) * centralAngle
}

// Distance calculates the straight-line distance between two GeoPoint objects.
// It takes another GeoPoint as input and returns the distance in meters.
//
// The distance is calculated using the haversine formula, which considers the
// spherical shape of the Earth. The function takes into account the latitude and
// longitude values of the two GeoPoints to compute the straight-line distance
// between them.
//
// The returned distance represents the shortest path between the two GeoPoints
// assuming a spherical Earth. It is not affected by the terrain or obstacles.
//
// Example:
//
//	point1 := GeoPoint{Latitude: 52.5200, Longitude: 13.4050} // Berlin, Germany
//	point2 := GeoPoint{Latitude: 48.8566, Longitude: 2.3522}  // Paris, France
//	distance := point1.Distance(&point2)
//	fmt.Printf("The straight-line distance between Berlin and Paris is: %.2f meters\n", distance)
func (g *GeoPoint) Distance(gp *GeoPoint) float64 {
	// This line declares a constant PI with a value of approximately
	// 3.141592653589793. It represents the ratio of a circle's circumference
	// to its diameter and is commonly used in trigonometric calculations. In
	// this code, PI is used for converting degrees to radians by multiplying
	// with PI / 180 later in the code.
	const PI float64 = 3.141592653589793

	// These lines convert the latitude values of both GeoPoints (g and gp) from
	// degrees to radians. The reason for this conversion is that most
	// trigonometric functions in Go's math package work with radian values. By
	// multiplying the latitude values by PI / 180, which is the conversion
	// factor from degrees to radians, we obtain the equivalent latitude values
	// in radians. These converted values are used in subsequent trigonometric
	// calculations.
	latitude1Rad := float64(PI * g.Latitude / 180)
	latitude2Rad := float64(PI * gp.Latitude / 180)

	// Here, the difference in longitude between the two GeoPoints (g and gp) is
	// calculated. The difference (g.Longitude - gp.Longitude) represents the
	// angular separation between the longitudes. It is then multiplied by PI /
	// 180 to convert the difference from degrees to radians, similar to the
	// latitude conversion. The resulting deltaLongitude represents the
	// difference in longitude in radians and is used in further calculations.
	deltaLongitude := float64(g.Longitude - gp.Longitude)
	longitudeRad := float64(PI * deltaLongitude / 180)

	// This line calculates the haversine of the angular distance between the
	// two GeoPoints. The haversine formula involves trigonometric functions
	// (sine and cosine) and is used to determine the angular separation
	// between two points on a sphere given their latitude and longitude
	// values. The haversine value is calculated based on the latitude and
	// longitude differences between g and gp in radians.
	haversine := math.Sin(latitude1Rad)*math.Sin(latitude2Rad) + math.Cos(latitude1Rad)*math.Cos(latitude2Rad)*math.Cos(longitudeRad)

	// This conditional check ensures that the haversine value stays within the
	// valid range of -1 to 1. Due to the nature of floating-point
	// calculations, the haversine value may slightly exceed 1 or fall below
	// -1, which can cause errors when attempting to calculate the inverse
	// cosine (math.Acos). Therefore, if the calculated haversine value exceeds
	// 1, it is capped at 1 to ensure the subsequent calculations proceed
	// without errors.
	if haversine > 1 {
		haversine = 1
	}

	// These lines calculate the angular distance between the two GeoPoints.
	// The math.Acos function is used to compute the arccosine of the haversine
	// value, which gives the angular distance in radians. This value is then
	// multiplied by 180 / PI to convert it from radians to degrees. Since
	// there are 60 nautical miles in one degree, the angular distance is
	// further multiplied by 60. Lastly, to convert nautical miles to statute
	// miles, the angular distance is multiplied by the conversion factor
	// 1.1515.
	//
	// The conversion factor of 1.1515 is used to convert the angular distance
	// from nautical miles to statute miles. Let's dive into the rationale
	// behind this conversion:
	//
	// In the code, the angular distance is initially calculated in nautical
	// miles. A nautical mile is a unit of measurement commonly used in
	// navigation and is defined as one minute of latitude along any meridian.
	// However, for many practical purposes, it is often desirable to express
	// distances in statute miles, which are more commonly used in everyday
	// contexts.
	//
	// To convert from nautical miles to statute miles, we multiply the angular
	// distance by the conversion factor 1.1515. This conversion factor arises
	// from the relationship between nautical miles and statute miles.
	//
	// one nautical mile is approximately equal to 1.15078 statute miles.
	// However, for simplicity and convenience, a commonly used rounded value
	// of 1.1515 is often employed in conversions.
	//
	// By multiplying the angular distance in nautical miles by 1.1515, we
	// obtain an approximate equivalent value in statute miles. The resulting
	// distance is then further multiplied by conversion factors (1.609344 and
	// 1000) to convert it to kilometers and meters, respectively, as specified
	// by the function's requirement to return the distance in meters.
	//
	// It's important to note that the exact conversion factor between nautical
	// miles and statute miles is slightly different from 1.1515, but for most
	// practical purposes, the rounded value of 1.1515 provides a reasonable
	// approximation.
	//
	// In summary, the conversion factor of 1.1515 is used to convert the
	// angular distance from nautical miles to statute miles, allowing the
	// final result to be expressed in meters as required by the function's
	// return type.
	angularDistance := math.Acos(haversine)
	angularDistance = angularDistance * 180 / PI
	angularDistance = angularDistance * 60 * 1.1515

	// This line returns the straight-line distance between the two GeoPoints
	// in meters. The angularDistance is multiplied by 1.609344 to convert
	// statute miles to kilometers, and then by 1000 to convert kilometers to
	// meters, resulting in the final distance value in meters.
	return angularDistance * 1.609344 * 1000
}
