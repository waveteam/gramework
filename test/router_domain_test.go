package test

import (
	"testing"

	"github.com/waveteam/gramework"
)

func TestDomainShouldNeverReturnNil(t *testing.T) {
	app := gramework.New()
	if app.Domain("test") == nil {
		t.FailNow()
	}
}
