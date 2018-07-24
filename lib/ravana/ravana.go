package ravana

import (
	"fmt"
	"io/ioutil"
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

	// greatwhite: any errors go here
	serr error
}

// the module being interacted with
var globalContext string

// New() creates a new ravana
// param proxy_port: port for the proxy to listen on
func New(proxyPort int) (Ravana, error) {
	r := Ravana{}
	globalContext = ""

	//r.pro, r.serr = proxy.New(proxy_port)
	//r.inr, r.serr = intruder.New()
	//r.rep, r.serr = repeater.New()

	return r, r.serr
}

// starts the run loop
func (r Ravana) Run() {
	cmd := console.GetNopAction()
	exitact := console.GetExitAction()

	showBanner()
	fmt.Println("Type 'help' to begin")

	for cmd != exitact {
		// fetch the command
		cmd = console.Prompt(globalContext)

		// handle the action
		handler(cmd)
	}

	// user exited
	shutdown(r.serr)
}

// private functions

func handler(cmd console.Action) {
	if console.IsInvalid(cmd) {
		printInvalid(cmd.Reason)
	} else if cmd == console.GetNopAction() {
		// do nothing
	} else if cmd == console.GetHelpAction() {
		showUsage()
	} else {
		if cmd.Command == "use" {
			globalContext = cmd.Module
		} else if cmd.Command == "back" {
			globalContext = ""
		}
	}
}

func showUsage() {
	// TODO(greatwhite): generate the help message dynamically.
	help := "Help\n"
	help += "----\n\n"
	help += "COMMAND\t\tDESCRIPTION\n"
	help += "help\t\tdisplay this message\n"
	help += "use\t\tUse a module\n"
	help += "back\t\tSwitch back to the main context\n"
	fmt.Println(help)
}

func printInvalid(reason string) {
	fmt.Println("INVALID COMMAND: " + reason)
}

func shutdown(err error) {
	// TODO(greatwhite)
	os.Exit(0)
}

func authorInfo() string {
	info := "\n====| Ravana version 0.1BETA Copyleft RavanaTeam\n"
	info += "====| Github: https://github.com/RavanaTeam/Ravana\n"
	return info
}
func showBanner() {
	art, _ := ioutil.ReadFile("banner.txt")
	banner := string(art)
	banner += authorInfo()
	fmt.Println(banner)
}
