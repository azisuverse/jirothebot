package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
  Client *openai.Client
}

func NewOpenAIService(apiKey string) *OpenAIService {
  return &OpenAIService{
    Client: openai.NewClient(apiKey),
  }
}

func (oai *OpenAIService) GetResponse(prompt string) (string, error ){
  res, err := oai.Client.CreateChatCompletion(context.Background(),
    openai.ChatCompletionRequest{
      Model: openai.GPT4oMini,
      Store: true,
      Messages: []openai.ChatCompletionMessage{
        {
          Role: openai.ChatMessageRoleSystem,
          Content: "You are Jiro the bot, and you builded and developed by @showmegoods. Answer all the question using informal lang or slang if you can.",
        },
        {
          Role: openai.ChatMessageRoleUser,
          Content: prompt,
        },
      },
    })
  if err != nil {
    return "", err
  }
  return res.Choices[0].Message.Content, nil
}
