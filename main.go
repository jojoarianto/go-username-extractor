package main

import (
	"strings"
	"reflect"
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`@([a-z0-9_]+)`)

	text := "pinter lu wkwk. tadi aku @bingung gw @nirianto kenapa binya gw gede banget. Thanks masukkannyacolek @nanang @zeihan"
	matches := r.FindAllString(text, -1)

	// print type of matches
	fmt.Println(reflect.TypeOf(matches))

	for _, username := range matches {
		username = strings.Replace(username, "@", "", -1)
		fmt.Println(username)
	}
} 
