package factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var factory AbstractFactory
	var tv TV

	factory = &MiFactory{}
	tv = factory.CreateTV()
	tv.Watch()

	factory = &HuaweiFactory{}
	tv = factory.CreateTV()
	tv.Watch()
}
