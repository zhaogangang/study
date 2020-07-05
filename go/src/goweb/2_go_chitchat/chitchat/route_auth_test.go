package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Login(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("Get", "/login", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	if strings.Contains(body, "sign in") == false {
		t.Errorf("Body does not contain Sign in")
	}
}
