package loghook

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	LOG_MESSAGE_PREFIX = "gautocloud"
	DEBUG_MODE_ENV_VAR = "GAUTOCLOUD_DEBUG"
	BUF_SIZE           = 35
)

type GautocloudHook struct {
	entries []*logrus.Entry
	nbWrite int
	buf     *bytes.Buffer
}

func NewGautocloudHook(buf *bytes.Buffer) *GautocloudHook {
	return &GautocloudHook{
		entries: make([]*logrus.Entry, 0),
		nbWrite: 0,
		buf:     buf,
	}
}
func (h GautocloudHook) IsDebugMode() bool {
	return os.Getenv(DEBUG_MODE_ENV_VAR) != ""
}
func (h GautocloudHook) TraceDebugMode(entry *logrus.Entry) {
	stdLogger := logrus.StandardLogger()
	currentLvl := stdLogger.Level
	stdLogger.Level = logrus.DebugLevel
	h.Trace(entry)
	stdLogger.Level = currentLvl
}
func (h GautocloudHook) Trace(entry *logrus.Entry) {
	stdLogger := logrus.StandardLogger()
	currentOut := entry.Logger.Out
	entry.Logger.Out = stdLogger.Out
	b, _ := stdLogger.Formatter.Format(entry)
	fmt.Fprint(stdLogger.Out, string(b))
	entry.Logger.Out = currentOut
}
func (h *GautocloudHook) Fire(entry *logrus.Entry) error {
	defer h.buf.Reset()
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
	if h.IsDebugMode() {
		h.TraceDebugMode(entry)
		return nil
	}

	if h.nbWrite == BUF_SIZE {
		h.entries = make([]*logrus.Entry, 0)
		h.nbWrite = 0
	}
	h.entries = append(h.entries, entry)
	h.nbWrite++
	return nil
}
func (h GautocloudHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *GautocloudHook) ShowPreviousLog() {
	newEntries := make([]*logrus.Entry, 0)
	stdLogger := logrus.StandardLogger()
	if len(h.entries) == 0 {
		return
	}
	stdLogger.Warn("")
	stdLogger.Warnf(
		"%s: Show previous log was called, next logs was stored between '%s' and '%s'.",
		LOG_MESSAGE_PREFIX,
		h.entries[0].Time.Format("15:04:05.999999999"),
		h.entries[len(h.entries)-1].Time.Format("15:04:05.999999999"),
	)
	for i := len(h.entries) - 1; i >= 0; i-- {
		entry := h.entries[i]
		if entry.Level > logrus.GetLevel() {
			newEntries = append(newEntries, entry)
			continue
		}
		h.Trace(entry)
	}
	h.entries = newEntries
	h.nbWrite = len(newEntries)
	stdLogger.Warnf(
		"%s: Finished to show previous logs.",
		LOG_MESSAGE_PREFIX,
	)
	stdLogger.Warn("")
}
