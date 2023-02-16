
package llms

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

var tokenLimitsPerModel = map[string]int{
	"code-davinci-002":       4096,
	"gpt-3.5-turbo-0301":     4096,
	"gpt-3.5-turbo-0613":     4096,