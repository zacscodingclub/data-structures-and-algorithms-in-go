package main

import (
	"fmt"
	"strconv"
)

type Expression struct {
	input         string
	output        string
	operatorStack *StringStack
	operandStack  *Stack
}

func RunExpressionExample() {
	e := NewExpression("3*(4+5)-6/(1+2)")
	f := NewExpression("1+2")

	fmt.Printf("%s evaluates to: %s, expected: 3\n", f.String(), f.Evaluate())
	fmt.Printf("%s evaluates to: %s, expected: 25\n", e.String(), e.Evaluate())
}

func RunExpressionInfixToPostfixExamples() {
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
	operatorStack := OperatorStack(stackSize)
	operandStack := NewStack(stackSize)

	return &Expression{
		input:         s,
		output:        "",
		operatorStack: operatorStack,
		operandStack:  operandStack,
	}
}

func OperatorStack(m int) *StringStack {
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

func (e *Expression) Evaluate() string {
	pf := e.infixToPostFix()
	value := e.evaluatePostFix(pf)
	return strconv.Itoa(value)
}

func (e *Expression) evaluatePostFix(pf string) int {
	var ans int

	for _, c := range pf {
		sc := string(c)

		if sc >= "0" && sc <= "9" {
			ic, err := strconv.Atoi(sc)
			if err != nil {
				fmt.Printf("Could not convert ( %s ) into an integer because: \n", sc, err)
			}
			e.operandStack.Push(ic)
		} else {
			n2 := e.operandStack.Pop()
			n1 := e.operandStack.Pop()

			switch sc {
			case "+":
				ans = n1 + n2
				break
			case "-":
				ans = n1 - n2
				break
			case "*":
				ans = n1 * n2
				break
			case "/":
				ans = n1 / n2
				break
			default:
				ans = 0
			}
			e.operandStack.Push(ans)
		}
	}
	ans = e.operandStack.Pop()
	return ans
}

func (e *Expression) gotOper(c string, p1 int) {
	for !e.operatorStack.IsEmpty() {
		t := e.operatorStack.Pop()
		if t == "(" {
			e.operatorStack.Push(t)
			break
		} else {
			var p2 int
			if t == "+" || t == "-" {
				p2 = 1
			} else {
				p2 = 2
			}
			if p2 < p1 {
				e.operatorStack.Push(t)
				break
			} else {
				e.output = e.output + t
			}
		}
	}

	e.operatorStack.Push(c)
}

func (e *Expression) gotParen(c string) {
	for !e.operatorStack.IsEmpty() {
		chx := e.operatorStack.Pop()
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
			e.operatorStack.Push(sc)
			break
		case ")":
			e.gotParen(sc)
			break
		default:
			e.output = e.output + sc
			break
		}
	}
	for !e.operatorStack.IsEmpty() {
		e.output = e.output + e.operatorStack.Pop()
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
