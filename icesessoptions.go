package gopjnath

/*
#cgo pkg-config: libpjnath
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/

import "C"

import (
    "sync"
    "syscall"
    "time"
    "unsafe"
    )

type IceSessOptions struct {
    Aggressive                    bool
    NominatedCheckDelay           uint
    ControlledAgentWantNomTimeout int
}
