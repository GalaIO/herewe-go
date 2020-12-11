package main

import "net/http"

func main() {
	http.HandleFunc("/data/price", btcprice)
	if err := http.ListenAndServe("127.0.0.1:80", nil); err != nil {
		panic(err)
	}
}

func btcprice(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"USD": 10000}`))
}
