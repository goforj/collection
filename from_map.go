package collection

// FromMap materializes a map into a collection of key/value pairs.
// @group Maps
// @behavior immutable
// @fluent true
//
// The iteration order of the resulting collection is unspecified,
// matching Go's map iteration semantics.
//
// This function does not mutate the input map.
//
// Example: basic usage
//
//	m := map[string]int{
//		"a": 1,
//		"b": 2,
//		"c": 3,
//	}
//
//	c := collection.FromMap(m)
//	collection.Dump(c.Items())
//
//	// #[]collection.Pair[string,int] [
//	//   0 => {Key:"a" Value:1}
//	//   1 => {Key:"b" Value:2}
//	//   2 => {Key:"c" Value:3}
//	// ]
//
// Example: filtering map entries
//
//	type Config struct {
//		Enabled bool
//		Timeout int
//	}
//
//	configs := map[string]Config{
//		"router-1": {Enabled: true,  Timeout: 30},
//		"router-2": {Enabled: false, Timeout: 10},
//		"router-3": {Enabled: true,  Timeout: 45},
//	}
//
//	out := collection.
//		FromMap(configs).
//		Filter(func(p collection.Pair[string, Config]) bool {
//			return p.Value.Enabled
//		}).
//		Items()
//
//	collection.Dump(out)
//
//	// #[]collection.Pair[string,collection.Config] [
//	//   0 => {Key:"router-1" Value:{Enabled:true Timeout:30}}
//	//   1 => {Key:"router-3" Value:{Enabled:true Timeout:45}}
//	// ]
//
// Example: map → collection → map
//
//	users := map[string]int{
//		"alice": 1,
//		"bob":   2,
//	}
//
//	c2 := collection.FromMap(users)
//	out2 := collection.ToMapKV(c2)
//
//	collection.Dump(out2)
//
//	// #map[string]int [
//	//   "alice" => 1
//	//   "bob"   => 2
//	// ]
func FromMap[K comparable, V any](m map[K]V) *Collection[Pair[K, V]] {
	items := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		items = append(items, Pair[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return &Collection[Pair[K, V]]{items: items}
}
