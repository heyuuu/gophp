package sapi

import (
	"os"
	"testing"
)

func TestRun1(t *testing.T) {
	_ = os.Chdir("../")
	args := []string{"php", "-r", "echo 1;"}
	Run(args)
}
