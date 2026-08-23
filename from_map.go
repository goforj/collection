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
//		return a.First < b.First
//	})
//	collection.Dump(c)
//	// #[]collection.Pair[string,int] [
//	//   0 => #collection.Pair[string,int] {
//	//     +First  => "a" #string
//	//     +Second => 1 #int
//	//   }
//	//   1 => #collection.Pair[string,int] {
//	//     +First  => "b" #string
//	//     +Second => 2 #int
//	//   }
//	//   2 => #collection.Pair[string,int] {
//	//     +First  => "c" #string
//	//     +Second => 3 #int
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
//		"router-1": {Enabled: true, Timeout: 30},
//		"router-2": {Enabled: false, Timeout: 10},
//		"router-3": {Enabled: true, Timeout: 45},
//	}
//
//	out := collection.
//		FromMap(configs).
//		Filter(func(p collection.Pair[string, Config]) bool {
//			return p.Second.Enabled
//		}).
//		Sort(func(a, b collection.Pair[string, Config]) bool {
//			return a.First < b.First
//		})
//
//	collection.Dump(out)
//	// #[]collection.Pair[string,main.Config·1] [
//	//   0 => #collection.Pair[string,main.Config·1] {
//	//     +First     => "router-1" #string
//	//     +Second    => #main.Config {
//	//       +Enabled => true #bool
//	//       +Timeout => 30 #int
//	//     }
//	//   }
//	//   1 => #collection.Pair[string,main.Config·1] {
//	//     +First     => "router-3" #string
//	//     +Second    => #main.Config {
//	//       +Enabled => true #bool
//	//       +Timeout => 45 #int
//	//     }
//	//   }
//	// ]
func FromMap[K comparable, V any](m map[K]V) Slice[Pair[K, V]] {
	items := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		items = append(items, Pair[K, V]{
			First:  k,
			Second: v,
		})
	}
	return New(items)
}
