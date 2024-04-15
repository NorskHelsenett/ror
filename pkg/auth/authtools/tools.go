package authtools

import (
	"fmt"
	"strings"
)

func SplitUserId(userId string) (string, string, error) {
	parts := strings.Split(userId, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid userId: %s", userId)
	}
	return parts[0], parts[1], nil
}
