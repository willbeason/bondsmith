package fileio

import (
	"fmt"
	"io"
	"os"
)

// Close closes closer. If this returns an error, calls each handle function on the resulting error.
func Close(closer io.Closer, handles ...func(error)) {
	err := closer.Close()
	if err != nil {
		for _, handle := range handles {
			handle(err)
		}
	}
}

func Print(w io.Writer, err error) {
	_, _ = fmt.Fprintln(w, err)
}

func PrintStdErr(err error) {
	Print(os.Stderr, err)
}
