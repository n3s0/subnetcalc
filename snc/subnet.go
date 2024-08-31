package snc

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Process that this needs to perform is below. Code will need to do a
number of things. But, we'll start with some of the basics first.

TODO

1. Convert CIDR into subnet mask.
2. Convert subnet mask to CIDR.
3. Determine host bits and network bits of subnet mask/cidr
4. Determine network address if not given.
5. Determine broadcast address if not given.
6. Full number of addresses in subnet with count
6. Available address range and count
*/

func SubnetMaskToCidr(netmask string) string {
	octets := strings.Split(netmask, ".")
	if len(octets) != 4 {
		fmt.Println("Invalid number of octets provided...")
		os.Exit(-1)
	}

	mask := 0
	for _, octet := range octets {
		part, err := strconv.Atoi(octet)
		if err != nil || part < 0 || part > 255 {
			fmt.Println("Not within valid network range or non existent...")
		}
		mask = (mask << 8) | part
	}

	binaryStr := fmt.Sprintf("%032b", mask)
	count := 0
	for _, bit := range binaryStr {
		if bit == '1' {
			count++
		} else {
			break
		}
	}
	cidr := strconv.Itoa(count)

	return cidr
}

func CidrToSubnetMask(cidr string) string {
	cidrInt, err := strconv.Atoi(cidr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	if cidrInt < 0 || cidrInt > 32 {
		fmt.Println("Invalid CIDR entry...")
		os.Exit(-1)
	}

	octets := [4]byte{}

	netBits := cidrInt / 8
	hostBits := cidrInt % 8

	for i := 0; i < netBits; i++ {
		octets[i] = 255
	}

	if netBits < 4 {
		octets[netBits] = (255 << (8 - hostBits)) & 255
	}

	netmask := fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3])

	return netmask
}

func CidrHostCount(cidr string) string {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	hosts := int(math.Pow(float64(2), float64(netBits)))

	return fmt.Sprintf("%v", hosts)
}

func CidrAvailableHostCount(cidr string) string {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	fullHosts := int(math.Pow(float64(2), float64(netBits)))

	availHosts := fullHosts - 2

	return fmt.Sprintf("%v", availHosts)
}
