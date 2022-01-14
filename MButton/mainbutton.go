package MButton

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type Calculator interface {
	add()
	multiply()
	divide()
	sub()
}
