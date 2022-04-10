
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/assistants"
	"github.com/feiskyer/openai-copilot/pkg/consts"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	// global flags
	model, prompt string
	maxTokens     int
	countTokens   bool
	verbose       bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "openai-copilot",
		Short: "OpenAI Copilot",
		Run: func(cmd *cobra.Command, args []string) {
			chat()
		},
	}
)

func chat() {
	var err error
	var response string
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,