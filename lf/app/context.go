package app

import (
	"context"
	"time"
)

// 带error的context实现
type mContext struct {
	c context.Context
}

func (mContext) Deadline() (deadline time.Time, ok bool) {

}

func (mContext) Done() <-chan struct{} {

}

func (mContext) Err() error {

}

func (mContext) Value(key interface{}) interface{} {

}

var _ context.Context = &mContext{}
