package atdd_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginFailedWithWrongPassword(t *testing.T) {
	expectedRespBody := []byte(`{"status":"wrong password"}`)
	reqBody := url.Values{}
	reqBody.Add("username", "imkk-000")
	reqBody.Add("password", "wrongP@ssw0rd")

	resp, err := http.PostForm(LOGIN_URL, reqBody)
	if err != nil {
		t.Fatal(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.NotEmpty(t, respBody)
	assert.Equal(t, expectedRespBody, respBody)
}
