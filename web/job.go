package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/francoishill/leeroyci/database"
	"github.com/francoishill/leeroyci/runner"
)

const limit = 20

// viewListJobs shows a paginated list of all jobs.
func viewListJobs(w http.ResponseWriter, r *http.Request) {
	template := "job/list.html"
	ctx := make(responseContext)

	offset := 0
	paramOffset := r.URL.Query().Get("offset")

	if len(paramOffset) > 0 {
		offset = stringToInt(paramOffset)
	}

	ctx["jobs"] = database.GetJobs(offset, limit)

	prev, next, first := previousNextNumber(offset)

	ctx["previous_offset"] = prev
	ctx["next_offset"] = next
	ctx["first_page"] = first

	render(w, r, template, ctx)
}

// viewJobDetail shows a specific job with all related information.
func viewDetailJob(w http.ResponseWriter, r *http.Request) {
	template := "job/detail.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	jobID, _ := strconv.Atoi(vars["jid"])

	job := database.GetJob(int64(jobID))
	ctx["job"] = job

	render(w, r, template, ctx)
}

// viewCancelJob cancels a job.
func viewCancelJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID, _ := strconv.Atoi(vars["jid"])

	job := database.GetJob(int64(jobID))
	job.Cancel()

	http.Redirect(w, r, "/", 302)
}

// viewRerunJob resets a job status and enqueues it agian.
func viewRerunJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID, _ := strconv.Atoi(vars["jid"])

	old := database.GetJob(int64(jobID))
	job := database.CreateJob(
		&old.Repository,
		old.Branch,
		old.Commit,
		old.CommitURL,
		old.Name,
		old.Email,
	)

	queueJob := runner.QueueJob{
		JobID: job.ID,
	}

	queueJob.Enqueue()

	http.Redirect(w, r, "/", 302)
}

// viewSearchJobs returns a list of jobs filtered by the search string.
func viewSearchJobs(w http.ResponseWriter, r *http.Request) {
	template := "job/list.html"
	ctx := make(responseContext)

	query := r.URL.Query().Get("query")

	ctx["jobs"] = database.SearchJobs(query)
	ctx["query"] = query

	render(w, r, template, ctx)
}

// returns the offset for the previous and next page.
func previousNextNumber(offset int) (int, int, bool) {
	count := database.NumberOfJobs()
	prev := offset - limit
	next := offset + limit

	if prev < 0 {
		prev = 0
	}

	if next >= count {
		next = -1
	}

	first := false

	if count > limit && next != limit {
		first = true
	}

	return prev, next, first
}
