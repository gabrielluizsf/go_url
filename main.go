package main

import (
	"github.com/theGOURL/OS_Analyzer/analyst"
	"github.com/theGOURL/go_url/pkg/version"
)

// Copyright 2023 theGOURL . All rights reserved
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

func main() {
	version.PrintVersion();
  	analyst.OSAnalyzer();
}
