package resourceupdatev2

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients"
	"github.com/NorskHelsenett/ror/cmd/agentv2/services/authservice"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// Hashlist for use in agent communication
type hashList struct {
	Items []hashItem `json:"items"`
}

// Item for use in the hashlist
type hashItem struct {
	Uid    string `json:"uid"`
	Hash   string `json:"hash"`
	Active bool
}

func InitHashList() (*apicontractsv2resources.HashList, error) {
	var hashList apicontractsv2resources.HashList
	rorclient := clients.RorConfig.GetRorClient()
	apiHashList, err := rorclient.Resources().GetHashList(authservice.CreateRorOwnerref())
	if err != nil {
		fmt.Println("Error getting hashlist from api", err)
		return nil, err
	}
	for _, item := range apiHashList.Items {
		hashList.Items = append(hashList.Items, apicontractsv2resources.HashItem{
			Uid:  item.Uid,
			Hash: item.Hash,
		})
	}
	return &hashList, nil

}
func InitHashListv2() (*apicontractsv2resources.HashList, error) {
	var hashList apicontractsv2resources.HashList
	rorclient := clients.RorConfig.GetRorClient()
	hashList, err := rorclient.ResourceV2().GetOwnHashes()
	if err != nil {
		fmt.Println("Error getting hashlist from api", err)
		return nil, err
	}
	return &hashList, nil

}

func (hl hashList) getInactiveUid() []string {
	var ret []string
	if len(hl.Items) == 0 {
		return ret
	}
	for i := range hl.Items {
		if !hl.Items[i].Active {
			ret = append(ret, hl.Items[i].Uid)
		}
	}
	return ret
}

func (hl *hashList) markActive(uid string) {
	item, i := hl.getHashByUid(uid)
	if item.Uid != "" {
		hl.Items[i].Active = true
	}

}

// Returns a bool value of true if the resource need to be commited
func (rc hashList) checkUpdateNeeded(uid string, hash string) bool {
	hashitem, _ := rc.getHashByUid(uid)
	if hashitem.Hash == hash {
		rlog.Debug("No need to update, hash matched")
		return false
	} else {
		return true
	}
}
func (hl hashList) getHashByUid(uid string) (hashItem, int) {
	if len(hl.Items) > 0 {
		for i := range hl.Items {
			if hl.Items[i].Uid == uid {
				return hl.Items[i], i
			}
		}
	}
	return hashItem{}, 0
}

// updates hash in internal hashlist on update. The api will update its list on commiting the resource to its database.
func (hl *hashList) updateHash(uid string, hash string) {
	_, i := hl.getHashByUid(uid)
	if i != 0 {
		rlog.Debug("Update needed, hash updated", rlog.String("uid", uid))
		hl.Items[i].Hash = hash
		return
	}
	rlog.Debug("Uid not found in hashList, adding hash", rlog.String("uid", uid))

	newItem := hashItem{
		Uid:    uid,
		Hash:   hash,
		Active: true,
	}
	hl.Items = append(hl.Items, newItem)
}
