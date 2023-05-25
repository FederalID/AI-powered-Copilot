package tools

// Tool is a function that takes an input and returns an output.
type Tool func(input string) (string, error)

// CopilotTools is a map of tool names to too