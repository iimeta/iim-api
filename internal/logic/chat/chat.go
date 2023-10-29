package chat

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/iimeta/iim-api/internal/errors"
	"github.com/iimeta/iim-api/internal/model"
	"github.com/iimeta/iim-api/internal/service"
	"github.com/iimeta/iim-api/utility/logger"
	"github.com/iimeta/iim-api/utility/util"
	"github.com/iimeta/iim-sdk/sdk"
	"github.com/sashabaranov/go-openai"
	"reflect"
)

type sChat struct{}

func init() {
	service.RegisterChat(New())
}

func New() service.IChat {
	return &sChat{}
}

func (s *sChat) Completions(ctx context.Context, params model.CompletionsReq) (openai.ChatCompletionResponse, error) {

	chat := sdk.NewChat()
	chat.Corp = sdk.CORP_OPENAI
	chat.Model = params.Model
	chat.Messages = params.Messages

	response, err := sdk.Chat.Chat(ctx, chat)
	if err != nil {
		e := &openai.APIError{}
		if errors.As(err, &e) && !reflect.DeepEqual(response, openai.ChatCompletionResponse{}) {
			return response, nil
		}
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}

func (s *sChat) CompletionsStream(ctx context.Context, params model.CompletionsReq) (err error) {

	response := make(chan openai.ChatCompletionStreamResponse)
	defer close(response)

	if err = grpool.AddWithRecover(ctx, func(ctx context.Context) {
		chat := sdk.NewChat()
		chat.Corp = sdk.CORP_OPENAI
		chat.Model = params.Model
		chat.Messages = params.Messages
		chat.Stream = true
		err = sdk.Chat.ChatStream(ctx, chat, response)
		if err != nil {
			logger.Error(ctx, err)
			return
		}
	}, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	for {
		select {
		case response := <-response:

			if response.Choices[0].FinishReason == "stop" {

				if response.Choices[0].Delta.Content != "" {
					err = util.SSEServer(ctx, "", gjson.MustEncode(response))
					if err != nil {
						logger.Error(ctx, err)
						return err
					}
				}

				err = util.SSEServer(ctx, "", "[DONE]")
				if err != nil {
					logger.Error(ctx, err)
					return err
				}

				return nil
			}

			err = util.SSEServer(ctx, "", gjson.MustEncode(response))
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		default:
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}
}
