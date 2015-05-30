package trkr

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// Iteration represents a collection of stories completed at some specific point
type Iteration struct {
	Number       int        `json:"number,omitempty"`
	ProjectID    int64      `json:"project_id,omitempty"`
	Length       int        `json:"length,omitempty"`
	TeamStrength float64    `json:"team_strength,omitempty"`
	Start        *time.Time `json:"start,omitempty"`
	Finish       *time.Time `json:"finish,omitempty"`
	StoryIDs     []int64    `json:"story_ids,omitempty"`
	Stories      []*Story   `json:"stories,omitempty"`
}

// IterationsService interacts with Tracker iteration API
type IterationsService struct {
	client *Client
}

// IterationsListOptions allow the Request URL to be customized
type IterationsListOptions struct {
	Offset int    `url:"offset,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Label  string `url:"label,omitempty"`
	Scope  string `url:"scope,omitempty"`
}

// List returns iterations & stories that match a project ID
func (is *IterationsService) List(projectID int64, opts *IterationsListOptions) ([]Iteration, *http.Response, error) {
	u, _ := url.Parse(fmt.Sprintf("projects/%d/iterations", projectID))
	if opts != nil {
		v, err := query.Values(opts)
		if err != nil {
			return nil, nil, err
		}
		u.RawQuery = v.Encode()
	}

	req, err := is.client.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	iterations := make([]Iteration, 0, 0)

	resp, err := is.client.Do(req, &iterations)
	if err != nil {
		return nil, resp, err
	}
	return iterations, resp, nil
}
