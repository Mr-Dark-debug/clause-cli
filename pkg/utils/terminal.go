package utils

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// TerminalInfo contains information about the terminal.
type TerminalInfo struct {
	Width       int
	Height      int
	IsTerminal  bool
	IsDumbTerm  bool
	ColorDepth  ColorDepth
	SupportsUTF bool
	Platform    string
}

// ColorDepth represents the color support level.
type ColorDepth int

const (
	// ColorDepthNone means no color support.
	ColorDepthNone ColorDepth = iota
	// ColorDepth16 means 16-color support.
	ColorDepth16
	// ColorDepth256 means 256-color support.
	ColorDepth256
	// ColorDepthTrue means 24-bit true color support.
	ColorDepthTrue
)

// DetectTerminalInfo detects and returns terminal information.
func DetectTerminalInfo() TerminalInfo {
	info := TerminalInfo{
		Platform:   runtime.GOOS,
		IsTerminal: IsTerminal(),
		IsDumbTerm: IsDumbTerminal(),
		ColorDepth: DetectColorDepth(),
	}

	if width, height, err := GetTerminalSize(); err == nil {
		info.Width = width
		info.Height = height
	}

	// Assume UTF-8 support on modern terminals
	info.SupportsUTF = !info.IsDumbTerm && info.ColorDepth > ColorDepthNone

	return info
}

// IsTerminal checks if stdout is connected to a terminal.
func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// IsTerminalFd checks if the given file descriptor is a terminal.
func IsTerminalFd(fd int) bool {
	return term.IsTerminal(fd)
}

// IsDumbTerminal checks if we're in a dumb terminal.
func IsDumbTerminal() bool {
	term := os.Getenv("TERM")
	return term == "" || term == "dumb"
}

// GetTerminalSize returns the terminal dimensions.
func GetTerminalSize() (width, height int, err error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

// GetTerminalWidth returns the terminal width.
func GetTerminalWidth() int {
	width, _, err := GetTerminalSize()
	if err != nil {
		return 80 // Default width
	}
	return width
}

// GetTerminalHeight returns the terminal height.
func GetTerminalHeight() int {
	_, height, err := GetTerminalSize()
	if err != nil {
		return 24 // Default height
	}
	return height
}

// DetectColorDepth detects the color support level.
func DetectColorDepth() ColorDepth {
	// Check COLORTERM for true color
	colorterm := os.Getenv("COLORTERM")
	if strings.Contains(strings.ToLower(colorterm), "truecolor") ||
		strings.Contains(strings.ToLower(colorterm), "24bit") {
		return ColorDepthTrue
	}

	// Check TERM for 256-color
	termEnv := os.Getenv("TERM")
	termLower := strings.ToLower(termEnv)

	if strings.Contains(termLower, "256color") || strings.Contains(termLower, "256") {
		return ColorDepth256
	}

	// Check for basic color terminals
	colorTerms := []string{
		"xterm", "screen", "vt100", "vt200", "ansi",
		"color", "cygwin", "linux", "rxvt", "konsole",
		"gnome", "terminator", "alacritty", "kitty",
	}

	for _, ct := range colorTerms {
		if strings.Contains(termLower, ct) {
			// Check TERM_PROGRAM for more specific detection
			termProgram := strings.ToLower(os.Getenv("TERM_PROGRAM"))
			switch termProgram {
			case "iterm.app", "apple_terminal", "vscode", "hyper", "warp":
				return ColorDepthTrue
			}
			return ColorDepth16
		}
	}

	// Check for Windows Terminal
	if os.Getenv("WT_SESSION") != "" {
		return ColorDepthTrue
	}

	// Check for Windows
	if runtime.GOOS == "windows" {
		// Modern Windows terminals often support true color
		if os.Getenv("ANSICON") != "" {
			return ColorDepth16
		}
		// Windows 10+ with virtual terminal sequences
		return ColorDepthTrue
	}

	return ColorDepthNone
}

// SupportsTrueColor checks if the terminal supports 24-bit color.
func SupportsTrueColor() bool {
	return DetectColorDepth() >= ColorDepthTrue
}

// Supports256Colors checks if the terminal supports 256 colors.
func Supports256Colors() bool {
	return DetectColorDepth() >= ColorDepth256
}

// SupportsColor checks if the terminal supports any color.
func SupportsColor() bool {
	return DetectColorDepth() > ColorDepthNone
}

// SupportsUTF8 checks if the terminal supports UTF-8.
func SupportsUTF8() bool {
	// Check LANG or LC_CTYPE for UTF-8
	lang := os.Getenv("LANG")
	ctype := os.Getenv("LC_CTYPE")
	combined := lang + ctype

	return strings.Contains(strings.ToLower(combined), "utf") ||
		(!IsDumbTerminal() && runtime.GOOS != "windows")
}

// ClearScreen clears the terminal screen.
func ClearScreen() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.Write([]byte("\x1b[2J\x1b[H"))
}

