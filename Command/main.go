package main

import (
	"fmt"
)

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) turnOn() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) turnOff() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

// Command interface
type Command interface {
	execute()
}

// ConcreteCommand
type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) execute() {
	c.light.turnOn()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) execute() {
	c.light.turnOff()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) pressButton() {
	r.command.execute()
}

func main() {
	light := &Light{}
	turnOnCommand := &TurnOnCommand{light: light}
	turnOffCommand := &TurnOffCommand{light: light}

	remote := &RemoteControl{}

	var choice int
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Turn On the Light")
		fmt.Println("2. Turn Off the Light")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			remote.command = turnOnCommand
			remote.pressButton()
		case 2:
			remote.command = turnOffCommand
			remote.pressButton()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
