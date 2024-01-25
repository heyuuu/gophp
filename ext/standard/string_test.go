package standard

import (
	"testing"
)

func Test_charmaskEx(t *testing.T) {
	type args struct {
		input   string
		onError func(string)
	}
	tests := []struct {
		name  string
		want  string
		want1 bool
	}{
		{"abc", "abc", true},
		{"a..c", "abc", true},
		{"abc..cd", "abcd", true},
		{"abc..ad", "abc.ad", false},
		{"z..A", "z.A", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := charmaskEx(tt.name, func(s string) {})
			if got != tt.want {
				t.Errorf("charmaskEx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("charmaskEx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestZifBin2hex(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{string([]byte{128})}, "80"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZifBin2hex(tt.args.data); got != tt.want {
				t.Errorf("ZifBin2hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestPhpStrToStrI(t *testing.T) {
//	type args struct {
//		haystack   string
//		lcHaystack string
//		needle     string
//		str        string
//	}
//	tests := []struct {
//		name  string
//		args  args
//		want  string
//		want1 int
//	}{
//		{
//			"1",
//			args{
//				`/*<b> I am a comment</b>*/`,
//				`/*<B>`,
//				`<B>`,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := PhpStrToStrIEx(tt.args.haystack, tt.args.lcHaystack, tt.args.needle, tt.args.str)
//			if got != tt.want {
//				t.Errorf("PhpStrToStrIEx() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("PhpStrToStrIEx() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}

func Test_stringReplaceIgnoreCase(t *testing.T) {
	type args struct {
		s       string
		search  string
		replace string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			"1",
			args{
				`/*<b> I am a comment</b>*/`,
				`/*<B>`,
				`<B>`,
			},
			`<B> I am a comment</b>*/`,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := stringReplaceIgnoreCase(tt.args.s, tt.args.search, tt.args.replace)
			if got != tt.want {
				t.Errorf("stringReplaceIgnoreCase() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("stringReplaceIgnoreCase() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
