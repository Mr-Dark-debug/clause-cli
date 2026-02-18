//go:build windows
// +build windows

package utils

// NotifyResize returns a channel that receives terminal resize signals.
// On Windows, this is a no-op since SIGWINCH is not supported.
// Applications should poll for size changes on Windows.
func NotifyResize() <-chan struct{} {
	ch := make(chan struct{}, 1)
	// Windows doesn't support SIGWINCH, so we just return an empty channel
	// Applications on Windows should poll for size changes
	return ch
}

// StopResizeNotify stops listening for resize signals.
// On Windows, this is a no-op.
func StopResizeNotify(ch <-chan struct{}) {
	// No-op on Windows
}
