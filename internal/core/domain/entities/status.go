package entities

import "time"

type StaffStatus struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStatus struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StatusMapping struct {
	ID            uint        `json:"id"`
	StaffStatusID uint        `json:"staff_status_id"`
	StaffStatus   StaffStatus `json:"staff_status"`
	UserStatusID  uint        `json:"user_status_id"`
	UserStatus    UserStatus  `json:"user_status"`
}
