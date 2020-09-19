package helpers

import (
	"testing"
)

func Test_median(t *testing.T) {
	type args struct {
		vals []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Odd Number of Values",
			want: 6,
			args: args{
				vals: []float64{10, 2, 3, 4, 5, 6, 7, 8, 9, 1, 11},
			},
		},
		{
			name: "Even Number of Values",
			want: 7.5,
			args: args{
				vals: []float64{1, 2, 13, 4, 5, 6, 7, 8, 9, 10, 11, 12, 3, 14},
			},
		},
		{
			name: "Empty Input",
			want: 0,
			args: args{
				vals: []float64{},
			},
		},
		{
			name: "Single value in input",
			want: 7,
			args: args{
				vals: []float64{7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.args.vals); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slopeIntercept(t *testing.T) {
	type args struct {
		known1   float64
		voltage1 float64
		known2   float64
		voltage2 float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "basic",
			args: args{
				known1:   2,
				voltage1: 5,
				known2:   4,
				voltage2: 13,
			},
			want:  4,
			want1: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := slopeIntercept(tt.args.known1, tt.args.voltage1, tt.args.known2, tt.args.voltage2)
			if got != tt.want {
				t.Errorf("slopeIntercept() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("slopeIntercept() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
