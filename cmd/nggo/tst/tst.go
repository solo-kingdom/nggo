package main

import (
	"fmt"
)

func GetAddress(d interface{}) string {
	fmt.Printf("c: %p\n", &d)
	address := fmt.Sprint(&d)
	//fmt.Printf("Address of var argument: %s\n", address)
	//var adr uint64
	//adr, err := strconv.ParseUint(address, 0, 64)
	//if err != nil {
	//	panic(err)
	//}
	//var ptr = uintptr(adr)
	//goland:noinspection GoVetUnsafePointer
	//p := unsafe.Pointer(ptr)
	//return *(*string)(p)
	return address
}

func PointArg(l interface{}) {
	fmt.Printf("%p %p\n", l, &l)
}

type Person struct {
	name string
}

func main() {
	p := Person{name: "wii"}
	fmt.Printf("%T %v %p\n", p, p, &p)
	PointArg(p)

	l := make([]int, 0, 10)
	fmt.Printf("%T %p\n", l, l)
	PointArg(l)

	// 设置容量
	hi := "hi"
	s := make([]int, 0, 10)
	s = append(s, 1)
	s = append(s, 1)
	fmt.Printf("s: id=%s, len=%d, cap=%d\n", GetAddress(s), len(s), cap(s))
	fmt.Printf("a: id=%v, %p\n", s, s)
	s = append(s, 10)
	fmt.Printf("after append, s: id=%s, len=%d, cap=%d\n", GetAddress(&s), len(s), cap(s))
	fmt.Printf("b: id=%s\n", fmt.Sprint(&hi))

	// 不设置容量
	var f []string
	fmt.Printf("f: id=%d, len=%d, cap=%d\n", &f, len(f), cap(f))
	f = append(f, "A")
	fmt.Printf("f: id=%d, len=%d, cap=%d\n", &f, len(f), cap(f))
}
