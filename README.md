# lmp-cli

## Dependency

Binary program *bpftool*, *ecc* and *ecli* builded from (eunomia-bpf)[https://github.com/eunomia-bpf/eunomia-bpf], 
and *wasm-to-oci* builded from (wasm-to-oci)[https://github.com/engineerd/wasm-to-oci]. After getting the binary 
programs, you should move them in *~/.eunomia/bin/* for the calling of lmp.

*golang* with package *docopt* for building lmp.

## Build

In the main directory of the project.

```bash
cd src && go build
```

## Install

### Use binary programs builded by yourself

In the main directory of the project.

```bash
sudo cp -i src/lmp /usr/bin/ && \
cp -r bin/ build-wasm/ include/ libc-buildin-sysroot/ wasi-sdk-16.0/ ~/.eunomia/
```

###  Use the binary programs we provided

You could download our release package for Ubuntu x86 if your system is adoptable.

```bash
wget https://github.com/linuxkerneltravel/lmp_cli/releases/download/lmp_go/lmp_cli_go_docopt_v1.0.tar.gz && \
mkdir ~/.eunomia/ && \
tar -zxf lmp_cli_go_docopt_v1.0.tar.gz -C ~/.eunomia/ && \
sudo cp ~/.eunomia/bin/lmp /usr/bin/
```

## Quick Start

This is a LMP command line initiator to invoke ecli commands:

### User

First, we have a use case for developers who want to use an ebpf binary or program but doesn't know how/where to find it:
Run directly.

```bash
# Use a name and run it directly. If it's not available locally, download it from the corresponding repo on the web
$ lmp run opensnoop
...

# Use a name and version number
$ lmp run opensnoop:latest
...

# Use an http API
$ lmp run https://github.com/ebpf-io/raw/master/examples/opensnoop/app.wasm
...

# Use a local path
$ lmp run ./opensnoop/package.json
...

# With parameters
$ lmp run sigsnoop -h
Usage: sigsnoop [-h] [-x] [-k] [-n] [-p PID] [-s SIGNAL]
Trace standard and real-time signals.


    -h, --help  show this help message and exit
    -x, --failed  failed signals only
    -k, --killed  kill only
    -p, --pid=<int>  target pid
    -s, --signal=<int>  target signal
```

In fact, the run command contains the pull command. If the local ebpf file is not available, it will be downloaded from the Internet. If the local EBPF file is not available, it will be directly used:

```bash
$ ecli pull opensnoop
$ ecli run opensnoop
...
```
 
Or he or she can search the web and download it (see the next chapter).

### Developer

Our second role is a developer who wants to create a universal eBPF/WASM precompiled binary, distribute it on any machine and operating system, and load it dynamically to run. This is useful for command-line tools or anything you can run directly in the Shell, or as a plug-in for large projects.

Generate ebpf data files.

```bash
# Generate a project template
$ lmp init opensnoop
init project opensnoop success.

# A new directory is created
$ cd opensnoop
$ ls # The following files are generated
opensnoop.bpf.c opensnoop.bpf.h README.md config.json .gitignore

# Build a kernel-state program
$ lmp build
$ ls # The following files are generated. package.json is the result of compilation
opensnoop.bpf.c opensnoop.bpf.h README.md config.json package.json .gitignore

# Generate a wasm user mode project template
$ lmp gen-wasm-skel
make
  GENERATE_PACKAGE_JSON
  GEN-WASM-SKEL
$ ls # The following files are generated
app.c eunomia-include ewasm-skel.h package.json README.md  opensnoop.bpf.c  opensnoop.bpf.h

$ lmp build-wasm
make
  GENERATE_PACKAGE_JSON
  BUILD-WASM
build app.wasm success

$ lmp run app.wasm -h
Usage: opensnoop [-h] [-x] [-k] [-n] [-p PID]

```

gen-wasm-skel provide the C language version of the WASM development framework, it contains the following files:

- ewasm-skel.h：The header file of the user-mode WebAssembly development framework contains the precompiled bytecode of eBPF program and the auxiliary information of eBPF program framework for dynamic loading.
- eunomia-include：Some header-only library functions and auxiliary files to aid development.
- app.c：The main code of user mode WebAssembly program contains the main logic of eBPF program and the data processing flow of eBPF program.

The code that users need to write in app.c should be pure, normal C code, without any knowledge of the underlying WASM implementation. You can develop with the framework without knowing anything about WASM.

After building the wasm package, you could push your wasm package to LMP package repo. Before doing it, you should obtain a personal access token and use it to login oras.

```bash
$ lmp pull bootstrap:latest
```
As well, you can push your work to the repo.

```bash
$ lmp push app.wasm <work_name>:<version>
```

## Recommend

The relevant framework eunomia-bpf: An eBPF program Dynamic Loading Framework: https://github.com/eunomia-bpf/eunomia-bpf

The details for Linux Microscope LMP project: https://github.com/linuxkerneltravel/lmp

The details for clipp project: https://github.com/muellan/clipp#an-example-from-docopt
