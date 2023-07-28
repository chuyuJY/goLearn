package util

import (
	"fmt"
	"log"
	"strings"
	"trlab-backend-go/logger"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

// (query text) token limit: 4,096 for gpt-3.5-turbo or 8,192 for gpt-4
// (embedding text) token limit: 8191 for text-embedding-ada-002 and cl100k_base

var encoders = map[string]*tiktoken.Tiktoken{}

// CountTokens calculate text's token
func CountTokens(text string, model string) int {
	return len(GetToken(text, model))
}

// CountMessagesTokens based on "How_to_count_tokens_with_tiktoken.ipynb"
func CountMessagesTokens(model string, messages ...openai.ChatCompletionMessage) (numTokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("encoding for model: %v", err)
		log.Println(err)
		return
	}

	var tokensPerMessage, tokensPerName int
	switch model {
	case "gpt-3.5-turbo-0613",
		"gpt-3.5-turbo-16k-0613",
		"gpt-4-0314",
		"gpt-4-32k-0314",
		"gpt-4-0613",
		"gpt-4-32k-0613":
		tokensPerMessage = 3
		tokensPerName = 1
	case "gpt-3.5-turbo-0301":
		tokensPerMessage = 4 // every message follows <|start|>{role/name}\n{content}<|end|>\n
		tokensPerName = -1   // if there's a name, the role is omitted
	default:
		if strings.Contains(model, "gpt-3.5-turbo") {
			//logger.ZapLogger.Warn("warning: gpt-3.5-turbo may update over time. Returning num tokens assuming gpt-3.5-turbo-0613.")
			return CountMessagesTokens("gpt-3.5-turbo-0613", messages...)
		} else if strings.Contains(model, "gpt-4") {
			//logger.ZapLogger.Warn("warning: gpt-4 may update over time. Returning num tokens assuming gpt-4-0613.")
			return CountMessagesTokens("gpt-4-0613", messages...)
		} else {
			err = fmt.Errorf("num_tokens_from_messages() is not implemented for model %s. See https://github.com/openai/openai-python/blob/main/chatml.md for information on how messages are converted to tokens.", model)
			logger.ZapLogger.Error(err.Error())
			return
		}
	}

	for _, message := range messages {
		numTokens += tokensPerMessage
		numTokens += len(tkm.Encode(message.Content, nil, nil))
		numTokens += len(tkm.Encode(message.Role, nil, nil))
		numTokens += len(tkm.Encode(message.Name, nil, nil))
		if message.Name != "" {
			numTokens += tokensPerName
		}
	}
	numTokens += 3 // every reply is primed with <|start|>assistant<|message|>
	return numTokens
}

func ChunkedText(text string, model string, chunkLength int) []string {
	var texts []string
	for _, token := range ChunkedTokens(text, model, chunkLength) {
		texts = append(texts, ParseToken(token, model))
	}

	return texts
}

// ChunkedTokens encodes a string into tokens and then breaks it up into chunks
func ChunkedTokens(text string, model string, chunkLength int) [][]int {
	if chunkLength < 1 {
		logger.ZapLogger.Error("chunkLength must be at least one")
		return nil
	}

	return Batched(GetToken(text, model), chunkLength)
}

// Batched batch data into slices of length n. The last batch may be shorter.
func Batched(data []int, n int) [][]int {
	var batches [][]int

	for len(data) > 0 {
		if len(data) < n {
			n = len(data)
		}

		batch := make([]int, n)
		copy(batch, data[:n])
		batches = append(batches, batch)
		data = data[n:]
	}
	return batches
}

func GetToken(text string, model string) []int {
	enc := RegisterEncoder(model)
	return enc.Encode(text, nil, nil)
}

func ParseToken(token []int, model string) string {
	enc := RegisterEncoder(model)
	return enc.Decode(token)
}

func RegisterEncoder(model string) *tiktoken.Tiktoken {
	enc, ok := encoders[model]
	if !ok {
		enc, _ = tiktoken.EncodingForModel(model)
		encoders[model] = enc
	}
	return enc
}
