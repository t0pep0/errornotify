package errornotify

import (
	redmine "github.com/mattn/go-redmine"
	"os"
	"strconv"
	"time"
)

var Redmine RedmineOptions

func init() {
	Redmine.Project = new(redmine.IdName)
	Redmine.Tracker = new(redmine.IdName)
	Redmine.Assigned = new(redmine.IdName)
	Redmine.Status = new(redmine.IdName)
	Redmine.Author = new(redmine.IdName)
	Redmine.Priority = new(redmine.IdName)
}

type Nerror struct {
	timestamp time.Time
	pid       int
	gid       int
	uid       int
	message   string
	level     string
	env       []string
	cur_dir   string
	hostname  string
}

func (e *Nerror) Set(lvl string, msg string) {
	e.timestamp = time.Now()
	e.env = os.Environ()
	e.gid = os.Getgid()
	e.uid = os.Getuid()
	e.pid = os.Getpid()
	e.cur_dir, _ = os.Getwd()
	e.hostname, _ = os.Hostname()
	e.level = lvl
	e.message = msg
	if Redmine.Use {
		RedmineClient := redmine.NewClient(Redmine.Url, Redmine.ApiKey)
		RedmineClient.CreateIssue(redmine.Issue{
			Subject:     "Error on " + Redmine.Project.Name + " " + e.timestamp.Format(time.RFC1123),
			Description: e.Error(),
			ProjectId:   Redmine.Project.Id,
			Project:     Redmine.Project,
			Tracker:     Redmine.Tracker,
			Assigned:    Redmine.Assigned,
		})
	}
}

func (e *Nerror) Error() (msg string) {
	msg = e.timestamp.Format(time.RFC1123) + ": Level: " + e.level + " " + e.message + "\n\r"
	msg += "Environment:\n\r"
	for _, env := range e.env {
		msg += "      " + env + "\n\r"
	}
	msg += "Hostname: " + e.hostname + "\n\r"
	msg += "Work dir: " + e.cur_dir + "\n\r"
	msg += "uid: " + strconv.Itoa(e.uid) + " gid: " + strconv.Itoa(e.gid) + " pid: " + strconv.Itoa(e.pid) + "\n\r"
	return msg
}
