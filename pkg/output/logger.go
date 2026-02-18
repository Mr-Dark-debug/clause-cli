package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/clause-cli/clause/pkg/utils"
)

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	// LevelDebug is for detailed debugging information.
	LevelDebug LogLevel = iota
	// LevelInfo is for general informational messages.
	LevelInfo
	// LevelWarn is for warning messages.
	LevelWarn
	// LevelError is for error messages.
	LevelError
	// LevelFatal is for fatal errors that should terminate the program.
	LevelFatal
)

// String returns the string representation of a log level.
func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger provides structured logging capabilities.
type Logger struct {
	mu          sync.Mutex
	level       LogLevel
	writer      io.Writer
	timeFormat  string
	showTime    bool
	showLevel   bool
	colorOutput bool
	prefix      string

	// Styles for different levels
	debugStyle, infoStyle, warnStyle, errorStyle, fatalStyle string
}

// LoggerOption is a functional option for configuring a Logger.
type LoggerOption func(*Logger)

// WithLevel sets the minimum log level.
func WithLevel(level LogLevel) LoggerOption {
	return func(l *Logger) {
		l.level = level
	}
}

// WithWriter sets the output writer.
func WithWriter(w io.Writer) LoggerOption {
	return func(l *Logger) {
		l.writer = w
	}
}

// WithTimeFormat sets the time format.
func WithTimeFormat(format string) LoggerOption {
	return func(l *Logger) {
		l.timeFormat = format
		l.showTime = true
	}
}

// WithShowTime enables or disables time display.
func WithShowTime(show bool) LoggerOption {
	return func(l *Logger) {
		l.showTime = show
	}
}

// WithShowLevel enables or disables level display.
func WithShowLevel(show bool) LoggerOption {
	return func(l *Logger) {
		l.showLevel = show
	}
}

// WithColor enables or disables colored output.
func WithColor(enabled bool) LoggerOption {
	return func(l *Logger) {
		l.colorOutput = enabled
	}
}

// WithPrefix sets a prefix for all log messages.
func WithPrefix(prefix string) LoggerOption {
	return func(l *Logger) {
		l.prefix = prefix
	}
}

// NewLogger creates a new logger with the given options.
func NewLogger(opts ...LoggerOption) *Logger {
	l := &Logger{
		level:       LevelInfo,
		writer:      os.Stderr,
		timeFormat:  "15:04:05",
		showTime:    true,
		showLevel:   true,
		colorOutput: utils.IsTerminal(),
	}

	for _, opt := range opts {
		opt(l)
	}

	// Set up colors if enabled
	if l.colorOutput {
		l.debugStyle = "\x1b[36m" // Cyan
		l.infoStyle = "\x1b[32m"  // Green
		l.warnStyle = "\x1b[33m"  // Yellow
		l.errorStyle = "\x1b[31m" // Red
		l.fatalStyle = "\x1b[35m" // Magenta
	}

	return l
}

// DefaultLogger is the default logger instance.
var DefaultLogger = NewLogger()

// log writes a log message at the specified level.
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var parts []string

	// Add timestamp
	if l.showTime {
		parts = append(parts, time.Now().Format(l.timeFormat))
	}

	// Add level
	if l.showLevel {
		levelStr := level.String()
		if l.colorOutput {
			levelStr = l.colorizeLevel(level, levelStr)
		}
		parts = append(parts, levelStr)
	}

	// Add prefix
	if l.prefix != "" {
		parts = append(parts, l.prefix)
	}

	// Add message
	msg := fmt.Sprintf(format, args...)
	parts = append(parts, msg)

	// Write the log line
	output := strings.Join(parts, " ")
	fmt.Fprintln(l.writer, output)

	// Handle fatal
	if level == LevelFatal {
		os.Exit(1)
	}
}

