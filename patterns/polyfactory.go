package main

import "fmt"

type Shape interface {
	Draw()
}

type PolyFactory interface {
	Create() Shape
}

var shapeFactoryMap map[string]PolyFactory = make(map[string]PolyFactory, 0)

func RegisterShapeFactory(name string, factory PolyFactory) error {
	shapeFactoryMap[name] = factory
	return nil
}

func ShapeFactoryV2(name string) Shape {
	factory := shapeFactoryMap[name]
	return factory.Create()
}

type Cube struct {
}

func (c *Cube) Draw() {
	fmt.Println("it's cube~~")
}

type CubeCreater struct {
}

func (c *CubeCreater) Create() Shape {
	return &Cube{}
}

func main() {
	RegisterShapeFactory("cube", &CubeCreater{})
	ShapeFactoryV2("cube").Draw()
}