package solutions

import (
	utils "adventofcode/utils"
	"testing"
)

func Test_ship_Move(t *testing.T) {
	type args struct {
		d utils.Vector
		v int
	}
	tests := []struct {
		name             string
		s                ship
		args             args
		expectedPosition utils.Vector
	}{
		{
			"zero direction",
			NewShip(),
			args{
				utils.NewZeroVector(),
				0,
			},
			utils.NewZeroVector(),
		},
		{
			"zero direction, movement",
			NewShip(),
			args{
				utils.NewZeroVector(),
				10,
			},
			utils.NewZeroVector(),
		},
		{
			"slide north",
			NewShip(),
			args{
				North,
				10,
			},
			utils.NewVector(0, 10, 0),
		},
		{
			"slide south",
			NewShip(),
			args{
				South,
				5,
			},
			utils.NewVector(0, -5, 0),
		},
		{
			"slide east",
			NewShip(),
			args{
				East,
				12,
			},
			utils.NewVector(12, 0, 0),
		},
		{
			"slide west",
			NewShip(),
			args{
				West,
				2,
			},
			utils.NewVector(-2, 0, 0),
		},
		{
			"slide neagtive north",
			NewShip(),
			args{
				North,
				-2,
			},
			utils.NewVector(0, -2, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Move(tt.args.d, tt.args.v)
			if !utils.VecEqual(tt.s.Position(), tt.expectedPosition) {
				t.Errorf("Unexpected Position")
			}
			if !utils.VecEqual(tt.s.Facing(), utils.NewVector(1, 0, 0)) {
				t.Errorf("Unexpected Facing")
			}
		})
	}
}

func Test_ship_Forward(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name             string
		s                ship
		args             args
		expectedPosition utils.Vector
	}{
		{
			"forward zero",
			NewShip(),
			args{
				0,
			},
			utils.NewVector(0, 0, 0),
		},
		{
			"forward some",
			NewShip(),
			args{
				10,
			},
			utils.NewVector(10, 0, 0),
		},
		{
			"forward negative",
			NewShip(),
			args{
				-6,
			},
			utils.NewVector(-6, 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Forward(tt.args.v)
			if !utils.VecEqual(tt.s.Position(), tt.expectedPosition) {
				t.Errorf("Unexpected Position")
			}
			if !utils.VecEqual(tt.s.Facing(), utils.NewVector(1, 0, 0)) {
				t.Errorf("Unexpected Facing")
			}
		})
	}
}

func Test_ship_Turn(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name           string
		s              ship
		args           args
		expectedFacing utils.Vector
	}{
		{
			"zero turn",
			NewShip(),
			args{
				0,
			},
			East,
		},
		{
			"90 degrees",
			NewShip(),
			args{
				90,
			},
			South,
		},
		{
			"180 degrees",
			NewShip(),
			args{
				180,
			},
			West,
		},
		{
			"270 degrees",
			NewShip(),
			args{
				270,
			},
			North,
		},
		{
			"360 degrees",
			NewShip(),
			args{
				360,
			},
			East,
		},
		{
			"-90 degrees",
			NewShip(),
			args{
				-90,
			},
			North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Turn(tt.args.v)
			if !utils.VecEqual(tt.s.Position(), utils.NewZeroVector()) {
				t.Errorf("Unexpected Position")
			}
			if !utils.VecEqual(tt.s.Facing(), tt.expectedFacing) {
				t.Errorf("Unexpected Facing")
			}
		})
	}
}
