package usecase

import "github.com/soluty/x/clean/entity"

type Profile struct {
	entity.User
	isFollow bool
}

type ProfileInput interface {
	ProfileGet(requestingUserName, userName string) error
	ProfileUpdateFollow(loggedInUsername, username string, follow bool) error
}

type ProfileOutput interface {
	ProfileGet(profile Profile)
	ProfileUpdateFollow(profile Profile)
}

func (i *interactor) ProfileGet(requestingUserName, userName string) error {
	panic("implement me")
}

func (i *interactor) ProfileUpdateFollow(loggedInUsername, username string, follow bool) error {
	panic("implement me")
}
