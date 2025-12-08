package main

import "github.com/goforj/collection"

func main() {


	  c := collection.New([]string{"a", "b"})
	  out, err := c.ToPrettyJSON()
	Example (error):
	  type Bad struct{}
	  func (Bad) MarshalJSON() ([]byte, error) {
	      return nil, fmt.Errorf("marshal failure")
	  }
	  c := collection.New([]Bad{{}})
	  out, err := c.ToPrettyJSON()
	Returns:
	  - string: the pretty-printed JSON representation
	  - error : nil on success, or the unwrapped marshalling error
	  // out:
	  // [
	  //   "a",
	  //   "b"
	  // ]
	  // err: nil



	  // out: ""
	  // err.Error(): "marshal failure"

}
