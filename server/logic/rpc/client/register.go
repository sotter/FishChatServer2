package client

import (
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/protocol/rpc"
	"github.com/oikomi/FishChatServer2/server/logic/conf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RegisterRPCCli struct {
	conn *grpc.ClientConn
}

func NewRegisterRPCCli() (registerRPCCli *RegisterRPCCli, err error) {
	conn, err := grpc.Dial(conf.Conf.RPCClient.RegisterClient.Addr, grpc.WithInsecure())
	if err != nil {
		glog.Error(err)
		return
	}
	registerRPCCli = &RegisterRPCCli{
		conn: conn,
	}
	return
}

func (routerRPCCli *RegisterRPCCli) Online(ctx context.Context, uid int64) (res *rpc.RGOnlineRes, err error) {
	r := rpc.NewRegisterServerRPCClient(routerRPCCli.conn)
	if res, err = r.Online(ctx, &rpc.RGOnlineReq{UID: uid}); err != nil {
		glog.Error(err)
	}
	return
}

func (routerRPCCli *RegisterRPCCli) Ping(ctx context.Context, uid int64) (res *rpc.RGPingRes, err error) {
	r := rpc.NewRegisterServerRPCClient(routerRPCCli.conn)
	if res, err = r.Ping(ctx, &rpc.RGPingReq{UID: uid}); err != nil {
		glog.Error(err)
	}
	return
}
