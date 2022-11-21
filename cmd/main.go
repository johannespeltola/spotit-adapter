package main

import (
	"entsoe/src/config"
	routine "entsoe/src/event"
	"entsoe/src/global"
)

func init() {
	config.ReadConfigJSON()
}

func main() {
	env := global.Init()
	go routine.Routine(30, env)
	select {} // this will cause the program to run forever
}
