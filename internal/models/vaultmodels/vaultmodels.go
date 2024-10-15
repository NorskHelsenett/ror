package vaultmodels

type VaultClusterModel struct {
	Data Data `json:"data"`
}

type Data struct {
	RorClientSecret string `json:"rorClientSecret"`
}

type VaultDexModel struct {
	Data DexData `json:"data"`
}

type DexData struct {
	DexSecret string `json:"dexSecret"`
}
