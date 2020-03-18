package usecase

import "github.com/soluty/x/clean/entity"

type ArticleInput interface {
	ArticlesFeed(username string, limit, offset int) error
	GetArticles(username string, limit, offset int, filters []entity.ArticleFilter) error
	ArticleGet(slug, username string) error
	ArticlePost(username string, article entity.Article) error
	ArticlePut(username, slug string, newArticle *entity.Article) error
	ArticleDelete(username, slug string) error
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


func (i *interactor) ArticlesFeed(username string, limit, offset int) error {
	panic("implement me")
}

func (i *interactor) GetArticles(username string, limit, offset int, filters []entity.ArticleFilter) error {
	panic("implement me")
}

func (i *interactor) ArticleGet(slug, username string) error {
	panic("implement me")
}

func (i *interactor) ArticlePost(username string, article entity.Article) error {
	panic("implement me")
}

func (i *interactor) ArticlePut(username, slug string, newArticle *entity.Article) error {
	panic("implement me")
}

func (i *interactor) ArticleDelete(username, slug string) error {
	panic("implement me")
}
