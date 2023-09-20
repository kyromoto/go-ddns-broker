package infrastructures

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kyromoto/go-ddns-broker/ClientManager/entities"
	"github.com/kyromoto/go-ddns-broker/ClientManager/usecases"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const filename = "client.db.yaml"

type BaseModel struct {
	UUID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (m *BaseModel) BeforeCreate(db *gorm.DB) error {
	m.UUID = uuid.New()
	return nil
}

type clientModel struct {
	BaseModel
	Description string
	Password    string
}

type clientRepositoryImpl struct{}

func (cr *clientRepositoryImpl) getDB(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		log.Error().Err(err)
		return nil, fmt.Errorf("load db failed")
	}

	return db, nil
}

func (cr *clientRepositoryImpl) Add(ctx *context.Context, client entities.Client) error {
	db, err := cr.getDB(filename)

	if err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return fmt.Errorf("access db failed")
	}

	if err = db.WithContext(*ctx).Save(client).Error; err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return fmt.Errorf("save client failed")
	}

	return nil
}

func (cr *clientRepositoryImpl) Delete(ctx *context.Context, uuid uuid.UUID) error {
	db, err := cr.getDB(filename)

	if err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return fmt.Errorf("access db failed")
	}

	if err = db.WithContext(*ctx).Delete(&clientModel{}, uuid).Error; err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return fmt.Errorf("delete client failed")
	}

	return nil
}

func (cr *clientRepositoryImpl) List(ctx *context.Context) ([]entities.Client, error) {
	db, err := cr.getDB(filename)

	if err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return nil, fmt.Errorf("access db failed")
	}

	var clientModels []clientModel

	if err = db.WithContext(*ctx).Find(&clientModel{}).Error; err != nil {
		log.Error().Err(err).Ctx(*ctx).Send()
		return nil, fmt.Errorf("get all clients failed")
	}

	var clients []entities.Client

	for _, c := range clientModels {
		clients = append(clients, entities.NewClient(c.UUID, c.Description, c.Password))
	}

	return clients, nil
}

func (cr *clientRepositoryImpl) FindByUuid(ctx *context.Context, uuid uuid.UUID) (*entities.Client, error) {
	db, err := cr.getDB(filename)

	if err != nil {
		return nil, err
	}

	var clientModel clientModel

	if err = db.WithContext(*ctx).First(&clientModel, "UUID = ?", uuid.String()).Error; err != nil {
		return nil, fmt.Errorf("not found")
	}

	client := entities.NewClient(clientModel.UUID, clientModel.Description, clientModel.Password)

	return &client, nil
}

func NewClientRepository() usecases.ClientRepository {
	return &clientRepositoryImpl{}
}
