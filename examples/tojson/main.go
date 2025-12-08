package main

import "github.com/goforj/collection"

func main() {


	  c := collection.New([]int{1, 2, 3})
	  out, err := c.ToJSON()
	Example (error):
	  type Bad struct{}
	  func (Bad) MarshalJSON() ([]byte, error) {
	      return nil, fmt.Errorf("marshal failure")
	  }
	  c := collection.New([]Bad{{}})
	  out, err := c.ToJSON()
	Returns:
	  - string: the JSON-encoded representation of the collection
	  - error : nil on success, or the unwrapped marshalling error
	  // out: "[1,2,3]"
	  // err: nil



	  // out: ""
	  // err.Error(): "marshal failure"

}
