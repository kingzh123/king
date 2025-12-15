package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3" // imported as openai
	"github.com/openai/openai-go/v3/option"
)

func main() {
	apiKey := "sk-b45d6d7110b14f2c8889de29212e79fa"
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
	)
	completion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("中国 现在几点"),
			},
			Model: "qwen-plus",
		})
	if err != nil {
		panic(err)
	}
	fmt.Println(completion.Choices[0].Message.Content)
}
