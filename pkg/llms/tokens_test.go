package llms

import (
	"testing"
)

func TestGetTokenLimits(t *testing.T) {
	type args struct {
		model string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gpt-3.5-