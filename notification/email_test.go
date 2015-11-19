package notification

import (
	"testing"

	"github.com/francoishill/leeroyci/database"
)

func TestEmailSubject(t *testing.T) {
	repo, _ := database.CreateRepository("repo", "bar", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "1234", "commitURL", "foo", "bar")
	job.TasksDone()

	build := emailSubject(job, EventBuild)
	test := emailSubject(job, EventTest)
	deployStart := emailSubject(job, EventDeployStart)
	deployEnd := emailSubject(job, EventDeployEnd)

	if build != "repo/branch build" {
		t.Error("Wrong message", build)
	}

	if test != "repo/branch tests" {
		t.Error("Wrong message", test)
	}

	if deployStart != "repo/branch deployment started" {
		t.Error("Wrong message", deployStart)
	}

	if deployEnd != "repo/branch deploy success" {
		t.Error("Wrong message", deployEnd)
	}
}
