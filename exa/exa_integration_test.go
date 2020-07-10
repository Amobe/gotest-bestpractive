// +build integration

package exa_test

import (
	"testing"

	"github.com/amobe/gotest-bestpractice/exa"
)

// f(x) = x^2 + 2x + 1 = (x + 1)^2
func TestFormula(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "f(x) = x^2 + 2x + 1 = (x + 1)^2",
			args: args{
				x: 5,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formula_expand := func(x int) int {
				return exa.Pow(x, 2) + exa.Times(x, 2) + 1
			}
			formula_short := func(x int) int {
				return exa.Pow(x+1, 2)
			}
			if got := formula_expand(tt.args.x); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
			if got := formula_short(tt.args.x); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
