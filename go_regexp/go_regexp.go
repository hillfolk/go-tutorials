package main

import (
	"fmt"
	"regexp"
)



func main(){


	if matched, _ := regexp.MatchString("devices/", "/devces/123123");matched {
		
	fmt.Println(matched) // true
	}
}
	
