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
			args: args{
				model: "gpt-4",
			},
			want: 8192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenLimits(tt.args.model); got != tt.want {
				t.Errorf("GetTokenLimits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstrictPrompt(t *testing.T) {
	type args struct {
		prompt    string
		model     string
		maxTokens int
	}
	tests := []struct {
		name string
		args args
		want s