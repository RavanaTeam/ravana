package ravana

import (
	"fmt"
	"os"
	"ravana/lib/console"
	//"ravana/lib/core/intruder"
	//"ravana/lib/core/proxy"
	//"ravana/lib/core/repeater"
)

// the basic struct
type Ravana struct {
	//pro proxy.Proxy
	//inr intruder.Intruder
	//rep repeater.Repeater
	cli console.Console

	// greatwhite: any errors go here
	serr error
}

// New() creates a new ravana
// param proxy_port: port for the proxy to listen on
func New(proxyPort int) (Ravana, error) {
	r := Ravana{}

	//r.pro, r.serr = proxy.New(proxy_port)
	//r.inr, r.serr = intruder.New()
	//r.rep, r.serr = repeater.New()
	r.cli, r.serr = console.New()

	return r, r.serr
}

// starts the run loop
func (r Ravana) Run() {
	cmd := console.GetNopAction()

	for cmd != console.GetExitAction() {
		// fetch the command
		cmd = r.cli.Prompt()

		// handle the action
		handler(cmd)
	}

	// user exited
	shutdown(r.serr)
}

// private functions

func handler(cmd console.Action) {
	// TODO(greatwhite)
	fmt.Println(cmd)
}

func shutdown(err error) {
	// TODO(greatwhite)
	os.Exit(0)
}
