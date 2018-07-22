package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type struct for action
type Action struct {
	Command string
	Module  string
	Args    [10]string
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
	a := Action{}
	a.Command = "exit"
	return a
}

// GetHelpActiion returns the help function
func GetHelpAction() Action {
	a := Action{}
	a.Command = "help"
	return a
}

// private function

func parse(text string) Action {
	/*
	* Actions are made up of
	* 1. The Command: start, stop, status, restart or config
	* 2. The Module: proxy, intruder or repeater
	* 3. The Args: Everything else
	*
	* An Action will look like so: [Command] [Module] [[arg1], [arg2]..[arg10]]
	 */

	// TODO(greatwhite): better parsing than just splits
	// FIXME(greatwhite): doesn't return proper values
	a := Action{}
	broken := strings.Split(text, " ")

	// fmt.Println(len(broken))
	if len(broken) >= 3 {
		a.Command = broken[0]
		a.Module = broken[1]
		i := 0
		for i = 2; i < len(broken); i++ {
			a.Args[i-2] = broken[i]
		}
		return a
	} else if len(broken) == 1 {
		if broken[0] == "exit" {
			return GetExitAction()
		} else if broken[0] == "help" {
			return GetHelpAction()
		}
	}
	return GetNopAction()
}
