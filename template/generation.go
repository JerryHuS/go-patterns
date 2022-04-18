/**
 * @Author: alessonhu
 * @Description:
 * @File:  generation.go
 * @Version: 1.0.0
 * @Date: 2022/3/29 8:03 下午
 */

package PACKAGE_NAME

import "fmt"

//go:generate ./gen.sh ./con.tmp.go gen uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./con.tmp.go gen string container
func generateStringExample() {
	var s string = "Hello"
	c := NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateExample: %s (%T)\n", v, v)
}
