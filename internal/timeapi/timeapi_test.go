// timeapi/timeapi_test.go
package timeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTime(t *testing.T) {
	tests := []struct {
		name       string
		location   string
		statusCode int
		wantError  bool
	}{
		{
			name:       "Valid timezone",
			location:   "America/New_York",
			statusCode: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Invalid timezone",
			location:   "Invalid/Timezone",
			statusCode: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:       "Missing location",
			location:   "",
			statusCode: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/time", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.location != "" {
				q := req.URL.Query()
				q.Add("location", tt.location)
				req.URL.RawQuery = q.Encode()
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetTime)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.statusCode)
			}

			var resp TimeResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
				t.Errorf("could not unmarshal response body: %v", err)
			}

			if tt.wantError {
				if resp.Error == "" {
					t.Errorf("expected an error, but got none")
				}
				if resp.Time != "" {
					t.Errorf("expected no time, but got: %s", resp.Time)
				}
			} else {
				if resp.Error != "" {
					t.Errorf("expected no error, but got: %s", resp.Error)
				}
				if resp.Time == "" {
					t.Errorf("expected time, but got an empty string")
				}
			}
		})
	}
}
