package model

import (
	"github.com/google/uuid"
)

type Auth struct {
	ID   uuid.UUID `gorm:"type:varchar(255);default:uuid_generate_v4();primary_key" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`
}

type Keypair struct {
	ID      uuid.UUID `gorm:"type:varchar(255);default:uuid_generate_v4();primary_key" json:"id"`
	Key     string    `gorm:"type:varchar(255);not null" json:"key"`
	Value   string    `gorm:"type:varchar(255);not null" json:"value"`
	OwnedBy uuid.UUID `gorm:"type:varchar(255);not null" json:"ownedBy"`
}
