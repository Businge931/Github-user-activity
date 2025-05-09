package github

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchUserActivity(t *testing.T) {
	tests := []struct {
		name           string
		mockStatusCode int
		mockBody       any
		wantErr        bool
		wantStatus     string
		wantCount      int
	}{
		{
			name:           "success",
			mockStatusCode: http.StatusOK,
			mockBody:       []UserActivity{{ID: "1", Type: PushEvent}},
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:           "not found",
			mockStatusCode: http.StatusNotFound,
			mockBody: ErrMessage{
				Message:          "Not Found",
				DocumentationURL: "https://docs.github.com/rest/activity/events#list-events-for-the-authenticated-user",
				Status:           "404",
			},
			wantErr:    true,
			wantStatus: "404",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseBody, _ := json.Marshal(tc.mockBody)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.mockStatusCode)
				w.Write(responseBody)
			}))
			defer ts.Close()

			oldGet := httpGet
			httpGet = func(url string) (resp *http.Response, err error) {
				return http.Get(ts.URL)
			}
			defer func() { httpGet = oldGet }()

			activity, err := FetchUserActivity("testuser")
			if tc.wantErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if tc.wantStatus != "" {
					apiErr, ok := err.(ErrMessage)
					if !ok || apiErr.Status != tc.wantStatus {
						t.Errorf("Expected status %s error, got %v", tc.wantStatus, err)
					}
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				if len(activity) != tc.wantCount {
					t.Errorf("Expected %d activities, got %d", tc.wantCount, len(activity))
				}
			}
		})
	}
}
