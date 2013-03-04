// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the untagged struct literal checker.

// +build vet_test

// This file contains the test for untagged struct literals.

package main

import (
	"flag"
	"go/scanner"
)

// Testing is awkward because we need to reference things from a separate package
// to trigger the warnings.

var BadStructLiteralUsedInTests = flag.Flag{ // ERROR "untagged fields"
	"Name",
	"Usage",
	nil, // Value
	"DefValue",
}

// Used to test the check for slices and arrays: If that test is disabled and
// vet is run with --compositewhitelist=false, this line triggers an error.
// Clumsy but sufficient.
var scannerErrorListTest = scanner.ErrorList{nil, nil}
