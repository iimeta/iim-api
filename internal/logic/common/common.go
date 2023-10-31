package common

import (
	"context"
	"fmt"
	"github.com/iimeta/iim-api/internal/consts"
	"github.com/iimeta/iim-api/internal/service"
	"github.com/iimeta/iim-api/utility/logger"
	"github.com/iimeta/iim-api/utility/redis"
	"time"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() service.ICommon {
	return &sCommon{}
}

func (s *sCommon) GetUidUsageKey(ctx context.Context) string {
	return fmt.Sprintf(consts.UID_USAGE_KEY, service.Auth().GetUid(ctx), time.Now().Format("20060102"))
}

func (s *sCommon) RecordUsage(ctx context.Context, totalTokens int) error {

	if _, err := redis.HIncrBy(ctx, s.GetUidUsageKey(ctx), consts.USAGE_COUNT_FIELD, 1); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err := redis.HIncrBy(ctx, s.GetUidUsageKey(ctx), consts.USED_TOKENS_FIELD, int64(totalTokens)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *sCommon) GetUsedTokens(ctx context.Context) (int, error) {
	return redis.HGetInt(ctx, s.GetUidUsageKey(ctx), consts.USAGE_COUNT_FIELD)
}

func (s *sCommon) GetTotalTokens(ctx context.Context) (int, error) {
	return redis.HGetInt(ctx, s.GetUidUsageKey(ctx), consts.TOTAL_TOKENS_FIELD)
}
