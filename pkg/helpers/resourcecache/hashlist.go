package resourcecache

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/hashlist"
)

// func InitHashList() (*hashlist.HashList, error) {
// 	rorclient := clients.RorConfig.GetRorClient()
// 	hashList, err := rorclient.ResourceV2().GetOwnHashes(context.TODO(), clients.RorConfig.GetClusterId())
// 	if err != nil {
// 		fmt.Println("Error getting hashlist from api", err)
// 		return nil, err
// 	}
// 	return hashList, nil

// }

func NewEmptyHashList() *hashlist.HashList {
	return &hashlist.HashList{Items: []hashlist.HashItem{}}
}
