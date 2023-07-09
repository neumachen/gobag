package gobag

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeoPoint_GreatCircleDistance(t *testing.T) {
	// Define test cases with expected distances in meters.
	testCases := []struct {
		name     string
		point1   GeoPoint
		point2   GeoPoint
		expected float64
	}{
		{
			name:     "Berlin to Paris",
			point1:   GeoPoint{Latitude: 52.5200, Longitude: 13.4050}, // Berlin, Germany
			point2:   GeoPoint{Latitude: 48.8566, Longitude: 2.3522},  // Paris, France
			expected: 878441.19,
		},
		{
			name:     "New York to Los Angeles",
			point1:   GeoPoint{Latitude: 40.7128, Longitude: -74.0060},  // New York, USA
			point2:   GeoPoint{Latitude: 34.0522, Longitude: -118.2437}, // Los Angeles, USA
			expected: 3940132.35,
		},
		{
			name:     "Tokyo to Sydney",
			point1:   GeoPoint{Latitude: 35.6895, Longitude: 139.6917},  // Tokyo, Japan
			point2:   GeoPoint{Latitude: -33.8651, Longitude: 151.2099}, // Sydney, Australia
			expected: 7834941.05,
		},
		{
			name:     "London to Moscow",
			point1:   GeoPoint{Latitude: 51.5074, Longitude: -0.1278}, // London, UK
			point2:   GeoPoint{Latitude: 55.7558, Longitude: 37.6176}, // Moscow, Russia
			expected: 2503341.40,
		},
		{
			name:     "Cape Town to Cairo",
			point1:   GeoPoint{Latitude: -33.9249, Longitude: 18.4241}, // Cape Town, South Africa
			point2:   GeoPoint{Latitude: 30.0444, Longitude: 31.2357},  // Cairo, Egypt
			expected: 7247314.54,
		},
		{
			name:     "Rio de Janeiro to Buenos Aires",
			point1:   GeoPoint{Latitude: -22.9068, Longitude: -43.1729}, // Rio de Janeiro, Brazil
			point2:   GeoPoint{Latitude: -34.6037, Longitude: -58.3816}, // Buenos Aires, Argentina
			expected: 1969966.27,
		},
		{
			name:     "Sydney to Auckland",
			point1:   GeoPoint{Latitude: -33.8651, Longitude: 151.2099}, // Sydney, Australia
			point2:   GeoPoint{Latitude: -36.8485, Longitude: 174.7633}, // Auckland, New Zealand
			expected: 2158358.12,
		},
		{
			name:     "Santiago to Mexico City",
			point1:   GeoPoint{Latitude: -33.4489, Longitude: -70.6693}, // Santiago, Chile
			point2:   GeoPoint{Latitude: 19.4326, Longitude: -99.1332},  // Mexico City, Mexico
			expected: 6617464.43,
		},
		{
			name:     "Moscow to Beijing",
			point1:   GeoPoint{Latitude: 55.7558, Longitude: 37.6176},  // Moscow, Russia
			point2:   GeoPoint{Latitude: 39.9042, Longitude: 116.4074}, // Beijing, China
			expected: 5800239.16,
		},
		{
			name:     "Nairobi to Johannesburg",
			point1:   GeoPoint{Latitude: -1.2864, Longitude: 36.8172},  // Nairobi, Kenya
			point2:   GeoPoint{Latitude: -26.2041, Longitude: 28.0473}, // Johannesburg, South Africa
			expected: 2928545.00,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distance := tc.point1.GreatCircleDistance(&tc.point2)

			// Compare the calculated distance with the expected distance, allowing a tolerance of 0.01 meters.
			if math.Abs(distance-tc.expected) > 0.01 {
				require.Failf(t, "Unexpected distance", "Expected: %.2f meters, Got: %.2f meters", tc.expected, distance)
			}
		})
	}
}

func TestGeoPoint_Distance(t *testing.T) {
	// Define test cases with expected distances in meters.
	testCases := []struct {
		name     string
		point1   GeoPoint
		point2   GeoPoint
		expected float64
	}{
		{
			name:     "Berlin to Paris",
			point1:   GeoPoint{Latitude: 52.5200, Longitude: 13.4050}, // Berlin, Germany
			point2:   GeoPoint{Latitude: 48.8566, Longitude: 2.3522},  // Paris, France
			expected: 877421.11,
		},
		{
			name:     "New York to Los Angeles",
			point1:   GeoPoint{Latitude: 40.7128, Longitude: -74.0060},  // New York, USA
			point2:   GeoPoint{Latitude: 34.0522, Longitude: -118.2437}, // Los Angeles, USA
			expected: 3935556.90,
		},
		{
			name:     "Tokyo to Sydney",
			point1:   GeoPoint{Latitude: 35.6895, Longitude: 139.6917},  // Tokyo, Japan
			point2:   GeoPoint{Latitude: -33.8651, Longitude: 151.2099}, // Sydney, Australia
			expected: 7825842.79,
		},
		{
			name:     "London to Moscow",
			point1:   GeoPoint{Latitude: 51.5074, Longitude: -0.1278}, // London, UK
			point2:   GeoPoint{Latitude: 55.7558, Longitude: 37.6176}, // Moscow, Russia
			expected: 2500434.42,
		},
		{
			name:     "Cape Town to Cairo",
			point1:   GeoPoint{Latitude: -33.9249, Longitude: 18.4241}, // Cape Town, South Africa
			point2:   GeoPoint{Latitude: 30.0444, Longitude: 31.2357},  // Cairo, Egypt
			expected: 7238898.66,
		},
		{
			name:     "Rio de Janeiro to Buenos Aires",
			point1:   GeoPoint{Latitude: -22.9068, Longitude: -43.1729}, // Rio de Janeiro, Brazil
			point2:   GeoPoint{Latitude: -34.6037, Longitude: -58.3816}, // Buenos Aires, Argentina
			expected: 1967678.66,
		},
		{
			name:     "Sydney to Auckland",
			point1:   GeoPoint{Latitude: -33.8651, Longitude: 151.2099}, // Sydney, Australia
			point2:   GeoPoint{Latitude: -36.8485, Longitude: 174.7633}, // Auckland, New Zealand
			expected: 2155851.75,
		},
		{
			name:     "Santiago to Mexico City",
			point1:   GeoPoint{Latitude: -33.4489, Longitude: -70.6693}, // Santiago, Chile
			point2:   GeoPoint{Latitude: 19.4326, Longitude: -99.1332},  // Mexico City, Mexico
			expected: 6609779.95,
		},
		{
			name:     "Moscow to Beijing",
			point1:   GeoPoint{Latitude: 55.7558, Longitude: 37.6176},  // Moscow, Russia
			point2:   GeoPoint{Latitude: 39.9042, Longitude: 116.4074}, // Beijing, China
			expected: 5793503.68,
		},
		{
			name:     "Nairobi to Johannesburg",
			point1:   GeoPoint{Latitude: -1.2864, Longitude: 36.8172},  // Nairobi, Kenya
			point2:   GeoPoint{Latitude: -26.2041, Longitude: 28.0473}, // Johannesburg, South Africa
			expected: 2925144.2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			distance := tc.point1.Distance(&tc.point2)

			// Compare the calculated distance with the expected distance, allowing a tolerance of 0.01 meters.
			if math.Abs(distance-tc.expected) > 0.01 {
				require.Failf(t, "Unexpected distance.", "Expected: %.2f meters, Got: %.2f meters", tc.expected, distance)
			}
		})
	}
}
