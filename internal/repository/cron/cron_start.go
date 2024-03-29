package cron

import (
	"fmt"
	"math"

	"github.com/nxsre/go-gin-api/internal/repository/mysql"
	"github.com/nxsre/go-gin-api/internal/repository/mysql/cron_task"

	"go.uber.org/zap"
)

func (s *server) Start() {
	s.cron.Start()
	go s.taskCount.Wait()

	qb := cron_task.NewQueryBuilder()
	qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)

	totalNum, err := qb.Count(s.db.GetDbR())
	if err != nil {
		s.logger.Fatal("cron initialize tasks count err", zap.Error(err))
	}

	pageSize := 50
	maxPage := int(math.Ceil(float64(totalNum) / float64(pageSize)))

	taskNum := 0
	s.logger.Info("开始初始化后台任务")

	for page := 1; page <= maxPage; page++ {
		qb = cron_task.NewQueryBuilder()
		qb.WhereIsUsed(mysql.EqualPredicate, cron_task.IsUsedYES)
		listData, err := qb.
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			OrderById(false).
			QueryAll(s.db.GetDbR())
		if err != nil {
			s.logger.Fatal("cron initialize tasks list err", zap.Error(err))
		}

		s.logger.Info("listData", zap.String("listData", fmt.Sprint(listData)))
		for _, item := range listData {
			entryID, err := s.AddTask(item)
			if err != nil {
				continue
			}
			taskNum++
			qb = cron_task.NewQueryBuilder()
			qb.WhereId(mysql.EqualPredicate, item.Id)
			if err := qb.Updates(s.db.GetDbW(), map[string]interface{}{"entry_id": entryID}); err != nil {
				s.RemoveTask(entryID)
			}
		}
	}

	s.logger.Info(fmt.Sprintf("后台任务初始化完成，总数量：%d", taskNum))
}
