package views

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	var s []string
	te, _ := json.Marshal(s)
	fmt.Println(string(te))
}
