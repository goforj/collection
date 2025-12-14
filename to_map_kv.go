package collection

// ToMapKV converts a collection of key/value pairs into a map.
// @group Maps
//
// If multiple pairs contain the same key, the last value wins.
//
// This operation allocates a map sized to the collection length.
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
//	out := collection.ToMapKV(c)
//
//	collection.Dump(out)
//
//	// #map[string]int [
//	//   "a" => 1
//	//   "b" => 2
//	//   "c" => 3
//	// ]
//
// Example: filtering before conversion
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
//	c2 := collection.
//		FromMap(configs).
//		Filter(func(p collection.Pair[string, Config]) bool {
//			return p.Value.Enabled
//		})
//
//	out2 := collection.ToMapKV(c2)
//
//	collection.Dump(out2)
//
//	// #map[string]collection.Config [
//	//   "router-1" => {Enabled:true Timeout:30}
//	//   "router-3" => {Enabled:true Timeout:45}
//	// ]
func ToMapKV[K comparable, V any](c *Collection[Pair[K, V]]) map[K]V {
	out := make(map[K]V, len(c.items))
	for _, p := range c.items {
		out[p.Key] = p.Value
	}
	return out
}
