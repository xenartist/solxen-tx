package logic

import (
	"context"
	"solxen-tx/internal/svc"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type Producer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	all            int
	mux            sync.RWMutex
	ProgramIdMiner solana.PublicKeySlice
}

func NewProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Producer {
	return &Producer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		all:    0,
		mux:    sync.RWMutex{},
		ProgramIdMiner: solana.PublicKeySlice{
			solana.MustPublicKeyFromBase58("H4Nk2SDQncEv5Cc6GAbradB4WLrHn7pi9VByFL9zYZcA"),
			solana.MustPublicKeyFromBase58("58UESDt7K7GqutuHBYRuskSgX6XoFe8HXjwrAtyeDULM"),
			solana.MustPublicKeyFromBase58("B1Dw79PE8dzpHPKjiQ8HYUBZ995hL1U32bUTRdNVtRbr"),
			solana.MustPublicKeyFromBase58("7ukQWD7UqoC61eATrBMrdfMrJMUuY1wuPTk4m4noZpsH"),
		},
	}
}

func (l *Producer) Start() {
	logx.Infof("start  miner")

	// var subscription pb.SubscribeRequest
	// subscription.Blocks = make(map[string]*pb.SubscribeRequestFilterBlocks)
	// subscription.Blocks["blocks"] = &pb.SubscribeRequestFilterBlocks{}
	//
	// subscriptionJson, err := json.Marshal(&subscription)
	// if err != nil {
	// 	logx.Errorf("Failed to marshal subscription request: %v", subscriptionJson)
	// }
	// logx.Infof("Subscription request: %s", string(subscriptionJson))
	//
	// stream, err := l.svcCtx.GrpcCli.Subscribe(context.Background())
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// err = stream.Send(&subscription)
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }

	// for {
	// 	resp, err := stream.Recv()
	// 	// timestamp := time.Now().UnixNano()
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Error occurred in receiving update: %v", err)
	// 	}
	// 	block := resp.GetBlock()
	// 	// log.Printf("timestamp: %v block:%v %v %v", timestamp, block.GetBlockhash(), block.GetSlot(), block.GetBlockHeight())
	// 	Blockhash := block.GetBlockhash()
	// 	slot := block.GetSlot()
	// 	if Blockhash == "" || slot == 0 {
	// 		continue
	// 	}
	//
	// 	err = l.Mint()
	// 	if err != nil {
	// 		logx.Errorf("Mint err:%v", err)
	// 		continue
	// 	}
	//
	// }

	for {
		// 1. CheckAddressBalance
		err := l.CheckAddressBalance()
		if err != nil {
			logx.Errorf("%v", err)
			return
		}
		// todo 2.QueryNetWorkGas
		// err = l.QueryNetWorkGas()
		// if err != nil {
		// 	logx.Errorf("%v", err)
		// 	return
		// }

		// 3.Miner

		err = l.Miner()
		if err != nil {
			logx.Errorf("Mint err:%v", err)
			continue
		}

		time.Sleep(time.Duration(l.svcCtx.Config.Sol.Time) * time.Millisecond)
	}

}

func (l *Producer) Stop() {
	logx.Infof("stop Producer \n")
}
