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

type GobBlockEncoder struct{
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	return &GobBlockEncoder{
		w: w,
	}
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {
	return &GobBlockDecoder{
		r: r,
	}
}

type GobBlockDecoder struct{
	r io.Reader
}

func (enc *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(enc.w).Encode(b)
}

func (dec *GobBlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(dec.r).Decode(b)
}

func init(){
	gob.Register(elliptic.P256())
}


