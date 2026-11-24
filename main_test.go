package main

import "testing"

func TestCountAndSay(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{0, ""},
	}
	for _, c := range cases {
		if got := CountAndSay(c.n); got != c.want {
			t.Errorf("CountAndSay(%d) = %q, want %q", c.n, got, c.want)
		}
	}
}
