package global 

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"domain"
)


var Globalenv  domain.EnviromentData

//
func  ReadConfg(Enviromet string) (domain.ResultModel) { 	
	var filename =fmt.Sprintf("config.%s.json",Enviromet) 
	rtmodel :=domain.ResultModel {} 
	if _, err := os.Stat(filename); err == nil { 
		rtmodel.Success=true;
		rtmodel.Message="성공";
		dat, _ := ioutil.ReadFile(filename)		 
        json.Unmarshal(dat, &Globalenv)  
		return rtmodel
     }else {
		rtmodel.Success=true;
		rtmodel.Message="File Not Found"; 
		return rtmodel
    } 
}