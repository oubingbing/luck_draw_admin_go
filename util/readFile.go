package util

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

func GetHtml(path string) template.HTML {
	data, err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Printf("读html模板出错：%v\n",err.Error())
	}
	return template.HTML(string(data))
}
