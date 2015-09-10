package loader

import (
	"encoding/binary"
)

type Segment struct {
	Addr uint64
	Data []byte
}

type Loader interface {
	Arch() string
	Bits() int
	ByteOrder() binary.ByteOrder
	OS() string
	Entry() uint64
	Type() int
	Interp() string
	Header() ([]byte, int)
	Symbolicate(addr uint64) (string, error)
	Segments() ([]Segment, error)
	DataSegment() (uint64, uint64)
}

type LoaderHeader struct {
	arch      string
	bits      int
	byteOrder binary.ByteOrder
	os        string
	entry     uint64
}

func (l *LoaderHeader) Arch() string {
	return l.arch
}

func (l *LoaderHeader) Bits() int {
	return l.bits
}

func (l *LoaderHeader) ByteOrder() binary.ByteOrder {
	if l.byteOrder == nil {
		return binary.LittleEndian
	}
	return l.byteOrder
}

func (l *LoaderHeader) OS() string {
	return l.os
}

func (l *LoaderHeader) Entry() uint64 {
	return l.entry
}
