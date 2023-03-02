package middleware

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-chat/pkg/errorm"
	"go-zero-chat/pkg/result"
	"net/http"
)

const (
	burst = 200
	rate  = 40
)

type LimitMiddleware struct {
	store *redis.Redis
}

func NewLimitMiddleware(store *redis.Redis) *LimitMiddleware {
	return &LimitMiddleware{
		store: store,
	}
}

// Handle go-zero提供了两种限流方式 令牌桶限流&&固定时间窗口限流 这里使用令牌桶限流
func (m *LimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path
		l := limit.NewTokenLimiter(rate, burst, m.store, key)

		if !l.Allow() {
			result.HttpResult(r, w, nil, errorm.NewError(errorm.ErrRequestLimit, "", nil))
			return
		}

		next(w, r)
	}
}
