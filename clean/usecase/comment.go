package usecase

import "github.com/soluty/x/clean/entity"

type CommentInput interface {
	CommentsSelect(articleSlug string) error
	CommentsCreate(userId int, slug, comment string) error
	CommentsDelete(userId int, slug string, id int) error
}

type CommentOutput interface {
	CommentsSelectOut([]entity.Comment)
	CommentsCreateOut(*entity.Comment)
	CommentsDeleteOut()
}

type CommentRepo interface {
	CreateComment(comment entity.Comment) (*entity.Comment, error)
	GetByID(id int) (*entity.Comment, error)
	DeleteComment(id int) error
}

func (i *interactor) CommentsSelect(slug string) error {
	panic("implement me")
}

func (i *interactor) CommentsCreate(userId int, slug, comment string) error {
	panic("implement me")
}

func (i *interactor) CommentsDelete(userId int, slug string, id int) error {
	comment, err := i.CommentRepo.GetByID(id)
	if err != nil {
		return err
	}
	if comment.Author.Id != userId {
		return todoErr
	}
	if err := i.CommentRepo.DeleteComment(id); err != nil {
		return err
	}
	article, err := i.ArticleRepo.GetBySlug(slug)
	if err != nil {
		return err
	}
	//article.UpdateComments(*comment, false)
	if _, err := i.ArticleRepo.Save(*article); err != nil {
		return err
	}
	return nil
}
