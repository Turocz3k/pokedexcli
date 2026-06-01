package main

func main() {
	configuration := config{
		Next:     nil,
		Previous: nil,
	}
	startRepl(&configuration)
}
