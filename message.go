package main

var bus = make(chan Message, 1)

type Message struct {
}
