package service

import (
	"fmt"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/pb"
	"reflect"
)

// 房间 黑板 游戏 玩家 游戏中的一方Side

// 简单工厂模式, 创建一个Game对象
var gameTypes = map[pb.GameType]reflect.Type{}
var sideTypes = map[pb.GameType]reflect.Type{}

func Register(head pb.GameType, gameType reflect.Type, sideType reflect.Type) {
	// 确保注册进来的类型正确
	_ = reflect.New(gameType).Interface().(entity.Game)
	_ = reflect.New(sideType).Interface().(entity.Side)
	gameTypes[head] = gameType
	sideTypes[head] = sideType
}

func CreateGame(gameType pb.GameType) (entity.Game, error) {
	tpy, ok := gameTypes[gameType]
	if !ok {
		return nil, fmt.Errorf("CreateGame type %v not exist", gameType)
	}
	game, ok := reflect.New(tpy).Interface().(entity.Game)
	if !ok {
		return nil, fmt.Errorf("CreateGame type %v is not interface Game, it is %v", gameType, tpy)
	}
	return game, nil
}

func CreateSide(gameType pb.GameType) (entity.Side, error) {
	tpy, ok := sideTypes[gameType]
	if !ok {
		return nil, fmt.Errorf("CreateSide type %v not exist", gameType)
	}
	side, ok := reflect.New(tpy).Interface().(entity.Side)
	if !ok {
		return nil, fmt.Errorf("CreateSide type %v is not interface Side, it is %v", gameType, tpy)
	}
	return side, nil
}