# ip-country
An offline tool to get country by IP


## Requirements
- IP2Location™ LITE IP-COUNTRY-REGION-CITY Database

## Installation

1. Download IP2Location database for IPv4 from [https://lite.ip2location.com/database/ip-country-region-city](here)

2. Install:

   ```
   go get https://github.com/dmitriiweb/go-ip-country
   ```

   

## Usage

```
package main

import (
	"fmt"
	"github.com/dmitriiweb/ip_country"
)


func main() {
	path := "/home/di/Downloads/ips.csv"
	ips := []string{"5.34.169.32", "4.69.200.225"}

	res := ipCountry.GetIPData(ips, path)
	for _, row := range res {
		fmt.Printf("IP: %s\nCountry: %s\nCity: %s\n\n", row.IP, row.CountryName, row.CityName)
	}
}
```

