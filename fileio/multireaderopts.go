package fileio

import "io/fs"

type MultiReaderOpt func(*MultiReader)

func OptFilesystem(fs fs.FS) MultiReaderOpt {
	return func(mr *MultiReader) {
		mr.filesystem = fs
	}
}
