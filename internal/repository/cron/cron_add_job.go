package cron

import (
	"fmt"
	"go.uber.org/zap"
	"os/exec"

	"github.com/mattn/go-shellwords"
	"github.com/nxsre/go-gin-api/internal/repository/mysql/cron_task"
	"github.com/robfig/cron/v3"
)

// AddJob 执行任务内容
func (s *server) AddJob(task *cron_task.CronTask) cron.FuncJob {
	return func() {
		s.taskCount.Add()
		defer s.taskCount.Done()

		// TODO: 执行任务内容
		// 将 task 信息写入到 Kafka Topic 中，任务执行器订阅 Topic 如果为符合条件的任务并进行执行，反之不执行
		// 为了便于演示，不写入到 Kafka 中，仅记录日志
		switch task.Protocol {
		// shell
		case cron_task.ProtocolShell:
			strs, err := shellwords.Parse(task.Command)
			if err != nil {
				s.logger.Error(err.Error())
			}
			cmd := exec.Command(strs[0], strs[1:]...)
			bs, err := cmd.CombinedOutput()
			if err != nil {
				s.logger.Error(err.Error())
			}
			s.logger.Info(string(bs))
		// http
		case cron_task.ProtocolHTTP:
			switch task.HttpMethod {
			// get
			case cron_task.HttpMethodGet:

			// post
			case cron_task.HttpMethodPost:
				s.logger.Info(task.ReqURL, zap.String("body", task.ReqBody))
			}
		}
		msg := fmt.Sprintf("执行任务：(%d)%s [%s]", task.Id, task.Name, task.Spec)
		s.logger.Info(msg)
	}
}
