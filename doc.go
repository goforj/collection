// Package collection provides fluent, explicit pipelines over slices.
//
// New borrows slices by default, while pure transformations return independent
// results and view operations document their shared storage. Use Clone to make
// ownership independent explicitly.
package collection
