package trkr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStoriesService_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/stories/21", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"id": 21}`)
	})

	s := &StoriesService{client}
	story, _, err := s.Get(21)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if story == nil {
		t.Error("Story is nil")
	}
}

func TestStoriesService_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/projects/23/stories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DE")

		fmt.Fprint(w, `[{"id": 21}]`)
	})

	s := &StoriesService{client}
	stories, _, err := s.List(23)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if stories == nil {
		t.Error("Stories is nil")
	}

	if len(stories) == 0 {
		t.Error("Stories are empty")
	}
}

func TestStoriesService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &StoryRequest{Type: String("bug")}

	mux.HandleFunc("/projects/45/stories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		defer r.Body.Close()

		output := new(StoryRequest)
		err := json.NewDecoder(r.Body).Decode(output)
		if err != nil {
			t.Errorf("Reading error %v", err)
		}

		if !reflect.DeepEqual(input, output) {
			t.Errorf("deep equal %+v, %+v", input, output)
		}

		fmt.Fprint(w, `{"id": 21}`)
	})

	s := &StoriesService{client}
	story, resp, err := s.Create(45, input)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if story == nil {
		t.Error("Story should not be empty")
	}

	if resp == nil {
		t.Errorf("Response should not be empty")
	}
}

func TestStoriesService_Put(t *testing.T) {
	setup()
	defer teardown()

	input := &StoryRequest{Type: String("bug")}

	mux.HandleFunc("/stories/21", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		defer r.Body.Close()

		output := new(StoryRequest)
		err := json.NewDecoder(r.Body).Decode(output)
		if err != nil {
			t.Errorf("Reading error %v", err)
		}

		if !reflect.DeepEqual(input, output) {
			t.Errorf("deep equal %+v, %+v", input, output)
		}

		fmt.Fprint(w, `{"id": 21}`)
	})

	s := &StoriesService{client}
	story, resp, err := s.Put(21, input)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if story == nil {
		t.Error("Story should not be empty")
	}

	if resp == nil {
		t.Errorf("Response should not be empty")
	}
}

func TestStoriesService_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/stories/21", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		fmt.Fprint(w, ``)
	})

	s := &StoriesService{client}
	resp, err := s.Delete(21)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if resp == nil {
		t.Error("Response should not be empty")
	}
}
