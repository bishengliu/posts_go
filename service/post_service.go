package service

import (
	"errors"
	"math/rand"

	"github.com/bishehngliu/posts/entity"
	"github.com/bishehngliu/posts/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

var (
	repo repository.PostRepository
)

type service struct{}

func NewPostService(postRepository repository.PostRepository) PostService {
	repo = postRepository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
