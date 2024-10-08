// Package protoio defines methods for reading and writing to streams of protocol
// buffers.
//
// The official protocol buffer documentation does not define a means of
// serializing a sequence of protocol buffers, so streams written with this
// library may not be readable by other libraries, and vice versa.
//
// For the purposes of this package, a stream of protocol buffers takes the form
// of alternating Uvarints and protocol buffers. The Uvarint specifies the
// length of the succeeding protocol buffer in bytes.
package protoio
