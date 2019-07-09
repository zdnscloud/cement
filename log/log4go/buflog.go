package log4go

type BufLogWriter struct {
	format       string
	rec          chan *LogRecord
	bufCloseSync chan struct{}
}

func NewBufLogWriter(fmt string, bufCh chan string) *BufLogWriter {
	w := &BufLogWriter{
		format:       fmt,
		rec:          make(chan *LogRecord, LogBufferLength),
		bufCloseSync: make(chan struct{}),
	}
	go w.run(bufCh)
	return w
}

func (w *BufLogWriter) run(out chan string) {
	for {
		rec, ok := <-w.rec
		if !ok {
			close(w.bufCloseSync)
			return
		}

		out <- FormatLogRecord(w.format, rec)

	}
}

func (w *BufLogWriter) LogWrite(rec *LogRecord) {
	w.rec <- rec
}

func (w *BufLogWriter) Close() {
	close(w.rec)
	<-w.bufCloseSync
}

func NewBufLogger(buf chan string, lvl level, fmt string) Logger {
	return Logger{
		"buf": &Filter{lvl, NewBufLogWriter(fmt, buf)},
	}
}
