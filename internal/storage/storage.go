package storage

import "github.com/c-kissel/SellerMDM.git/internal/service/seller"

type storage struct {
	seller.SellerStorage
}

func NewStorage() *storage {
	return &storage{}
}
