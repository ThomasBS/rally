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

type PersistableObject struct {
	CreationDate string
}

type DomainObject struct {
	PersistableObject
}

type WorkspaceDomainObject struct {
	DomainObject
}

type Artifact struct {
	WorkspaceDomainObject

	Name string
}

type SchedulableArtifact struct {
	Artifact
}

type Requirement struct {
	SchedulableArtifact
}

type HierarchicalRequirement struct {
	Requirement
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
