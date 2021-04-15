package main

import (
	"testing"

	"github.com/ahuigo/requests"
	"gopkg.in/h2non/gock.v1"
)

func TestFoo(t *testing.T) {
	defer gock.Off()

	gock.New("http://m.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})
	resp, err := requests.Get("http://m.com/bar")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.Text())
}
