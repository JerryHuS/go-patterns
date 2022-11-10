package factory

import "fmt"

type AbstractFactory interface {
	CreateTV() TV
	CreateAir() Air
}

type TV interface {
	Watch()
}

type Air interface {
	SetTmp(int)
}

// MiFactory xiaomi
type MiFactory struct{}

func (m *MiFactory) CreateTV() TV {
	return &MiTV{}
}

func (m *MiFactory) CreateAir() Air {
	return &MiAir{}
}

type MiTV struct{}

func (m *MiTV) Watch() {
	fmt.Println("watch Mi TV")
}

type MiAir struct{}

func (m *MiAir) SetTmp(t int) {
	fmt.Printf("set tmp %d\n", t)
}

// HuaweiFactory huawei
type HuaweiFactory struct{}

func (h *HuaweiFactory) CreateTV() TV {
	return &HuaweiTV{}
}

func (h *HuaweiFactory) CreateAir() Air {
	return &HuaweiAir{}
}

type HuaweiTV struct{}

func (h *HuaweiTV) Watch() {
	fmt.Println("watch Huawei TV")
}

type HuaweiAir struct{}

func (h *HuaweiAir) SetTmp(t int) {
	fmt.Sprintf("set tmp %d\n", t)
}
