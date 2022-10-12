package main

import (
	"fmt"
	"github.com/go-turk/adiyaman/calctree"
)

func main() {
	Generate()
	//TestX()
}

func TestX() {
	root, err := calctree.LoadTree("after.txt")
	if err != nil {
		panic(err)
	}
	root.CheckHealth()

}

func Generate() {
	startValue := 35
	root := calctree.NewNode(startValue)
	root.SharewithChildren()
	fmt.Println("before force:")
	root.CheckHealth()
	forceCount := root.ForceToTree(100)
	root.SaveTree("ornekAgac.json")
	fmt.Println("Force count:", forceCount)
	repeairCount := root.Repair()
	root.SaveTree("ornekAgac-cozumu.json")
	fmt.Println("Repair count: ", repeairCount)
	if repeairCount == forceCount {
		fmt.Println("Repair is successfull")
	}
}
