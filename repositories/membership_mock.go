package repositories

import (
	"fmt"
	"injection-go/entities"
	"time"
)

type membershipRepositoryMock struct {
	memberships []entities.Membership
}

func NewMembershipRepositoryMock() *membershipRepositoryMock {

	membershipsMock := []entities.Membership{
		{
			MemberID:    1,
			Username:    "test01",
			Password:    "test001",
			Email:       "test001@gmail.com",
			MemberSince: time.Now(),
			LastVisit:   time.Now(),
		},
		{
			MemberID:    2,
			Username:    "test02",
			Password:    "test002",
			Email:       "test002@gmail.com",
			MemberSince: time.Now(),
			LastVisit:   time.Now(),
		},
		{
			MemberID:    3,
			Username:    "test03",
			Password:    "test003",
			Email:       "test003@gmail.com",
			MemberSince: time.Now(),
			LastVisit:   time.Now(),
		},
	}
	return &membershipRepositoryMock{
		memberships: membershipsMock,
	}
}

func (r *membershipRepositoryMock) GetAll() ([]entities.Membership, error) {
	// Return mock data

	return r.memberships, nil
}

func (r *membershipRepositoryMock) GetById(id int) (*entities.Membership, error) {

	// Return mock data
	for _, membership := range r.memberships {
		fmt.Println(membership)
		fmt.Println(membership.MemberID)
		if membership.MemberID == id {

			return &membership, nil
		}
	}
	return nil, nil
}
