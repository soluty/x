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

type Logger interface {
	Log(...interface{})
}

// create
// update
// delete
// select

type interactor struct {
	Logger           Logger

	ProfileOutput ProfileOutput
	ArticleOutput ArticleOutput
	TagOutput TagOutput
	UserOutput UserOutput
	CommentOutput CommentOutput

	TagRepo TagRepo
	UserRepo UserRepo
	CommentRepo CommentRepo
	ArticleRepo ArticleRepo
}



