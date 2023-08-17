package network

import (
	"fmt"
	"time"

	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/crypto"
)

type ServerOpts struct{
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime time.Duration
}


type Server struct{
	ServerOpts
	rpcCh chan RPC
	memPool *TxPool
	blockTime time.Duration
	quitCh chan struct{}
	isValidator bool
}


func NewServer(opts ServerOpts) *Server{
	return &Server{
		ServerOpts: opts,
		memPool: NewTxPool(),
		blockTime: opts.BlockTime,
		rpcCh: make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh: make(chan struct{}, 1),
	}
}

func (s *Server) CreateNewBlock() error {
	return nil
}

func (s *Server) Start(){
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)
free:
	for {
		select{
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <- ticker.C:
			if s.isValidator{
			fmt.Println("Creating a new block!!")
			s.CreateNewBlock()
			}
		}
	}
}

func (s *Server) handleTransactions( tx *core.Transaction) error {
	if err := tx.Verify(); err != nil{
		return err
	}
	if s.memPool.Has(tx.Hash(core.TxHasher{})){
		fmt.Println("Already has the tx")
		return nil
	}

	return s.memPool.Add(tx)
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports{
		go func(tr Transport){
			for rpc := range tr.Consume(){
				s.rpcCh <- rpc
			}
		}(tr)
	}
}

