package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

func HandleErr(err error, when string)  {
	if err!= nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	fetUrl := "https://www.haomagujia.com"
	resp, err := http.Get(fetUrl)
	HandleErr(err, "fetch phone")
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, "get response")
	html := string(bytes)
	fmt.Println(html)
	// 匹配 模板
	/**
	<div class="jg666">[求购]<a href="/15555555545" title="手机号码15555555545估价评估_值多少钱_归属地查询_测吉凶_数字含义_求购转让信息" class="lj44">15555555545</a>由 张新 10000元求购,QQ：6461991,联系电话：15555555545,缺钱</div>
	pattern <div class="jg666">[求购]<a href="/([1-9].*?)" title="(.*?)" class="lj44">(.*?)</a>(.*?)</div>
	 */
	pattern := `<div class="jg666">[求购]<a href="/([1-9].*?)" title="(.*?)" class="lj44">(.*?)</a>(.*?)</div>`;
	re, e := regexp.Compile(pattern)
	HandleErr(e, "complie")
	submatch := re.FindAllStringSubmatch(html, -1)
	fmt.Println(submatch)
	for _, match := range submatch {
		fmt.Printf("phton number: %s, msg: %s\r\n", match[0], match[3])
	}
}