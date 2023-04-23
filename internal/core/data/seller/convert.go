package seller

import (
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

// Returns specs Seller from data Seller
func (s *SellerModel) ToSpecs() specs.Seller {

	id := s.ID
	name := s.Name
	ogrn := s.OGRN
	inn := s.INN
	city := s.City
	created := s.CreatedAt.String()
	updated := s.UpdatedAt.String()
	// descr := s.Description

	result := &specs.Seller{
		Id:      &id,
		City:    &city,
		Created: &created,
		Inn:     &inn,
		Name:    &name,
		Ogrn:    &ogrn,
		Updated: &updated,
		// Description: &descr,
	}

	return *result
}

func FromSpecs(s specs.Seller) SellerModel {
	var id uuid.UUID
	var name string
	// var img []string = make([]string, 0)
	var ogrn string
	var inn string
	var city string
	// var descr string

	if s.Id == nil {
		id = uuid.New()
	} else {
		id = *s.Id
	}
	if s.Name != nil {
		name = *s.Name
	}
	// if s.Description != nil {
	// 	descr = *s.Description
	// }
	// if s.ImageNames != nil {
	// 	img = append(img, *s.ImageNames...)
	// }
	if s.Ogrn != nil {
		ogrn = *s.Ogrn
	}
	if s.Inn != nil {
		inn = *s.Inn
	}
	if s.City != nil {
		city = *s.City
	}

	return SellerModel{
		ID:   id,
		Name: name,
		OGRN: ogrn,
		INN:  inn,
		// ImageNames: img,
		City: city,
		// Description: descr,
	}
}
