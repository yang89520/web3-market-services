package services

import (
	"context"
	"fmt"
	"github.com/yang89520/web3-merket-services/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/ethereum/go-ethereum/log"
	"github.com/yang89520/web3-merket-services/proto/market"
	"sync/atomic"
)

const MaxRecvMessageSize = 1024 * 1024 * 30000

type MarketRpcConfig struct {
	Host string
	Port int
}
type MarketRpcService struct {
	*MarketRpcConfig

	db *database.DB

	market.UnimplementedMarketServicesServer
	stopped atomic.Bool
}

func NewMarketRpcService(config *MarketRpcConfig) (*MarketRpcService, error) {
	return &MarketRpcService{
		MarketRpcConfig: config,
	}, nil
}

func (ms *MarketRpcService) Start(ctx context.Context) error {
	go func(ms *MarketRpcService) {
		rpcAddr := fmt.Sprintf("%s:%d", ms.Host, ms.Port)
		listener, err := net.Listen("tcp", rpcAddr)
		if err != nil {
			log.Error("failed to listen", "addr", rpcAddr, "err", err)
		}

		//
		gs := grpc.NewServer(
			grpc.MaxRecvMsgSize(MaxRecvMessageSize),
			grpc.ChainUnaryInterceptor(
				nil,
			),
		)

		// 为什么这行拿掉 不能访问该server
		reflection.Register(gs)
		market.RegisterMarketServicesServer(gs, ms)

		log.Info("start market rpc service", "addr", rpcAddr)
		log.Info("grpc info", "addr", listener.Addr())
		if err := gs.Serve(listener); err != nil {
			log.Error("Start grpc server failed", "err", err)
		}
	}(ms)
	return nil
}

func (ms *MarketRpcService) Stop(ctx context.Context) error {
	ms.stopped.Store(true)
	return nil
}

func (ms *MarketRpcService) Stopped() bool {
	return ms.stopped.Load()
}
