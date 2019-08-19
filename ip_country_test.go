package ipCountry

import (
	"testing"
)

const path = "/home/di/Downloads/ips.csv"

func TestGetDB(t *testing.T) {
	df, err := getDB(path)
	if err != nil {
		t.Errorf("Can't open file with db %s", path)
	}
	dfFirst := df[1]
	if dfFirst.CountryCode != "US" {
		t.Errorf("CountryCode should equal US")
	}
}

func TestOneIP(t *testing.T) {
	ips := []string{"5.34.169.32", "5.34.169.32"}
	res := GetIPData(ips, path)
	if res[0].CountryName != "Angola" {
		t.Errorf("Country name 0 should be Angola")
	}
	if res[1].CountryName != "Angola" {
		t.Errorf("Country name 1 should be Angola")
	}
}



