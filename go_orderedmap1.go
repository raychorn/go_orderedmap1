package main

import (
	"fmt"
	"runtime"

	"github.com/elliotchance/orderedmap"
)

func main () {
	goos := runtime.GOOS
	fmt.Printf("runtime.GOOS = %s\n", goos)
    switch goos {
    case "windows":
        //kp.WinKey()
    case "darwin":
        //kp.MacKey()
    case "linux":
        //kp.UnixKey()
    default:
        //kp.TryKey()
    }
	
	m := orderedmap.NewOrderedMap()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set(123, true)

	fmt.Println("BEFORE:")
	for _, key := range m.Keys() {
		value, _:= m.Get(key)
		fmt.Println(key, value)
	}
	fmt.Println()
	// Iterate through all elements from oldest to newest:
	fmt.Println("FORWARD:")
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
	fmt.Println()
	// You can also use Back and Prev to iterate in reverse:
	fmt.Println("REVERSE:")
	for el := m.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Key, el.Value)
	}
	fmt.Println()

	m.Delete("qux")

	fmt.Println("AFTER:")
	for _, key := range m.Keys() {
		value, _:= m.Get(key)
		fmt.Println(key, value)
	}
	fmt.Println()

}