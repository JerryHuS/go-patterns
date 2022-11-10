package factory

type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	Compute() int
}

type OperatorFactory interface {
	Create() MathOperator
}

type basicOperator struct {
	a, b int
}

func (b *basicOperator) SetOperandA(num int) {
	b.a = num
}

func (b *basicOperator) SetOperandB(num int) {
	b.b = num
}

type PlusOperator struct {
	*basicOperator
}

func (p *PlusOperator) Compute() int {
	return p.a + p.b
}

type PlusOperatorFactory struct{}

func (p *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{basicOperator: &basicOperator{}}
}
