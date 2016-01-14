package waterwheel

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var smallMsg = "this is a tiny message."
var largeMsg = `
I would fain prove so. But what might you think,
When I had seen this hot love on the wing
(As I perceived it, I must tell you that,
Before my daughter told me), what might you,
Or my dear Majesty your queen here, think,
If I had played the desk or table-book Or given my heart a text from the Folio not found in the Second Quartowinking,text from the Folio not found in the Second Quarto mute and dumb,
Or looked upon this love with idle sight?
What might you think? No, I went round to work,
And my young mistress thus I did bespeak:
“Lord Hamlet is a prince, out of thy star.
This must not be.” And then I prescripts gave her,
That she should lock herself from text from the Folio not found in the Second Quartohistext from the Folio not found in the Second Quarto resort,
Admit no messengers, receive no tokens;
Which done, she took the fruits of my advice,
And he, repelled (a short tale to make),
Fell into a sadness, then into a fast,
Thence to a watch, thence into a weakness,
Thence to text from the Folio not found in the Second Quartoatext from the Folio not found in the Second Quarto lightness, and, by this declension,
Into the madness wherein now he raves
And all we mourn for.
` // from Hamlet

func BenchmarkStandardLogSmallMessageToDiscard(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for n := 0; n < b.N; n++ {
		log.Println(smallMsg)
	}
}

func BenchmarkStandardLogLargeMessageToDiscard(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for n := 0; n < b.N; n++ {
		log.Println(largeMsg)
	}
}

func BenchmarkStandardLogSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	for n := 0; n < b.N; n++ {
		log.Println(smallMsg)
	}
	os.Remove("/tmp/bench.log")
}

func BenchmarkStandardLogLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	for n := 0; n < b.N; n++ {
		log.Println(largeMsg)
	}
	os.Remove("/tmp/bench.log")
}

func benchmarkAsync(wc io.WriteCloser, size int, msg string, b *testing.B) {
	logger := NewAsyncLogger(wc, SimpleFormatter, size, "debug")
	defer logger.Close()

	for n := 0; n < b.N; n++ {
		logger.Debug(msg)
	}
}

func BenchmarkAsyncLoggerSize1WithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 1, smallMsg, b)
}

func BenchmarkAsyncLoggerSize10WithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 10, smallMsg, b)
}

func BenchmarkAsyncLoggerSize100WithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 100, smallMsg, b)
}

func BenchmarkAsyncLoggerSize1000WithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 1000, smallMsg, b)
}

func BenchmarkAsyncLoggerSize10000WithoutBufferedWriterSmallMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 10000, smallMsg, b)
}

func BenchmarkAsyncLoggerSize1WithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 1, largeMsg, b)
}

func BenchmarkAsyncLoggerSize10WithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 10, largeMsg, b)
}

func BenchmarkAsyncLoggerSize100WithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 100, largeMsg, b)
}

func BenchmarkAsyncLoggerSize1000WithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 1000, largeMsg, b)
}

func BenchmarkAsyncLoggerSize10000WithoutBufferedWriterLargeMessageToDiscard(b *testing.B) {
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(buf, 10000, largeMsg, b)
}

func BenchmarkAsyncLoggerSize1WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 1
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, smallMsg, b)
}

func BenchmarkAsyncLoggerSize10WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 10
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, smallMsg, b)
}

func BenchmarkAsyncLoggerSize100WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 100
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, smallMsg, b)
}

func BenchmarkAsyncLoggerSize1000WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 1000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, smallMsg, b)
}

func BenchmarkAsyncLoggerSize10000WithBufferedWriterSmallMessageToDiscard(b *testing.B) {
	size := 10000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, smallMsg, b)
}

func BenchmarkAsyncLoggerSize1WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 1
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, largeMsg, b)
}

func BenchmarkAsyncLoggerSize10WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 10
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, largeMsg, b)
}

func BenchmarkAsyncLoggerSize100WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 100
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, largeMsg, b)
}

func BenchmarkAsyncLoggerSize1000WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 1000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, largeMsg, b)
}

func BenchmarkAsyncLoggerSize10000WithBufferedWriterLargeMessageToDiscard(b *testing.B) {
	size := 10000
	buf := &NopCloserBuf{ioutil.Discard}
	benchmarkAsync(NewBufferedWriteCloser(size, buf), size, largeMsg, b)
}

func BenchmarkAsyncLoggerSize1WithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 1, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10WithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 10, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize100WithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 100, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1000WithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 1000, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10000WithoutBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 10000, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1WithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 1, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10WithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 10, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize100WithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 100, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1000WithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 1000, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10000WithoutBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	benchmarkAsync(f, 10000, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize100WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 100
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1000WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1000
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10000WithBufferedWriterSmallMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10000
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, smallMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize100WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 100
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize1000WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 1000
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, largeMsg, b)
	os.Remove("/tmp/bench.log")
}

func BenchmarkAsyncLoggerSize10000WithBufferedWriterLargeMessageToFile(b *testing.B) {
	f, err := os.OpenFile("/tmp/bench.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	size := 10000
	benchmarkAsync(NewBufferedWriteCloser(size, f), size, largeMsg, b)
	os.Remove("/tmp/bench.log")
}
