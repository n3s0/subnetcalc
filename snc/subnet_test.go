package snc

import (
	"regexp"
	"testing"
)

func TestSubnetMaskToCidr(t *testing.T) {
	testStr := "255.255.128.0"
	correct := "17"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg := SubnetMaskToCidr(testStr)
	if !want.MatchString(msg) {
		t.Fatalf(`SubnetMaskToCidr(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}

func TestCidrToSubnetMask(t *testing.T) {
	testStr := "21"
	correct := "255.255.248.0"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg := CidrToSubnetMask(testStr)
	if !want.MatchString(msg) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}

func TestCidrHostCount(t *testing.T) {
	testStr := "25"
	correct := "128"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg := CidrHostCount(testStr)
	if !want.MatchString(msg) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}

func TestCidrAvailableHostCount(t *testing.T) {
	testStr := "21"
	correct := "2046"
	want := regexp.MustCompile(`\b` + correct + `\b`)
	msg := CidrAvailableHostCount(testStr)
	if !want.MatchString(msg) {
		t.Fatalf(`CidrToSubnetMask(%s) = %q, want match for %#q`, testStr, msg, want)
	}
}
