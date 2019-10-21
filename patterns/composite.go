package main

import "fmt"

type Category interface {
	Name() string
	SubCates() []Category
}

type Book struct {
	subCate []Category
}

func (b *Book) Name() string {
	return "Book"
}

func (b *Book) SubCates() []Category {
	return b.subCate
}

type Novel struct {
	subCate []Category
}

func (b *Novel) Name() string {
	return "Novel"
}

func (b *Novel) SubCates() []Category {
	return nil
}

type IT struct {
	subCate []Category
}

func (b *IT) Name() string {
	return "IT"
}

func (b *IT) SubCates() []Category {
	return []Category{&C{}}
}

type C struct {
	subCate []Category
}

func (b *C) Name() string {
	return "C"
}

func (b *C) SubCates() []Category {
	return nil
}

func NewBookCate(cates ... Category) Category {
	book := Book{
		subCate: make([]Category, len(cates)),
	}
	for i, cate := range cates {
		book.subCate[i] = cate
	}
	return &book
}

func main() {
	b := NewBookCate(&Novel{}, &IT{})
	printCate("", b)
}

func printCate(prefix string, b Category) {
	fmt.Println(prefix + b.Name() + "/")
	prefix = "\t" + prefix
	for _, cate := range b.SubCates() {
		if len(cate.SubCates()) > 0 {
			printCate(prefix, cate)
			continue
		}
		fmt.Println(prefix + cate.Name())
	}
}


