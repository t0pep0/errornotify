package errornotify

import (
	redmine "github.com/mattn/go-redmine"
)

type RedmineOptions struct {
	Use      bool
	Url      string
	ApiKey   string
	Project  *redmine.IdName
	Tracker  *redmine.IdName
	Assigned *redmine.IdName
	StatusId int
	Status   *redmine.IdName
	Author   *redmine.IdName
	Priority *redmine.IdName
}
