package helper

import (
	"strings"
)

func ExtractingArgumentsFromCommands(instruction string, args string) map[string]string {
	m := make(map[string]string)
	switch instruction {
	case "LOGIN":
		return extractForLogin(args)
	default:
		m["error"] = "Command not found"
		return m
	}
}

func extractForLogin(args string) map[string]string {
	commands := strings.Split(args, "|")
	m := make(map[string]string)
	for _, command := range commands {
		cmd := strings.Split(command, " ")
		argument := cmd[0]
		data := strings.Join(cmd[1:], " ")
		m[argument] = data
	}
	return m
}
