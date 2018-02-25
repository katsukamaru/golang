package main

type Client interface {
	list() []string
	cd()
}

type ClientStruct struct {
	name string
}

func (c *ClientStruct) list() []string {
	return []string{"aaa", "bbb"}
}

func (c *ClientStruct) cd() {

}
