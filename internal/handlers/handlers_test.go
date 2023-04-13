package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "get", []postData{}, http.StatusOK},
	{"about", "/about", "get", []postData{}, http.StatusOK},
	{"generals", "/generals-quarters", "get", []postData{}, http.StatusOK},
	{"majors", "/majors-suite", "get", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "get", []postData{}, http.StatusOK},
	{"res", "/make-reservations", "get", []postData{}, http.StatusOK},
	{"contact", "/contact", "get", []postData{}, http.StatusOK},
	{"post-sa", "/search-availability", "post", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "end", value: "2021-01-02"},
	}, http.StatusOK},
	{"post-sa-json", "/search-availability-json", "post", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "end", value: "2021-01-02"},
	}, http.StatusOK},
	{"post-res", "/make-reservations", "post", []postData{
		{key: "first_name", value: "Avishka"},
		{key: "last_name", value: "Umesh"},
		{key: "email", value: "test@test.com"},
		{key: "phone", value: "0771234567"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, e := range theTest {

		if e.method == "get" {
			resp, err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := testServer.Client().PostForm(testServer.URL+e.url, values)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}

}
