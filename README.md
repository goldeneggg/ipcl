# ipcl [![Build Status](https://drone.io/github.com/jpshadowapps/ipcl/status.png)](https://drone.io/github.com/jpshadowapps/ipcl/latest)
* __ipcl__ is IP addresses calculator from CIDR.

## Install

```
% go get github.com/jpshadowapps/ipcl
% ls $GOBIN/ipcl
ipcl
```

## Usage

```
Usage of ipcl:
  -csv=false: Output format is csv
  -f="": Filepath listed target CIDR
  -tsv=false: Output format is tsv
  -version=false: Print version
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

* Multi CIDR strings from file using `-f` option

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

* You can use CSV or TSV format using `-csv` or `-tsv` option

```
% ipcl -f cidrs.txt -csv
source_cidr,network,mask,host_num,min_address,max_address,broadcast
192.168.1.0/24,192.168.1.0,255.255.255.0,254,192.168.1.1,192.168.1.254,192.168.1.255
192.168.1.0/28,192.168.1.0,255.255.255.240,14,192.168.1.1,192.168.1.14,192.168.1.15
192.168.1.0/2,192.0.0.0,192.0.0.0,1073741822,192.0.0.1,255.255.255.254,255.255.255.255
```
```
% ipcl -f cidrs.txt -tsv
source_cidr     network mask    host_num        min_address     max_address     broadcast
192.168.1.0/24  192.168.1.0     255.255.255.0   254     192.168.1.1     192.168.1.254   192.168.1.255
192.168.1.0/28  192.168.1.0     255.255.255.240 14      192.168.1.1     192.168.1.14    192.168.1.15
192.168.1.0/2   192.0.0.0       192.0.0.0       1073741822      192.0.0.1       255.255.255.254 255.255.255.255
```
