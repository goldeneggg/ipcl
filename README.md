# ipcl [![Build Status](https://drone.io/github.com/goldeneggg/ipcl/status.png)](https://drone.io/github.com/goldeneggg/ipcl/latest) [![GoDoc](https://godoc.org/github.com/goldeneggg/ipcl?status.png)](https://godoc.org/github.com/goldeneggg/ipcl)
* __ipcl__ is IP addresses calculator from CIDR.

## Install

```
% go get github.com/goldeneggg/ipcl
% ls $GOBIN/ipcl
ipcl
```

## Binary download link
* [linux_amd64](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/linux_amd64/ipcl)
* [linux_386](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/linux_386/ipcl)
* [darwin_amd64](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/darwin_amd64/ipcl)
* [darwin_386](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/darwin_386/ipcl)
* [windows_amd64](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/windows_amd64/ipcl.exe)
* [windows_386](https://drone.io/github.com/goldeneggg/ipcl/files/artifacts/bin/windows_386/ipcl.exe)

## Usage

```
Usage:
  ipcl [OPTIONS] [CIDR TEXT]

Application Options:
  -f, --file=    Filepath listed target CIDR
  -c, --csv=     Output format is csv
  -t, --tsv=     Output format is tsv
  -v, --version  Print version

Help Options:
  -h, --help     Show this help message
```

* Single argument of CIDR string

```
% ipcl 192.168.1.0/24
source_cidr : 192.168.1.0/24
network     : 192.168.1.0
mask        : 255.255.255.0
host_num    : 254
min_address : 192.168.1.1
max_address : 192.168.1.254
broadcast   : 192.168.1.255
```

* Multi CIDR strings from file using `-f``--file` option

```
% cat cidrs.txt
192.168.1.0/24
192.168.1.0/28
192.168.1.0/2

% ipcl -f cidrs.txt
source_cidr : 192.168.1.0/24
network     : 192.168.1.0
mask        : 255.255.255.0
host_num    : 254
min_address : 192.168.1.1
max_address : 192.168.1.254
broadcast   : 192.168.1.255

source_cidr : 192.168.1.0/28
network     : 192.168.1.0
mask        : 255.255.255.240
host_num    : 14
min_address : 192.168.1.1
max_address : 192.168.1.14
broadcast   : 192.168.1.15

source_cidr : 192.168.1.0/2
network     : 192.0.0.0
mask        : 192.0.0.0
host_num    : 1073741822
min_address : 192.0.0.1
max_address : 255.255.255.254
broadcast   : 255.255.255.255
```

* You can use CSV or TSV format using `-c``--csv` or `-t``--tsv` option

```
% ipcl -f cidrs.txt -c
source_cidr,network,mask,host_num,min_address,max_address,broadcast
192.168.1.0/24,192.168.1.0,255.255.255.0,254,192.168.1.1,192.168.1.254,192.168.1.255
192.168.1.0/28,192.168.1.0,255.255.255.240,14,192.168.1.1,192.168.1.14,192.168.1.15
192.168.1.0/2,192.0.0.0,192.0.0.0,1073741822,192.0.0.1,255.255.255.254,255.255.255.255
```
```
% ipcl -f cidrs.txt -t
source_cidr     network mask    host_num        min_address     max_address     broadcast
192.168.1.0/24  192.168.1.0     255.255.255.0   254     192.168.1.1     192.168.1.254   192.168.1.255
192.168.1.0/28  192.168.1.0     255.255.255.240 14      192.168.1.1     192.168.1.14    192.168.1.15
192.168.1.0/2   192.0.0.0       192.0.0.0       1073741822      192.0.0.1       255.255.255.254 255.255.255.255
```
