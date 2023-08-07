# AI-powered Copilot

Your life Copilot powered by OpenAI (CLI interface for OpenAI with searching).

Features:

* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Human interactions on uncertain instructions to avoid inappropriate operations.

## Install

Install the copilot with the commands below:

```sh
go install github.com/FederalID/AI-powered-Copilot/cmd/AI-powered-Copilot
```

## Setup

* OpenAI API key should be set to `OPENAI_API_KEY` environment variable to enable the ChatGPT feature.
  * `OPENAI_API_BASE` should be set as well for Azure OpenAI service and other self-hosted OpenAI services.
* Google Search API key and CSE ID should be set to `GOOGLE_API_KEY` and `GOOGLE_CSE_ID`.

## How to use

```sh
Usage:
  AI-powered-Copilot [flags]

Flags:
  -c, --count-tokens     Print tokens count
  -h, --help             help for AI-powered-Copilot
  -t, --max-tokens int   Max tokens for the GPT model (default 1024)
  -m, --model string     OpenAI model to use (default "gpt-4")
  -p, --prompt string    Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used
  -v, --verbose          Enable verbose output (default true)
```

### Interactive mode

Here is a conversation sample (user inputs are after `You:`)):

```sh
$ AI-powered-Copilot --verbose=false
You: What is OpenAI?
AI: OpenAI is an artificial intelligence research lab, which includes a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to build safe and beneficial AGI, and are also committed to aiding others in achieving this outcome.

You:
```

### Non-interactive mode

```sh
$ AI-powered-Copilot -p 'What is OpenAI?'
Initial response from LLM:
{
 "question": "What is OpenAI?",
 "thought": "OpenAI is a well-known organization in the field of artificial intelligence. I should provide a brief description of it.",
 "action": {
  "name": "search",
  "input": "OpenAI"
 },
 "observation": "OpenAI is an artificial intelligence research lab consisting of the for-profit arm OpenAI LP and its parent company, the non-profit OpenAI Inc. OpenAI's mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to build safe and beneficial AGI directly, but are also committed to aiding others in achieving this outcome.",
 "final_answer": "OpenAI is an artificial intelligence research lab made up of a for-profit arm, OpenAI LP, and its parent company, the non-profit OpenAI Inc. Their mission is to ensure that artificial general intelligence (AGI) benefits all of humanity. They aim to directly build safe and beneficial AGI, but are also committed to aiding others in achieving this outcome."
}

Thought: OpenAI is a well-known organization in the field of artificial intelligence. I should provide a brief description of it.

Final answer: OpenAI is an artificial intelligence research lab made up of a for-prof