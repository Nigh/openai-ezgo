package openaiezgo

import (
	"context"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client
var config NodeConfig

type NodeConfig struct {
	APIKey          string
	BaseURL         string
	Timeout         int
	TimeoutCallback func(string, int)
	MaxTokens       int
	HistoryLimit    int
}

func DefaultConfig(authToken string) NodeConfig {
	return NodeConfig{
		APIKey:       authToken,
		BaseURL:      "",
		Timeout:      300,
		MaxTokens:    768,
		HistoryLimit: 10,
	}
}
func NewClientWithConfig(config NodeConfig) {
	config = DefaultConfig(config.APIKey)
	gptConfig := openai.DefaultConfig(config.APIKey)
	if config.BaseURL != "" {
		gptConfig.BaseURL = config.BaseURL
	}
	client = openai.NewClientWithConfig(gptConfig)
}

func NewClient(apiKey string) {
	NewClientWithConfig(DefaultConfig(apiKey))
}

func getChat(from string) {
	if _, ok := Chats[from]; !ok {
		Chats[from] = ChatInstance{
			Timeout:   config.Timeout,
			History:   &chatHistory{[]openai.ChatCompletionMessage{}, []openai.ChatCompletionMessage{}},
			TokenUsed: 0,
		}
	}
}
func NewCharacterSet(from string, words string) string {
	getChat(from)
	thisChat := Chats[from]
	thisChat.History.BaseMemory = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: words,
		},
	}
	return "调教指令已保存，最后一次调教设置将会持续保留并置于对话记忆的最开始处。直到对话重置。"
}
func NewSpeech(from string, words string) string {
	getChat(from)

	thisChat := Chats[from]
	thisChat.History.WorkMemory = append(thisChat.History.WorkMemory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: words,
	})
	thisChat.Timeout = config.Timeout
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens: config.MaxTokens,
			Model:     openai.GPT3Dot5Turbo,
			Messages:  thisChat.getAllHistory(),
		},
	)
	for len(thisChat.History.WorkMemory) > config.HistoryLimit {
		thisChat.History.WorkMemory = thisChat.History.WorkMemory[1:]
	}
	if err != nil {
		return "报个错先：" + err.Error()
	}
	thisChat.TokenUsed += resp.Usage.TotalTokens
	return resp.Choices[0].Message.Content
}

func EndSpeech(from string) string {
	TokenUsed := Chats[from].TokenUsed
	delete(Chats, from)
	return "对话已重置，记忆已清空。共消耗token：" + strconv.Itoa(TokenUsed)
}
