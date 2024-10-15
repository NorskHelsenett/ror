package clusters

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	viper.Set(configconsts.OIDC_PROVIDER, testServer.URL)
	viper.Set(configconsts.VAULT_URL, testServer.URL)

	// mtest.Setup()
}

func setup() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, ginEngine := gin.CreateTestContext(recorder)

	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.POST("/v1/clusters/filter", ClusterByFilter())

	return context, ginEngine, recorder
}

func setupWithIdentity(identity identitymodels.Identity) (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, ginEngine := gin.CreateTestContext(recorder)

	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	ginEngine.Use(func(ctx *gin.Context) {
		ctx.Set("identity", identity)
		// TODO: Remove user
		ctx.Set("user", *identity.User)
	})

	ginEngine.POST("/v1/clusters/filter", ClusterByFilter())

	return context, ginEngine, recorder
}

func Test_ClustersController_WithoutUser(t *testing.T) {
	_, ginEngine, recorder := setup()
	t.Run("NoAccessToken_Returns401", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodPost, "/v1/clusters/filter", nil)
		ginEngine.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusUnauthorized, recorder.Result().StatusCode)
	})
}

func Test_ClustersController_WithMockUserWithoutGroups(t *testing.T) {
	user := identitymodels.User{
		Name:            "Test",
		IsEmailVerified: true,
		Email:           "test@ror",
	}
	identity := identitymodels.Identity{
		Type: identitymodels.IdentityTypeUser,
		User: &user,
	}
	context, ginEngine, recorder := setupWithIdentity(identity)
	t.Run("UserInContextSetButMissingFilterBody_ReturnsBadRequest", func(t *testing.T) {
		context.Request, _ = http.NewRequest(http.MethodPost, "/v1/clusters/filter", nil)
		ginEngine.ServeHTTP(recorder, context.Request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func Test_ClustersController_WithMockUserWithGroups(t *testing.T) {
	user := identitymodels.User{
		Name:            "Test",
		IsEmailVerified: true,
		Email:           "test@ror",
		Groups: []string{
			"Developers@ror",
		},
	}
	identity := identitymodels.Identity{
		Type: identitymodels.IdentityTypeUser,
		User: &user,
	}
	context, ginEngine, recorder := setupWithIdentity(identity)
	t.Run("UserInContextSetButMissingFilterBody_ReturnsBadRequest", func(t *testing.T) {
		context.Request, _ = http.NewRequest(http.MethodPost, "/v1/clusters/filter", nil)
		ginEngine.ServeHTTP(recorder, context.Request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

// TODO: Mock mongodb clusters (or should we do this another way?!)
// func Test_ClustersController_WithMockUserWithGroups2(t *testing.T) {
// 	user := identitymodels.User{
// 		Name:            "Test",
// 		IsEmailVerified: true,
// 		Email:           "test@ror",
// 		Groups: []string{
// 			"ShouldNotGetAnythingGroup",
// 		},
// 	}

// 	context, ginEngine, recorder := setupWithUser(user)
// 	t.Run("UserInContextSetButMissingFilterBody_ReturnsBadRequest", func(t *testing.T) {
// 		filter := apiModels.Filter{
// 			Skip:  0,
// 			Limit: 10,
// 			SortList: []apiModels.Sort{
// 				{
// 					By:  "clustername",
// 					Asc: true,
// 				},
// 			},
// 		}

// 		bodyByte, _ := json.Marshal(filter)

// 		context.Request, _ = http.NewRequest(http.MethodPost, "/v1/clusters/filter", bytes.NewBuffer(bodyByte))
// 		ginEngine.ServeHTTP(recorder, context.Request)
// 		assert.Equal(t, http.StatusBadRequest, recorder.Code)
// 	})
// }
