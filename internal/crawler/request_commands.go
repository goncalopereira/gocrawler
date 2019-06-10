package crawler

import (
	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//RequestCommand list of actions for new requests
type RequestCommand interface {
	Command(req *data.Request) (ok bool)
}

//ExecuteCommands check all actions
func ExecuteCommands(req *data.Request, commands ...RequestCommand) {
	for _, command := range commands {
		ok := command.Command(req)
		if !ok {
			break
		}
	}
}
