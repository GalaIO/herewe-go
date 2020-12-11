package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	request.Header.Set("Authorization", "token "+"4e897044cf20bcbc47d2434e9363d662f98a7534")
	panicErr(err)
	response, err := http.DefaultClient.Do(request)
	panicErr(err)
	bytes, err := ioutil.ReadAll(response.Body)
	panicErr(err)
	fmt.Println(string(bytes))
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
