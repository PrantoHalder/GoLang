package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input2()  {
	fmt.Println("this is a input function")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("there is an error")
	}else{
		fmt.Println("the value of input",input)
		fmt.Printf("the type pf the string %T",input)
		
	}
	change,err2:=  strconv.ParseFloat(strings.TrimSpace(input),64)
	if err2 != nil{
		fmt.Println("there is a error")
	}else{
		fmt.Println("the value is ",change)
		fmt.Printf("the type of the value is %T\n",change)
	}


}