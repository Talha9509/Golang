package main

import "fmt"

var mm = map[string]vertex{
	"b": vertex{5934.3, 234.23},
	"c": {5934.3, 234.23},
	"d": {594.3, 4.23},
}

func main(){
	fmt.Println(mm)
}