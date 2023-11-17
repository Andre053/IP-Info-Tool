package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type GatherIpTests struct {
	name   string
	data   []byte
	result []ipIds
}

type ServeTests struct {
	name          string
	server        *httptest.Server
	response      *Response
	expectedError error
}

// tests for http request function
func TestIpInfoLookup(t *testing.T) {

	tests := []ServeTests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{ "ip": "8.8.8.8", "hostname": "google.com", "city": "blah", "region": "reg", "country": "country", "loc": "0, 0", "org": "Google", "postal": "12345", "timezone": "EST"}`))
			})),
			response: &Response{
				Ip:       "8.8.8.8",
				Hostname: "google.com",
				City:     "blah",
				Region:   "reg",
				Country:  "country",
				Loc:      "0, 0",
				Org:      "Google",
				Postal:   "12345",
				Timezone: "EST",
			},
			expectedError: nil,
		},
		{
			name: "missing-data-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{ "ip": "8.8.8.8", "region": "reg", "country": "country", "loc": "0, 0", "org": "Google", "postal": "12345", "timezone": "EST"}`))
			})),
			response:      &Response{},
			expectedError: &MissingJsonValuesError{Line: 51, Col: 4},
		},
		{
			name: "no-data-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})),
			response:      &Response{},
			expectedError: &NoHttpData{Line: 43, Col: 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			url := test.server.URL
			time.Sleep(time.Millisecond * 100)
			res, err := handleHTTPRequest(url)

			if !reflect.DeepEqual(res, test.response) {
				t.Errorf("TEST FAILED: Expected %v, got %v\n", test.response, res)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("TEST FAILED: Expected %v, got %v\n", test.expectedError.Error(), err)
			}

		})
	}
}

// none of these tests should return nil error
func FuzzBadGetIpsFromUser(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}, []byte("123.123.123.1233"), []byte("0.0.0.0"), []byte("00000")} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		_, err := getIpsFromUser()
		if err == nil {
			t.Fatalf("%v: parse success: %v", in, err)
		}
	})
}
