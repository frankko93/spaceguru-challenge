package utils

import "fmt"

// Tags ...
type Tags map[string]string

// BuildTags ...
func BuildTags(tags Tags) []string {
	taglist := []string{}

	for k, v := range tags {
		taglist = append(taglist, fmt.Sprintf("%s:%s", k, v))
	}

	return taglist
}

// AppendTags ...
func AppendTags(builtTags *[]string, newTags Tags) {
	newTaglist := BuildTags(newTags)

	for _, tag := range newTaglist {
		*builtTags = append(*builtTags, tag)
	}
}
