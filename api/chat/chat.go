// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chat

import (
	"context"
	
	"github.com/iimeta/iim-api/api/chat/v1"
)

type IChatV1 interface {
	Completions(ctx context.Context, req *v1.CompletionsReq) (res *v1.CompletionsRes, err error)
}


