package cache

import (
	"net/url"
	"strings"

	"github.com/dwisarut/dealmaxxingCLI/internal/model"
)

func InitIDCache() map[string][]model.GameIdentifier {
	idCache := make(map[string][]model.GameIdentifier)
	return idCache
}

func InitGetCache() map[string]model.GetGameID {
	getCache := make(map[string]model.GetGameID)
	return getCache
}

func IdentifierCaching(results []model.GameIdentifier, cacheKey string, idCache map[string][]model.GameIdentifier) map[string][]model.GameIdentifier {
	idCache[cacheKey] = results
	return idCache
}

func MakeCachingKey(input string) string {
	decoded, _ := url.QueryUnescape(input)
	normalized := strings.ToLower(decoded)
	normalized = strings.TrimSpace(normalized)
	normalized = strings.Join(strings.Fields(normalized), " ")

	return normalized
}

func GetCaching(results model.GetGameID, cacheKey string, getCache map[string]model.GetGameID) map[string]model.GetGameID {
	getCache[cacheKey] = results
	return getCache
}
