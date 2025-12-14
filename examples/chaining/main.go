package main

import "github.com/goforj/collection"

type DeviceEvent struct {
	Device string
	Region string
	Errors int
}

func main() {
	events := []DeviceEvent{
		{Device: "router-1", Region: "us-east", Errors: 3},
		{Device: "router-2", Region: "us-east", Errors: 15},
		{Device: "router-3", Region: "us-west", Errors: 22},
		{Device: "router-4", Region: "us-west", Errors: 5},
		{Device: "router-5", Region: "us-east", Errors: 9},
		{Device: "router-6", Region: "us-west", Errors: 30},
	}

	// Fluent slice pipeline
	result := collection.
		New(events). // Construction

		// Ordering
		Shuffle().

		// Querying / filtering
		Filter(func(e DeviceEvent) bool {
			return e.Errors > 5
		}).

		// Ordering
		Sort(func(a, b DeviceEvent) bool {
			return a.Errors > b.Errors
		}).

		// Slicing
		Take(5).
		// Slicing (stop when predicate becomes true)
		TakeUntilFn(func(e DeviceEvent) bool {
			return e.Errors < 10
		}).
		SkipLast(1)

	collection.Dump(result.Items())

	// []main.DeviceEvent [
	//  0 => #main.DeviceEvent {
	//    +Device => "router-6" #string
	//    +Region => "us-west" #string
	//    +Errors => 30 #int
	//  }
	//  1 => #main.DeviceEvent {
	//    +Device => "router-3" #string
	//    +Region => "us-west" #string
	//    +Errors => 22 #int
	//  }
	// ]
}
