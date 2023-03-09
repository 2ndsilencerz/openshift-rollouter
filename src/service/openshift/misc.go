package openshift

import (
	"strings"
)

func ReadYaml(fileContent []byte) []string {
	allContent := string(fileContent)
	var contents []string
	contents = strings.Split(allContent, "---")
	strippedContents := *new([]string)
	for _, v := range contents {
		if len(strings.TrimSpace(v)) > 1 && strings.TrimSpace(v)[0] != '#' {
			strippedContents = append(strippedContents, v)
		}
	}
	return strippedContents
}
