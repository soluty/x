package usecase

import "github.com/soluty/x/clean/entity"

type UserInput interface {
	// post /users/login
	UserCreate(username, email, password string) error
	// get /user
	UserSelect(userId int) error
	// put /user
	UserUpdate(userId int, toUpdate *entity.User) error
	// post /users/login
	UserLogin(email, password string) error
	// POST|DELETE /articles/:slug/favorite
	FavoritesUpdate(userId int, slug string, favortie bool) error
}

type UserOutput interface {
	UserLoginOut(user *entity.User)
	UserCreateOut(user *entity.User)
	UserSelectOut(user *entity.User)
	UserUpdateOut(user *entity.User)
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
	user, err := i.UserRepo.GetByEmailAndPassword(email, password)
	if err != nil {
		return err
	}
	i.UserOutput.UserLoginOut(user)
	return nil
}

func (i *interactor) UserSelect(userId int) error {
	user, err := i.UserRepo.GetById(userId)
	if err != nil {
		return err
	}
	i.UserOutput.UserSelectOut(user)
	return nil
}

func (i *interactor) UserUpdate(userId int, toUpdate *entity.User) error {
	user, err := i.UserRepo.GetById(userId)
	if err != nil {
		return err
	}
	user.Update(toUpdate)
	if err := i.UserRepo.Save(*user); err != nil {
		return err
	}
	i.UserOutput.UserUpdateOut(user)
	return nil
}

func (i *interactor) FavoritesUpdate(userId int, slug string, favortie bool) error {
	panic("implement me")
}
