// Copyright 2015 Mark Haferkamp. This code is licensed under the
// license found in the LICENSE file.

// fdupes-dedupe - Use fdupes to find duplicate files and then
// deduplicate their storage via 'cp --reflink'.

package main

import (
	"fmt"     // output
	"os"      // for getting and setting file metadata
	"os/exec" // running fdupes and cp
	//"strings", "io", or "buffio" // Save fdupes output and process it to get file lists
	"io/ioutil" // write plan file before actions, read logs in recovery mode
	// "os" or "buffio"? // write action log file one line at a time
)
