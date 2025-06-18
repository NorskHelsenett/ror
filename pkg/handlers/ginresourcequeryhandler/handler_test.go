package ginresourcequeryhandler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseResourceQuery(t *testing.T) {
	// Create a gin test context with our sample query parameters
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Full example URL from the request
	// GET /v2/resources?q=""&apiversion=general.ror.internal/v1alpha1&kind=KubernetesCluster&ownerrefs=[{"scope":"cluster","subject":"t-test-001-1337"}]&uids=uid1,uid2&fields=name,status&sort=+clustername,-clusterid&filters=[{"field":"status","value":"Running","type":"string","operator":"eq"}]&offset=0&limit=50

	// Set up the request with query parameters
	req := httptest.NewRequest("GET", "/v2/resources?q=&apiversion=general.ror.internal/v1alpha1&kind=KubernetesCluster&ownerrefs=[{\"scope\":\"cluster\",\"subject\":\"t-test-001-1337\"}]&uids=uid1,uid2&fields=name,status&sort=+clustername,-clusterid&filters=[{\"field\":\"status\",\"value\":\"Running\",\"type\":\"string\",\"operator\":\"eq\"}]&offset=0&limit=50", nil)
	c.Request = req

	// Parse the query
	query := ParseResourceQuery(c)

	// Validate the parsed query
	assert.Equal(t, "general.ror.internal", query.VersionKind.Group)
	assert.Equal(t, "v1alpha1", query.VersionKind.Version)
	assert.Equal(t, "KubernetesCluster", query.VersionKind.Kind)

	assert.Equal(t, 2, len(query.Uids))
	assert.Contains(t, query.Uids, "uid1")
	assert.Contains(t, query.Uids, "uid2")

	assert.Equal(t, 1, len(query.OwnerRefs))
	assert.Equal(t, "cluster", string(query.OwnerRefs[0].Scope))
	assert.Equal(t, "t-test-001-1337", string(query.OwnerRefs[0].Subject))

	assert.Equal(t, 2, len(query.Fields))
	assert.Contains(t, query.Fields, "name")
	assert.Contains(t, query.Fields, "status")

	assert.Equal(t, 2, len(query.Order))
	assert.Equal(t, "clustername", query.Order[0].Field)
	assert.False(t, query.Order[0].Descending)
	assert.Equal(t, "clusterid", query.Order[1].Field)
	assert.True(t, query.Order[1].Descending)

	assert.Equal(t, 1, len(query.Filters))
	assert.Equal(t, "status", query.Filters[0].Field)
	assert.Equal(t, "Running", query.Filters[0].Value)
	assert.Equal(t, "string", string(query.Filters[0].Type))
	assert.Equal(t, "eq", string(query.Filters[0].Operator))

	assert.Equal(t, 0, query.Offset)
	assert.Equal(t, 50, query.Limit)
}

func TestParseResourceQueryFromURL(t *testing.T) {
	urlStr := "/v2/resources?q=&apiversion=general.ror.internal/v1alpha1&kind=KubernetesCluster&ownerrefs=[{\"scope\":\"cluster\",\"subject\":\"t-test-001-1337\"}]&uids=uid1,uid2&fields=name,status&sort=+clustername,-clusterid&filters=[{\"field\":\"status\",\"value\":\"Running\",\"type\":\"string\",\"operator\":\"eq\"}]&offset=0&limit=50"

	query, err := ParseResourceQueryFromURL(urlStr)
	assert.NoError(t, err)

	// Validate the parsed query
	assert.Equal(t, "general.ror.internal", query.VersionKind.Group)
	assert.Equal(t, "v1alpha1", query.VersionKind.Version)
	assert.Equal(t, "KubernetesCluster", query.VersionKind.Kind)

	assert.Equal(t, 2, len(query.Uids))
	assert.Equal(t, 1, len(query.OwnerRefs))
	assert.Equal(t, 2, len(query.Fields))
	assert.Equal(t, 2, len(query.Order))
	assert.Equal(t, 1, len(query.Filters))
	assert.Equal(t, 0, query.Offset)
	assert.Equal(t, 50, query.Limit)
}

// // Example handler that demonstrates how to use the parser in a Gin handler function
// func ExampleHandler(c *gin.Context) {
// 	query := ParseResourceQuery(c)

// 	// Use the query to fetch resources
// 	// resourceSet := resourcesv2service.GetResourceByQuery(c, query)

// 	// Return the results
// 	c.JSON(http.StatusOK, gin.H{
// 		"query": query,
// 		// "resources": resourceSet,
// 	})
// }

// Example showing how to use with custom URL generation
// func ExampleURLGeneration() {
// 	// Build a URL string with query parameters
// 	apiVersion := "general.ror.internal/v1alpha1"
// 	kind := "KubernetesCluster"

// 	ownerRefs := []map[string]string{
// 		{
// 			"scope":   "cluster",
// 			"subject": "t-test-001-1337",
// 		},
// 	}
// 	ownerRefsJSON, _ := json.Marshal(ownerRefs)

// 	filters := []map[string]string{
// 		{
// 			"field":    "status",
// 			"value":    "Running",
// 			"type":     "string",
// 			"operator": "eq",
// 		},
// 	}
// 	filtersJSON, _ := json.Marshal(filters)

// 	url := "/v2/resources?q=&apiversion=" + apiVersion +
// 		"&kind=" + kind +
// 		"&ownerrefs=" + string(ownerRefsJSON) +
// 		"&uids=uid1,uid2" +
// 		"&fields=name,status" +
// 		"&sort=+clustername,-clusterid" +
// 		"&filters=" + string(filtersJSON) +
// 		"&offset=0&limit=50"

// 	// Parse the URL into a ResourceQuery object
// 	query, _ := ParseResourceQueryFromURL(url)

// 	// Use the query...
// 	_ = query
// }
