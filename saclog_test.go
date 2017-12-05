package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Saclog(t *testing.T) {
	t.Run("parseGolangSyslog", func(t *testing.T) {
		p := parseGolangSyslog("2017-12-04T19:26:45-08:00 goji.mailsac.com inbound[61504]: Starting server on port 25")
		assert.Equal(t, "goji.mailsac.com", p["hostname"], "wrong hostname")
		assert.Equal(t, "inbound", p["tag"], "wrong tag")
		assert.Equal(t, "Starting server on port 25", p["content"], "wrong content")
	})
}
