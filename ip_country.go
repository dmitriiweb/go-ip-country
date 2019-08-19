package ipCountry

import (
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"net"
	"os"
	"strconv"
)

type IP struct {
	IPFrom      uint32
	IPTo        uint32
	CountryCode string
	CountryName string
	RegionName  string
	CityName    string
}

type IPData struct {
	IP string
	CountryCode string
	CountryName string
	RegionName  string
	CityName    string
	Error string
}

func GetIPData(ips []string, pathToDB string) []*IPData {
	df, err := getDB(pathToDB)
	if err != nil {
		panic(err)
	}
	var ipDatas []*IPData
	for _, ip := range ips {
		ipUint, err := ipToUint(ip)
		newRecord := IPData{}
		if err != nil {
			newRecord.IP = ip
			newRecord.Error = err.Error()
		} else {
			id := searchIPData(ipUint, df)
			newRecord.IP = ip
			newRecord.RegionName = id.RegionName
			newRecord.CityName = id.CityName
			newRecord.CountryName = id.CountryName
			newRecord.CountryCode = id.CountryCode
		}
		ipDatas = append(ipDatas, &newRecord)

	}
	return ipDatas
}

func searchIPData(ip uint32, df []*IP) *IP {
	for _, row := range df {
		if row.IPFrom <= ip && row.IPTo >= ip {
			return row
		}
	}
	return &IP{0, 0, "", "", "", ""}
}

func ipToUint(ip string) (uint32, error) {
	var ipUint uint32
	err := binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &ipUint)
	if err != nil {
		return 0, err
	}
	return ipUint, err
}

func getDB(path string) ([]*IP, error) {
	var rows []*IP
	csvfile, err := os.Open(path)
	if err != nil {
		return rows, err
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		return rows, err
	}

	for _, row := range rawCSVData {
		newRecord := &IP{strToUint(row[0]), strToUint(row[1]), row[2], row[3],
			row[4], row[5]}
		rows = append(rows, newRecord)
	}

	return rows, err
}

func strToUint(txt string) uint32 {
	u, err := strconv.ParseUint(txt, 10, 32)
	if err != nil {
		u = 0
	}

	return uint32(u)
}
