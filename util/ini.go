package util

import (
	"fmt"
	"gopkg.in/ini.v1"
)

/**
 * 获取adm.ini
 */
func GetIni(sectionStr string) (*ini.Section,error) {
	iniCfg, err := ini.Load("./adm.ini")
	if err != nil {
		return nil,err
	}

	section,getScErr := iniCfg.GetSection(sectionStr)
	if getScErr != nil {
		return nil,getScErr
	}

	return section,nil
}

func GetIniKey(sectionStr string ,key string) (string,error) {
	section,err := GetIni(sectionStr)
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	keyObj,err := section.GetKey(key)
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	return keyObj.Value(),nil
}

func GetCosIni(key string) (string,error) {
	return GetIniKey("cos",key)
}

func GetRedisIni(key string) (string,error) {
	return GetIniKey("redis",key)
}
