package service

import (
	"reflect"
	"testing"
)

var (
	getLocationtests = []struct {
		x        float64
		y        float64
		z        float64
		vel      float64
		sectorID int
		expected float32
	}{
		{123.12, 456.56, 789.89, 20.0, 1, 1389.57},
		{123.12, 456.56, 789.89, 10.0, 1, 1379.57},
	}
	responseTests = []struct {
		input interface{}
		want  interface{}
	}{
		{1389.57, AtlasCorp{1389.57}},
	}
)

func TestAtlasCorp_GetLocation(t *testing.T) {
	corp := Location{LocationService: AtlasCorp{}}
	for _, tt := range getLocationtests {
		got := corp.GetLocation(tt.x, tt.y, tt.z, tt.vel, tt.sectorID)
		if got != tt.expected {
			t.Errorf("AtlasCorp.GetLocation() = %v, want %v", got, tt.expected)
		}
	}
}

func TestAtlasCorp_GetResponse(t *testing.T) {

	corp := Location{LocationService: AtlasCorp{}}
	for _, tt := range responseTests {
		if got := corp.GetResponse(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("AtlasCorp.GetResponse() = %v, want %v", got, tt.want)
		}
	}
}
