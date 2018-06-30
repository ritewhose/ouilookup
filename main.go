// Package ouilookup provides a command line interface for looking up the manufacturer
// of a device by its MAC address's Organizational Unit Identifer.
package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"

	"github.com/google/gopacket/macs"
)

func main() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintf(os.Stderr, "Usage: %s <MAC Address...>\n", os.Args[0])
		os.Exit(1)
	}

	for _, macstr := range os.Args[1:] {
		mac, err := parseMACHex(macstr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var oui [3]byte
		copy(oui[:], mac[:3])

		mfr, ok := macs.ValidMACPrefixMap[oui]
		if !ok {
			mfr = "N/A"
		}
		fmt.Printf("%v\t%s\n", mac, mfr)
	}
}

// parseMACHex decodes MAC addresses from the formats supported by net.ParseMAC, plus
// from a plain hex string.
func parseMACHex(s string) (net.HardwareAddr, error) {
	nbytes := len(s) / 2
	if nbytes != 6 && nbytes != 12 && nbytes != 20 {
		return net.ParseMAC(s)
	}
	h, err := hex.DecodeString(s)
	if err != nil {
		err = fmt.Errorf("address %v: invalid MAC address", s)
		return nil, err
	}
	hw := make(net.HardwareAddr, nbytes)
	copy(hw, h)
	return hw, nil
}
