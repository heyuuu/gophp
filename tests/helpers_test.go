package tests

import (
	"testing"
	"time"
)

func Test_timeFormat(t *testing.T) {
	tim := time.Date(2010, 5, 20, 15, 44, 55, 0, time.Local)
	type args struct {
		t   time.Time
		fmt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{t: tim, fmt: "Ymd_His"},
			"20100520_154455",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeFormat(tt.args.t, tt.args.fmt); got != tt.want {
				t.Errorf("timeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
