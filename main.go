package main

import (
	"fmt"
	"global"
	"os"

	"domain"
	"images"
)

func main() {

	rtmodel := domain.ResultModel{}

	fmt.Println("모듈을 시작합니다.")
	if len(os.Args) == 1 {
		rtmodel = global.ReadConfg("dev") //환경설정
	} else {
		rtmodel = global.ReadConfg(os.Args[1])
	}
	if !rtmodel.Success {
		os.Exit(3)
	}

	images.Imagedown("https://1.bp.blogspot.com/-cNRX4ory3lI/XmBCIvthS7I/AAAAAAAOrZ4/ZqfENsHP8AAeaZakN4I5jV3SqJAr4mfJQCLcBGAsYHQ/s1600/2.gif")

}
