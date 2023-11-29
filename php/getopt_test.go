package php

import (
	"github.com/heyuuu/gophp/shim/maps"
	"reflect"
	"sort"
	"testing"
)

func TestOptsParser_Each(t *testing.T) {
	type fields struct {
		args     []string
		opts     []Opt
		idx      int
		startIdx int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string][]string
	}{
		{
			"base",
			fields{
				args: []string{"-r", "echo 1;"},
				opts: []Opt{
					{'r', 1, "run"},
				},
			},
			map[string][]string{
				"run": {"echo 1;"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := map[string][]string{}
			handler := func(opt *Opt, optArg string) {
				name := opt.Name()
				result[name] = append(result[name], optArg)
			}

			p := NewOptsParser(tt.fields.args, tt.fields.opts, tt.fields.startIdx)
			p.Each(handler)
			diffOptsParserResult(t, result, tt.want)
		})
	}
}

func diffOptsParserResult(t *testing.T, result map[string][]string, want map[string][]string) {
	// sort uniq key
	keySet := map[string]bool{}
	for key := range result {
		keySet[key] = true
	}
	for key := range want {
		keySet[key] = true
	}
	keys := maps.Keys(keySet)
	sort.Strings(keys)

	// echo key
	for _, key := range keys {
		resultLine := result[key]
		wantLine := want[key]
		if !reflect.DeepEqual(resultLine, wantLine) {
			t.Errorf("opts parse error, key = %s, result = %v, want = %v", key, resultLine, wantLine)
		}
	}
}
