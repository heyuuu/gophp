package tests

import (
	"reflect"
	"testing"
)

func Test_runPhpScript(t *testing.T) {
	type args struct {
		testFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *scriptResult
		wantErr bool
	}{
		{
			"1",
			args{
				testFile: "/Users/heyu/Code/src/php-7.4.33/tests/run-test/bug75042-2.phpt",
			},
			&scriptResult{
				Output: "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runPhpScript(tt.args.testFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("runPhpScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runPhpScript() got = %v, want %v", got, tt.want)
			}
		})
	}
}
