package main

import "fmt"

type Expression struct {
	input           string
	output          string
	expressionStack *StringStack
}

func RunExpressionExample() {
	tests := map[string]string{
		"A+B-C":           "AB+C-",
		"A*B/C":           "AB*C/",
		"A+B*C":           "ABC*+",
		"A*B+C":           "AB*C+",
		"A*(B+C)":         "ABC+*",
		"A*B+C*D":         "AB*CD*+",
		"(A+B)*(C-D)":     "AB+CD-*",
		"((A+B)*C-D)":     "AB+C*D-",
		"A+B*(C-D/(E+F))": "ABCDEF+/-*+",
	}

	for tk, tv := range tests {
		e := NewExpression(tk)
		o := e.infixToPostFix()
		areEqual := tv == o
		fmt.Printf("Now comparing ( %s to %s ) : %t\n", o, tv, areEqual)
	}

}

func NewExpression(s string) *Expression {
	stackSize := len(s)
	es := ExpressionStack(stackSize)

	return &Expression{
		input:           s,
		output:          "",
		expressionStack: es,
	}
}

func ExpressionStack(m int) *StringStack {
	items := make([]string, m, m)

	return &StringStack{
		maxSize: m,
		items:   items,
		top:     0,
	}
}

func (e *Expression) String() string {
	return fmt.Sprintf("%s", e.input)
}

func (e *Expression) Evaluate() {
	fmt.Printf("Expression Evaluated: '%s' ", e.String())
	/*
		pf := e.toPostFix()
		return e.evaluatePostFix(pf)
	*/
}

func (e *Expression) gotOper(c string, p1 int) {
	for !e.expressionStack.IsEmpty() {
		t := e.expressionStack.Pop()
		if t == "(" {
			e.expressionStack.Push(t)
			break
		} else {
			var p2 int
			if t == "+" || t == "-" {
				p2 = 1
			} else {
				p2 = 2
			}
			if p2 < p1 {
				e.expressionStack.Push(t)
				break
			} else {
				e.output = e.output + t
			}
		}
	}

	e.expressionStack.Push(c)
}

func (e *Expression) gotParen(c string) {
	for !e.expressionStack.IsEmpty() {
		chx := e.expressionStack.Pop()
		if chx == "(" {
			break
		} else {
			e.output = e.output + chx
		}
	}
}

func (e *Expression) infixToPostFix() string {
	for _, c := range e.input {
		sc := string(c)
		switch sc {
		case "+", "-":
			e.gotOper(sc, 1)
			break
		case "*", "/":
			e.gotOper(sc, 2)
			break
		case "(":
			e.expressionStack.Push(sc)
			break
		case ")":
			e.gotParen(sc)
			break
		default:
			e.output = e.output + sc
			break
		}
	}
	for !e.expressionStack.IsEmpty() {
		e.output = e.output + e.expressionStack.Pop()
	}
	return e.output
}

/*
1. Transform expression into postfix notation
2. Evaluate that expression

Example Postfix Transformations:
A+B-C => AB+C-
A*B/C => AB*C/
A+B*C => ABC*+
A*B+C => AB*C+
A*(B+C) => ABC+*
A*B+C*D => AB*CD*+
(A+B)*(C-D) => AB+CD-*
((A+B)*C-D) => AB+C*D-
A+B*(C-D/(E+F)) => ABCDEF+/-*+
*/
