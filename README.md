# ouilookup
A command line tool for looking up the manufacturer of a network interface by its MAC address's Organizationally Unique Identifer (OUI).

## Usage
```
$ ouilookup -h
Usage: ouilookup <MAC Address...>
$ ouilookup 00000C382A4C 00-10-fa-c2-bf-d5 D8B1.22a3.53FE c0:f3:2d:3e:88:ee
00:00:0c:38:2a:4c	Cisco Systems, Inc
00:10:fa:c2:bf:d5	Apple, Inc.
d8:b1:22:a3:53:fe	Juniper Networks
c0:f3:2d:3e:88:ee	N/A
```

## Installation
In addition to [setting up your Go environment](https://golang.org/doc/install)...

``` go get github.com/ritewhose/ouilookup```
