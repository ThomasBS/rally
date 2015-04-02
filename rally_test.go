package rally

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const token = "token"

func TestGet(t *testing.T) {
	expected := "/hierarchicalrequirement/id"
	response := "testdata/hierarchical-requirement.json"

	s := setUpServer(t, expected, response)

	defer s.Close()

	rally := New(token)
	rally.url = s.URL + "/"

	var hr HierarchicalRequirement
	rally.Read(&hr, "id")

	assert.Equal(t, "User Story Title", hr.Name)
	assert.Equal(t, "2015-01-01T12:00:00.000Z", hr.CreationDate)
	assert.Equal(t, 5, hr.TasksQueryReference.Count)
}

func TestFetch(t *testing.T) {
	expected := "/hierarchicalrequirement/id/workspace"
	response := "testdata/workspace.json"

	s := setUpServer(t, expected, response)
	defer s.Close()

	rally := New(token)

	var hr HierarchicalRequirement
	hr.WorkspaceReference.ReferenceUrl = s.URL + expected

	var w Workspace
	rally.Fetch(&w, &hr.WorkspaceReference)

	assert.Equal(t, "Workspace Title", w.Name)
}

func TestQueryFetch(t *testing.T) {
	expected := "/hierarchicalrequirement/id/tasks"
	response := "testdata/task-query.json"

	s := setUpServer(t, expected, response)
	defer s.Close()

	rally := New(token)

	var hr HierarchicalRequirement
	hr.TasksQueryReference.ReferenceUrl = s.URL + expected

	var tq TaskQuery
	rally.QueryFetch(&tq, &hr.TasksQueryReference)

	assert.Equal(t, 5, tq.TotalResultCount)
	assert.Equal(t, "Task 1", tq.Results[0].Name)
	assert.Equal(t, "Task 5", tq.Results[4].Name)
}

func TestFailCreatingNewRequest(t *testing.T) {

	rally := New(token)
	rally.url = ":"

	var a Artifact
	err := rally.Read(&a, "")
	assert.NotEmpty(t, err)
}

func TestFailClientDoRequest(t *testing.T) {

	rally := New(token)
	rally.url = ""

	var a Artifact
	var r reference

	err := rally.Fetch(&a, &r)
	assert.NotEmpty(t, err)
}

func TestFailStatusCodeNotOK(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	s := httptest.NewServer(handler)
	defer s.Close()

	rally := New(token)

	var a Artifact
	var q queryReference
	q.ReferenceUrl = s.URL

	err := rally.QueryFetch(&a, &q)
	assert.NotEmpty(t, err)
}

func TestFailJsonUnmarshalResponse(t *testing.T) {

	s := setUpServer(t, "/artifact/id", "")
	defer s.Close()

	rally := New(token)
	rally.url = s.URL + "/"

	var a Artifact

	err := rally.Read(&a, "id")
	assert.NotEmpty(t, err)
}

func TestFailJsonUnmarshalOntoStruct(t *testing.T) {

	s := setUpServer(t, "/artifact/id", "testdata/fail-artifact.json")
	defer s.Close()

	rally := New(token)
	rally.url = s.URL + "/"

	var a Artifact

	err := rally.Read(&a, "id")
	assert.NotEmpty(t, err)
}

func setUpServer(t *testing.T, uri string, file string) *httptest.Server {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, uri, r.URL.Path)
		assert.Equal(t, token, r.Header["Zsessionid"][0])

		data, _ := ioutil.ReadFile(file)
		buffer := new(bytes.Buffer)

		json.Compact(buffer, data)
		fmt.Fprintln(w, buffer)
	})

	return httptest.NewServer(handler)
}
