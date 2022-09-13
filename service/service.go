package service

import (
	constants "pipeline/const"
	"pipeline/dto"
)

func Pipeline(pipeline chan dto.Message, result chan dto.Error) {
	var lastMsg map[string]dto.Message = make(map[string]dto.Message)
	var flap map[string]*dto.Flap = make(map[string]*dto.Flap)
	for input := range pipeline {
		message, ok := lastMsg[input.Ip]

		if !ok {
			lastMsg[input.Ip] = input
			continue
		}
		if message.Timestamp-input.Timestamp < constants.MAX_DELAY && message.Timestamp-input.Timestamp > 0 { //message.Timestamp > input.Timestamp no need to check this, message.Timestamp-input.Timestamp < 0 then
			//swap
			var temp dto.Message = message
			message = input
			input = temp
		}
		lastMsg[input.Ip] = input

		if message.Status != constants.GONE ||
			input.Status != constants.AVAILABLE {
			continue
		}

		if input.Timestamp-message.Timestamp > constants.TIME_FOR_LOST {
			result <- dto.Error{
				Ip:    input.Ip,
				Error: "Lost",
			}
			continue
		}

		currentFlap, ok := flap[input.Ip]

		if !ok { //first flap
			flap[input.Ip] = &dto.Flap{
				Start:  message.Timestamp, //time first GONE
				Amount: 1,
			}
			continue
		}

		if input.Timestamp-currentFlap.Start > constants.TIME_FOR_FLAP {
			flap[input.Ip].Amount = 0
			continue
		}

		if currentFlap.Amount == constants.AMOUNT_FOR_FLAP-1 {
			result <- dto.Error{
				Ip:    input.Ip,
				Error: "Flap",
			}
			flap[input.Ip].Amount = 0
			continue
		}

		flap[input.Ip].Amount++
	}

	close(result)
}
