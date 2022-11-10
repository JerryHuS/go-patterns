package factory

import (
	"fmt"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	printer := NewPrinter("cn")
	res := printer.Print("hello")
	fmt.Println(res)
}
