// Config takes care of the whole configuration.
package config

import (
	"errors"
)

type Config struct {
	Secret       string
	Repositories []Repository
}

type Repository struct {
	URL      string
	Commands []Command
	Notify   []Notify
}

type Command struct {
	Name    string
	Execute string
}

type Notify struct {
	Name  string
	Email string
}

// ConfigForRepo returns the configuration for a repository that matches
// the URL.
func (c *Config) ConfigForRepo(url string) (r Repository, err error) {
	r = Repository{}

	for _, repo := range c.Repositories {
		if repo.URL == url {
			r = repo
			return
		}
	}

	msg := "Could not find repository with URL: " + url
	err = errors.New(msg)
	return
}
