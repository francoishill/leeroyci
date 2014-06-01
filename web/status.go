// Expose the build status over http.
package web

import (
	"encoding/hex"
	"html/template"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
	"strings"
)

// View that shows all builds ever done.
func Status(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	blog.Sort()
	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": blog.Jobs,
		},
	)
}

// View to show builds for a specific repository.
func Repo(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)

	j := blog.JobsForRepo(r)

	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": j,
		},
	)
}

// View to show builds for a specific repository and branch.
func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	b := splitSecond(req.URL.Path)

	j := blog.JobsForRepoBranch(r, b)

	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": j,
		},
	)
}

// View to show the build for a commit in a repository.
func Commit(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	co := splitSecond(req.URL.Path)

	j := blog.JobByCommit(r, co)

	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": []logging.Job{j},
		},
	)
}

// Splits a request path and returns the first part after the endpoint.
// This is usually the hex string of the repository.
func splitFirst(path string) string {
	s := strings.Split(path, "/")[3]
	r, err := hex.DecodeString(s)

	if err != nil {
		log.Println(err)
	}

	return string(r)
}

// Splits a request path and returns the first part after the endpoint.
// This is likely the branch name or commit sha1.
func splitSecond(path string) string {
	return strings.Split(path, "/")[4]
}
