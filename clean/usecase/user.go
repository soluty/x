package usecase

import "github.com/soluty/x/clean/entity"

type UserInput interface {
	UserCreate(username, email, password string) error
	UserGet(userName string) error
	UserEdit(userName string, toUpdate *entity.User) error

	UserLogin(email, password string) error
	FavoritesUpdate(username, slug string, favortie bool) error
}

type UserOutput interface {
	UserCreateOp(user *entity.User, token string)
	UserLoginOp(user *entity.User, token string)
	UserGetOp(user *entity.User, token string)
	UserEditOp(user *entity.User, token string)
	FavoritesUpdateOp() (*entity.User, *entity.Article)
}

type UserRepo interface {
	Create(username, email, password string) (*entity.User, error)
	GetById(userName int) (*entity.User, error)
	GetByEmailAndPassword(email, password string) (*entity.User, error)
	Save(user entity.User) error
}

func (i *interactor) UserCreate(username, email, password string) error {
	panic("implement me")
}

func (i *interactor) UserLogin(email, password string) error {
	panic("implement me")
}

func (i *interactor) UserGet(userName string) error {
	panic("implement me")
}

func (i *interactor) UserEdit(userName string, toUpdate *entity.User) error {
	panic("implement me")
}

func (i *interactor) FavoritesUpdate(username, slug string, favortie bool) error {
	panic("implement me")
}