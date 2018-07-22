package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type struct for action
type Action struct {
	command string
	module  string
	args    [10]string
}

// type struct for console
type Console struct {
	// TODO(greatwhite): what does a CLI need?
}

// New() returns a new console instance
func New() (Console, error) {

	return Console{}, nil
}

// Prompt() prompt's the user for input and takes it
func (cli Console) Prompt() Action {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ravana>")

	// TODO(greatwhite): Actually handle the error
	text, _ := reader.ReadString('\n')

	return parse(text)
}

// static functions

// GetNopAction returns an empty Action
func GetNopAction() Action {
	return Action{}
}

// GetExitAction returns the exit action
func GetExitAction() Action {
	return Action{
		command: "exit",
	}
}

// private function

func parse(text string) Action {
	/*
	* Actions are made up of
	* 1. The command: start, stop, status, restart or config
	* 2. The module: proxy, intruder or repeater
	* 3. The args: Everything else
	*
	* An Action will look like so: [command] [module] [[arg1], [arg2]..[arg10]]
	 */

	// TODO(greatwhite): better parsing than just splits
	a := Action{}
	broken := strings.Split(text, " ")

	a.command = broken[0]
	a.module = broken[1]

	i := 0
	for i = 2; i < len(broken); i++ {
		a.args[i-2] = broken[i]
	}

	return a
}
