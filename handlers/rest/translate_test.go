package rest_test

import (
	"encoding/json"
	"hello-api/handlers/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          404,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	// Arrange

	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		// Act
		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf("expected status %d but received %d",
				test.StatusCode, rr.Code)
		}

		// Assert
		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf("expected language %s but received %q", test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation %s but received %s`, test.ExpectedTranslation, resp.Translation)
		}
	}
}
