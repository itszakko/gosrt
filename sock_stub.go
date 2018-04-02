// +build darwin dragonfly freebsd netbsd openbsd linux plan9 windows nacl solaris

package gosrt

import "syscall"

func maxListenerBacklog() int {
	// TODO: Implement this
	// NOTE: Never return a number bigger than 1<<16 - 1. See issue 5030.
	return syscall.SOMAXCONN
}