package entities

type MemberUsecase struct {
	MemberID int    `json:"member_id" db:"member_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type MemberUsecaseRepository interface {
	GetAll() ([]MemberUsecase, error)
	GetById(id int) (*MemberUsecase, error)
}
