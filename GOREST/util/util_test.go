package util_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"testing"
)

type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

func New(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (r *FakeResponse) Assert(status int, body string) {
	if r.status != status {
		r.t.Errorf("expected status %+v to equal %+v", r.status, status)
	}
	if string(r.body) != body {
		r.t.Errorf("expected body %+v to equal %+v", string(r.body), body)
	}
}

func TestWriteJSON(t *testing.T) {

	var response interface{}
	w := New(t)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		t.Errorf("[ERROR] TestWriteJSON : %d\n", err)
	}

}

func TestRequireJSON(t *testing.T) {

	var refer interface{}

	os.Chdir("../config")
	pwd, _ := os.Getwd()
	path := pwd + "/config.json"

	raw, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("[ERROR] TestRequireJSON : %d\n", err)
	}
	re := regexp.MustCompile("(?s)[^https?:]//.*?\n|/\\*.*?\\*/")
	valid := re.ReplaceAll(raw, nil)
	if err := json.Unmarshal(valid, &refer); err != nil {
		t.Errorf("[ERROR] TestRequireJSON : %d\n", err)
	}

	return

}

func TestIntInSlice(t *testing.T) {

	num := 1
	list := []int{1, 2, 3, 4, 5}

	for _, b := range list {
		if b == num {
			return
		}
	}

	t.Errorf("[ERROR] TestIntInSlice ")

}
