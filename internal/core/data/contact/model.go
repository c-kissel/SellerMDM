package contact

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"firstname"`
	LastName  string    `db:"lastname"`
	Email     string    `db:"email"`
	Mobile    string    `db:"email"`
	Messenger string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
