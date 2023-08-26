package network

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/crypto"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct{
	RPCDecodeFunc RPCDecodeFunc
	RPCProcessor RPCProcessor
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime time.Duration
}


type Server struct{
	ServerOpts
	rpcCh chan RPC
	memPool *TxPool
	quitCh chan struct{}
	isValidator bool
}


func NewServer(opts ServerOpts) *Server{
	
	if opts.BlockTime == time.Duration(0){
		opts.BlockTime = defaultBlockTime
	}

	if opts.RPCDecodeFunc == nil{
		opts.RPCDecodeFunc = DefaultPRCDecodeFunc
	}

	s := &Server{
		ServerOpts: opts,
		memPool: NewTxPool(),
		rpcCh: make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh: make(chan struct{}, 1),
	}
	//if no process is provided in server opts, default processor is assumed to be server itself.
	if opts.RPCProcessor == nil {
		s.RPCProcessor = s
	}
	return s
}

func (s *Server) CreateNewBlock() error {
	return nil
}

func (s *Server) Start(){
	s.initTransports()
	ticker := time.NewTicker(s.BlockTime)
free:
	for {
		select{
		case rpc := <-s.rpcCh:
			msg , err := s.ServerOpts.RPCDecodeFunc(rpc)
			if err != nil{
				fmt.Errorf(err.Error())
			}
			if err:= s.RPCProcessor.ProcessMessage(msg); err != nil{
				fmt.Errorf(err.Error())
			}
		case <-s.quitCh:
			break free
		case <- ticker.C:
			if s.isValidator{
			fmt.Println("Creating a new block!!")
			s.CreateNewBlock()
			} else {
				fmt.Println("Hello from Node!!")
			}
		}
	}
}

func (s *Server) broadcast(payload []byte) error {
	for _, tr := range s.Transports{
		if err := tr.Broadcast(payload); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) broadcastTx(tx *core.Transaction) error {
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	msg := NewMessage(MessageTypeTx, buf.Bytes())
	return s.broadcast(msg.Bytes())
}

func (s *Server) ProcessMessage(msg *DecodedMessage) error {
	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.ProcessTransaction(t)
	}
	return nil
}

func (s *Server) ProcessTransaction(tx *core.Transaction) error {
	
	if s.memPool.Has(tx.Hash(core.TxHasher{})){
		fmt.Println("Already has the tx")
		return nil
	}
	
	if err := tx.Verify(); err != nil{
		return err
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	fmt.Println("HASH: " + tx.Hash(core.TxHasher{}).String())
	
	go s.broadcastTx(tx)

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

