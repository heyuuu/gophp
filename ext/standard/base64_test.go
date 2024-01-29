package standard

import "testing"

func TestPhpBase64Decode(t *testing.T) {
	type args struct {
		str    string
		strict bool
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PhpBase64Decode(tt.args.str, tt.args.strict)
			if got != tt.want {
				t.Errorf("PhpBase64Decode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PhpBase64Decode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
