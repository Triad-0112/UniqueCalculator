package main

import (
	"main.go/MButton"
)

func main() {

	calculator := &MButton.ScreenCalc{
		A: 700,
		B: 50,
	}
	addCommand := &MButton.AddCommand{
		Calculator: calculator,
	}
	subCommand := &MButton.SubCommand{
		Calculator: calculator,
	}
	multiplyCommand := &MButton.MultiplyCommand{
		Calculator: calculator,
	}
	divideCommand := &MButton.DivideCommand{
		Calculator: calculator,
	}
	addButton := &MButton.Button{
		Command: addCommand,
	}
	addButton.Press()
	subButton := &MButton.Button{
		Command: subCommand,
	}
	subButton.Press()
	multiplyButton := &MButton.Button{
		Command: multiplyCommand,
	}
	multiplyButton.Press()
	divideButton := &MButton.Button{
		Command: divideCommand,
	}
	divideButton.Press()
}
