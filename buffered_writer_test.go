package waterwheel

import (
	"bytes"
	"io"
	"testing"
)

type NopCloserBuf struct {
	w io.Writer
}

func (b *NopCloserBuf) Write(p []byte) (int, error) {
	return b.w.Write(p)
}

func (b *NopCloserBuf) Close() error {
	return nil
}

func TestBufferedWriteCloser(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	b := NopCloserBuf{buf}
	wc := NewBufferedWriteCloser(2, &b)
	wc.Write([]byte("1"))
	if len(buf.Bytes()) != 0 {
		t.Fatalf(
			"Writing to BufferedWriteCloser should be buffered. expecting 0 len, got %d len",
			len(buf.Bytes()),
		)
	}

	wc.Write([]byte("2"))
	if len(buf.Bytes()) != 0 {
		t.Fatalf(
			"Writing to BufferedWriteCloser should be buffered. expecting 0 len, got %d len",
			len(buf.Bytes()),
		)
	}

	wc.Write([]byte("3"))
	if len(buf.Bytes()) != 2 {
		t.Fatalf(
			"BufferedWriteCloser should be flushing the buffer over fixed size. expecting 2 len, got %d len",
			len(buf.Bytes()),
		)
	}

	wc.Close()
	if len(buf.Bytes()) != 3 {
		t.Fatalf(
			"BufferedWriteCloser should be flushing the buffer over when closing. expecting 3 len, got %d len",
			len(buf.Bytes()),
		)
	}
}
