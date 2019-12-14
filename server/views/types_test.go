package views

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	// var s []string
	var g [0]int
	te, _ := json.Marshal(g)
	fmt.Println(string(te))
}
