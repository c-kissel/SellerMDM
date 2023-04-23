package seller

import (
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

// Returns specs Seller from data Seller
func (s *SellerModel) ToSpecs() specs.SellerResponse {
	// City
	var city string
	if s.City != nil {
		city = *s.City
	}

	// Id
	id := s.Id

	// INN
	var inn string
	if s.INN != nil {
		inn = *s.INN
	}

	// Logo
	var logo string
	if s.Logo != nil {
		logo = *s.Logo
	}

	// Memo
	var memo string
	if s.Memo != nil {
		memo = *s.Memo
	}

	// Name
	name := s.Name

	// OGRN
	var ogrn string
	if s.OGRN != nil {
		ogrn = *s.OGRN
	}

	// Site
	var site string
	if s.Site != nil {
		site = *s.Site
	}

	// YML
	var yml string
	if s.YML != nil {
		yml = *s.YML
	}

	created := s.CreatedAt.String()
	updated := s.UpdatedAt.String()

	result := &specs.SellerResponse{
		City:    &city,
		Created: &created,
		Id:      &id,
		Inn:     &inn,
		Logo:    &logo,
		Memo:    &memo,
		Name:    &name,
		Ogrn:    &ogrn,
		Site:    &site,
		Updated: &updated,
		Yml:     &yml,
	}

	return *result
}

func FromNewRequest(s specs.NewSellerRequest) SellerModel {
	// City
	var city string
	if s.City != nil {
		city = *s.City
	}

	// INN
	var inn string
	if s.Inn != nil {
		inn = *s.Inn
	}

	// Logo
	var logo string
	if s.Logo != nil {
		logo = *s.Logo
	}

	// Memo
	var memo string
	if s.Memo != nil {
		memo = *s.Memo
	}

	// Name
	var name string
	if s.Name != nil {
		name = *s.Name
	}

	// OGRN
	var ogrn string
	if s.Ogrn != nil {
		ogrn = *s.Ogrn
	}

	// Site
	var site string
	if s.Site != nil {
		site = *s.Site
	}

	// YML
	var yml string
	if s.Yml != nil {
		yml = *s.Yml
	}

	return SellerModel{
		City: &city,
		INN:  &inn,
		Logo: &logo,
		Memo: &memo,
		Name: name,
		OGRN: &ogrn,
		Site: &site,
		YML:  &yml,
	}
}

func FromEditRequest(s specs.EditSellerRequest) SellerModel {
	// City
	var city string
	if s.City != nil {
		city = *s.City
	}

	// Id
	var id uuid.UUID
	if s.Id == nil {
		id = uuid.New()
	} else {
		id = *s.Id
	}

	// INN
	var inn string
	if s.Inn != nil {
		inn = *s.Inn
	}

	// Logo
	var logo string
	if s.Logo != nil {
		logo = *s.Logo
	}

	// Memo
	var memo string
	if s.Memo != nil {
		memo = *s.Memo
	}

	// Name
	var name string
	if s.Name != nil {
		name = *s.Name
	}

	// OGRN
	var ogrn string
	if s.Ogrn != nil {
		ogrn = *s.Ogrn
	}

	// Site
	var site string
	if s.Site != nil {
		site = *s.Site
	}

	// YML
	var yml string
	if s.Yml != nil {
		yml = *s.Yml
	}

	return SellerModel{
		City: &city,
		Id:   id,
		INN:  &inn,
		Logo: &logo,
		Memo: &memo,
		Name: name,
		OGRN: &ogrn,
		Site: &site,
		YML:  &yml,
	}
}
