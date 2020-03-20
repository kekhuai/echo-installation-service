package store

import (
	"github.com/Kamva/mgm/v2"
	"github.com/kekhuay/echo-installation-service/model"
)

type PartnerStore struct{}

func NewPartnerStore() *PartnerStore {
	return &PartnerStore{}
}

// CreatePartner and persist
func (ps *PartnerStore) CreatePartner(p *model.Partner) error {
	if err := mgm.Coll(p).Create(p); nil != err {
		return err
	}
	return nil
}
