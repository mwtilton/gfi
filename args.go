package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// headerArgs is the type used to store the header arguments
type headerArgs []string

func (h *headerArgs) Set(val string) error {
	*h = append(*h, val)
	return nil
}

func (h headerArgs) String() string {
	return "string"
}

type saveStatusArgs []int

func (s *saveStatusArgs) Set(val string) error {
	i, _ := strconv.Atoi(val)
	*s = append(*s, i)
	return nil
}

func (s saveStatusArgs) String() string {
	return "string"
}

func (s saveStatusArgs) Includes(search int) bool {
	for _, status := range s {
		if status == search {
			return true
		}
	}
	return false
}

type config struct {
	body           string
	concurrency    int
	delay          int
	headers        headerArgs
	followLocation bool
	method         string
	saveStatus     saveStatusArgs
	timeout        int
	verbose        bool

	paths     string
	hosts     string
	output    string
	noHeaders bool

	requester requester
}

func processArgs() config {

	// default the output directory to ./out
	output := flag.Arg(2)
	if output == "" {
		output = defaultOutputDir
	}

	return config{
		body:           body,
		concurrency:    concurrency,
		delay:          delay,
		headers:        headers,
		followLocation: followLocation,
		method:         method,
		saveStatus:     saveStatus,
		timeout:        timeout,
		requester:      requesterFn,
		verbose:        verbose,
		paths:          paths,
		hosts:          hosts,
		output:         output,
		noHeaders:      noHeaders,
	}
}


func init() {
	flag.Usage = func() {
		h := "Get file Information (i.e. Hashes, ACL's)\n\n"

		h += "Usage:\n"
		h += "  gfi [directory|file] [outputDir]\n\n"

		h += "Options:\n"
		h += "  -d, --directory <val>           Set the target directory\n"
		h += "  -f, --file <val>           Set the target file\n"
		h += "  -o, --output <method>      Outputs a .txt file to specified directory.\n\n"

		h += "Defaults:\n"
		h += "  pathsFile: ./paths\n"
		h += "  outputDir:  ./out\n\n"

		h += "Examples:\n"
		h += "  gfi /dir /out/output.txt\n"
		h += "  gfi /dir1 /dir2 /dir3 /output\n"

		fmt.Fprintf(os.Stderr, h)
	}
}