package main

import (
	"fmt"
	"global"
	"os"

	"domain"
)




func main() {   
	 
	rtmodel :=domain.ResultModel {}
	

	fmt.Println("모듈을 시작합니다.")
	if len(os.Args) == 1 {
		rtmodel = global.ReadConfg("dev")
		
	} else	{
		rtmodel = global.ReadConfg(os.Args[1])				
	}

	if(!rtmodel.Success){
		os.Exit(3) 
	}
	//if rt ==true {
	//	fmt.Println(mes)
	//}
	//else {
	//	global.ReadConfg(os.Args[0])
	//}

	//argwithprog := os.Args
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	//arg := os.Args[3]
	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(arg)
	//fmt.Println("hi")
}
