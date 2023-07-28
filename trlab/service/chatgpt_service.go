package service

import (
	"context"
	"errors"
	"fmt"
	"trlab-backend-go/api_service"
	"trlab-backend-go/config"
	"trlab-backend-go/util"
	"trlab-backend-go/util/db"

	"github.com/sashabaranov/go-openai"
)

const (
	DefaultGPTModel    = openai.GPT3Dot5Turbo16K // nax tokens: 16384
	DefaultStream      = true
	TotalTokens        = 8192 // Tokens from the prompt(message) and the completion all together should not exceed the token limit of a particular GPT-3 model.
	ContextTokens      = 2000
	AnswerTokens       = 1000
	AssistantToken     = TotalTokens - ContextTokens - AnswerTokens
	DefaultTemperature = 0.8
	TopN               = 10
)

// TODO: wait GPT4.0
//const (
//	DefaultGPTModel    = openai.GPT4 // max tokens: 8192
//	DefaultStream      = true
//	TotalTokens        = 8192 // Tokens from the prompt(message) and the completion all together should not exceed the token limit of a particular GPT-3 model.
//	ContextTokens      = 2000
//	AnswerTokens       = 1000
//	AssistantToken     = TotalTokens - ContextTokens - AnswerTokens
//	DefaultTemperature = 0.8
//	TopN               = 5
//)

type ChatBot struct {
	Client     *openai.Client
	Session    *Session
	Stream     *openai.ChatCompletionStream
	ChatConfig ChatConfig
}

type ChatConfig struct {
	ApiKey      string
	Model       string
	Tokens      int
	Stream      bool
	Temperature float32
}

func NewChatBot(apiKey, model string, stream bool, tokens int, temperature float32) *ChatBot {
	chatConfig := ChatConfig{
		ApiKey:      apiKey,
		Model:       model,
		Tokens:      tokens,
		Stream:      stream,
		Temperature: temperature,
	}
	return &ChatBot{Client: openai.NewClient(apiKey), Session: NewSession(), ChatConfig: chatConfig}
}

// DefaultChatBot base on gpt-3.5
func DefaultChatBot() *ChatBot {
	return NewChatBot(config.GetConfig().OpenAiKey(), DefaultGPTModel, DefaultStream, AnswerTokens, DefaultTemperature)
}

func (chatBot *ChatBot) Receive() (resp openai.ChatCompletionStreamResponse, err error) {
	if chatBot.Stream == nil {
		err = errors.New("chat Stream is nil")
		return
	}

	return chatBot.Stream.Recv()
}

func (chatBot *ChatBot) Query(question string) (resp openai.ChatCompletionResponse, err error) {
	chatReq := openai.ChatCompletionRequest{
		Model:       chatBot.ChatConfig.Model,
		Messages:    chatBot.CreatePrompt(question),
		MaxTokens:   chatBot.ChatConfig.Tokens,
		Temperature: chatBot.ChatConfig.Temperature,
		Stream:      chatBot.ChatConfig.Stream,
	}

	if !chatBot.ChatConfig.Stream {
		return chatBot.Client.CreateChatCompletion(context.Background(), chatReq)
	}

	chatBot.Stream, err = chatBot.Client.CreateChatCompletionStream(context.Background(), chatReq)
	return
}

// CreatePrompt create query prompt with context
func (chatBot *ChatBot) CreatePrompt(question string) []openai.ChatCompletionMessage {
	// Add the two points, it worked wonderfully on the responses!!!
	preQMsg := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: api_service.PreQ,
	}

	preAMsg := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: api_service.PreA,
	}

	queryMsg := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	}
	tokenCount := util.CountMessagesTokens(chatBot.ChatConfig.Model, preQMsg, preAMsg, queryMsg)

	var chatMessages []openai.ChatCompletionMessage
	chatMessages = append(chatMessages, chatBot.LimitMsg(AssistantToken-tokenCount, chatBot.GetAssistantMsg(question, TopN))...)
	chatMessages = append(chatMessages, reverse(chatBot.LimitMsg(ContextTokens, chatBot.GetSessionMsg()))...)
	chatMessages = append(chatMessages, preQMsg, preAMsg, queryMsg)
	return chatMessages
}

func (chatBot *ChatBot) SetSession(sessionId string) {
	if session, exist := SessionManager.GetSession(sessionId); exist {
		chatBot.Session = session
	} else {
		chatBot.Session = SessionManager.CreateSession()
	}
}

func (chatBot *ChatBot) GetSessionMsg() []openai.ChatCompletionMessage {
	var chatMessages []openai.ChatCompletionMessage

	for _, message := range chatBot.Session.Messages {
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Context,
		})
	}
	return chatMessages
}

// GetAssistantMsg get similar prompts for a corresponding question.
func (chatBot *ChatBot) GetAssistantMsg(question string, topN int) []openai.ChatCompletionMessage {
	texts := chatBot.RankedByRelatedness(question, topN)

	var chatMessages []openai.ChatCompletionMessage
	for _, text := range texts {
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: fmt.Sprintf("Knowledge-Library: ```%v```", text),
		})
	}

	return chatMessages
}

func (chatBot *ChatBot) LimitMsg(limitTokens int, chatMessages []openai.ChatCompletionMessage) []openai.ChatCompletionMessage {
	index := 0
	for index < len(chatMessages) && limitTokens > 0 {
		limitTokens -= util.CountMessagesTokens(chatBot.ChatConfig.Model, chatMessages[index])
		index++
		if limitTokens < 0 {
			break
		}
	}
	return chatMessages[:index]
}

// RankedByRelatedness return topN sorted []string, notice: reverse order
func (chatBot *ChatBot) RankedByRelatedness(question string, topN int) []string {
	queryEmbedding := util.GetEmbedding(chatBot.Client, util.EmbeddingModel, question)
	texts, _ := GetSimilarTextFromMilvus(db.MilvusChatgptCollection, queryEmbedding, topN)
	return texts
}

func (chatBot *ChatBot) Close() {
	if chatBot.ChatConfig.Stream && chatBot.Stream != nil {
		chatBot.Stream.Close()
	}
}

func reverse(texts []openai.ChatCompletionMessage) []openai.ChatCompletionMessage {
	left, right := 0, len(texts)-1
	for left < right {
		texts[left], texts[right] = texts[right], texts[left]
		left++
		right--
	}
	return texts
}
