package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var oauthHtml = `<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>OAuth</title>
</head>
<body>
    <a href="https://github.com/login/oauth/authorize?client_id=30a277e5a1f01c1735fe&redirect_uri=http://localhost:9100/oauth/callback/github">跳转github登录</a>
</body>
</html>`

func main() {
	routers := http.NewServeMux()
	routers.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(oauthHtml))
	})
	routers.HandleFunc("/oauth/callback/github", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		err := request.ParseForm()
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		code := request.Form.Get("code")
		url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			"30a277e5a1f01c1735fe", "cb02ca3d87dc840fea03d9f15bbdb2bf609e5888", code)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		panicErr(err)
		req.Header.Set("Accept", "application/json")
		resp, err := http.DefaultClient.Do(req)
		panicErr(err)
		bytes, err := ioutil.ReadAll(resp.Body)
		panicErr(err)
		fmt.Println(string(bytes))
		tokenMap := new(map[string]interface{})
		err = json.Unmarshal(bytes, tokenMap)
		panicErr(err)
		fmt.Println(*tokenMap)
		fmt.Println((*tokenMap)["data"])
	})
	http.ListenAndServe(":9100", routers)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}