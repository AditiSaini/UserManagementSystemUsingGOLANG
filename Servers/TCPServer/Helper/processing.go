package helper

import (
	"strings"
)

func ExtractingArgumentsFromCommands(instruction string, args string) map[string]string {
	m := make(map[string]string)
	possibleInstruction := []string{"LOGIN", "LOGOUT", "SHOW_PROFILE", "UPDATE_PROFILE", "UPLOAD_PICTURE", "CHANGE_PASSWORD"}
	_, found := Find(possibleInstruction, instruction)

	if !found {
		m["error"] = "Command not found"
		return m
	}
	commands := strings.Split(args, "|")
	for _, command := range commands {
		cmd := strings.Split(command, " ")
		argument := cmd[0]
		data := strings.Join(cmd[1:], " ")
		m[argument] = data
	}
	return m
}

// Find takes a slice and looks for an element in it. If found it will return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
