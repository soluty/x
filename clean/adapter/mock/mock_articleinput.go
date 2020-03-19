// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/soluty/x/clean/usecase (interfaces: ArticleInput)

package mock

import (
	pegomock "github.com/petergtz/pegomock"
	entity "github.com/soluty/x/clean/entity"
	"reflect"
	"time"
)

type MockArticleInput struct {
	fail func(message string, callerSkip ...int)
}

func NewMockArticleInput(options ...pegomock.Option) *MockArticleInput {
	mock := &MockArticleInput{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockArticleInput) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockArticleInput) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockArticleInput) ArticlesFeed(userId int, limit int, offset int) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{userId, limit, offset}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticlesFeed", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) ArticlesSelect(userId int, limit int, offset int, filters []entity.ArticleFilter) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{userId, limit, offset, filters}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticlesSelect", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) ArticleSelect(slug string, username int) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{slug, username}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticleSelect", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) ArticleCreate(userId int, article entity.Article) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{userId, article}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticleCreate", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) ArticleUpdate(userId int, slug string, newArticle *entity.Article) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{userId, slug, newArticle}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticleUpdate", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) ArticleDelete(userId int, slug string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockArticleInput().")
	}
	params := []pegomock.Param{userId, slug}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ArticleDelete", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockArticleInput) VerifyWasCalledOnce() *VerifierMockArticleInput {
	return &VerifierMockArticleInput{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockArticleInput) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockArticleInput {
	return &VerifierMockArticleInput{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockArticleInput) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockArticleInput {
	return &VerifierMockArticleInput{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockArticleInput) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockArticleInput {
	return &VerifierMockArticleInput{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockArticleInput struct {
	mock                   *MockArticleInput
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockArticleInput) ArticlesFeed(userId int, limit int, offset int) *MockArticleInput_ArticlesFeed_OngoingVerification {
	params := []pegomock.Param{userId, limit, offset}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticlesFeed", params, verifier.timeout)
	return &MockArticleInput_ArticlesFeed_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticlesFeed_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticlesFeed_OngoingVerification) GetCapturedArguments() (int, int, int) {
	userId, limit, offset := c.GetAllCapturedArguments()
	return userId[len(userId)-1], limit[len(limit)-1], offset[len(offset)-1]
}

func (c *MockArticleInput_ArticlesFeed_OngoingVerification) GetAllCapturedArguments() (_param0 []int, _param1 []int, _param2 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]int, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockArticleInput) ArticlesSelect(userId int, limit int, offset int, filters []entity.ArticleFilter) *MockArticleInput_ArticlesSelect_OngoingVerification {
	params := []pegomock.Param{userId, limit, offset, filters}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticlesSelect", params, verifier.timeout)
	return &MockArticleInput_ArticlesSelect_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticlesSelect_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticlesSelect_OngoingVerification) GetCapturedArguments() (int, int, int, []entity.ArticleFilter) {
	userId, limit, offset, filters := c.GetAllCapturedArguments()
	return userId[len(userId)-1], limit[len(limit)-1], offset[len(offset)-1], filters[len(filters)-1]
}

func (c *MockArticleInput_ArticlesSelect_OngoingVerification) GetAllCapturedArguments() (_param0 []int, _param1 []int, _param2 []int, _param3 [][]entity.ArticleFilter) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]int, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
		_param3 = make([][]entity.ArticleFilter, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.([]entity.ArticleFilter)
		}
	}
	return
}

func (verifier *VerifierMockArticleInput) ArticleSelect(slug string, username int) *MockArticleInput_ArticleSelect_OngoingVerification {
	params := []pegomock.Param{slug, username}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticleSelect", params, verifier.timeout)
	return &MockArticleInput_ArticleSelect_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticleSelect_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticleSelect_OngoingVerification) GetCapturedArguments() (string, int) {
	slug, username := c.GetAllCapturedArguments()
	return slug[len(slug)-1], username[len(username)-1]
}

func (c *MockArticleInput_ArticleSelect_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockArticleInput) ArticleCreate(userId int, article entity.Article) *MockArticleInput_ArticleCreate_OngoingVerification {
	params := []pegomock.Param{userId, article}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticleCreate", params, verifier.timeout)
	return &MockArticleInput_ArticleCreate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticleCreate_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticleCreate_OngoingVerification) GetCapturedArguments() (int, entity.Article) {
	userId, article := c.GetAllCapturedArguments()
	return userId[len(userId)-1], article[len(article)-1]
}

func (c *MockArticleInput_ArticleCreate_OngoingVerification) GetAllCapturedArguments() (_param0 []int, _param1 []entity.Article) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
		_param1 = make([]entity.Article, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(entity.Article)
		}
	}
	return
}

func (verifier *VerifierMockArticleInput) ArticleUpdate(userId int, slug string, newArticle *entity.Article) *MockArticleInput_ArticleUpdate_OngoingVerification {
	params := []pegomock.Param{userId, slug, newArticle}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticleUpdate", params, verifier.timeout)
	return &MockArticleInput_ArticleUpdate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticleUpdate_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticleUpdate_OngoingVerification) GetCapturedArguments() (int, string, *entity.Article) {
	userId, slug, newArticle := c.GetAllCapturedArguments()
	return userId[len(userId)-1], slug[len(slug)-1], newArticle[len(newArticle)-1]
}

func (c *MockArticleInput_ArticleUpdate_OngoingVerification) GetAllCapturedArguments() (_param0 []int, _param1 []string, _param2 []*entity.Article) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]*entity.Article, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(*entity.Article)
		}
	}
	return
}

func (verifier *VerifierMockArticleInput) ArticleDelete(userId int, slug string) *MockArticleInput_ArticleDelete_OngoingVerification {
	params := []pegomock.Param{userId, slug}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ArticleDelete", params, verifier.timeout)
	return &MockArticleInput_ArticleDelete_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockArticleInput_ArticleDelete_OngoingVerification struct {
	mock              *MockArticleInput
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockArticleInput_ArticleDelete_OngoingVerification) GetCapturedArguments() (int, string) {
	userId, slug := c.GetAllCapturedArguments()
	return userId[len(userId)-1], slug[len(slug)-1]
}

func (c *MockArticleInput_ArticleDelete_OngoingVerification) GetAllCapturedArguments() (_param0 []int, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]int, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(int)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}
