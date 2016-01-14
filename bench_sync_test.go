package waterwheel

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func benchmarkSync(wc io.WriteCloser, msg string, b *testing.B) {
	logger := NewSyncLogger(wc, SimpleFormatter, "debug")
	defer logger.Close()

	for n := 0; n < b.N; n++ {
		logger.Debug(msg)
	}
}

func BenchmarkSyncLoggerWithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(buf, smallMsg, b)
}

func BenchmarkSyncLoggerWithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(buf, largeMsg, b)
}

func BenchmarkSyncLoggerSize1WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 1
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), smallMsg, b)
}

func BenchmarkSyncLoggerSize10WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 10
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), smallMsg, b)
}

func BenchmarkSyncLoggerSize100WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 100
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), smallMsg, b)
}

func BenchmarkSyncLoggerSize1000WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 1000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), smallMsg, b)
}

func BenchmarkSyncLoggerSize10000WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 10000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), smallMsg, b)
}

func BenchmarkSyncLoggerSize1WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 1
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), largeMsg, b)
}

func BenchmarkSyncLoggerSize10WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 10
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), largeMsg, b)
}

func BenchmarkSyncLoggerSize100WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 100
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), largeMsg, b)
}

func BenchmarkSyncLoggerSize1000WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 1000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), largeMsg, b)
}

func BenchmarkSyncLoggerSize10000WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 10000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkSync(NewBufferedWriteCloser(size, buf), largeMsg, b)
}

func BenchmarkSyncLoggerWithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkSync(f, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerWithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkSync(f, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize1WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1
	benchmarkSync(NewBufferedWriteCloser(size, f), smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize10WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10
	benchmarkSync(NewBufferedWriteCloser(size, f), smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize100WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 100
	benchmarkSync(NewBufferedWriteCloser(size, f), smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize1000WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1000
	benchmarkSync(NewBufferedWriteCloser(size, f), smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize10000WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10000
	benchmarkSync(NewBufferedWriteCloser(size, f), smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize1WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1
	benchmarkSync(NewBufferedWriteCloser(size, f), largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize10WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10
	benchmarkSync(NewBufferedWriteCloser(size, f), largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize100WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 100
	benchmarkSync(NewBufferedWriteCloser(size, f), largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize1000WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1000
	benchmarkSync(NewBufferedWriteCloser(size, f), largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkSyncLoggerSize10000WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10000
	benchmarkSync(NewBufferedWriteCloser(size, f), largeMsg, b)
	os.Remove("/tmp/bench.log")
}
