package utils

import (
	"testing"
)

func TestGetResourceName(t *testing.T) {
	contents := ReadYaml([]byte(yamlExample))
	resources := Categorize(contents)
	rscName := resources.StatefulSet[0].Metadata.Name
	if rscName != "openshift-rollouter" {
		t.Error("Failed")
	}
}
