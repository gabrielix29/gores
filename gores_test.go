package gores

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Name  string
	Email string
	Age   int
}

var user = User{
	Name:  "name",
	Email: "name@email.com",
	Age:   30,
}

func TestJSON(t *testing.T) {
	responseString := `{"Name":"name","Email":"name@email.com","Age":30}`
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	JSON(resp, responseCode, user)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationJSONCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}
