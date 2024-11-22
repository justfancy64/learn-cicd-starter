package main

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestApiKey(t *testing.T) {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("error in NewRequst: %v", err)
	}
	req.Header.Add("Authorization", "ApiKey e9c00484-c1ef-40bb-8901-4ab75eef02c6")
	expected := "e9c00484-c1ef-40bb-8901-4ab75eef02c6"
	out, err := auth.GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("error in GetApiKey: %v", err)
	}
	if out != expected {
		t.Fatalf("want: %v\nGot: %v\n", expected, out)
	}
}