// ClearLine clears the current line.
func ClearLine() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.Write([]byte("\x1b[2K\r"))
}

// ClearToEndOfLine clears from cursor to end of line.
func ClearToEndOfLine() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.Write([]byte("\x1b[K"))
}

// MoveCursor moves the cursor to the specified position (1-based).
func MoveCursor(row, col int) {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[" + strconv.Itoa(row) + ";" + strconv.Itoa(col) + "H")
}

// MoveCursorUp moves the cursor up n lines.
func MoveCursorUp(n int) {
	if IsDumbTerminal() || n <= 0 {
		return
	}
	os.Stdout.WriteString("\x1b[" + strconv.Itoa(n) + "A")
}

// MoveCursorDown moves the cursor down n lines.
func MoveCursorDown(n int) {
	if IsDumbTerminal() || n <= 0 {
		return
	}
	os.Stdout.WriteString("\x1b[" + strconv.Itoa(n) + "B")
}

// MoveCursorForward moves the cursor forward n columns.
func MoveCursorForward(n int) {
	if IsDumbTerminal() || n <= 0 {
		return
	}
	os.Stdout.WriteString("\x1b[" + strconv.Itoa(n) + "C")
}

// MoveCursorBack moves the cursor back n columns.
func MoveCursorBack(n int) {
	if IsDumbTerminal() || n <= 0 {
		return
	}
	os.Stdout.WriteString("\x1b[" + strconv.Itoa(n) + "D")
}

// SaveCursor saves the current cursor position.
func SaveCursor() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[s")
}

// RestoreCursor restores the saved cursor position.
func RestoreCursor() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[u")
}

// HideCursor hides the cursor.
func HideCursor() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?25l")
}

// ShowCursor shows the cursor.
func ShowCursor() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?25h")
}

// EnableMouse enables mouse tracking.
func EnableMouse() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1000h\x1b[?1002h\x1b[?1015h\x1b[?1006h")
}

// DisableMouse disables mouse tracking.
func DisableMouse() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1000l\x1b[?1002l\x1b[?1015l\x1b[?1006l")
}

// EnableAlternateScreen switches to the alternate screen buffer.
func EnableAlternateScreen() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1049h")
}

// DisableAlternateScreen switches back to the main screen buffer.
func DisableAlternateScreen() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1049l")
}

// SetWindowTitle sets the terminal window title.
func SetWindowTitle(title string) {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b]0;" + title + "\x07")
}

// EnableBracketedPaste enables bracketed paste mode.
func EnableBracketedPaste() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?2004h")
}

// DisableBracketedPaste disables bracketed paste mode.
func DisableBracketedPaste() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?2004l")
}

// EnableFocusReporting enables focus reporting.
func EnableFocusReporting() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1004h")
}

// DisableFocusReporting disables focus reporting.
func DisableFocusReporting() {
	if IsDumbTerminal() {
		return
	}
	os.Stdout.WriteString("\x1b[?1004l")
}

// StartRawMode puts the terminal in raw mode and returns a restore function.
func StartRawMode() (restore func(), err error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, err
	}

	restore = func() {
		term.Restore(fd, oldState)
	}

	return restore, nil
}

// IsInteractive checks if the terminal is interactive (supports TUI).
func IsInteractive() bool {
	return IsTerminal() && !IsDumbTerminal()
}

// GetShell returns the current shell.
func GetShell() string {
	if runtime.GOOS == "windows" {
		if shell := os.Getenv("COMSPEC"); shell != "" {
			return shell
		}
		return "cmd.exe"
	}

	if shell := os.Getenv("SHELL"); shell != "" {
		return shell
	}

	return "/bin/sh"
}

// GetEditor returns the preferred text editor.
func GetEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}
	if visual := os.Getenv("VISUAL"); visual != "" {
		return visual
	}

	if runtime.GOOS == "windows" {
		return "notepad"
	}

	return "vi"
}

// GetPager returns the preferred pager.
func GetPager() string {
	if pager := os.Getenv("PAGER"); pager != "" {
		return pager
	}

	if runtime.GOOS == "windows" {
		return "more"
	}

	return "less"
}

// IsColorForced checks if color output is forced via environment.
func IsColorForced() bool {
	return os.Getenv("CLICOLOR_FORCE") != "" ||
		os.Getenv("FORCE_COLOR") != ""
}

// IsColorDisabled checks if color output is disabled via environment.
func IsColorDisabled() bool {
	return os.Getenv("NO_COLOR") != "" ||
		os.Getenv("CLICOLOR") == "0"
}

// ShouldUseColor determines if color output should be used.
func ShouldUseColor() bool {
	if IsColorForced() {
		return true
	}
	if IsColorDisabled() {
		return false
	}
	return SupportsColor()
}
