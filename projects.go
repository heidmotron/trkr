package trkr

import (
	"fmt"
	"net/http"
	"time"
)

// ProjectsService handles the communication for fetching projects
type ProjectsService struct {
	client *Client
}

// Project represents a Tracker project
type Project struct {
	ID                     *int64     `json:"id,omitempty"`
	Version                *int64     `json:"version,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	Description            *string    `json:"description,omitempty"`
	WeekStartDay           *string    `json:"week_start_day,omitempty"`
	StartDate              *time.Time `json:"start_date,omitempty"`
	IterationLength        *int64     `json:"iteration_length,omitempty"`
	AutomaticPlanning      *bool      `json:"automatic_planning,omitempty"`
	EstimateBugsChores     *bool      `json:"bugs_and_chores_are_estimatable,omitempty"`
	PointScale             *string    `json:"point_scale,omitempty"`
	PointScaleIsCustom     *bool      `json:"point_scale_is_custom,omitempty"`
	TimeZone               *TimeZone  `json:"time_zone,omitempty"`
	EnableTasks            *bool      `json:"enable_tasks,omitempty"`
	VelocityAveragedOver   *int       `json:"velocity_averaged_over,omitempty"`
	NumberOfDoneIterations *int       `json:"number_of_done_iterations_to_show,omitempty"`
	HasGoogleDomain        *bool      `json:"has_google_domain,omitempty"`
	ProfileContent         *string    `json:"profile_content,omitempty"`
	EnableIncomingEmails   *bool      `json:"enable_incoming_emails,omitempty"`
	InitialVelocity        *int       `json:"initial_velocity,omitempty"`
	ProjectType            *string    `json:"project_type,omitempty"`
	Public                 *bool      `json:"public,omitempty"`
	AtomEnabled            *bool      `json:"atom_enabled,omitempty"`
	CurrentIterationNumber *int       `json:"current_iteration_number,omitempty"`
	CurrentVelocity        *int       `json:"current_velocity,omitempty"`
	CurrentVolatility      *float64   `json:"current_volatility,omitempty"`
	AccountID              *int64     `json:"account_id,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

// List returns all projects associated with the user
func (p *ProjectsService) List() ([]Project, *http.Response, error) {
	req, err := p.client.NewRequest("GET", "projects", nil)
	if err != nil {
		return nil, nil, err
	}
	projs := new([]Project)

	resp, err := p.client.Do(req, projs)
	if err != nil {
		return nil, resp, err
	}

	return *projs, resp, nil
}

// Get returns a single project by ID
func (p *ProjectsService) Get(id int64) (*Project, *http.Response, error) {
	req, err := p.client.NewRequest("GET", fmt.Sprintf("projects/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	proj := new(Project)
	resp, err := p.client.Do(req, proj)
	if err != nil {
		return nil, resp, err
	}

	return proj, resp, nil
}
