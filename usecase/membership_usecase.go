package usecase

import (
	"injection-go/entities"
)

// MembershipService represents a usecase for managing membership
type membershipService struct {
	memberRepo entities.MembershipRepository
}

// NewMembershipService creates a new membership usecase
func NewMembershipService(memberRepo entities.MembershipRepository) entities.MemberUsecaseRepository {
	return &membershipService{
		memberRepo: memberRepo,
	}
}

// GetAll returns all members
func (s *membershipService) GetAll() ([]entities.MemberUsecase, error) {
	// Get all members from the repository
	members, err := s.memberRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert members to member services
	var memberServices []entities.MemberUsecase
	for _, member := range members {
		memberServices = append(memberServices, entities.MemberUsecase{
			MemberID: member.MemberID,
			Username: member.Username,
			Email:    member.Email,
		})
	}
	return memberServices, nil
}

func (s *membershipService) GetById(id int) (*entities.MemberUsecase, error) {
	// Get member by ID from the repository

	membership, err := s.memberRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	memberServices := entities.MemberUsecase{
		MemberID: membership.MemberID,
		Username: membership.Username,
		Email:    membership.Email,
	}

	return &memberServices, nil
}
