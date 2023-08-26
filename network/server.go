package network

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/sharansh123/MyBlockChain/core"
	"github.com/sharansh123/MyBlockChain/crypto"
	"github.com/sharansh123/MyBlockChain/types"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct{
	ID 	string
	Logger log.Logger
	RPCDecodeFunc RPCDecodeFunc
	RPCProcessor RPCProcessor
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime time.Duration
}


type Server struct{
	ServerOpts
	rpcCh chan RPC
	chain *core.Blockchain
	memPool *TxPool
	quitCh chan struct{}
	isValidator bool
}


func NewServer(opts ServerOpts) (*Server, error){
	
	if opts.BlockTime == time.Duration(0){
		opts.BlockTime = defaultBlockTime
	}

	if opts.RPCDecodeFunc == nil{
		opts.RPCDecodeFunc = DefaultPRCDecodeFunc
	}

	if opts.Logger == nil {
		opts.Logger = log.NewLogfmtLogger(os.Stderr)
		opts.Logger = log.With(opts.Logger, "ID", opts.ID)
	}

	chain, err := core.NewBlockChain(genesisBlock())
	if err != nil{
		return nil, err
	}

	s := &Server{
		ServerOpts: opts,
		memPool: NewTxPool(),
		chain: chain,
		rpcCh: make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh: make(chan struct{}, 1),
	}
	//if no process is provided in server opts, default processor is assumed to be server itself.
	if opts.RPCProcessor == nil {
		s.RPCProcessor = s
	}

	if s.isValidator {
		go s.validatorLoop()
	}

	return s, nil
}

func (s *Server) CreateNewBlock() error {
	currentHeader, err := s.chain.GetHeader(s.chain.Height())
	if err != nil{
		return err
	}

	block, err := core.NewBlockFromPrevHeader(currentHeader, nil)
	if err != nil {
		return err
	}

	if err := block.Sign(*s.PrivateKey); err != nil{
		return err
	}

	if err := s.chain.AddBlock(block); err != nil {
		return err
	}

	s.Logger.Log("msg", "Successfully Added Block", "Hash", block.DataHash, "ChainLenght", s.chain.Height())

	return nil
}

func (s *Server) Start(){
	s.initTransports()
free:
	for {
		select{
		case rpc := <-s.rpcCh:
			msg , err := s.ServerOpts.RPCDecodeFunc(rpc)
			if err != nil{
				s.Logger.Log("error", err)
			}
			if err:= s.RPCProcessor.ProcessMessage(msg); err != nil{
				s.Logger.Log("error", err)
			}
		case <-s.quitCh:
			break free
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


	s.Logger.Log(
		"msg", "adding new tx to mempool",
		"hash", tx.Hash(core.TxHasher{}),
		"memPoolLength", s.memPool.Len(),
	)
	
	go s.broadcastTx(tx)

	return s.memPool.Add(tx)
}

func (s *Server) validatorLoop(){
	ticker := time.NewTicker(s.BlockTime)

	s.Logger.Log("msg", "Starting Validator Loop", "BlockTime", s.BlockTime)

	for {
		<- ticker.C
		s.CreateNewBlock()
	}
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

func genesisBlock() *core.Block {
	header:= &core.Header{
		Version: 1,
		DataHash: types.Hash{},
		Timestamp :time.Now().UnixNano(),
		Height: 0,
	}
	return core.NewBlock(header, nil)
}

