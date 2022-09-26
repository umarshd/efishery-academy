package main

import (
	"fmt"
	"tugas-week-1-2/api"
)

func main() {
	var cache api.Cache
	var aggregate api.Aggregation
	var err error

	for i := 1; i <= 6; i++ {
		fmt.Println("Attempt :", i)
		aggregate, err = cache.Aggregate()
		if err != nil {
			panic(err)
		}
	}

	_ = aggregate
}
