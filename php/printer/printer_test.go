package printer

import (
	"fmt"
	"github.com/heyuuu/gophp/php/parser"
	"log"
	"testing"
)

func init() {
	parser.SetProjRoot("../../")
}

func TestSprint(t *testing.T) {
	type args struct {
		node any
	}
	tests := []struct {
		code    string
		want    string
		wantErr bool
	}{
		{
			"<?php var_dump(1);",
			"",
			false,
		},
	}
	for i, tt := range tests {
		testName := fmt.Sprintf("case-%d", i)
		t.Run(testName, func(t *testing.T) {
			node, err := parser.ParseCode(tt.code)
			if err != nil {
				t.Errorf("Sprint() parse code error = %v", err)
				return
			}

			got, err := Sprint(node)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sprint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sprint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSprintDev(t *testing.T) {
	code := `<?php $a = 1;`
	nodes, err := parser.ParseCode(code)
	if err != nil {
		log.Fatal(err)
	}

	var p = &printer{}
	p.printNode(nodes)
	result, err := p.result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
