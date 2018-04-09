package utils

import "github.com/microcosm-cc/bluemonday"

func AvoidXSS(theHTML string) string {
	return bluemonday.UGCPolicy().Sanitize(theHTML)
}
