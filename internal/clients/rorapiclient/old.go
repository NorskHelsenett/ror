package rorapiclient

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var rorclient *resty.Client
var rorclientnonauth *resty.Client

// Deprecated: GetOrCreateRorRestyClient is deprecated. Use rorclient instead
func GetOrCreateRorRestyClient() (*resty.Client, error) {
	if rorclient != nil {
		return rorclient, nil
	}

	rorclient = resty.New()
	rorclient.SetBaseURL(viper.GetString(configconsts.API_ENDPOINT))
	rorclient.Header.Add("X-API-KEY", viper.GetString(configconsts.API_KEY))
	rorclient.Header.Set("User-Agent", fmt.Sprintf("ROR-Agent/%s", viper.GetString(configconsts.VERSION)))

	return rorclient, nil
}

// Deprecated: GetOrCreateRorRestyClientNonAuth is deprecated. Use rorclient instead
func GetOrCreateRorRestyClientNonAuth() (*resty.Client, error) {
	if rorclientnonauth != nil {
		return rorclientnonauth, nil
	}

	rorclientnonauth = resty.New()
	rorclientnonauth.SetBaseURL(viper.GetString(configconsts.API_ENDPOINT))
	rorclientnonauth.Header.Set("User-Agent", fmt.Sprintf("ROR-Agent/%s", viper.GetString(configconsts.VERSION)))
	return rorclientnonauth, nil
}
