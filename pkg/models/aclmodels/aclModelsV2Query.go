package aclmodels

// ClusterIdToUidResolver is an optional function that resolves a cluster ID
// (human-readable name) to its UID (UUID). Set this at init time in the
// application that has database access (e.g. ror-api).
// Returns the UID string, or empty string if not found.
var ClusterIdToUidResolver func(clusterID string) string
