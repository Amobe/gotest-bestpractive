package exa_test

import (
	"testing"

	"github.com/amobe/gotest-bestpractice/exa"
)

func TestTimes(t *testing.T) {
	type args struct {
		a     int
		times int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Zero times",
			args: args{
				a:     2,
				times: 0,
			},
			want: 0,
		},
		{
			name: "Many times",
			args: args{
				a:     2,
				times: 5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exa.Times(tt.args.a, tt.args.times); got != tt.want {
				t.Errorf("Times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		a   int
		pow int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Zero pow",
			args: args{
				a:   2,
				pow: 0,
			},
			want: 1,
		},
		{
			name: "10 pow",
			args: args{
				a:   2,
				pow: 10,
			},
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exa.Pow(tt.args.a, tt.args.pow); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
