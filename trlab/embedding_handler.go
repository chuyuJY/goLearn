package util

import (
	"context"
	"trlab-backend-go/entity/chatgpt"
	"trlab-backend-go/logger"

	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"gonum.org/v1/gonum/mat"
)

const (
	EmbeddingModel      = openai.AdaEmbeddingV2
	EmbeddingMaxToken   = 8191
	EmbeddingDimension  = 1536
	EmbeddingLimitToken = 1024
)

// SafeGetChunks safely handles embedding requests, even when the input text is longer than the maximum context length.
func SafeGetChunks(textBot *openai.Client, model openai.EmbeddingModel, text string, average bool, maxToken int) []*chatgpt.Chunk {
	var chunks []*chatgpt.Chunk
	var chunkLens []int
	var totalLen int

	for _, chunk := range ChunkedTokens(text, model.String(), maxToken) {
		t := ParseToken(chunk, model.String())
		chunks = append(chunks, &chatgpt.Chunk{
			Embedding: GetEmbedding(textBot, model, t),
			Content:   t,
		})
		chunkLens = append(chunkLens, len(chunk))
		totalLen += len(chunk)
	}

	// return the weighted average of the chunk embeddings
	if average {
		sum := make([]float64, len(chunks[0].Embedding))
		for i, chunk := range chunks {
			for j, val := range chunk.Embedding {
				sum[j] += float64(val) * float64(chunkLens[i])
			}
		}
		for i := 0; i < len(sum); i++ {
			sum[i] /= float64(totalLen)
		}
		return []*chatgpt.Chunk{{Embedding: NormalizeVector(sum), Content: text}}
	}

	return chunks
}

// GetEmbedding prepare embedding for content, return []float32 with 1536 dimensions
func GetEmbedding(textBot *openai.Client, model openai.EmbeddingModel, input any) []float32 {
	embeddingResp, err := textBot.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: input, // tokens or string
			Model: model,
		},
	)

	if err != nil {
		logger.ZapLogger.Error("get embedding failed", zap.Error(err))
		return nil
	}

	return embeddingResp.Data[0].Embedding
}

// NormalizeVector normalizes length to 1
func NormalizeVector(vec []float64) []float32 {
	embedding := make([]float32, len(vec))
	norm := mat.Norm(mat.NewVecDense(len(vec), vec), 2)
	for i, val := range vec {
		embedding[i] = float32(val / norm)
	}

	return embedding
}
