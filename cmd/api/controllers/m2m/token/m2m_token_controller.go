// TODO: Describe package
package token

import (
	"encoding/json"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/internal/models/m2mmodels"
	"github.com/NorskHelsenett/ror/internal/models/vaultmodels"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init m2m token controller")
	validate = validator.New()
}

// TODO: Describe function
//
// TODO: Add swagger
func SelfRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenModel m2mmodels.TokenModel
		ctx := c.Request.Context()
		//validate the request body
		if err := c.BindJSON(&tokenModel); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing body",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&tokenModel); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		secretPath := fmt.Sprintf("secret/data/v1.0/ror/clusters/%s", tokenModel.ClusterId)
		secret, err := apiconnections.VaultClient.GetSecret(secretPath)
		if err != nil {
			rlog.Errorc(ctx, "error getting secret", err, rlog.String("cluster id", tokenModel.ClusterId), rlog.String("path", secretPath))
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Error checking token",
			})
			return
		}

		byteArray, err := json.Marshal(secret)
		if err != nil {
			rlog.Errorc(ctx, "could not marshal", err)
		}

		var clusterSecret vaultmodels.VaultClusterModel
		err = json.Unmarshal(byteArray, &clusterSecret)
		if err != nil {
			rlog.Errorc(ctx, "could not unmarshal", err)
		}

		if len(clusterSecret.Data.RorClientSecret) > 0 {
			c.JSON(http.StatusUnauthorized, nil)
			return
		} else {
			newSecret := stringhelper.RandomString(20, stringhelper.StringTypeAlphaNum)
			clusterSecret.Data.RorClientSecret = newSecret
			secretByteArray, err := json.Marshal(clusterSecret)
			if err != nil {
				rlog.Errorc(ctx, "could not marshal clustersecret", err, rlog.String("cluster id", tokenModel.ClusterId))
				c.JSON(http.StatusInternalServerError, rorerror.RorError{
					Status:  http.StatusInternalServerError,
					Message: "A error occured",
				})
				return
			}

			_, err = apiconnections.VaultClient.SetSecret(secretPath, secretByteArray)
			if err != nil {
				rlog.Errorc(ctx, "Could not set secret", err, rlog.String("cluster id", tokenModel.ClusterId))
				c.JSON(http.StatusInternalServerError, rorerror.RorError{
					Status:  http.StatusInternalServerError,
					Message: "A error occured",
				})
				return
			}
			tokenModel.Token = clusterSecret.Data.RorClientSecret

			c.JSON(http.StatusOK, tokenModel)
			return
		}
	}
}
