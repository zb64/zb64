package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

type (
	newZipWriterFunc func(io.Writer) io.WriteCloser
	newZipReaderFunc func(io.Reader) (io.ReadCloser, error)
)

type zipper struct {
	name      string
	newWriter newZipWriterFunc
	newReader newZipReaderFunc
}

func newZipper(name string, writer newZipWriterFunc, reader newZipReaderFunc) *zipper {
	return &zipper{
		name:      name,
		newWriter: writer,
		newReader: reader,
	}
}

func (z *zipper) Name() string {
	return z.name
}

func (z *zipper) Encode(data []byte) (result string, err error) {
	var buf bytes.Buffer
	w := z.newWriter(&buf)
	if _, err = w.Write(data); err != nil {
		return
	}
	if err = w.Close(); err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

func (z *zipper) Decode(raw string) (result []byte, err error) {
	var (
		data   []byte
		outBuf bytes.Buffer
		r      io.ReadCloser
	)
	if data, err = base64.StdEncoding.DecodeString(raw); err != nil {
		return
	}

	if r, err = z.newReader(bytes.NewReader(data)); err != nil {
		return
	}
	if _, err = io.Copy(&outBuf, r); err != nil {
		return
	}
	if err = r.Close(); err != nil {
		return
	}
	result = outBuf.Bytes()
	return
}

var (
	base64Deflate = newZipper("flate",
		func(buf io.Writer) io.WriteCloser {
			w, _ := flate.NewWriter(buf, flate.BestCompression)
			return w
		}, func(buf io.Reader) (io.ReadCloser, error) {
			return flate.NewReader(buf), nil
		})
)
