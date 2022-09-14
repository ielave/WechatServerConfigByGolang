package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func main() {
	http.HandleFunc("/", GetHandle)
	fmt.Println("Running at port 54188 ...")

	err := http.ListenAndServe(":54188", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func GetHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	echostr := query.Get("echostr")
	token := "xxx" //你的token

	//将排序后的三个参数字符串拼接成一个字符串进行sha1加密
	slice := []string{timestamp, nonce, token}
	sort.Strings(slice)
	res := slice[0] + slice[1] + slice[2]
	h := sha1.New()
	io.WriteString(h, res)
	str := fmt.Sprintf("%x", h.Sum(nil))

	//将加密后的字符串与 signature 对比
	if str == signature {
		fmt.Fprintf(w, echostr)
	} else {
		fmt.Println("认证失败")
	}
}
