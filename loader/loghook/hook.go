package loghook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"strings"
)

const (
	LOG_MESSAGE_PREFIX = "gautocloud"
	BUF_SIZE           = 50
)

type GautocloudHook struct {
	entries []*logrus.Entry
	nbWrite int
}

func NewGautocloudHook() *GautocloudHook {
	return &GautocloudHook{
		entries: make([]*logrus.Entry, BUF_SIZE),
		nbWrite: 0,
	}
}

func (h *GautocloudHook) Fire(entry *logrus.Entry) error {
	stdLogger := logrus.StandardLogger()
	currentOut := entry.Logger.Out
	entry.Logger.Out = stdLogger.Out
	b, err := stdLogger.Formatter.Format(entry)
	if err != nil {
		return err
	}
	entry.Logger.Out = currentOut
	line := string(b)
	if !strings.HasPrefix(entry.Message, LOG_MESSAGE_PREFIX) {
		fmt.Fprint(stdLogger.Out, line)
		return nil
	}

	currentLvl := logrus.GetLevel()
	if entry.Level <= currentLvl {
		fmt.Fprint(stdLogger.Out, line)
		return nil
	}
	if h.nbWrite == BUF_SIZE {
		h.entries = make([]*logrus.Entry, BUF_SIZE)
		h.nbWrite = 0
	}
	h.entries[h.nbWrite] = entry
	h.nbWrite++
	return nil
}
func (h GautocloudHook) checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
func (h GautocloudHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h GautocloudHook) ShowPreviousLog() {
	stdLogger := logrus.StandardLogger()
	for i := h.nbWrite - 1; i >= 0; i-- {
		entry := h.entries[i]
		if entry.Level > logrus.GetLevel() {
			continue
		}
		entry.Logger.Out = stdLogger.Out
		b, _ := stdLogger.Formatter.Format(entry)
		fmt.Fprint(stdLogger.Out, string(b))

	}
}
