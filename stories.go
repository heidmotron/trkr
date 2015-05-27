package trkr

import (
	"fmt"
	"net/http"
)

// StoriesService retreives stories from the api
type StoriesService struct {
	client *Client
}

// Story model
type Story struct {
	ID          *int64    `json:"id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Estimate    *int      `json:"estimate,omitempty"`
	Type        *string   `json:"story_type,omitempty"`
	State       *string   `json:"current_state,omitempty"`
	Description *string   `json:"description,omitempty"`
	Labels      []Label   `json:"labels,omitempty"`
	Comments    []Comment `json:"comments,omitempty"`
}

// HasEstimate returns true if an estimate is set
func (s *Story) HasEstimate() bool {
	return s.Estimate != nil
}

// Label model
type Label struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Comment model
type Comment struct {
	ID          *int64       `json:"id,omitempty"`
	Text        *string      `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment model
type Attachment struct {
	ID            *int64  `json:"id,omitempty"`
	FileName      *string `json:"filename,omitempty"`
	Height        *int64  `json:"height,omitempty"`
	Width         *int64  `json:"width,omitempty"`
	DownloadURL   *string `json:"download_url,omitempty"`
	BigURL        *string `json:"big_url,omitempty"`
	ThumbnailURL  *string `json:"thumbnail_url,omitempty"`
	ContentType   *string `json:"content_type,omitempty"`
	UploaderID    *int64  `json:"uploader_id,omitempty"`
	Thumbnailable *bool   `json:"thumbnailable,omitempty"`
}

// StoryRequest is used when constructing or updating a story
type StoryRequest struct {
	Name        *string `json:"name,omitempty"`
	Estimate    *int    `json:"estimate,omitempty"`
	Type        *string `json:"story_type,omitempty"`
	State       *string `json:"current_state,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Get retreives an individual story
func (s *StoriesService) Get(id int64) (*Story, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("stories/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	story := new(Story)
	resp, err := s.client.Do(req, story)
	if err != nil {
		return nil, resp, err
	}

	return story, resp, nil
}

// Put updates an individual story
func (s *StoriesService) Put(id int64, r *StoryRequest) (*Story, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("stories/%d", id), r)
	if err != nil {
		return nil, nil, err
	}

	story := new(Story)

	resp, err := s.client.Do(req, story)
	if err != nil {
		return nil, resp, err
	}

	return story, resp, nil
}

// Create adds new stroies to a given project
func (s *StoriesService) Create(projectID int64, r *StoryRequest) (*Story, *http.Response, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("projects/%d/stories", projectID), r)
	if err != nil {
		return nil, nil, err
	}

	story := new(Story)

	resp, err := s.client.Do(req, story)
	if err != nil {
		return nil, resp, err
	}

	return story, resp, nil
}

// List returns a stories for a given project
func (s *StoriesService) List(projectID int64) ([]Story, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("projects/%d/stories", projectID), nil)
	if err != nil {
		return nil, nil, err
	}

	stories := make([]Story, 0, 10)

	resp, err := s.client.Do(req, stories)
	if err != nil {
		return nil, resp, err
	}

	return stories, resp, nil
}
