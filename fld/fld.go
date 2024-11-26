package fld

const FLD_TERMINATOR byte = 0x1E

type Field interface {
	Tag() string
	MARCrepr() []byte
}
