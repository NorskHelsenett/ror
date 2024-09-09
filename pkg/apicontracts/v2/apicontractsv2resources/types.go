package apicontractsv2resources

import "github.com/NorskHelsenett/ror/pkg/rlog"

type HashList struct {
	Items []HashItem `json:"items"`
}

// Item for use in the hashlist
// NB This struct has a counterpart in the agent.
type HashItem struct {
	Uid     string `json:"uid"`
	Hash    string `json:"hash"`
	Version string `json:"version"`
	active  bool
}

func (hl HashList) GetInactiveUid() []string {
	var ret []string
	if len(hl.Items) == 0 {
		return ret
	}
	for i := range hl.Items {
		if !hl.Items[i].active {
			ret = append(ret, hl.Items[i].Uid)
		}
	}
	return ret
}

func (hl *HashList) MarkActive(uid string) {
	item, i := hl.GetHashByUid(uid)
	if item.Uid != "" {
		hl.Items[i].active = true
	}

}

// Returns a bool value of true if the resource need to be committed
func (hl HashList) CheckUpdateNeeded(uid string, hash string) bool {
	hashitem, _ := hl.GetHashByUid(uid)
	if hashitem.Hash == hash {
		rlog.Debug("No need to update, hash matched")
		return false
	} else {
		return true
	}
}
func (hl HashList) GetHashByUid(uid string) (HashItem, int) {
	if len(hl.Items) > 0 {
		for i := range hl.Items {
			if hl.Items[i].Uid == uid {
				return hl.Items[i], i
			}
		}
	}
	return HashItem{}, 0
}

// updates hash in internal hashlist on update. The api will update its list on committing the resource to its database.
func (hl *HashList) UpdateHash(uid string, hash string) {
	_, i := hl.GetHashByUid(uid)
	if i != 0 {
		rlog.Debug("Update needed, hash updated", rlog.String("uid", uid))
		hl.Items[i].Hash = hash
		return
	}
	rlog.Debug("Uid not found in hashList, adding hash", rlog.String("uid", uid))

	newItem := HashItem{
		Uid:    uid,
		Hash:   hash,
		active: true,
	}
	hl.Items = append(hl.Items, newItem)
}
