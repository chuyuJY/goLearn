package main

type MilvusData struct {
	Rows []Row `json:"rows,omitempty"`
}

type Row struct {
	Slug      string `json:"slug,omitempty"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Url       string `json:"url,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Chunk
}

type Chunk struct {
	Embedding []float32 `json:"embedding,omitempty"`
	Content   string    `json:"content,omitempty"`
}
