// Package integrations handles receiving notifications from repository sources
// like GitHub.
package integrations

import (
	"leeroy/integrations/github"
	"log"
	"net/http"
	"strings"
)

// Callback handles callbacks and webhooks sent by code hosting serivces.
func Callback(rw http.ResponseWriter, req *http.Request) {
	s := getService(req)

	switch s {
	case "github":
		github.Parse(req)
	default:
		log.Println("serivce", s, "not supported")
	}
}

// Returns the name of the service and the secret key.
func getService(req *http.Request) string {
	return strings.Split(req.URL.Path, "/")[2]
}
