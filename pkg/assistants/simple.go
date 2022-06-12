package assistants

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/llms"
	"github.com/feiskyer/openai-copilot/pkg/tools"
	"github.com/sashabaranov/go-openai"
)

// Assistant is the simplest AI assistant.
func Assistant(model string, prompts []openai.ChatCompletionMessage, maxTokens int, countTokens bool, verbose bool) (result string, chatHistory []openai.ChatCompletionMessage, err error) {
	chatHistory = prompts
	if len(prompts) == 0 {
		return "", nil, fmt.Errorf("prompts cannot be empty")
	}

	client, err := llms.NewOpenAIClient()
	if err != nil {
		return "", nil, fmt.Errorf("unable to get OpenAI client: %v", err)
	}

	defer func() {
		if countTokens {
			count := llms.NumTokensFromMessages(chatHistory, model)
			color.Green("Total tokens: %d\n\n", count)
		}
	}()

	req := openai.ChatCompletionRequest{
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: math.SmallestNonzeroFloat32,
		Messages:    chatHistory,
	}
	resp, err := client.CreateChatCompletio