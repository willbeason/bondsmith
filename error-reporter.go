package bondsmith_io

import "io"

type ErrorReporter interface {
	Report(error)
}

type ErrorChanReporter struct {
	errs chan error
}

func NewErrorChanReporter(errs chan error) *ErrorChanReporter {
	return &ErrorChanReporter{errs: errs}
}

func (r *ErrorChanReporter) Report(err error) {
	r.errs <- err
}

type ErrorWriter struct {
	w io.Writer
}

func NEwErrorWriter(w io.Writer) *ErrorWriter {
	return &ErrorWriter{w: w}
}

func (r *ErrorWriter) Report(err error) {
	_, _ = io.WriteString(r.w, err.Error())
}

