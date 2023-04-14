package review

import "github.com/google/uuid"

type Review struct {
	ID         uuid.UUID `db:"id"`
	Seller     uuid.UUID `db:"seller"`
	Contact    uuid.UUID `db:"contact"`
	Rating     int       `db:"rating"`
	Commentary string    `db:"commentary"`
	Parent     uuid.UUID `db:"parent"`
}
