package utils

import (
	"regexp"
)

func EmojiFilter(s string) string {
	emojiRx := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	return emojiRx.ReplaceAllString(s, ``)
}

func EmojiCheck(s string) bool {
	emojiRx := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	return emojiRx.MatchString(s)
}
