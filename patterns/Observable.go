package main

import (
	"sync"
	"errors"
	"fmt"
)

type Observable struct {
	sync.Mutex
	observers []Listener
}

type Listener func(obj interface{})

func (o *Observable) AddObserver(ob Listener) error {
	if ob == nil {
		return errors.New("Observer cannot be null")
	}
	o.Lock()
	defer o.Unlock()
	if o.observers == nil {
		o.observers = []Listener{}
	}
	o.observers = append(o.observers, ob)
	return nil
}

func (o *Observable) NotifyObserver(obj interface{}) error {
	o.Lock()
	defer o.Unlock()
	for _, ob := range o.observers {
		ob(obj)
	}
	return nil
}


type Flower struct {
	Observable
}

func (f Flower) open() {
	f.NotifyObserver("open flower")
}

func (f Flower) close() {
	f.NotifyObserver("close flower")
}

func main() {
	flower := Flower{}
	flower.AddObserver(func(obj interface{}) {
		fmt.Printf("bee wait...%s\r\n", obj)
	})
	flower.AddObserver(func(obj interface{}) {
		fmt.Printf("bird wait...%s\r\n", obj)
	})
	flower.open()
	flower.close()
}
