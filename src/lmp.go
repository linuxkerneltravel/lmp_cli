package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/docopt/docopt-go"
)

const (
	CONSOLE = "/bin/bash"
	CONSOLE_OPTION = "-c"
	ECLI = "sudo EUNOMIA_REPOSITORY=https://linuxkerneltravel.github.io/lmp/ EUNOMIA_HOME=/home/ubuntu/.lmp/ ~/.eunomia/bin/ecli "
	WASM_TO_OCI = "sudo ~/.eunomia/bin/wasm-to-oci "
	LMP_ORAS = "ghcr.io/linuxkerneltravel/"
	INIT = "git clone https://github.com/eunomia-bpf/ebpm-template && mv ebpm-template "
	BUILD = "~/.eunomia/bin/ecc *.bpf.c *.bpf.h -p"
	MAKE = "SOURCE_DIR=`pwd` make -C ~/.eunomia/bin "
)

func main() {
	usage := `LMP ecli.

Usage:
  lmp run <app> [<options_for_app>...] 
  lmp pull <package_name:version>
  lmp push <app.wasm> <package_name:version>
  lmp init <app>
  lmp build
  lmp gen-wasm-skel
  lmp build-wasm

Options:
  --help     Show this screen.
  --version     Show version.`

	parser := &docopt.Parser{
		HelpHandler:  docopt.PrintHelpAndExit,
		OptionsFirst: true,
	}
	opts, _ := parser.ParseArgs(usage, os.Args[1:], "go-docopt version 0.1")
	if opts["run"] == true {
		system(ECLI + "run " + opts["<app>"].(string) + " " + 
			strings.Join(opts["<options_for_app>"].([]string), " "))
	} else if opts["pull"] == true {
		//system(ECLI + "pull " + opts["<app>"].(string))
		system(WASM_TO_OCI + "pull " + LMP_ORAS + opts["<package_name:version>"].(string) + 
			" --out " + strings.Split(opts["<package_name:version>"].(string), ":")[0] + ".wasm")
	} else if opts["push"] == true {
		system(WASM_TO_OCI + "push " + opts["<app.wasm>"].(string) +  " " + 
			LMP_ORAS + opts["<package_name:version>"].(string))
	} else if opts["init"] == true {
		system(INIT + opts["<app>"].(string))
	} else if opts["build"] == true {
		system(BUILD)
	} else if opts["gen-wasm-skel"] == true {
		system(MAKE + "generate")
	} else if opts["build-wasm"] == true {
		system(MAKE + "build")
	}
}

func system(arg string) bool {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(CONSOLE , append([]string{CONSOLE_OPTION}, arg)...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	//var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		//_, errStdout = io.Copy(stdout, stdoutIn)
		_, _ = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		//_, errStderr = io.Copy(stderr, stderrIn)
		_, _ = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//if errStdout != nil || errStderr != nil {
	//	log.Fatal("failed to capture stdout or stderr\n")
	//}
	// outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		return false
	} else {
		return true
	}
}
