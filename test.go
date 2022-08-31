package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func main() {
	resp, err := http.Get("http://47.100.21.147:5004/WeBASE-Sign/user/newUser?" +
		"appId=invoice&signUserId=20177830226&encryptType=0&returnPrivateKey=false")
	if err != nil {
		log.Fatalln(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	r := reflect.TypeOf(respBody)
	fmt.Println(r)
	fmt.Println(string(respBody))
}
