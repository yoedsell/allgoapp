package Controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmLogin(t *testing.T) {

	url := "http://localhost:8080/login"

	//data of type byte sclice
	var jsonStr = []byte(`{"email" :"yoedsellnamgyall@gmail.com", "password" : "wee"}`)

	//create http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	//set request header
	req.Header.Set("content-Type", "application/json")

	//create a pointer variable client which points to client
	client := &http.Client{}

	//client sends http request using Do() and gets http response
	resp, err := client.Do(req)

	//handle error if any
	if err != nil {
		panic(err)
	}
	// defer the closing of response body until function terminates
	defer resp.Body.Close()

	// get data from the response body
	body, _ := io.ReadAll(resp.Body)

	// validate if response status is same as expected status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expResp := `{"message":"Login success"}`
	// validate if response body is same as expected response body
	assert.JSONEq(t, expResp, string(body))

}
