package models

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

// Order table
type Order struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"not null" json:"title"`
	Note          string         `gorm:"not null" json:"note"`
	OrderMappings []OrderMapping `gorm:"foreignKey:OrderID" json:"order_mappings"`
}

func (o *Order) ToEntity() *entities.Order {
	return &entities.Order{
		ID:    o.ID,
		Title: o.Title,
		Note:  o.Note,
	}
}

func (o *Order) FromEntity(entity *entities.Order) {
	o.ID = entity.ID
	o.Title = entity.Title
	o.Note = entity.Note
}
