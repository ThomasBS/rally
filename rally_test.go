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

var key = "key"

func TestGetHierarchicalRequirement(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/hierarchicalrequirement/id", r.URL.Path)
		assert.Equal(t, key, r.Header["Zsessionid"][0])

		d, _ := ioutil.ReadFile("testdata/hierarchical-requirement.json")
		b := new(bytes.Buffer)

		json.Compact(b, d)
		fmt.Fprintln(w, b)
	}))
	defer ts.Close()

	rally := New(key)
	rally.api = ts.URL + "/"

	var hr HierarchicalRequirement
	rally.Get(&hr, "id")

	assert.Equal(t, "User Story Title", hr.Name)
	assert.Equal(t, "2015-01-01T12:00:00.000Z", hr.CreationDate)
}
