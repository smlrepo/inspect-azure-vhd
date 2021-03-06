// Package ext4 provides an API for reading ext4 filesystems through an
// io.ReadSeeker interface.
//
// Largely from https://ext4.wiki.kernel.org/index.php/Ext4_Disk_Layout
package ext4

import (
	"fmt"
)

var errNotImplemented = fmt.Errorf("Not implemented")

var ErrNotFound = fmt.Errorf("Not Found")

type UUID [16]byte

func (b UUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
