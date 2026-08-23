package collection

// Window returns overlapping (or stepped) windows of the collection.
// Each window is a slice of length size; iteration advances by step (default 1 if step <= 0).
// Windows that are shorter than size are omitted.
// @group Slicing
// @behavior readonly
// @chainable false
// @terminal true
//
// NOTE: windows share the backing array with the source collection.
//
// Example: integers - step 1
//
//	collection.Dump(collection.New([]int{1, 2, 3, 4, 5}).Window(3, 1))
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
//	collection.Dump(collection.New([]string{"a", "b", "c", "d", "e"}).Window(2, 2))
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
//	win3 := points.Window(2, 1)
//	collection.Dump(win3)
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
func (c Slice[T]) Window(size int, step int) [][]T {
	if size <= 0 {
		return nil
	}

	if step <= 0 {
		step = 1
	}

	n := len(c)
	if n < size {
		return nil
	}

	lastStart := n - size
	count := 1 + lastStart/step
	out := make([][]T, 0, count)

	for i := 0; i <= lastStart; {
		out = append(out, c[i:i+size:i+size])
		if step > lastStart-i {
			break
		}
		i += step
	}

	return out
}
