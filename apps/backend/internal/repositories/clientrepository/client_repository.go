package clientrepository

import (
	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/clientmanager"
	"gorm.io/gorm"
)

type Client struct {
	ID           uuid.UUID `gorm:"primaryKey;type:bytes"`
	PasswordHash []byte    `gorm:"type:bytes"`
}

func New(db *gorm.DB) clientmanager.ClientRepository {
	db.AutoMigrate(&Client{})

	return &clientRepository{}
}

type clientRepository struct {
	db *gorm.DB
}

func (r *clientRepository) FindById(clientid uuid.UUID) (clientmanager.Client, error) {
	var client Client

	if res := r.db.First(&client).Where("id = ?", clientid); res.Error != nil {
		return clientmanager.Client{}, res.Error
	}

	c := clientmanager.Client{
		Id:           client.ID,
		PasswordHash: string(client.PasswordHash),
	}

	return c, nil
}
