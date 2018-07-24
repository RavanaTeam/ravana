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
	Reason  string // reason for err msg
}

// Prompt prompt's the user for input and takes it
func Prompt(context string) Action {
	reader := bufio.NewReader(os.Stdin)
	prompt := "ravana"

	if context != "" {
		prompt += "(" + context + ")"
	}
	prompt += "> "

	fmt.Print(prompt)

	// TODO(greatwhite): Actually handle the error
	text, _ := reader.ReadString('\n')
	if strings.TrimSuffix(text, "\r\n") == "" {
		return GetNopAction()
	}

	return parse(text)
}

// static functions

// IsInvalid checks if an Action is invalid
func IsInvalid(act Action) bool {
	return act.Reason != ""
}

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

// GetInvalidAction returns the invalid action
func GetInvalidAction(data string) Action {
	return Action{
		Command: "invalid",
		Reason:  data,
	}
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
	a := Action{}
	broken := strings.Split(text, " ")
	if len(broken) > 12 {
		return GetInvalidAction("More than 10 arguments passed")
	}

	// get rid of all the newlines
	for index, element := range broken {
		broken[index] = strings.TrimSuffix(element, "\r\n")
	}

	// fmt.Println(len(broken))
	// fmt.Println(broken)

	if len(broken) >= 2 {
		a.Command = broken[0]
		a.Module = broken[1]
		i := len(broken) + 2
		for i = 2; i < len(broken); i++ {
			a.Args[i-2] = broken[i]
		}
		return a
	} else if len(broken) == 1 {
		if broken[0] == "exit" {
			return GetExitAction()
		} else if broken[0] == "help" {
			return GetHelpAction()
		} else if broken[0] == "back" {
			return Action{Command: "back"}
		}
	}
	return GetInvalidAction("Invalid")
}
