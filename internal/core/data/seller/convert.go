package seller

import "github.com/c-kissel/SellerMDM.git/specs"

// Returns specs Seller from data Seller
func (s *Seller) ToSpecs() specs.Seller {

	id := s.ID
	contact := s.Contact
	name := s.Name
	descr := s.Description
	ogrn := s.OGRN
	inn := s.INN
	city := s.City
	created := s.CreatedAt.String()
	updated := s.UpdatedAt.String()

	result := &specs.Seller{
		Id:          &id,
		Contact:     &contact,
		City:        &city,
		Created:     &created,
		Description: &descr,
		Inn:         &inn,
		Name:        &name,
		Ogrn:        &ogrn,
		Updated:     &updated,
	}

	return *result
}
