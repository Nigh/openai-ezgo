package openaiezgo

import (
	openai "github.com/sashabaranov/go-openai"
)

type chatHistory struct {
	BaseMemory []openai.ChatCompletionMessage
	WorkMemory []openai.ChatCompletionMessage
}

type ChatInstance struct {
	Timeout   int
	History   *chatHistory
	MaxTokens int
	TokenUsed int
}

var Chats map[string]ChatInstance

func init() {
	Chats = make(map[string]ChatInstance)
}

func (c *ChatInstance) getAllHistory() []openai.ChatCompletionMessage {
	return append(c.History.BaseMemory, c.History.WorkMemory...)
}
