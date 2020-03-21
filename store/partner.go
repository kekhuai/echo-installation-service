package store

import (
	"github.com/Kamva/mgm/v2"
	"github.com/kekhuay/echo-installation-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PartnerStore implementation
type PartnerStore struct{}

// NewPartnerStore return an instance of PartnerStore
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

// List all of partners
func (ps *PartnerStore) List(offset, limit int64) ([]model.Partner, int64, error) {
	partnerModel := &model.Partner{}
	context := mgm.Ctx()
	count, err := mgm.Coll(partnerModel).EstimatedDocumentCount(context)
	if nil != err {
		return nil, 0, err
	}
	partners := []model.Partner{}
	err = mgm.Coll(partnerModel).SimpleFind(&partners, bson.M{}, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})
	if nil != err {
		return nil, 0, err
	}
	return partners, count, nil
}

func (ps *PartnerStore) UpdatePartner(p *model.Partner) error {
	err := mgm.Coll(p).Update(p)
	if nil != err {
		return err
	}
	return nil
}

func (ps *PartnerStore) GetById(id string) (*model.Partner, error) {
	partner := &model.Partner{}
	err := mgm.Coll(partner).FindByID(id, partner)
	if nil != err {
		return nil, err
	}
	return partner, nil
}
