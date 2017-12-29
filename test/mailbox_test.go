package test

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/inwecrypto/mailbox"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {

	client := sling.New()

	user := &mailbox.User{}

	request, err := client.Post("http://localhost:8000/user").BodyJSON(user).Request()

	assert.NoError(t, err)

	var errmsg interface{}

	resp, err := client.Do(request, nil, &errmsg)

	assert.NoError(t, err)

	if assert.NoError(t, err) {
		assert.Equal(t, 200, resp.StatusCode)
	}
}
