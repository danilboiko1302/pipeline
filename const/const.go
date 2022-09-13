package constants

import "pipeline/dto"

const (
	AVAILABLE dto.Status = iota + 1
	GONE
)

const (
	TIME_FOR_LOST   int64 = 60 //seconds
	TIME_FOR_FLAP   int64 = 60 //seconds
	AMOUNT_FOR_FLAP int64 = 3
	MAX_DELAY       int64 = 30 //seconds
)
