package snc

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

/*
Process that this needs to perform is below. Code will need to do a
number of things. But, we'll start with some of the basics first.

# TODO

1. [x] Convert CIDR into subnet mask.
2. [x] Convert subnet mask to CIDR.
3. [half] Determine host bits and network bits of subnet mask/cidr
4. [] Determine network address if not given.
5. [] Determine broadcast address if not given.
6. [] Full number of addresses in subnet with count
6. [] Available address range and count
*/

/*
Generate the subnet calculator chart
*/
func SubetCalculatorResultsCidr(string) {

	ct := table.NewWriter()
	ct.SetOutputMirror(os.Stdout)
	ct.SetTitle("Subnet Calculator Results")
	ct.AppendHeader(table.Row{"CIDR", "Subnet Mask", "Address", "Wildcard"})
	ct.AppendRows([]table.Row{
		{"/32", "255.255.255.255", "1", "0.0.0.0"},
		{"/31", "255.255.255.254", "2", "0.0.0.1"},
	})
	ct.SetStyle(table.StyleLight)
	ct.Render()
}

/*
Used to separate the user input for a CIDR address and return it.
*/
func SeparateAddressCidrInput(userString string) (net.IP, *net.IPNet) {
	ip, subnet, err := net.ParseCIDR(userString)
	if err != nil {
		fmt.Printf("[!] Invalid CIDR string...\n")
		fmt.Printf("[!] Error: %v\n", err)
	}

	return ip, subnet
}

/*
Converts the CIDR notation to its subnet mask.
*/
func CidrToSubnetMask(cidrAddressString string) net.IPMask {
	_, subnet, err := net.ParseCIDR(cidrAddressString)
	if err != nil {
		fmt.Printf("[!] Invalid CIDR string...\n")
		fmt.Printf("[!] Error: %v\n", err)
	}

	subnetMask := subnet.Mask

	return subnetMask
}

/*
Converts subnet mask to CIDR notation for the subnet mask.
*/
func SubnetMaskToCidr(subnetMaskString string) (int, error) {
	netmask := net.ParseIP(subnetMaskString).To4()
	if netmask == nil {
		return 0, fmt.Errorf("[!] Invalid Subnet mask")
	}

	mask := net.IPv4Mask(netmask[0], netmask[1], netmask[2], netmask[3])

	ones, _ := mask.Size()

	return ones, nil
}

/*
Gets the full host count by subtracting it from the net bits.
*/
func SubnetHostCount(subnet string) int {
	netBits, err := SubnetMaskToCidr(subnet)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	hostBits := 32 - netBits

	return hostBits
}

/*
Gets the full host count from bit count.
*/
func CidrHostCount(cidr string) int {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	hosts := int(math.Pow(float64(2), float64(netBits)))

	return hosts
}

/*
Gets available host count from full host count.
*/
func CidrAvailableHostCount(cidr string) int {
	netBitStrConv, _ := strconv.Atoi(cidr)
	netBits := 32 - netBitStrConv
	fullHosts := int(math.Pow(float64(2), float64(netBits)))

	availHosts := fullHosts - 2

	return availHosts
}

/*
Checks the IP address accessible to the Internet and returns it.

DNS is required for this. It's also experimental. So it probably wont
get used.
*/
func GetPrefferedInternetRoutableIP() net.IP {
	//
	tcpConn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println(err)
	}

	defer tcpConn.Close()

	addr := tcpConn.LocalAddr().(*net.TCPAddr)

	return addr.IP
}
