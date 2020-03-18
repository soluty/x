package usecase

import "github.com/soluty/x/clean/entity"

type CommentInput interface {
	CommentsGet(slug string) error
	CommentsPost(username, slug, comment string) error
	CommentsDelete(username, slug string, id int) error
}

type CommentOutput interface {
	CommentsGetOp([]entity.Comment)
	CommentsPostOp(*entity.Comment)
	CommentsDeleteOp()
}

type CommentRepo interface {
	Create(comment entity.Comment) (*entity.Comment, error)
	GetByID(id int) (*entity.Comment, error)
	Delete(id int) error
}


func (i *interactor) CommentsGet(slug string) error {
	panic("implement me")
}

func (i *interactor) CommentsPost(username, slug, comment string) error {
	panic("implement me")
}

func (i *interactor) CommentsDelete(username, slug string, id int) error {
	panic("implement me")
}