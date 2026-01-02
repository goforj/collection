package collection

// FromMap materializes a map into a collection of key/value pairs.
// @group Maps
// @behavior immutable
// @chainable true
// @terminal false
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
//	c.Sort(func(a, b collection.Pair[string, int]) bool {
//		return a.Key < b.Key
//	})
//	collection.Dump(c.Items())
//
//	// #[]collection.Pair[string,int] [
//	//   0 => #collection.Pair[string,int] {
//	//     +Key   => "a" #string
//	//     +Value => 1 #int
//	//   }
//	//   1 => #collection.Pair[string,int] {
//	//     +Key   => "b" #string
//	//     +Value => 2 #int
//	//   }
//	//   2 => #collection.Pair[string,int] {
//	//     +Key   => "c" #string
//	//     +Value => 3 #int
//	//   }
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
//		Sort(func(a, b collection.Pair[string, Config]) bool {
//			return a.Key < b.Key
//		}).
//		Items()
//
//	collection.Dump(out)
//
//	// #[]collection.Pair[string,main.Config·1] [
//	//   0 => #collection.Pair[string,main.Config·1] {
//	//     +Key       => "router-1" #string
//	//     +Value     => #main.Config {
//	//       +Enabled => true #bool
//	//       +Timeout => 30 #int
//	//     }
//	//   }
//	//   1 => #collection.Pair[string,main.Config·1] {
//	//     +Key       => "router-3" #string
//	//     +Value     => #main.Config {
//	//       +Enabled => true #bool
//	//       +Timeout => 45 #int
//	//     }
//	//   }
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
//	// #map[string]int {
//	//   alice => 1 #int
//	//   bob => 2 #int
//	// }
func FromMap[K comparable, V any](m map[K]V) *Collection[Pair[K, V]] {
	items := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		items = append(items, Pair[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return New(items)
}
