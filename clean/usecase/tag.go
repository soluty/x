package usecase

type TagRepo interface {
	GetAll() ([]string, error)
	Add(newTags []string) error
}

type TagInput interface {
	Tags() error
}

type TagOutput interface {
	TagProc([]string)
}


func (i *interactor) Tags() error {
	panic("implement me")
}