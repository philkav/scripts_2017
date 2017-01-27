package main

import (
	"fmt";
	"os";
	)

func main(){
	var philvar int;
	philvar = 7;
	emilyvar := 8;

	fmt.Println("Phil:",philvar,"Emily:",emilyvar);
	for i:= 1; i<len(os.Args); i++ {
		fmt.Println(i, "Value is :", os.Args[i]);
	}
}