// colorizeLevel adds ANSI color codes to a level string.
func (l *Logger) colorizeLevel(level LogLevel, s string) string {
	var colorCode string
	switch level {
	case LevelDebug:
		colorCode = l.debugStyle
	case LevelInfo:
		colorCode = l.infoStyle
	case LevelWarn:
		colorCode = l.warnStyle
	case LevelError:
		colorCode = l.errorStyle
	case LevelFatal:
		colorCode = l.fatalStyle
	}

	reset := "\x1b[0m"
	return colorCode + s + reset
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

// Info logs an info message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(LevelInfo, format, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

// Error logs an error message.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

// Fatal logs a fatal message and exits.
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(LevelFatal, format, args...)
}

// SetLevel sets the minimum log level.
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// SetWriter sets the output writer.
func (l *Logger) SetWriter(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writer = w
}

// SetPrefix sets the log prefix.
func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// Debugf logs a formatted debug message.
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debug(format, args...)
}

// Infof logs a formatted info message.
func Infof(format string, args ...interface{}) {
	DefaultLogger.Info(format, args...)
}

// Warnf logs a formatted warning message.
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warn(format, args...)
}

// Errorf logs a formatted error message.
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Error(format, args...)
}

// Fatalf logs a formatted fatal message and exits.
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatal(format, args...)
}

// Fields represents key-value pairs for structured logging.
type Fields map[string]interface{}

// FieldLogger provides logging with fields.
type FieldLogger struct {
	logger *Logger
	fields Fields
}

// WithFields creates a logger with additional fields.
func (l *Logger) WithFields(fields Fields) *FieldLogger {
	return &FieldLogger{
		logger: l,
		fields: fields,
	}
}

// formatFields formats fields for output.
func formatFields(fields Fields) string {
	if len(fields) == 0 {
		return ""
	}

	var parts []string
	for k, v := range fields {
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(parts, " ")
}

// Debug logs a debug message with fields.
func (l *FieldLogger) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if fields := formatFields(l.fields); fields != "" {
		msg = msg + " " + fields
	}
	l.logger.Debug(msg)
}

// Info logs an info message with fields.
func (l *FieldLogger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if fields := formatFields(l.fields); fields != "" {
		msg = msg + " " + fields
	}
	l.logger.Info(msg)
}

// Warn logs a warning message with fields.
func (l *FieldLogger) Warn(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if fields := formatFields(l.fields); fields != "" {
		msg = msg + " " + fields
	}
	l.logger.Warn(msg)
}

// Error logs an error message with fields.
func (l *FieldLogger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if fields := formatFields(l.fields); fields != "" {
		msg = msg + " " + fields
	}
	l.logger.Error(msg)
}

// VerboseLogger provides verbose output controlled by a flag.
type VerboseLogger struct {
	logger  *Logger
	verbose bool
}

// NewVerboseLogger creates a verbose logger.
func NewVerboseLogger(verbose bool) *VerboseLogger {
	return &VerboseLogger{
		logger:  NewLogger(WithLevel(LevelDebug)),
		verbose: verbose,
	}
}

// V prints verbose output if verbose mode is enabled.
func (l *VerboseLogger) V(format string, args ...interface{}) {
	if l.verbose {
		l.logger.Info(format, args...)
	}
}

// VDebug prints verbose debug output if verbose mode is enabled.
func (l *VerboseLogger) VDebug(format string, args ...interface{}) {
	if l.verbose {
		l.logger.Debug(format, args...)
	}
}

// QuietLogger suppresses all output except errors.
type QuietLogger struct {
	logger *Logger
	quiet  bool
}

// NewQuietLogger creates a quiet logger.
func NewQuietLogger(quiet bool) *QuietLogger {
	level := LevelInfo
	if quiet {
		level = LevelError
	}
	return &QuietLogger{
		logger: NewLogger(WithLevel(level)),
		quiet:  quiet,
	}
}

// Info prints info if not quiet.
func (l *QuietLogger) Info(format string, args ...interface{}) {
	if !l.quiet {
		l.logger.Info(format, args...)
	}
}

// Error always prints errors.
func (l *QuietLogger) Error(format string, args ...interface{}) {
	l.logger.Error(format, args...)
}

// Warn prints warnings if not quiet.
func (l *QuietLogger) Warn(format string, args ...interface{}) {
	if !l.quiet {
		l.logger.Warn(format, args...)
	}
}
