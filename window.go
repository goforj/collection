package collection

// Window returns overlapping (or stepped) windows of the collection.
// Each window is a slice of length size; iteration advances by step (default 1 if step <= 0).
// Windows that are shorter than size are omitted.
// @group Slicing
// @behavior allocates
// @chainable true
// @terminal false
//
// NOTE: windows share the backing array with the source collection.
//
// Example: integers - step 1
//
//	nums := collection.New([]int{1, 2, 3, 4, 5})
//	win := collection.Window(nums, 3, 1)
//	collection.Dump(win.Items())
//	// #[][]int [
//	//   0 => #[]int [
//	//     0 => 1 #int
//	//     1 => 2 #int
//	//     2 => 3 #int
//	//   ]
//	//   1 => #[]int [
//	//     0 => 2 #int
//	//     1 => 3 #int
//	//     2 => 4 #int
//	//   ]
//	//   2 => #[]int [
//	//     0 => 3 #int
//	//     1 => 4 #int
//	//     2 => 5 #int
//	//   ]
//	// ]
//
// Example: strings - step 2
//
//	words := collection.New([]string{"a", "b", "c", "d", "e"})
//	win2 := collection.Window(words, 2, 2)
//	collection.Dump(win2.Items())
//	// #[][]string [
//	//   0 => #[]string [
//	//     0 => "a" #string
//	//     1 => "b" #string
//	//   ]
//	//   1 => #[]string [
//	//     0 => "c" #string
//	//     1 => "d" #string
//	//   ]
//	// ]
//
// Example: structs
//
//	type Point struct {
//		X int
//		Y int
//	}
//
//	points := collection.New([]Point{
//		{X: 0, Y: 0},
//		{X: 1, Y: 1},
//		{X: 2, Y: 4},
//		{X: 3, Y: 9},
//	})
//
//	win3 := collection.Window(points, 2, 1)
//	collection.Dump(win3.Items())
//	// #[][]main.Point [
//	//   0 => #[]main.Point [
//	//     0 => #main.Point {
//	//       +X => 0 #int
//	//       +Y => 0 #int
//	//     }
//	//     1 => #main.Point {
//	//       +X => 1 #int
//	//       +Y => 1 #int
//	//     }
//	//   ]
//	//   1 => #[]main.Point [
//	//     0 => #main.Point {
//	//       +X => 1 #int
//	//       +Y => 1 #int
//	//     }
//	//     1 => #main.Point {
//	//       +X => 2 #int
//	//       +Y => 4 #int
//	//     }
//	//   ]
//	//   2 => #[]main.Point [
//	//     0 => #main.Point {
//	//       +X => 2 #int
//	//       +Y => 4 #int
//	//     }
//	//     1 => #main.Point {
//	//       +X => 3 #int
//	//       +Y => 9 #int
//	//     }
//	//   ]
//	// ]
func Window[T any](c *Collection[T], size int, step int) *Collection[[]T] {
	if size <= 0 {
		return New([][]T(nil))
	}

	if step <= 0 {
		step = 1
	}

	n := len(c.items)
	if n < size {
		return New([][]T(nil))
	}

	// Compute number of windows.
	count := 1 + (n-size)/step
	out := make([][]T, 0, count)

	for i := 0; i+size <= n; i += step {
		out = append(out, c.items[i:i+size])
	}

	return New(out)
}
