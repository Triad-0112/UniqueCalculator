package MButton

import "fmt"

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Command interface {
	Execute()
}

type Calculator interface {
	Add()
	Multiply()
	Divide()
	Sub()
}
type AddCommand struct {
	Calculator Calculator
}

type DivideCommand struct {
	Calculator Calculator
}

type MultiplyCommand struct {
	Calculator Calculator
}

type SubCommand struct {
	Calculator Calculator
}

func (a *AddCommand) Execute() {
	a.Calculator.Add()
}

func (d *DivideCommand) Execute() {
	d.Calculator.Divide()
}

func (m *MultiplyCommand) Execute() {
	m.Calculator.Multiply()
}

func (s *SubCommand) Execute() {
	s.Calculator.Sub()
}

type ScreenCalc struct {
	A, B  int
	total int64
}

func (sC *ScreenCalc) Add() {
	sC.total = int64(sC.A + sC.B)
	fmt.Println(sC.total)
}

func (sC *ScreenCalc) Divide() {
	sC.total = int64(sC.A / sC.B)
	fmt.Println(sC.total)
}

func (sC *ScreenCalc) Multiply() {
	sC.total = int64(sC.A * sC.B)
	fmt.Println(sC.total)
}

func (sC *ScreenCalc) Sub() {
	sC.total = int64(sC.A - sC.B)
	fmt.Println(sC.total)
}
