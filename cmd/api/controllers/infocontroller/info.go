package infocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Version struct {
	Version string `json:"version"`
}

// TODO: Describe
//
//	@Summary	Get version
//	@Schemes
//	@Description	Get version
//	@Tags			info
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	map[string]interface{}
//	@Failure		403					{object}	map[string]interface{}
//	@Failure		401					{object}	map[string]interface{}
//	@Failure		500					{object}	map[string]interface{}
//	@Router			/v1/info/version	[get]
func GetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := Version{
			Version: viper.GetString(configconsts.VERSION),
		}
		output, err := json.Marshal(res)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "500: Could not marshal json")
			return
		}
		c.Data(http.StatusOK, "application/json", output)
	}
}
