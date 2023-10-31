package consts

const (
	UID_KEY    = "uid"
	SECRET_KEY = "sk"
)

const (
	LOCK_SK_KEY = "api:lock:sk:%s"

	UID_USAGE_KEY      = "api:%d:usage:%s"
	USAGE_COUNT_FIELD  = "usage_count"
	USED_TOKENS_FIELD  = "used_tokens"
	TOTAL_TOKENS_FIELD = "total_tokens"
)

const (
	RootStatusDeleted = -1
	RootStatusNormal  = 0
	RootStatusDisable = 1
)
