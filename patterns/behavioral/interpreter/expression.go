package interpreter

import "strings"

// the interface which all expressions must implement
type Expression interface {
	Interpret(string) bool
}

type AndExpression struct {
	exp1 Expression
	exp2 Expression
}

func NewAndExpression(exp1, exp2 Expression) *AndExpression {
	ae := new(AndExpression)
	ae.exp1 = exp1
	ae.exp2 = exp2
	return ae
}

func (e *AndExpression) Interpret(context string) bool {
	return e.exp1.Interpret(context) && e.exp2.Interpret(context)
}

type OrExpression struct {
	exp1 Expression
	exp2 Expression
}

func NewOrExpression(exp1, exp2 Expression) *OrExpression {
	ae := new(OrExpression)
	ae.exp1 = exp1
	ae.exp2 = exp2
	return ae
}

func (e *OrExpression) Interpret(context string) bool {
	return e.exp1.Interpret(context) || e.exp2.Interpret(context)
}

type TerminalExpression struct {
	str string
}

func (e *TerminalExpression) Interpret(context string) bool {
	words := strings.Split(context, " ")

	for _, word := range words {
		if word == e.str {
			return true
		}
	}
	return false
}

func NewTerminalExpression(str string) *TerminalExpression {
	ae := new(TerminalExpression)
	ae.str = str
	return ae
}
