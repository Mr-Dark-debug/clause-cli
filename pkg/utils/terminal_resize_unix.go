//go:build !windows
// +build !windows

package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// NotifyResize returns a channel that receives terminal resize signals.
func NotifyResize() <-chan struct{} {
	ch := make(chan struct{}, 1)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)

	go func() {
		for range sigCh {
			select {
			case ch <- struct{}{}:
			default:
			}
		}
	}()

	// Send initial size
	select {
	case ch <- struct{}{}:
	default:
	}

	return ch
}

// StopResizeNotify stops listening for resize signals.
func StopResizeNotify(ch <-chan struct{}) {
	signal.Reset(syscall.SIGWINCH)
}
