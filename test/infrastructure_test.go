package test

import (
	"testing"

	"github.com/waveteam/gramework"
	"github.com/waveteam/gramework/infrastructure"
)

func TestServeInfrastructure(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Logf("Serve Infrastructure should not panic, but %v", e)
			t.FailNow()
		}
	}()
	app := gramework.New()
	i := infrastructure.New()
	app.ServeInfrastructure(i)
}
