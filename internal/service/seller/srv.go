package seller

import (
	"github.com/c-kissel/SellerMDM.git/internal/core/data/seller"
	"github.com/c-kissel/SellerMDM.git/specs"
	"github.com/google/uuid"
)

type sellerSrv struct {
	SellerStorage
}

func NewSellerSrv(s *SellerStorage) *sellerSrv {
	return &sellerSrv{
		SellerStorage: *s,
	}
}

type SellerStorage interface {
	Get(id uuid.UUID) (seller.Seller, error)
}

func (s *sellerSrv) Get(id uuid.UUID) (specs.Seller, error) {
	var result specs.Seller

	data, err := s.SellerStorage.Get(id)
	if err != nil {
		return result, err
	}

	result = data.ToSpecs()
	return result, nil
}
