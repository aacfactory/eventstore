package events

import (
	"context"
	"github.com/aacfactory/errors"
)

// ConsumeArgument
// @title Consume events argument
// @description Consume events argument
type ConsumeArgument struct {
	// ConsumerId
	// @title Consumer id
	// @description Consumer id
	ConsumerId uint64 `json:"consumerId" validate:"required" message:"consumerId is invalid"`
	// AggregateName
	// @title Aggregate name
	// @description Which aggregate name to be consumed
	AggregateName string `json:"aggregateName" validate:"required" message:"aggregateName is invalid"`
	// Offset
	// @title Offset
	// @description Offset of aggregate domain events, 0 is next of last consumed,
	Offset uint64 `json:"offset" validate:"required" message:"offset is invalid"`
	// Mode
	// @title Consume mode
	// @enum PULL,PUSH
	// @description Consume mode
	Mode string `json:"mode" validate:"oneof=PULL PUSH" message:"mode is invalid"`
}

// consume
// @fn consume
// @validate true
// @authorization false
// @permission false
// @internal false
// @title Consume events
// @description >>>
// Consume aggregate domain events
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | events_consume_failed    | 500     | consume events failed         |
// <<<
func consume(ctx context.Context, argument ConsumeArgument) (result []*Event, err errors.CodeError) {

	return
}
