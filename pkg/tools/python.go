package tools

import (
	"os/exec"
	"strings"
)

// PythonREPL runs the given Python script and returns the output.
func PythonREPL(script string) (string, error) {
	cmd 