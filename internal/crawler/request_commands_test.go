package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goncalocool/coolcoolcool/internal/data"
)

type MockCommandPassesAndChangesKey struct {
}

func (c *MockCommandPassesAndChangesKey) Command(req *data.Request) (ok bool) {
	req.Key = "new"
	return true
}

type MockCommandFails struct {
}

func (c *MockCommandFails) Command(req *data.Request) (ok bool) {
	return false
}

func TestExecuteCommandAndChangeReq(t *testing.T) {
	req := data.Request{}

	commands := []RequestCommand{
		new(MockCommandPassesAndChangesKey)}
	ExecuteCommands(&req, commands...)

	assert.Equal(t, "new", req.Key, "should modify Referer by ref")
}

func TestBreakBeforeSecondCommand(t *testing.T) {
	req := data.Request{}

	commands := []RequestCommand{
		new(MockCommandFails),
		new(MockCommandPassesAndChangesKey)}
	ExecuteCommands(&req, commands...)

	assert.Equal(t, "", req.Key, "should not reach second command with Key")
}
