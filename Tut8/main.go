package main

import (
	"fmt"
	"sync"
	"time"
)

// go routine
var rm = sync.RWMutex{}
var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var v0data = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}
func main(){
	t0 := time.Now()
	for i := 0; i < len(v0data); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution is: %v", time.Since(t0))
	fmt.Printf("\nResult is: %v", result)
}

func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(v0data[i])
	log()
	wg.Done()
}

func save(data string){
	m.Lock()
	result = append(result, data)
	m.Unlock()
}

func log(){
	rm.RLock()
	fmt.Printf("\nThe result from the database is: %v", result)
	rm.RUnlock()
}
