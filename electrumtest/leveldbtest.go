package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"encoding/json"
)

func main() {
	db, err := leveldb.OpenFile("/Users/galaio/PycharmProjects/electrumx/electrumxDir/db/utxo", nil)
	if err!= nil {
		panic(err)
	}
	defer db.Close()

	value, err := db.Get([]byte("state"), nil)
	if err != nil {
		panic(value)
	}
	message := new(map[string]json.RawMessage)
	if err := json.Unmarshal(value, message); err != nil {
		panic(err)
	}
	fmt.Println(value)
	fmt.Println(message)

	iterator := db.NewIterator(nil, nil)
	for iterator.Next() {
		fmt.Printf("key: %s, val: %s", iterator.Key(), iterator.Value())
	}
	iterator.Release()
}
