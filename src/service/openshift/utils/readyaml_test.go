package utils

import (
	"testing"
)

func TestReadYaml(t *testing.T) {
	contents := ReadYaml([]byte(yamlExample))
	if len(contents) <= 0 {
		t.Error("Failed")
		return
	}
	//for k, v := range contents {
	//	index := strconv.Itoa(k)
	//	log.Println("index " + index)
	//	log.Println(v)
	//}
	//log.Println()
}
