package interceptor

import (
	"encoding/json"
	"net/http"

	"github.com/nxsre/go-gin-api/configs"
	"github.com/nxsre/go-gin-api/internal/code"
	"github.com/nxsre/go-gin-api/internal/pkg/core"
	"github.com/nxsre/go-gin-api/internal/proposal"
	"github.com/nxsre/go-gin-api/internal/repository/redis"
	"github.com/nxsre/go-gin-api/pkg/errors"
)

func (i *interceptor) CheckLoginCookie(ctx core.Context) (sessionUserInfo proposal.SessionUserInfo, err core.BusinessError) {
	var token = ""
	cookies := ctx.Request().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "_login_token_" {
			token = cookie.Value
			break
		}
	}

	if token == "" {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Token 参数"))

		return
	}

	if !i.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("请先登录"))

		return
	}

	cacheData, cacheErr := i.cache.Get(configs.RedisKeyPrefixLoginUser+token, redis.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(cacheErr)

		return
	}

	jsonErr := json.Unmarshal([]byte(cacheData), &sessionUserInfo)
	if jsonErr != nil {
		core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(jsonErr)

		return
	}

	return
}
