package services

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"io/ioutil"
	"luck-admin/util"
	"net/http"
)

const PHONE_BILL_URL = "http://op.tianjurenhe.com/ofpay/mobile/onlineorder"

func JuHePhoneBill()  {
	singStr := "JH746601b8ed0b2cea45cb544643a4b310"+"86ef5c727da2ea3961d754fbbf3808fe"+"13425144866"+"1"+"12345789"

	h := md5.New()
	h.Write([]byte(singStr))
	cipherStr := h.Sum(nil)
	sign := hex.EncodeToString(cipherStr)


	url := PHONE_BILL_URL+"?phoneno=13425144866&cardnum=1&orderid=12345789&key=86ef5c727da2ea3961d754fbbf3808fe&sign="+sign
	httpClient := util.HttpClient{}
	err := httpClient.Get(url,nil, func(resp *http.Response) {
		body,readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			fmt.Println(readErr)
		}

		fmt.Println(string(body))
	})

	if err != nil {
		fmt.Println(err)
	}
}
