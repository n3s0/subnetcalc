package snc

import (
	"regexp"
	"testing"
    "strconv"
)

/*
func TestSubnetMaskToCidr(t *testing.T) {
	testStr := "255.255.128.0"
	correct := "17"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg, err := SubnetMaskToCidr(testStr)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !want.MatchString(msg) {
		t.Fatalf(`SubnetMaskToCidr(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}

func TestCidrToSubnetMask(t *testing.T) {
	testStr := "192.168.0.4/21"
	correct := "255.255.248.0"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	cidr := CidrToSubnetMask(testStr)
	msg := strconv.Itoa(cidr)
	if !want.MatchString(msg.String()) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}
*/

/*
func TestCidrHostCount(t *testing.T) {
	testStr := "192.168.1.0/25"
	correct := "128"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg := CidrToSubnetMask(testStr)
	if !want.MatchString(msg.String()) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}
*/

func TestCidrAvailableHostCount(t *testing.T) {
	testStr := "21"
	correct := "2046"
	want := regexp.MustCompile(`\b` + correct + `\b`)
    result := CidrAvailableHostCount(testStr)
    msg := strconv.Itoa(result)
	if !want.MatchString(msg) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}
