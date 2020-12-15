package solutions

import (
	"reflect"
	"testing"
)

func Test_seatPlan_setState(t *testing.T) {
	type args struct {
		x        int
		y        int
		newState state
	}
	tests := []struct {
		name string
		s    *seatPlan
		args args
	}{
		{
			"Set Occupied",
			NewSeatPlan(5, 5),
			args{
				1,
				2,
				Occupied,
			},
		},
		{
			"Set Empty",
			NewSeatPlan(5, 5),
			args{
				0,
				0,
				Empty,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.s.state(tt.args.x, tt.args.y) != Unset {
				t.Errorf("Init error")
			}
			tt.s.setState(tt.args.x, tt.args.y, tt.args.newState)
			if tt.s.state(tt.args.x, tt.args.y) != tt.args.newState {
				t.Errorf("Did not set state")
			}
		})
	}
}

func Test_valueToState(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name  string
		args  args
		wantS state
	}{
		{
			"occupied",
			args{
				'#',
			},
			Occupied,
		},
		{
			"floor",
			args{
				'.',
			},
			Floor,
		},
		{
			"empty",
			args{
				'L',
			},
			Empty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := valueToState(tt.args.r); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("valueToState() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
