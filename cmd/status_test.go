package cmd

import "testing"

func TestFormatMinute(t *testing.T) {
	cases := []struct {
		name string
		args float64
		want string
	}{
		{
			name: "Must be passed",
			args: 60,
			want: "1 hours 0 minutes",
		},
		{
			name: "Must be passed",
			args: 90,
			want: "1 hours 30 minutes",
		},
		{
			name: "less than 1 hour",
			args: 10,
			want: "0 hours 10 minutes",
		},
		{
			name: "negative",
			args: -90,
			want: "-1 hours 30 minutes",
		},
		{
			name: "less than 1 minute",
			args: 0.5,
			want: "0 hours 0.5 minutes",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := formatMinute(c.args)
			if got != c.want {
				t.Errorf("want '%s' but got '%s'", c.want, got)
			}
		})
	}
}
