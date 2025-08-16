package main

import "fmt"

func main(){
	x:=map[string]int{}
	t:="hello world"
	for _,i:=range t{
		x[string(i)]=x[string(i)]+1
	}


	
	fmt.Println(x)
}