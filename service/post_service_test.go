package service

import (
	"testing"

	"github.com/bishehngliu/posts/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := m.Called()

	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockRepository) FindAll() ([]entity.Post, error) {

	args := m.Called()
	return args.Get(0).([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {

	post := entity.Post{Id: 1, Title: "", Text: "Text"}

	testService := NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Save").Return(&entity.Post{Id: 1, Title: "Title", Text: "Text"}, nil)

	testService := NewPostService(mockRepo)
	post, err := testService.Create(&entity.Post{Id: 1, Title: "Title", Text: "Text"})

	mockRepo.AssertExpectations(t)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, post)

	assert.Equal(t, "Title", post.Title)
	assert.Equal(t, "Text", post.Text)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return([]entity.Post{{Id: 1, Title: "Title", Text: "Text"}}, nil)

	testService := NewPostService(mockRepo)
	posts, err := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, posts)
	assert.Equal(t, []entity.Post{{Id: 1, Title: "Title", Text: "Text"}}, posts)
}
