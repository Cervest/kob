package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Cervest/kob/k8s"
)

var usage string = `Usage:
  with-args --file path/to/spec <ARGS>
`

func main() {

	withArgsCmd := flag.NewFlagSet("with-args", flag.ExitOnError)

	var withArgsFile string
	withArgsDefault := "job.yml"
	withArgsUsage := "Path to Job spec. Spec file must contain a single Job with a single container."
	withArgsCmd.StringVar(&withArgsFile, "file", withArgsDefault, withArgsUsage)
	withArgsCmd.StringVar(&withArgsFile, "f", withArgsDefault, withArgsUsage)

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "with-args":
		withArgsCmd.Parse(os.Args[2:])
		k8s.RunJobWithArgs(withArgsFile, withArgsCmd.Args())
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
