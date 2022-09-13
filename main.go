package main

import (
	"fmt"
	constants "pipeline/const"
	"pipeline/dto"
	"pipeline/service"
	"time"
)

func main() {
	var pipeline chan dto.Message = make(chan dto.Message)
	var result chan dto.Error = make(chan dto.Error)

	fmt.Println("testScenario1")

	go service.Pipeline(pipeline, result)
	go testScenario1(pipeline)

	for input := range result {
		fmt.Printf("Ip: %v, Error: %v\n", input.Ip, input.Error)
	}

	time.Sleep(time.Second)

	pipeline = make(chan dto.Message)
	result = make(chan dto.Error)

	fmt.Println("testScenario1Swap")

	go service.Pipeline(pipeline, result)
	go testScenario1Swap(pipeline)

	for input := range result {
		fmt.Printf("Ip: %v, Error: %v\n", input.Ip, input.Error)
	}

	time.Sleep(time.Second)

	pipeline = make(chan dto.Message)
	result = make(chan dto.Error)

	fmt.Println("testScenario2")

	go service.Pipeline(pipeline, result)
	go testScenario2(pipeline)

	for input := range result {
		fmt.Printf("Ip: %v, Error: %v\n", input.Ip, input.Error)
	}

	time.Sleep(time.Second)

	pipeline = make(chan dto.Message)
	result = make(chan dto.Error)

	fmt.Println("testScenario2Swap")

	go service.Pipeline(pipeline, result)
	go testScenario2Swap(pipeline)

	for input := range result {
		fmt.Printf("Ip: %v, Error: %v\n", input.Ip, input.Error)
	}
}

func testScenario2Swap(pipeline chan dto.Message) {
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 1,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 2,
		Status:    constants.AVAILABLE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 5,
		Status:    constants.AVAILABLE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 4,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 5,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 60,
		Status:    constants.AVAILABLE,
	}

	close(pipeline)
}

func testScenario2(pipeline chan dto.Message) {
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 1,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 2,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 3,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 4,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 5,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 60,
		Status:    constants.AVAILABLE,
	}

	close(pipeline)
}

func testScenario1Swap(pipeline chan dto.Message) {
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 1,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 10,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 16,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 12,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 20,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 100,
		Status:    constants.AVAILABLE,
	}

	close(pipeline)
}

func testScenario1(pipeline chan dto.Message) {
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 1,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 10,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 11,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 12,
		Status:    constants.AVAILABLE,
	}
	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 16,
		Status:    constants.GONE,
	}

	pipeline <- dto.Message{
		Ip:        "1",
		Timestamp: 100,
		Status:    constants.AVAILABLE,
	}

	close(pipeline)
}
