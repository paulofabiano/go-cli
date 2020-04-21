// Generator
// $ go get -u github.com/josharian/impl
// $ go generate ./mylib >> mylib/mylib_gen.go

package mylib

//go:generate impl 'f *File' io.ReadWriteCloser

type File struct{}
