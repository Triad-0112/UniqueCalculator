package main

import (
	"fmt"
	"github.com/Triad-0112/UniqueCalculator/MButton"
)

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type addCommand struct {
	calculator calculator
}

type divideCommand struct {
	calculator calculator
}

type multiplyCommand struct {
	calculator calculator
}

type subCommand struct {
	calculator calculator
}

func (a *addCommand) execute() {
	a.calculator.add()
}

func (d *divideCommand) execute() {
	d.calculator.divide()
}

func (m *multiplyCommand) execute() {
	m.calculator.multiply()
}

func (s *subCommand) execute() {
	s.calculator.sub()
}

type screenCalc struct {
	a, b  int
	total int64
}

func (sC *screenCalc) add() {
	sC.total = int64(sC.a + sC.b)
	fmt.Println(sC.total)
}

func (sC *screenCalc) divide() {
	sC.total = int64(sC.a / sC.b)
	fmt.Println(sC.total)
}

func (sC *screenCalc) multiply() {
	sC.total = int64(sC.a * sC.b)
	fmt.Println(sC.total)
}

func (sC *screenCalc) sub() {
	sC.total = int64(sC.a - sC.b)
	fmt.Println(sC.total)
}

type calculator interface {
	add()
	multiply()
	divide()
	sub()
}

func main() {
	calculator := &screenCalc{
		a: 700,
		b: 50,
	}
	addCommand := &addCommand{
		calculator: calculator,
	}
	subCommand := &subCommand{
		calculator: calculator,
	}
	multiplyCommand := &multiplyCommand{
		calculator: calculator,
	}
	divideCommand := &divideCommand{
		calculator: calculator,
	}
	addButton := &button{
		command: addCommand,
	}
	addButton.press()
	subButton := &button{
		command: subCommand,
	}
	subButton.press()
	multiplyButton := &button{
		command: multiplyCommand,
	}
	multiplyButton.press()
	divideButton := &button{
		command: divideCommand,
	}
	divideButton.press()
}
