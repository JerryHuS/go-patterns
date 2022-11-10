package factory

import "fmt"

type Printer interface {
	Print(string) string
}

type CnPrinter struct {
}

func (cn *CnPrinter) Print(str string) string {
	return fmt.Sprintf("Cn Print %s", str)
}

type EnPrinter struct {
}

func (en *EnPrinter) Print(str string) string {
	return fmt.Sprintf("En Print %s", str)
}

func NewPrinter(str string) Printer {
	switch str {
	case "cn":
		return &CnPrinter{}
	default:
		return &EnPrinter{}
	}
}
