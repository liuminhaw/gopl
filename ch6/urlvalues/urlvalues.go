package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // Direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"
	fmt.Println(m["item"])     // "[1, 2]"

}
