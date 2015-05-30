package trkr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIterationsService_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/projects/21/iterations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `[{"number": 21}]`)
	})

	is := &IterationService{client}
	iterations, _, err := is.List(21, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if len(iterations) == 0 {
		t.Error("Empty iteration List")
	}
}

func TestIterationsService_ListwithOptions(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/projects/21/iterations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if got, wants := r.FormValue("scope"), "current"; got != wants {
			t.Errorf("Scope is %v, expected %v", got, wants)
		}
		fmt.Fprint(w, `[{"number": 21}]`)
	})

	is := &IterationService{client}
	opts := &IterationListOptions{Scope: "current"}

	iterations, _, err := is.List(21, opts)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if len(iterations) == 0 {
		t.Error("Empty iteration List")
	}
}
