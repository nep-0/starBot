package main

func patternHelp(message string) bool {
	if message[0:4] == "help" {
		return true
	}
	return message[0:5] == " help"
}

func patternR1(message string) bool {
	if message[0:2] == "r1" {
		return true
	}
	return message[0:3] == " r1"
}

func trimR1(message string) string {
	if message[0:2] == "r1" {
		return message[2:]
	}
	return message[3:]
}

func patternSim(message string) bool {
	if message[0:3] == "sim" {
		return true
	}
	return message[0:4] == " sim"
}

func trimSim(message string) string {
	if message[0:3] == "sim" {
		return message[3:]
	}
	return message[4:]
}
