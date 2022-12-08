package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(){

	rating := bufio.NewReader(os.Stdin)
	input,err := rating.ReadString('\n')
	if err != nil {
		fmt.Println("there is a error")
	}else {
        fmt.Printf("the value of the ratng is %T \n",input)
		fmt.Println(input)
	}
	ratting,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil {
		fmt.Println("there is an error")
	}else{
		fmt.Printf("the type of the rating %T" ,ratting)
	}
}

