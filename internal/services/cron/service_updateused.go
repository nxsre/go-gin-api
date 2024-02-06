package cron

import (
	"github.com/nxsre/go-gin-api/internal/pkg/core"
	"github.com/nxsre/go-gin-api/internal/repository/mysql"
	"github.com/nxsre/go-gin-api/internal/repository/mysql/cron_task"
	"github.com/robfig/cron/v3"
	"log"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := cron_task.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	// region 操作定时任务 避免主从同步延迟，在这需要查询主库
	qb = cron_task.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	info, err := qb.QueryOne(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	s.cronServer.RemoveTask(cron.EntryID(info.EntryID))

	if used != cron_task.IsUsedNo {
		entryID, err := s.cronServer.AddTask(info)
		if err != nil {
			return err
		}
		qb = cron_task.NewQueryBuilder()
		qb.WhereId(mysql.EqualPredicate, id)
		if qb.Updates(s.db.GetDbW(), map[string]interface{}{"entry_id": entryID}) != nil {
			log.Println("更新失败，删除任务")
			s.cronServer.RemoveTask(entryID)
		}
	}
	// endregion

	return
}
