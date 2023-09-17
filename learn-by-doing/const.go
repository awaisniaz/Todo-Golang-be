package main

import (
	"fmt"
)


type employee struct {
	name string
	age int
}
func main(){
with_const_map()
with_const_struct()
ElseIf()
}

func with_const_map(){
   var e = map[string]int{
		"a":1,
	}

	fmt.Println(e)
}

func with_const_struct(){
	var e = employee {
		name:"Awais Niaz",
		age:32,
	}
	fmt.Println(e.name)
}

func ElseIf(){
	count := 10
	
	if(count == 1){
		fmt.Println("I am if Block")
	}else{
		fmt.Println("I am Else Block")
	}
}