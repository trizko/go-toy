package strrev

import (
	"fmt"
	"testing"
)

func TestStringReverse(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"tony", "ynot"},
		{"danny", "ynnad"},
		{"joanna", "annaoj"},
		{"joanna and stephanie", "einahpets dna annaoj"},
		{"?tony. and danny!", "!ynnad dna .ynot?"},
		{"bæ", "æb"},
	}

	for _, c := range cases {
		got := StringReverse(c.in)
		if got != c.want {
			t.Errorf("FAIL StringReverse(%v)\treturns: %v\t| wanted: %v", c.in, got, c.want)
		} else {
			fmt.Printf("PASS StringReverese(%v)\t returns: %v\n", c.in, got)
		}
	}
}
