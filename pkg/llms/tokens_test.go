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
			name: "gpt-3.5-turbo-0613",
			args: args{
				model: "gpt-3.5-turbo-0613",
			},
			want: 4096,
		},
		{
			name: "gpt-4",
			args: