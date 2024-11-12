package providers

import (
	"github.com/NorskHelsenett/ror/cmd/api/provider"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Summary	Get providers
// @Schemes
// @Description	Get providers
// @Tags			providers
// @Accept			application/json
// @Produce		application/json
// @Success		200				{array}		providers.Provider
// @Failure		403				{string}	Forbidden
// @Failure		401				{string}	Unauthorized
// @Failure		500				{string}	Failure	message
// @Router			/v1/providers [get]
// @Security		ApiKey || AccessToken
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var providerlist []providers.Provider
		providerlist = append(providerlist,
			providers.Provider{
				Name:     "Tanzu",
				Type:     providers.ProviderTypeTanzu,
				Disabled: false,
			},
			// providers.Provider{
			// 	Name:     "Azure",
			// 	Type:     providers.ProviderTypeAks,
			// 	Disabled: true,
			// },
			// providers.Provider{
			// 	Name:     "Google",
			// 	Type:     providers.ProviderTypeGke,
			// 	Disabled: true,
			// },
		)

		if viper.GetBool(configconsts.DEVELOPMENT) {
			providerlist = append(providerlist, providers.Provider{
				Name:     "Kind",
				Type:     providers.ProviderTypeKind,
				Disabled: false,
			})
			providerlist = append(providerlist, providers.Provider{
				Name:     "Talos",
				Type:     providers.ProviderTypeTalos,
				Disabled: false,
			})
		}

		c.JSON(http.StatusOK, providerlist)
	}
}

// @Summary	Get kuberntes versions by provider
// @Schemes
// @Description	Get supported kubernetes versions by provider
// @Tags			providers
// @Accept			application/json
// @Produce		application/json
//
//	@Param			providerType	path		string	true	"providerType"
//
// @Success		200				{array}		providers.Provider
// @Failure		403				{string}	Forbidden
// @Failure		401				{string}	Unauthorized
// @Failure		500				{string}	Failure	message
// @Router			/v1/providers/{providerType}/kubernetes/versions [get]
// @Security		ApiKey || AccessToken
func GetKubernetesVersionByProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := gincontext.GetRorContextFromGinContext(c)
		providerType := c.Param("providerType")
		defer cancel()

		if providerType == "" || len(providerType) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid provider",
			})
			return
		}

		switch providerType {
		case string(providers.ProviderTypeTanzu):
			c.JSON(http.StatusOK, getTanzuVersion())
			return
		case string(providers.ProviderTypeAks):
			c.JSON(http.StatusOK, make([]providers.ProviderKubernetesVersion, 0))
			return
		case string(providers.ProviderTypeKind):
			c.JSON(http.StatusOK, getKindVersions())
			return
		case string(providers.ProviderTypeTalos):
			c.JSON(http.StatusOK, getTalosVersions())
			return
		default:
			if viper.GetBool(configconsts.DEVELOPMENT) {
				c.JSON(http.StatusOK, getTanzuVersion())
				return
			}
			c.JSON(http.StatusOK, make([]providers.ProviderKubernetesVersion, 0))
			return
		}
	}
}

// @Summary	Get kuberntes versions by provider
// @Schemes
// @Description	Get supported kubernetes versions by provider
// @Tags			providers
// @Accept			application/json
// @Produce		application/json
//
//	@Param			providerType	path		string	true	"providerType"
//
// @Success		200				{array}		providers.Provider
// @Failure		403				{string}	Forbidden
// @Failure		401				{string}	Unauthorized
// @Failure		500				{string}	Failure	message
// @Router			/v1/providers/{providerType}/configs/params [get]
// @Security		ApiKey || AccessToken
func GetConfigParametersByProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := gincontext.GetRorContextFromGinContext(c)
		providerType := c.Param("providerType")
		defer cancel()

		if providerType == "" || len(providerType) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid provider",
			})
			return
		}

		provids := []providers.ProviderType{providers.ProviderType(providers.ProviderTypeTanzu)}

		providerloader := provider.NewProviderloader(provids)

		if !providerloader.IsProviderLoaded(providers.ProviderType(providerType)) {
			k8sprovider, _ := providerloader.GetProvider(providers.ProviderType(providerType))
			c.JSON(http.StatusOK, k8sprovider.GetConfigurations("k8sversion"))

			switch providerType {
			case string(providers.ProviderTypeTanzu):
				c.JSON(http.StatusOK, getTanzuVersion())
				return
			case string(providers.ProviderTypeAks):
				c.JSON(http.StatusOK, make([]providers.ProviderKubernetesVersion, 0))
				return
			case string(providers.ProviderTypeKind):
				c.JSON(http.StatusOK, getKindVersions())
				return
			default:
				if viper.GetBool(configconsts.DEVELOPMENT) {
					c.JSON(http.StatusOK, getTanzuVersion())
					return
				}
				c.JSON(http.StatusOK, make([]providers.ProviderKubernetesVersion, 0))
				return
			}
		}
	}
}

func getTanzuVersion() []providers.ProviderKubernetesVersion {
	var kubernetesVersions []providers.ProviderKubernetesVersion
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.28.7",
		Version:  "v1.28.7---vmware.1-fips.1-tkg.1",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.27.10",
		Version:  "v1.27.10---vmware.1-fips.1-tkg.1",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.26.12",
		Version:  "v1.26.12---vmware.2-fips.1-tkg.2",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.25.13",
		Version:  "v1.25.13---vmware.1-fips.1-tkg.1",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.24.11",
		Version:  "v1.24.11---vmware.1-fips.1-tkg.1",
		Disabled: false,
	})
	return kubernetesVersions
}

func getKindVersions() []providers.ProviderKubernetesVersion {
	var kubernetesVersions []providers.ProviderKubernetesVersion
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.31.0",
		Version:  "kindest/node:v1.31.0@sha256:53df588e04085fd41ae12de0c3fe4c72f7013bba32a20e7325357a1ac94ba865",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.30.4",
		Version:  "kindest/node:v1.30.4@sha256:976ea815844d5fa93be213437e3ff5754cd599b040946b5cca43ca45c2047114",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.29.8",
		Version:  "kindest/node:v1.29.8@sha256:d46b7aa29567e93b27f7531d258c372e829d7224b25e3fc6ffdefed12476d3aa",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.28.13",
		Version:  "kindest/node:v1.28.13@sha256:45d319897776e11167e4698f6b14938eb4d52eb381d9e3d7a9086c16c69a8110",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.27.17",
		Version:  "kindest/node:v1.27.17@sha256:3fd82731af34efe19cd54ea5c25e882985bafa2c9baefe14f8deab1737d9fabe",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.26.15",
		Version:  "kindest/node:v1.26.15@sha256:1cc15d7b1edd2126ef051e359bf864f37bbcf1568e61be4d2ed1df7a3e87b354",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.25.16",
		Version:  "kindest/node:v1.25.16@sha256:6110314339b3b44d10da7d27881849a87e092124afab5956f2e10ecdb463b025",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.24.17",
		Version:  "kindest/node:v1.24.17@sha256:bad10f9b98d54586cba05a7eaa1b61c6b90bfc4ee174fdc43a7b75ca75c95e51",
		Disabled: false,
	}, providers.ProviderKubernetesVersion{
		Name:     "v1.23.17",
		Version:  "kindest/node:v1.23.17@sha256:14d0a9a892b943866d7e6be119a06871291c517d279aedb816a4b4bc0ec0a5b3",
		Disabled: false,
	},
	)
	return kubernetesVersions
}

func getTalosVersions() []providers.ProviderKubernetesVersion {
	var kubernetesVersions []providers.ProviderKubernetesVersion
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.31.1",
		Version:  "v1.31.1",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.30.3",
		Version:  "v1.30.3",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.29.6",
		Version:  "v1.29.6",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.28.11",
		Version:  "v1.28.11",
		Disabled: false,
	})
	kubernetesVersions = append(kubernetesVersions, providers.ProviderKubernetesVersion{
		Name:     "v1.27.15",
		Version:  "v1.27.15",
		Disabled: false,
	})
	return kubernetesVersions
}
