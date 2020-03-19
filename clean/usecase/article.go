package usecase

import "github.com/soluty/x/clean/entity"

type ArticleInput interface {
	ArticlesFeed(userId int, limit, offset int) error
	ArticlesSelect(userId int, limit, offset int, filters []entity.ArticleFilter) error

	ArticleSelect(slug string, username int) error
	ArticleCreate(userId int, article entity.Article) error
	ArticleUpdate(userId int, slug string, newArticle *entity.Article) error
	ArticleDelete(userId int, slug string) error
}

type ArticleOutput interface {
	ArticlesFeedOp(requestingUser *entity.User, articles entity.ArticleList, totalArticleCount int)
	GetArticlesOp(requestingUser *entity.User, articles entity.ArticleList, totalArticleCount int)

	ArticleGetOp(*entity.User, *entity.Article)
	ArticlePostOp(*entity.User, *entity.Article)
	ArticlePutOp(*entity.User, *entity.Article)
	ArticleDeleteOp()
}

type ArticleRepo interface {
	Create(entity.Article) (*entity.Article, error)
	Save(entity.Article) (*entity.Article, error)
	GetBySlug(slug string) (*entity.Article, error)
	GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]entity.Article, error)
	GetRecentFiltered(filters []entity.ArticleFilter) ([]entity.Article, error)
	Delete(slug string) error
}

func (i *interactor) ArticlesFeed(userId int, limit, offset int) error {
	panic("implement me")
}

func (i *interactor) ArticlesSelect(userId int, limit, offset int, filters []entity.ArticleFilter) error {
	panic("implement me")
}

func (i *interactor) ArticleSelect(slug string, userId int) error {
	var user *entity.User
	if userId != 0 {
		var err error
		user, err = i.UserRepo.GetById(userId)
		if err != nil {
			return err
		}
	}
	article, err := i.ArticleRepo.GetBySlug(slug)
	if err != nil {
		return err
	}
	i.ArticleOutput.ArticleGetOp(user, article)
	return nil
}

func (i *interactor) ArticleCreate(userId int, article entity.Article) error {
	panic("implement me")
}

func (i *interactor) ArticleUpdate(userId int, slug string, newArticle *entity.Article) error {
	panic("implement me")
}

func (i *interactor) ArticleDelete(userId int, slug string) error {
	panic("implement me")
}
