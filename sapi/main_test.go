package sapi

import (
	"reflect"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    *OptArgs
		wantErr bool
	}{
		{
			"base",
			[]string{"-d", "error_reporting=32767", "-r", "var_dump(1);"},
			&OptArgs{
				mode:       modeCliCode,
				IniAppend:  []string{"error_reporting=32767"},
				ScriptCode: "var_dump(1);",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunOpt(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"1", []string{"php", "-r", "echo 1;"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
