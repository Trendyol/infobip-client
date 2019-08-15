package infobip

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestInfoBipSendMessage(t *testing.T) {
	request := []Message{
		Message{
			Sender: "test sender",
		},
	}

	response := Response{
		Result: []Result{
			Result{
				Status:      "1",
				MessageID:   "2",
				Destination: "3",
			},

			Result{
				Status:      "3",
				MessageID:   "4",
				Destination: "5",
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}))

	defer ts.Close()

	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	infoBipClient := New(ts.URL, "test-username", "test-password", httpClient)

	resp, err := infoBipClient.Send(request)
	if err != nil {
		t.Fatalf("unexpected error. err: %s", err)
	}

	if !reflect.DeepEqual(&response, resp) {
		t.Errorf("expected response: %s but got :%s", response, resp)
	}

}
