//go:build ignore
// +build ignore

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
	}

	// Fluent slice pipeline
	collection.
		New(events). // Construction
		Shuffle(). // Ordering
		Filter(func(e DeviceEvent) bool {
			return e.Errors > 5
		}). // Slicing
		Sort(func(a, b DeviceEvent) bool {
			return a.Errors > b.Errors
		}). // Ordering
		Take(5). // Slicing
		TakeUntilFn(func(e DeviceEvent) bool {
			return e.Errors < 10
		}). // Slicing (stop when predicate becomes true)
		SkipLast(1). // Slicing
		Dump() // Debugging

	// []main.DeviceEvent [
	//  0 => #main.DeviceEvent {
	//    +Device => "router-3" #string
	//    +Region => "us-west" #string
	//    +Errors => 22 #int
	//  }
	// ]
}
