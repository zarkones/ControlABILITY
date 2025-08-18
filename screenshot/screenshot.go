//go:build !windows && !linux

package screenshot

func Capture() ([]byte, error) {
	return nil, ErrNotImplemented
}
