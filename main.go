package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var debugOn = false
var debugOut []string

type MyEvent struct {
	Nums  []int64 `json:"nums,omitempty"`
	Debug bool    `json:"debugon,omitempty"`
}

type MyResponse struct {
	Message     string   `json:"result"`
	DebugOutput []string `json:"debugOutput,omitempty"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {

	debugOn = event.Debug
	var resp string
	debugOut = []string{}

	//check debug
	if debugOn {
		debugOut = []string{"INFO: Start debug", string(strings.Join(strings.Fields(fmt.Sprint(event.Nums)), ","))}
	}

	return MyResponse{Message: string(strings.Join(strings.Fields(fmt.Sprint(productNN(event.Nums))), ",")), DebugOutput: debugOut}, nil

	//return with or without debug info
	if debugOn {
		return MyResponse{Message: resp, DebugOutput: debugOut}, nil
	}
	return MyResponse{Message: resp}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
	// fmt.Println("Hello, World!")

	//test multiplication
	// N := 70000
	// mySlice := make([]int64, N)
	// for i := 0; i < N; i++ {
	// 	mySlice[i] = 1
	// }
	// mySlice := []int64{1, 2, 3, 4}
	// product2N(mySlice)
	// productNN(mySlice)
	// productN(mySlice)

	// //test addition
	// // var M int64 = 125000
	// // mySlice = make([]int64, M)
	// // for i := int64(0); i < M; i++ {
	// // 	mySlice[i] = M
	// // }
	// addition2N(mySlice)
	// additionNN(mySlice)

}

func productNN(array []int64) []int64 {

	result := make([]int64, len(array))
	var prod int64 = 1

	for i, _ := range result {
		for j, itemj := range array {
			if i != j {
				prod = prod * itemj
			}
		}
		result[i] = prod
		prod = 1
	}
	return result
}

func product2N(array []int64) {
	start := time.Now()
	result := make([]int64, len(array))
	var prod int64 = 1

	for _, item := range array {
		prod = prod * item
	}

	for i, item := range array {
		result[i] = prod / item
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("PRODUCT 2N DONE:", elapsed)
}

func additionNN(array []int64) {
	start := time.Now()

	result := make([]int64, len(array))
	var sum int64 = 0

	for i, _ := range result {
		for j, itemj := range array {
			if i != j {
				sum = sum + itemj
			}
		}
		result[i] = sum
		sum = 0
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("ADDITION NN DONE:", elapsed)
}

func addition2N(array []int64) {
	start := time.Now()
	result := make([]int64, len(array))
	var sum int64 = 0

	for _, item := range array {
		sum = sum + item
	}

	for i, item := range array {
		result[i] = sum - item
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("ADDITION 2N DONE:", elapsed)
}

func productN(array []int64) {
	start := time.Now()
	//result := make([]int64, len(array))
	left := make([]int64, len(array))
	right := make([]int64, len(array))
	fmt.Println(array)

	var prod int64 = 1
	left[0] = array[0]

	for i, _ := range array {
		left[i] = prod
		if i+1 < len(array) {
			prod = prod * array[i+1]
		}
	}

	prod = 1

	for i, _ := range array {
		j := len(array) - 1 - i
		right[j] = prod
		if j > 0 {
			fmt.Println(j, "::", prod, "*", array[j])
			prod = prod * array[j]
		}

	}

	fmt.Println(left, right)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Product no divistion N DONE:", elapsed)
}
