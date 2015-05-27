package trkr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProjectsService_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/projects/21", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"id": 21}`)
	})

	s := &ProjectsService{client}
	project, _, err := s.Get(21)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if project == nil {
		t.Error("Project is nil")
	}
}

func TestProjectsService_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `[{"id": 21}]`)
	})

	s := &ProjectsService{client}
	projects, _, err := s.List()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if len(projects) == 0 {
		t.Error("Empty Project List")
	}
}
