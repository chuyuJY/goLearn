package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	milvus "github.com/milvus-io/milvus-sdk-go/v2/client"
)

type Attributes struct {
	Title          string `json:"title,omitempty"`
	Content        string `json:"content,omitempty"`
	Author         string `json:"author,omitempty"`
	Feature        bool   `json:"feature,omitempty"`
	SeoTitle       string `json:"seo_title,omitempty"`
	SeoDescription string `json:"seo_description,omitempty"`
}

type Data struct {
	Id         int        `json:"id,omitempty"`
	Attributes Attributes `json:"attributes"`
}

type Editorial struct {
	Data []Data `json:"data"`
}

type ChatContent struct {
	Text      string    `json:"text"`
	Embedding []float32 `json:"embedding,omitempty"`
}

var (
	getAllEditorials = "https://strapi.trlab.com/api/editorials?populate=*&sort[0]=date:desc&pagination[start]=0&pagination[limit]=100"
	collName         = "hello"

	editorial Editorial
	client    milvus.Client
)

func main() {
	var err error
	if client, err = milvus.NewDefaultGrpcClientWithURI(context.Background(), "https://in01-e11ef590049c1c3.aws-us-west-2.vectordb.zillizcloud.com:19535", "chuyu", "19990404Aa_"); err != nil {
		log.Println(err)
		return
	}
	log.Println("log in success")
	defer client.Close()

	texts := CrawRaw(getAllEditorials)
	if err = UploadToZilliz(texts); err != nil {
		log.Println(err)
	}
}

func CrawRaw(url string) []string {
	var texts []string

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(response *colly.Response) {
		if err := json.Unmarshal(response.Body, &editorial); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Visited:", len(editorial.Data))

		re := regexp.MustCompile("<[^>]*>")
		for i := 0; i < 2; i++ {
			title := re.ReplaceAllString(editorial.Data[i].Attributes.Title, "")
			author := re.ReplaceAllString(editorial.Data[i].Attributes.Author, "")
			text := re.ReplaceAllString(editorial.Data[i].Attributes.Content, "")
			texts = append(texts, fmt.Sprintf("Title: %v\nAuthor: %v\nContent: %v\n", title, author, text))
		}
	})
	c.Visit(url)

	return texts
}

func UploadToZilliz(texts []string) error {
	if exist, err := client.HasCollection(context.Background(), collName); !exist || err != nil {
		return errors.New(collName + " is not exist")
	}
	log.Println(collName + " is exist")

	chatContents := make([]ChatContent, len(texts))
	for i := 0; i < len(chatContents); i++ {
		chatContents[i] = ChatContent{
			Text: texts[i],
			//Embedding: ,
		}
	}

	return nil
}
