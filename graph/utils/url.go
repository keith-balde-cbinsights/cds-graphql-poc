package utils

import (
	"fmt"
	"strings"

	hashids "github.com/speps/go-hashids/v2"
)

func GenerateProfileURL(orgID int) (string, error) {
	baseURL := "https://example.com"

	if orgID <= 0 {
		return "", fmt.Errorf("invalid organization ID: %d", orgID)
	}

	hd := hashids.NewData()
	hd.Salt = "CBI Profiles"
	hasher, err := hashids.NewWithData(hd)
	if err != nil {
		return "", fmt.Errorf("failed to initialize hasher: %w", err)
	}

	encoded, err := hasher.Encode([]int{orgID})
	if err != nil {
		return "", fmt.Errorf("failed to encode org ID: %w", err)
	}

	baseURL = strings.TrimRight(baseURL, "/")
	return fmt.Sprintf("%s/profiles/e/%s", baseURL, encoded), nil
}
