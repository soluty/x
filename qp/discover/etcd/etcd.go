package etcd

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/soluty/x/qp"
	"strings"
)

type DiscoverEtcd struct {
	client *clientv3.Client
}

func (e *DiscoverEtcd) GetPlayerGameUrl(uid string) (string, error) {
	panic("implement me")
}

var _ qp.Discover = new(DiscoverEtcd)

func New(url string) *DiscoverEtcd {
	etcd := &DiscoverEtcd{}
	client, err := clientv3.NewFromURL(url)
	if err != nil {
		panic(err)
	}
	etcd.client = client
	return etcd
}

func (e *DiscoverEtcd) Get(key string) string {
	rsp, err := e.client.Get(context.Background(), key)
	if err != nil {
		panic(err)
		return ""
	}
	if len(rsp.Kvs) == 0 {
		return ""
	}
	return string(rsp.Kvs[0].Value)
}

func (e *DiscoverEtcd) OnStop() {
	e.client.Close()
}

// 这个一定获取到的是内网ip
func (e *DiscoverEtcd) GetPlayerGateUrl(uid string) (string, error) {
	// /connection/${uid} -> gateId   代表玩家当前所连接的网关id
	// /gate/url/${gateId} -> 外网 内网  网关的外网和内网地址
	var ip string
	txRsp, err := concurrency.NewSTM(e.client, func(stm concurrency.STM) error {
		gateId := stm.Get(fmt.Sprintf("/connection/%s", uid))
		gateIps := stm.Get(fmt.Sprintf("/gate/url/%s", gateId))
		if gateIps == "" {
			return errors.New("GetPlayerGateUrl ip err")
		}
		ips := strings.Split(gateIps, " ")
		if len(ips) == 1 {
			ip = ips[0]
		}
		ip = ips[1]
		return nil
	}, concurrency.WithIsolation(concurrency.ReadCommitted))
	if err != nil {
		return "", fmt.Errorf("GetPlayerGateUrl err: %w", err)
	}
	if !txRsp.Succeeded {
		return "", errors.New("GetPlayerGateUrl stm err")
	}
	return ip, nil
}

// 玩家登陆验证通过以后，需要根据网关负载, 给客户端选择一个负载最低的服务器(max-current) > 0 并且最低
func (e *DiscoverEtcd) AssignPlayerGate(uid string) (string, error) {
	// /payload/gate/${gateId} -> current:max
	// /payload/game/${gameId} -> current:max
}

func (e *DiscoverEtcd) AckPlayerGate(uid string) {

}
