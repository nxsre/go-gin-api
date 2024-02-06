package cron

import "github.com/robfig/cron/v3"

func (s *server) RemoveTask(taskId cron.EntryID) {
	s.cron.Remove(taskId)
}
