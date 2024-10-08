package snc

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func SubnetChart() {
	sct := table.NewWriter()
	sct.SetOutputMirror(os.Stdout)
	sct.SetTitle("Subnet Chart")
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

func DecimalToBinary() {
	btbt := table.NewWriter()
	btbt.SetOutputMirror(os.Stdout)
	btbt.SetTitle("Decimal To Binary")
	btbt.AppendHeader(table.Row{"Subnet Mask", "Wildcard"})
	btbt.AppendRows([]table.Row{
		{"255 1111 1111", "0 0000 0000"},
		{"254 1111 1110", "1 0000 0001"},
		{"252 1111 1100", "3 0000 0011"},
		{"248 1111 1000", "7 0000 0111"},
		{"240 1111 0000", "15 0000 1111"},
		{"224 1110 0000", "31 0001 1111"},
		{"192 1100 0000", "63 0011 1111"},
		{"128 1000 0000", "127 0111 1111"},
		{"0 0000 0000", "255 1111 1111"},
	})
	btbt.SetStyle(table.StyleLight)
	btbt.Render()
}

func ClassfulRanges() {
	crt := table.NewWriter()
	crt.SetOutputMirror(os.Stdout)
	crt.SetTitle("Classful Ranges")
	crt.AppendHeader(table.Row{"Class", "Start", "End"})
	crt.AppendRows([]table.Row{
		{"A", "0.0.0.0", "127.255.255.255"},
		{"B", "128.0.0.0", "191.255.255.255"},
		{"C", "192.0.0.0", "223.255.255.255"},
		{"D", "224.0.0.0", "239.255.255.255"},
		{"E", "240.0.0.0", "255.255.255.255"},
	})
	crt.SetStyle(table.StyleLight)
	crt.Render()
}

func ReservedRanges() {
	rrt := table.NewWriter()
	rrt.SetOutputMirror(os.Stdout)
	rrt.SetTitle("Reserved Ranges")
	rrt.AppendHeader(table.Row{"Reference", "Range"})
	rrt.AppendRows([]table.Row{
		{"rfc1918", "10.0.0.0 - 10.255.255.255"},
		{"localhost", "127.0.0.0 - 127.255.255.255"},
		{"rfc1918", "172.16.0.0 - 172.31.255.255"},
		{"rfc1918", "192.168.0.0 - 192.168.255.255"},
	})
	rrt.SetStyle(table.StyleLight)
	rrt.Render()
}

func EthernetCableLength() {
	eclt := table.NewWriter()
	eclt.SetOutputMirror(os.Stdout)
	eclt.SetTitle("Ethernet Cable Length")
	eclt.AppendHeader(table.Row{"Category", "Max. Data Rate", "Max. Distance"})
	eclt.AppendRows([]table.Row{
		{"Category 1", "1 Mbps", ""},
		{"Category 2", "4 Mbps", ""},
		{"Category 3", "10 Mbps", "100 m (328 ft.)"},
		{"Category 4", "16 Mbps", "100 m (328 ft.)"},
		{"Category 5", "100 Mbps", "100 m (328 ft.)"},
		{"Category 5e", "1 Gbps", "100 m (328 ft.)"},
		{"Category 6", "1 Gbps", "100 m (328 ft.)"},
		{"", "", "10Gb at 37 m (121 ft.)"},
		{"Category 6a", "10 Gbps", "100 m (328 ft.)"},
		{"Category 7", "10 Gbps", "100 m (328 ft.)"},
		{"Category 7a", "10 Gbps", "100 m (328 ft.)"},
		{"", "", "40 Gb at 50 m (164 ft.)"},
		{"Category 8", "1 Gbps", "100 m (328 ft.)"},
	})
	eclt.SetStyle(table.StyleLight)
	eclt.Render()
}
