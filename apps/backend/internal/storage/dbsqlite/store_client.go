package dbsqlite

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns/internal/services/client"
	"gorm.io/gorm"
)

type Client struct {
	ID           uuid.UUID `gorm:"primaryKey;type:bytes"`
	Description  string    `gorm:"type:string"`
	PasswordHash []byte    `gorm:"type:bytes"`
}

func NewClientStore(db *gorm.DB) client.ClientRepository {
	db.AutoMigrate(&Client{})

	return &clientRepository{}
}

type clientRepository struct {
	db *gorm.DB
}

func (r *clientRepository) Save(client client.ClientRepositoryDTO) error {
	result := r.db.Save(Client{
		ID:           client.ID,
		Description:  client.Description,
		PasswordHash: []byte(client.Password),
	})

	if result.Error != nil {
		return fmt.Errorf("save client to db failed")
	}

	return nil
}

func (r *clientRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Client{}).Where("ID = ?", []byte(id.String())).Error
}

func (r *clientRepository) Read() ([]client.ClientRepositoryDTO, error) {
	clientModels := make([]Client, 0)
	clientDTOs := make([]client.ClientRepositoryDTO, 0)

	if res := r.db.Find(&clientModels); res.Error != nil {
		return clientDTOs, res.Error
	}

	for _, c := range clientModels {
		clientDTOs = append(clientDTOs, client.ClientRepositoryDTO{
			ID:          c.ID,
			Description: c.Description,
			Password:    string(c.PasswordHash),
		})
	}

	return clientDTOs, nil
}

func (r *clientRepository) ReadById(id uuid.UUID) (client.ClientRepositoryDTO, error) {
	var clientModel Client

	if res := r.db.First(&clientModel).Where("ID = ?", []byte(id.String())); res.Error != nil {
		return client.ClientRepositoryDTO{}, res.Error
	}

	return client.ClientRepositoryDTO{
		ID:          clientModel.ID,
		Description: clientModel.Description,
		Password:    string(clientModel.PasswordHash),
	}, nil
}
