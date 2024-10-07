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
4. Determine network address if not given.
5. Determine broadcast address if not given.
6. Full number of addresses in subnet with count
6. Available address range and count
*/

/*
Generate the subnet calculator chart
*/
func SubetCalculatorResultsCidr(string) {

	sct := table.NewWriter()
	sct.SetOutputMirror(os.Stdout)
	sct.SetTitle("Subnet Calculator Results")
	sct.AppendHeader(table.Row{"CIDR", "Subnet Mask", "Address", "Wildcard"})
	sct.AppendRows([]table.Row{
		{"/32", "255.255.255.255", "1", "0.0.0.0"},
		{"/31", "255.255.255.254", "2", "0.0.0.1"},
		{"/30", "255.255.255.252", "4", "0.0.0.3"},
		{"/29", "255.255.255.248", "8", "0.0.0.7"},
		{"/28", "255.255.255.240", "16", "0.0.0.15"},
		{"/27", "255.255.255.224", "32", "0.0.0.31"},
		{"/26", "255.255.255.192", "64", "0.0.0.63"},
		{"/25", "255.255.255.128", "128", "0.0.0.127"},
		{"/24", "255.255.255.0", "256", "0.0.0.255"},
		{"/23", "255.255.254.0", "512", "0.0.1.255"},
		{"/22", "255.255.252.0", "1,024", "0.0.3.255"},
		{"/21", "255.255.248.0", "2,048", "0.0.7.255"},
		{"/20", "255.255.240.0", "4,096", "0.0.15.255"},
		{"/19", "255.255.224.0", "8,192", "0.0.31.255"},
		{"/18", "255.255.192.0", "16,384", "0.0.63.255"},
		{"/17", "255.255.128.0", "32,768", "0.0.127.255"},
		{"/16", "255.255.0.0", "65,536", "0.0.255.255"},
		{"/15", "255.254.0.0", "131,072", "0.1.255.255"},
		{"/14", "255.252.0.0", "262,144", "0.3.255.255"},
		{"/13", "255.248.0.0", "524,288", "0.7.255.255"},
		{"/12", "255.240.0.0", "1,048,576", "0.15.255.255"},
		{"/11", "255.224.0.0", "2,097,152", "0.31.255.255"},
		{"/10", "255.192.0.0", "4,194,304", "0.63.255.255"},
		{"/9", "255.128.0.0", "8,388,608", "0.127.255.255"},
		{"/8", "255.0.0.0", "16,777,216", "0.255.255.255"},
		{"/7", "254.0.0.0", "33,554,432", "1.255.255.255"},
		{"/6", "252.0.0.0", "67,108,864", "3.255.255.255"},
		{"/5", "248.0.0.0", "134,217,728", "7.255.255.255"},
		{"/4", "240.0.0.0", "268,435,456", "15.255.255.255"},
		{"/3", "224.0.0.0", "536,870,912", "31.255.255.255"},
		{"/2", "192.0.0.0", "1,073,741,824", "63.255.255.255"},
		{"/1", "128.0.0.0", "2,147,483,647", "127.255.255.255"},
		{"/0", "0.0.0.0", "4,294,967,296", "255.255.255.255"},
	})
	sct.SetStyle(table.StyleLight)
	sct.Render()
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
