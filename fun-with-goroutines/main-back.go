package main

import (
	"fmt"
	"sync"
)

type timerPayload struct {
	terminated chan bool
	token      chan string
}

var connections map[uint64]chan timerPayload
var wg sync.WaitGroup

func somefunc() {

	token := "Some token"
	agent := uint64(25)

	connections = make(map[uint64]chan timerPayload)

	tPayload := timerPayload{
		terminated: make(chan bool, 2),
		token:      make(chan string, 2),
	}
	tPayload.terminated <- false
	tPayload.token <- token

	payloadCh := make(chan timerPayload, 2)
	payloadCh <- tPayload

	connections[agent] = make(chan timerPayload)
	connections[agent] = payloadCh

	wg.Add(1)
	go func(agentId uint64) {
		tPayload := <-connections[agentId]
		fmt.Println(<-tPayload.token)
		wg.Done()
	}(agent)
	wg.Wait()
}

func main2() {
	agentId := 5
	agentId2 := 6

	token := "test"
	token2 := "test2"

	conn := make(map[int]chan string)

	conn[agentId] =token <-
}
