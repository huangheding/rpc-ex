package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println(grpc.BackoffConfig{0})

}
