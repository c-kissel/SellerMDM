package service

import (
	api "github.com/c-kissel/SellerMDM.git/internal/api/v1/seller"
	srv "github.com/c-kissel/SellerMDM.git/internal/service/seller"
)

type service struct {
	api.SellerServer
}

type Storer interface {
	srv.SellerStorage
}

func NewService(s *Storer) *service {

	st := (*s).(srv.SellerStorage)

	return &service{
		SellerServer: srv.NewSellerSrv(&st),
	}
}
