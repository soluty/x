package usecase

// repo validator resp

// interactor 实现 Input接口      controller 调用 interactor
// interactor 调用 OutPut接口     view 实现 output   repo和

func init()  {
	var _ ProfileInput = &interactor{}
	var _ CommentInput = &interactor{}
	var _ ArticleInput = &interactor{}
	var _ TagInput = &interactor{}
	var _ UserInput = &interactor{}
}


type interactor struct {
	ProfileOutput ProfileOutput
}

type Logger interface {
	Log(...interface{})
}


