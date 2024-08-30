package snc

import (
	"fmt"
	"math"
	"strconv"
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

func CIDRHostCount(cidr string) string {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	fullHosts := int(math.Pow(float64(2), float64(netBits)))

	return fmt.Sprintf("%v", fullHosts)
}

func CIDRAvailableHostCount(cidr string) string {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	fullHosts := int(math.Pow(float64(2), float64(netBits)))

	availHosts := fullHosts - 2

	return fmt.Sprintf("%v", availHosts)
}
