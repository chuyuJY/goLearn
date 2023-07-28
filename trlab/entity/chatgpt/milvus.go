package chatgpt

import "trlab-backend-go/payload"

type MilvusData struct {
	*payload.Article
	Chunks []Chunk
}

type Chunk struct {
	Embedding []float32
	Content   string
}
