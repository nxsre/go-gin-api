package admin

import (
	"github.com/nxsre/go-gin-api/configs"
	"github.com/nxsre/go-gin-api/internal/pkg/core"
	"github.com/nxsre/go-gin-api/internal/pkg/password"
	"github.com/nxsre/go-gin-api/internal/repository/mysql"
	"github.com/nxsre/go-gin-api/internal/repository/mysql/admin"
	"github.com/nxsre/go-gin-api/internal/repository/redis"
)

func (s *service) ResetPassword(ctx core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"password":     password.ResetPassword(),
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
