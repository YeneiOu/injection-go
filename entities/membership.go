package entities

import "time"

type Membership struct {
	MemberID    int       `json:"member_id" db:"member_id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	Email       string    `json:"email" db:"email"`
	MemberSince time.Time `json:"member_since" db:"member_since"`
	LastVisit   time.Time `json:"last_visit" db:"last_visit"`
}

type MembershipRepository interface {
	GetAll() ([]Membership, error)
	GetById(id int) (*Membership, error)
}
