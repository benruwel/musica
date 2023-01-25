package utils

import (
	"strings"

	"github.com/google/uuid"
)

// given a domain prefix such as 'album'
// it should generate id of 'album_efaa0ad'
func GenerateID(prefix string) string {
	id := uuid.New()
	firstSect := strings.Split(id.String(), "-")[0]
	return prefix + "_" + firstSect
}
