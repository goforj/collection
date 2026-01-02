package collection

import "testing"

var (
	allocCollection *Collection[int]
	allocNumeric    *NumericCollection[int]

	allocBool  bool
	allocInt   int
	allocValue int
)

func allocPred(v int) bool {
	return v%2 == 0
}

func allocReduce(acc, v int) int {
	return acc + v
}

func allocAny() {
	allocBool = allocCollection.Any(allocPred)
}

func allocAll() {
	allocBool = allocCollection.All(allocPred)
}

func allocNone() {
	allocBool = allocCollection.None(allocPred)
}

func allocFirst() {
	var ok bool
	allocValue, ok = allocCollection.First()
	allocBool = ok
}

func allocLast() {
	var ok bool
	allocValue, ok = allocCollection.Last()
	allocBool = ok
}

func allocFirstWhere() {
	var ok bool
	allocValue, ok = allocCollection.FirstWhere(allocPred)
	allocBool = ok
}

func allocIndexWhere() {
	var ok bool
	allocInt, ok = allocCollection.IndexWhere(allocPred)
	allocBool = ok
}

func allocReduceSum() {
	allocInt = allocCollection.Reduce(0, allocReduce)
}

func allocSum() {
	allocInt = allocNumeric.Sum()
}

func allocMin() {
	var ok bool
	allocInt, ok = allocNumeric.Min()
	allocBool = ok
}

func allocMax() {
	var ok bool
	allocInt, ok = allocNumeric.Max()
	allocBool = ok
}

func allocContains() {
	allocBool = Contains(allocCollection, 4)
}

func allocIsEmpty() {
	allocBool = allocCollection.IsEmpty()
}

func TestTerminalOps_ZeroAlloc(t *testing.T) {
	allocCollection = New([]int{1, 2, 3, 4, 5, 6})
	allocNumeric = NewNumeric([]int{1, 2, 3, 4, 5, 6})

	cases := []struct {
		name string
		fn   func()
	}{
		{name: "Any", fn: allocAny},
		{name: "All", fn: allocAll},
		{name: "None", fn: allocNone},
		{name: "First", fn: allocFirst},
		{name: "Last", fn: allocLast},
		{name: "FirstWhere", fn: allocFirstWhere},
		{name: "IndexWhere", fn: allocIndexWhere},
		{name: "Reduce", fn: allocReduceSum},
		{name: "Sum", fn: allocSum},
		{name: "Min", fn: allocMin},
		{name: "Max", fn: allocMax},
		{name: "Contains", fn: allocContains},
		{name: "IsEmpty", fn: allocIsEmpty},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			allocs := testing.AllocsPerRun(1000, tc.fn)
			if allocs != 0 {
				t.Fatalf("%s should allocate 0 times; got %v", tc.name, allocs)
			}
		})
	}
}
