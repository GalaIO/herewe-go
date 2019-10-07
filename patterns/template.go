package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type HttpStringHandler func(*string) error

type GetHttpString struct {
	handler HttpStringHandler
}

func (g *GetHttpString) Do(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(bytes)
	return g.handler(&s)
}

func main() {
	httpString := GetHttpString{
		handler: func(s *string) error {
			fmt.Println(*s)
			return nil
		},
	}
	httpString.Do("http://www.baidu.com")
}
