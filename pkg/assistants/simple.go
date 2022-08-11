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
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", chatHistory, fmt.Errorf("chat completion error: %v", err)
	}
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: string(resp.Choices[0].Message.Content),
	})

	if verbose {
		color.Cyan("Initial response from LLM:\n%s\n\n", resp.Choices[0].Message.Content)
	}

	var toolPrompt tools.ToolPrompt
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &toolPrompt); err != nil {
		if verbose {
			color.Cyan("Unable to parse tool from prompts (%s), assuming got final answer", resp.Choices[0].Message.Content)
		}
		return resp.Choices[0].Message.Content, chatHistory, nil
	}

	iterations := 0
	maxIterations := 10
	