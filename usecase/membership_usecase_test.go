package usecase

import (
	"injection-go/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Initialize mock repository
	mockRepo := repositories.NewMembershipRepositoryMock()

	// Initialize usecase
	memberService := NewMembershipService(mockRepo)

	// Test GetAll
	members, err := memberService.GetAll()

	// ตรวจสอบผลลัพธ์
	assert.NoError(t, err)
	assert.Equal(t, 3, len(members)) // ตรวจสอบจำนวนข้อมูลที่ return

	// ตรวจสอบข้อมูล mock รายการแรก
	assert.Equal(t, 1, members[0].MemberID)
	assert.Equal(t, "test01", members[0].Username)
	assert.Equal(t, "test001@gmail.com", members[0].Email)

	// ตรวจสอบข้อมูล mock รายการที่สอง
	assert.Equal(t, 2, members[1].MemberID)
	assert.Equal(t, "test02", members[1].Username)
	assert.Equal(t, "test002@gmail.com", members[1].Email)

	// ตรวจสอบข้อมูล mock รายการที่สาม
	assert.Equal(t, 3, members[2].MemberID)
	assert.Equal(t, "test03", members[2].Username)
	assert.Equal(t, "test003@gmail.com", members[2].Email)
}
