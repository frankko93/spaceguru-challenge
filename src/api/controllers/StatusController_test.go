package controllers_test

import (
	"testing"

	"github.com/frankko93/spaceguru-challenge/src/api/app"
	"github.com/frankko93/spaceguru-challenge/src/api/controllers"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	assert := assert.New(t)
	t.Run(
		"GET_PING_200",
		func(t *testing.T) {
			res := controllers.PerformRequest("GET", "/ping", nil, nil, app.Router)
			assert.Equal(200, res.Code)
		})
}
