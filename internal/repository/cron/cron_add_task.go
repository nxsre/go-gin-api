package cron

import (
	"strings"

	"github.com/nxsre/go-gin-api/internal/repository/mysql/cron_task"

	"github.com/robfig/cron/v3"
)

func (s *server) AddTask(task *cron_task.CronTask) (cron.EntryID, error) {
	// TODO: 支持秒级任务
	spec := "0 " + strings.TrimSpace(task.Spec)

	return s.cron.AddFunc(spec, s.AddJob(task))
}
