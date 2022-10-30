package flighttracker

import (
	"reflect"
	"testing"
)

func Test_calculateImpl(t *testing.T) {

	tests := []struct {
		name    string
		args    []FlightRoute
		want    FlightRoute
		wantErr bool
	}{
		{
			name: "scenario-1: one input",
			args: []FlightRoute{{"SFO", "EWR"}},
			want: []string{"SFO", "EWR"},
		},
		{
			name: "scenario-1: one input",
			args: []FlightRoute{{"ATL", "EWR"}, {"SFO", "ATL"}},
			want: []string{"SFO", "EWR"},
		},
		{
			name: "scenario-1: one input",
			args: []FlightRoute{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			want: []string{"SFO", "EWR"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateImpl(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
