package main

import (
	"fmt"

	"github.com/dullgiulio/pingo"
)

func runPlugin(proto, path string) {
	p := pingo.NewPlugin(proto, path)
	p.Start()
	defer p.Stop()

	objs, err := p.Objects()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Objects: %s\n", objs)

	var msg map[string]string

	if err := p.Call("Plugin.UserDefineMap", "val", &msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", msg)
	}

	var res string
	if err := p.Call("Plugin.SayHello", "val", &res); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", res)
	}
}

func main() {
	fmt.Println("Running plugin that fails to register in time")
	runPlugin("tcp", "userDefineMap/userDefineMap")
	fmt.Println("Plugin terminated.")
}
