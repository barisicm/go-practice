package timer_manager

import (
	"fmt"
	"time"
)

type timerMetaData struct {
	RefreshPeriod time.Time
	// AgentID       uint64
	UUID string
	// Payload       interface{}
}

//connections map[UUID]terminated
// var connections map[string]bool
var connections = make(map[string]chan bool)

//data map[UUID]payload
// var data map[string]interface{}
var timers = make(map[string]chan time.Time)

func (tm timerMetaData) Start() {
	//Nejasno ?
	//Mutex unlock
	UUID := tm.UUID
	connections[tm.UUID] <- false
	data[tm.UUID] <- tm.Payload
	//Mutex lock

	go func(UUID string) {
		fmt.Printf("Emmiting message for agent ID: %s", tm)
	}(UUID)
}

func (tm timerMetaData) Stop() {
	connections[tm.UUID] <- true

}

func main() {

}
