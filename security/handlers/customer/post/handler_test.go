package post

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/a-h/go-workshop-102/security/models"
	"github.com/google/go-cmp/cmp"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name           string
		r              *http.Request
		cp             CustomerPutter
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "posting invalid JSON returns a bad request status code",
			r:              httptest.NewRequest("POST", "/customer", strings.NewReader(`{`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "bad request\n",
		},
		{
			name:           "posting invalid data returns a bad request status code",
			r:              httptest.NewRequest("POST", "/customer", strings.NewReader(`{ "name": "Wolfgang", "company": "" }`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "bad request\n",
		},
		{
			name: "database failures return a 500 status",
			r:    httptest.NewRequest("POST", "/customer", strings.NewReader(`{ "name": "Wolfgang", "surname": "Mozart", "company": "" }`)),
			cp: newMockCustomerPutter(func(c models.Customer) error {
				return errors.New("database error")
			}),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "internal server error\n",
		},
		{
			name: "posting valid data returns a created status code",
			r:    httptest.NewRequest("POST", "/customer", strings.NewReader(`{ "name": "Wolfgang", "surname": "Mozart", "company": "" }`)),
			cp: newMockCustomerPutter(func(c models.Customer) error {
				// No error.
				return nil
			}),
			expectedStatus: http.StatusCreated,
			expectedBody:   "",
		},
	}

	log := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{}))
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := New(log, test.cp)

			w := httptest.NewRecorder()
			h.ServeHTTP(w, test.r)

			if w.Code != test.expectedStatus {
				t.Errorf("unexpected status - want %d, got %d", test.expectedStatus, w.Code)
			}
			if diff := cmp.Diff(test.expectedBody, w.Body.String()); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func newMockCustomerPutter(f func(c models.Customer) error) CustomerPutter {
	return mockCustomerPutter{f: f}
}

type mockCustomerPutter struct {
	f func(c models.Customer) error
}

func (cp mockCustomerPutter) PutCustomer(c models.Customer) error {
	return cp.f(c)
}
