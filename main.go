package main

import (
	"log"
	"time"
)

func main() {
	arr := sliceFilledWithString(10000000, "nemo")
	findNemo(arr)

}

func findNemo(arr []string) {
	t0 := time.Now().UnixNano() / int64(time.Millisecond)
	for i := range arr {
		if arr[i] == "nemo" {
			log.Println("Find nemo")
		}
	}
	t1 := time.Now().UnixNano() / int64(time.Millisecond)
	d := t1 - t0
	log.Printf("Call to find nemo took %d millisecond", d)
}

func sliceFilledWithString(size int, str string) []string {
	data := make([]string, size)
	for i := 0; i < size; i++ {
		data[i] = str
	}
	return data
}
