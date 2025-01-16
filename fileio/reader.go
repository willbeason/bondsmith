package fileio

type Reader interface {
	Read([]byte) (int, error)
	ReadByte() (byte, error)
}
