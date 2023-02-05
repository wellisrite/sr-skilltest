package constant

import "time"

const (
	MODE_PRODUCTION = "production"
)

const (
	YYYY_MM_DD = "2006-01-02"
)
const (
	CONTEXT_LOCALS_KEY_TRACE_ID = "TraceID"
)

const (
	ERR_EXPIRED_PRODUCT = "customer trying to buy expired product"
)

const (
	PAGINATION_CACHE_EXP_TIME = time.Second * 5
	ENTITY_CACHE_EXP_TIME     = time.Hour * 6
	ENTITY_CACHE_EXP_TIME_1   = time.Minute * 1
)
