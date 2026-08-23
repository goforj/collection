//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/goforj/collection/v3"
)

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
		{Device: "router-4", Region: "us-west", Errors: 9},
		{Device: "router-5", Region: "eu-west", Errors: 7},
	}

	// Clone creates a top-level ownership boundary before the mutable stages.
	collection.
		New(events).                                                      // Construction
		Clone().                                                          // Construction
		Retain(func(e DeviceEvent) bool { return e.Errors > 5 }).         // Mutation
		Sort(func(a, b DeviceEvent) bool { return a.Errors > b.Errors }). // Ordering
		Take(3).                                                          // Slicing
		TakeUntil(func(e DeviceEvent) bool { return e.Errors <= 9 }).     // Slicing (stop when predicate becomes true)
		Reverse().                                                        // Ordering
		Dump()                                                            // Debugging

	// #[]main.DeviceEvent [
	//  0 => #main.DeviceEvent {
	//    +Device => "router-2" #string
	//    +Region => "us-east" #string
	//    +Errors => 15 #int
	//  }
	//  1 => #main.DeviceEvent {
	//    +Device => "router-3" #string
	//    +Region => "us-west" #string
	//    +Errors => 22 #int
	//  }
	// ]

	regionsByPrefix := collection.New(events).
		Map(func(event DeviceEvent) string { return event.Region }).
		UniqueBy(func(region string) string { return region }).
		GroupBy(func(region string) string { return region[:2] })
	fmt.Println(len(regionsByPrefix))
	// 2
}
