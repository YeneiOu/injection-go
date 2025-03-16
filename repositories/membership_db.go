package repositories

import (
	"injection-go/entities"

	"github.com/jmoiron/sqlx"
)

type membershipRepositoryDB struct {
	db *sqlx.DB
}

func NewMembershipRepository(db *sqlx.DB) entities.MembershipRepository {
	return &membershipRepositoryDB{
		db: db,
	}
}

func (r *membershipRepositoryDB) GetAll() ([]entities.Membership, error) {
	query := `SELECT member_id, username, password, email, member_since, last_visit
	FROM membership;`
	var memberships []entities.Membership
	err := r.db.Select(&memberships, query)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

func (r *membershipRepositoryDB) GetById(id int) (*entities.Membership, error) {

	var membership entities.Membership
	query := `SELECT member_id, username, password, email, member_since, last_visit
	FROM membership WHERE member_id = $1;`

	err := r.db.Get(&membership, query, id)
	if err != nil {
		return nil, err
	}
	return &membership, nil
}
