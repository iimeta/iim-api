package chat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/iim-api/internal/service"

	"github.com/iimeta/iim-api/api/chat/v1"
)

func (c *ControllerV1) Completions(ctx context.Context, req *v1.CompletionsReq) (res *v1.CompletionsRes, err error) {

	if req.Stream {
		err = service.Chat().CompletionsStream(ctx, req.CompletionsReq)
		if err != nil {
			return nil, err
		}
	} else {
		response, err := service.Chat().Completions(ctx, req.CompletionsReq)
		if err != nil {
			return nil, err
		}
		g.RequestFromCtx(ctx).Response.WriteJson(response)
	}

	return
}
