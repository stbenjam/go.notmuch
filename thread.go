package notmuch

// Copyright © 2015 The go.notmuch Authors. Authors can be found in the AUTHORS file.
// Licensed under the GPLv3 or later.
// See COPYING at the root of the repository for details.

// #cgo LDFLAGS: -lnotmuch
// #include <stdlib.h>
// #include <notmuch.h>
import "C"

// Thread represents a notmuch thread.
type Thread struct {
	cptr    *C.notmuch_thread_t
	threads *Threads
}

// GetSubject returns the subject of a thread.
func (t *Thread) GetSubject() string {
	cstr := C.notmuch_thread_get_subject(t.cptr)
	str := C.GoString(cstr)
	return str
}

// GetID returns the ID of the thread.
func (t *Thread) GetID() string {
	return C.GoString(C.notmuch_thread_get_thread_id(t.cptr))
}
