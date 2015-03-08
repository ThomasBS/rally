package rally

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var api = "https://rally1.rallydev.com/slm/webservice/v2.0/"
var headerRequestKey = "ZSESSIONID"

type rally struct {
	api string
	key string
}

func New(k string) *rally {
	return &rally{key: k, api: api}
}

type Reference struct {
	ReferenceUrl string `json:"_ref"`
}

type QueryReference struct {
	ReferenceUrl string `json:"_ref"`
	Count        int
}

type QueryResult struct {
	TotalResultCount int
	StartIndex       int
	PageSize         int
}

type PersistableObject struct {
	CreationDate string
}

type DomainObject struct {
	PersistableObject
}

type Workspace struct {
	Name string
}

type WorkspaceDomainObject struct {
	DomainObject

	WorkspaceReference Reference `json:"Workspace"`
}

type Artifact struct {
	WorkspaceDomainObject

	Name string
}

type SchedulableArtifact struct {
	Artifact
}

type Task struct {
	Artifact

	Name string
}

type TaskQuery struct {
	QueryResult
	Results []Task
}

type Requirement struct {
	SchedulableArtifact
}

type HierarchicalRequirement struct {
	Requirement

	TasksQueryReference QueryReference `json:"Tasks"`
}

func (r *rally) Fetch(object interface{}, ref Reference) {

	name := fmt.Sprintf("%T", object)
	splitted := strings.Split(name, ".")
	kind := splitted[len(splitted)-1]

	client := &http.Client{}
	request, _ := http.NewRequest("GET", ref.ReferenceUrl, nil)
	request.Header.Add(headerRequestKey, r.key)

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var j map[string]*json.RawMessage
	json.Unmarshal(body, &j)
	json.Unmarshal(*j[kind], &object)
}

func (r *rally) QueryFetch(object interface{}, ref QueryReference) {

	client := &http.Client{}
	request, _ := http.NewRequest("GET", ref.ReferenceUrl, nil)
	request.Header.Add(headerRequestKey, r.key)

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var j map[string]*json.RawMessage
	json.Unmarshal(body, &j)
	json.Unmarshal(*j["QueryResult"], &object)
}

func (r *rally) Get(object interface{}, id string) {

	name := fmt.Sprintf("%T", object)
	splitted := strings.Split(name, ".")
	kind := splitted[len(splitted)-1]

	parts := []string{r.api, strings.ToLower(kind), "/", id}
	url := strings.Join(parts, "")

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add(headerRequestKey, r.key)

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var j map[string]*json.RawMessage
	json.Unmarshal(body, &j)
	json.Unmarshal(*j[kind], &object)
}
