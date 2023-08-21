package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)


type Encoder[T any] interface{
	Encode(T) error
}

type Decoder[T any] interface{
	Decode(T) error
}

type GobTxEncoder struct{
	w io.Writer
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder{
	gob.Register(elliptic.P256())
	return &GobTxEncoder{w}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error{
	enc := gob.NewEncoder(e.w)
	return enc.Encode(tx)

}

type GobTxDecoder struct{
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder{
	gob.Register(elliptic.P256())
	return &GobTxDecoder{r}
}

func (e *GobTxDecoder) Decode(tx *Transaction) error{
	enc := gob.NewDecoder(e.r)
	return enc.Decode(tx)

}