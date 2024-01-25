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

func TestStripTagsEx(t *testing.T) {
	type args struct {
		str       string
		state     uint8
		allowTags string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 uint8
	}{
		{
			"1",
			args{
				`<a href="test?test\!!!test">test</a><!-- test -->`,
				0,
				"",
			},
			"test",
			0,
		},
		{
			"2",
			args{
				`<a href="test?test\!!!test">test</a><!-- test -->`,
				0,
				"<a>",
			},
			`<a href="test?test\!!!test">test</a>`,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StripTagsEx(tt.args.str, tt.args.state, tt.args.allowTags)
			if got != tt.want {
				t.Errorf("StripTagsEx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StripTagsEx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_phpTagFind(t *testing.T) {
	type args struct {
		tag string
		set string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				`<a href="test?test\!!!test">`,
				`<a>`,
			},
			true,
		},
		{
			"2",
			args{
				`</a>`,
				`<a>`,
			},
			true,
		},
		{
			"3",
			args{
				`<a>`,
				`<a>`,
			},
			true,
		},
		{
			"bad case",
			args{
				`/a>`,
				`<a>`,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := phpTagFind(tt.args.tag, tt.args.set); got != tt.want {
				t.Errorf("phpTagFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrtr(t *testing.T) {
	type args struct {
		str  string
		from string
		to   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"bug-01",
			args{
				"Dot in brackets [.]\\n",
				".",
				"0",
			},
			"Dot in brackets [0]\\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strtr(tt.args.str, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Strtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
