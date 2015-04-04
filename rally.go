package rally

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const apiUrl = "https://rally1.rallydev.com/slm/webservice/v2.0/"
const apiSessionKey = "ZSESSIONID"

type rally struct {
	url   string
	token string
}

// New returns a rally instance storing the necessary data. The token string
// must be obtained from rally via: https://rally1.rallydev.com/login/
//
// Basic auth with username and password is not supported at the moment
func New(token string) *rally {
	return &rally{url: apiUrl, token: token}
}

// Get a populated instance of the supplied object given the id
func (r *rally) Get(object interface{}, id string) error {
	kind := getStructType(object)

	parts := []string{strings.ToLower(kind), "/", id}
	url := r.generateUrl(strings.Join(parts, ""))

	body, err := r.get(url)
	if err != nil {
		return err
	}

	return unmarshal(object, body, kind)
}

// Fetch a populated instance of the supplied object given the reference
// supplied
func (r *rally) Fetch(object interface{}, ref reference) error {
	kind := getStructType(object)

	if len(ref.ReferenceUrl) == 0 {
		message := fmt.Sprintf("null reference for: %s", kind)
		return errors.New(message)
	}

	body, err := r.get(ref.ReferenceUrl)
	if err != nil {
		return err
	}

	return unmarshal(object, body, kind)
}

// Fetch a populated query instance of the supplied query object given the
// query reference supplied
func (r *rally) QueryFetch(object interface{}, ref queryReference) error {
	if len(ref.ReferenceUrl) == 0 {
		kind := getStructType(object)
		message := fmt.Sprintf("null query reference for: %s", kind)
		return errors.New(message)
	}

	body, err := r.get(ref.ReferenceUrl)
	if err != nil {
		return err
	}

	return unmarshal(object, body, "QueryResult")
}

// Concatinates the api url with the supplied uri
func (r *rally) generateUrl(uri string) string {
	var buffer bytes.Buffer

	buffer.WriteString(r.url)
	buffer.WriteString(uri)

	return buffer.String()
}

// Convenience method for HTTP GET requests
func (r *rally) get(url string) ([]byte, error) {
	return r.call("GET", url)
}

// Call the supplied url using the HTTP method
func (r *rally) call(method string, url string) (body []byte, err error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add(apiSessionKey, r.token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message := fmt.Sprintf("http status code: %d", response.StatusCode)
		return nil, errors.New(message)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Get the name of the struct
func getStructType(object interface{}) string {
	name := fmt.Sprintf("%T", object)
	splitted := strings.Split(name, ".")

	return splitted[len(splitted)-1]
}

// Unmarshal the body onto the supplied object
func unmarshal(object interface{}, body []byte, key string) (err error) {
	var data map[string]*json.RawMessage

	if err = json.Unmarshal(body, &data); err != nil {
		return err
	}

	errorResponseKey := "OperationResult"
	if _, isset := data[errorResponseKey]; isset {
		if err = checkResponse(data, errorResponseKey); err != nil {
			return err
		}
	}

	if err = json.Unmarshal(*data[key], &object); err != nil {
		return err
	}

	if err = checkResponse(data, key); err != nil {
		return err
	}

	return nil
}

// Check the response for any errors or warnings. Warnings will be logged
// and errors will be returned
func checkResponse(data map[string]*json.RawMessage, key string) error {
	type result struct {
		Errors   []string
		Warnings []string
	}

	var r result
	json.Unmarshal(*data[key], &r)

	if len(r.Warnings) > 0 {
		message := strings.Join(r.Warnings, "\"; \"")
		log.Println("rally api warnings: \"" + message + "\"")
	}

	if len(r.Errors) > 0 {
		message := strings.Join(r.Errors, "\"; \"")
		return errors.New("rally api errors: \"" + message + "\"")
	}

	return nil
}
