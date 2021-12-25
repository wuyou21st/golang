package main

import (
	"fmt"
)

func main() {
    myArray := [5]string{"I", "am", "stupid", "and", "weak"}
    mySlice := myArray[:]
    for index,value := range mySlice {
    	if value == "stupid" {
			mySlice[index] = "smart"
			continue
		}
		if value == "weak" {
			mySlice[index] = "strong"
		}
	}
	fmt.Println(mySlice)
}