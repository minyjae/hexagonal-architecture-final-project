package entities

import "time"

type Order struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	Note          string         `json:"note"`
	OrderMappings []OrderMapping `json:"order_mappings"`
}

// OrderMapping table (many-to-many between ListQueue and Order)
type OrderMapping struct {
	ID          uint      `json:"id"`
	ListQueueID uint      `json:"list_queue_id"`
	OrderID     uint      `json:"order_id"`
	Checked     bool      `json:"checked"`
	ListQueue   ListQueue `json:"-"`
	Order       Order     `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
