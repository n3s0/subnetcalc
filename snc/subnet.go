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
func SubnetCalculatorResultsCidr(cidrInput string) {
	ipAddressOutput := GetIpAddress(cidrInput)
	subnetMaskOutput := CidrToSubnetMask(cidrInput)
	cidrNotationOutput := GetCidrNotation(cidrInput)
	ipClassOutput := "null"
	networkAddressOutput := GetNetworkAddress(cidrInput)
	usableHostRangeOutput := "null"
	broadcastAddressOutput := GetBroadcastAddress(cidrInput)
	totalHostsOutput := CidrHostCount(cidrNotationOutput)
	usableHostsOutput := CidrAvailableHostCount(cidrNotationOutput)
	ipTypeOutput := "null"
	ipVersionOutput := "null"
	arpaNameOutput := "null"

	ct := table.NewWriter()
	ct.SetOutputMirror(os.Stdout)
	ct.SetTitle("Subnet Calculator Results")
	ct.AppendRows([]table.Row{
		{"IP Address", ipAddressOutput},
		{"Subnet Mask", subnetMaskOutput},
		{"CIDR Notation", cidrNotationOutput},
		{"IP Class", ipClassOutput},
		{"Network Address", networkAddressOutput},
		{"Usable Host IP Range", usableHostRangeOutput},
		{"Broadcast Address", broadcastAddressOutput},
		{"Total Hosts", totalHostsOutput},
		{"Usable Hosts", usableHostsOutput},
		{"IP Version", ipVersionOutput},
		{"IP Type", ipTypeOutput},
		{"in-addr.arpa", arpaNameOutput},
	})
	ct.SetStyle(table.StyleLight)
	ct.Render()
}

/*
Pulls the IP Address given by user input
*/
func GetIpAddress(userString string) string {
	ip, _, err := net.ParseCIDR(userString)
	if err != nil {
		fmt.Printf("[!] Invalid CIDR string...\n")
		fmt.Printf("[!] Error: %v\n", err)
	}

	address := ip.String()

	return address
}

/*
Pulls the CIDR notation for the subnet from the user input
Will be adding this soon.
*/
func GetCidrNotation(userString string) string {
	toSubnet := CidrToSubnetMask(userString)

	cidr, err := SubnetMaskToCidr(toSubnet)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return cidr
}

/*
Converts the CIDR notation to its subnet mask.
*/
func CidrToSubnetMask(cidrAddressString string) string {
	_, subnet, err := net.ParseCIDR(cidrAddressString)
	if err != nil {
		fmt.Printf("[!] Invalid CIDR string...\n")
		fmt.Printf("[!] Error: %v\n", err)
	}

	mask := subnet.Mask

	netmask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])

	return netmask
}

/*
Converts subnet mask to CIDR notation for the subnet mask.
*/
func SubnetMaskToCidr(subnetMaskString string) (string, error) {
	netmask := net.ParseIP(subnetMaskString).To4()
	if netmask == nil {
		fmt.Printf("[!] Invalid Subnet mask")
	}

	mask := net.IPv4Mask(netmask[0], netmask[1], netmask[2], netmask[3])

	ones, _ := mask.Size()

	cidr := strconv.Itoa(ones)

	return cidr, nil
}

/*
Gets the full host count by subtracting it from the net bits.
*/
func SubnetHostCount(subnetMaskString string) int {
	bits, err := SubnetMaskToCidr(subnetMaskString)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	netBits, err := strconv.Atoi(bits)
	if err != nil {
		fmt.Printf("%v", err)
	}

	hostBits := 32 - netBits

	return hostBits
}

/*
Gets the full host count from bit count.
*/
func CidrHostCount(cidrAddressString string) string {
	netBitStrConv, _ := strconv.Atoi(cidrAddressString)
	netBits := 32 - netBitStrConv
	hostsResult := int(math.Pow(float64(2), float64(netBits)))

	hosts := strconv.Itoa(hostsResult)

	return hosts
}

/*
Gets available host count from full host count.
*/
func CidrAvailableHostCount(cidrAddressString string) int {
	netBitStrConv, _ := strconv.Atoi(cidrAddressString)
	netBits := 32 - netBitStrConv
	fullHosts := int(math.Pow(float64(2), float64(netBits)))

	availHosts := fullHosts - 2

	return availHosts
}

func GetNetworkAddress(cidrAddressString string) string {
	_, subnet, err := net.ParseCIDR(cidrAddressString)
	if err != nil {
		fmt.Printf("%v", err)
	}

	netAddr := subnet.IP

	network := netAddr.String()

	return network
}

func GetBroadcastAddress(cidrAddressString string) string {
	_, ipSubnet, err := net.ParseCIDR(cidrAddressString)
	if err != nil {
		fmt.Printf("Error with address given\n")
		fmt.Printf("%v", err)
	}

	ip := ipSubnet.IP
	mask := ipSubnet.Mask

	bcAddress := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		bcAddress[i] = ip[i] | ^mask[i]
	}

	broadcast := bcAddress.String()

	return broadcast
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
