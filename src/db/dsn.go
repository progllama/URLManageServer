package db

import "fmt"

func BuildDNS(params map[string]string) string {
	isValid := validate(params)
	if !isValid {
		panic("DNS paramerter is not valid.")
	} else {
		return build(params)
	}
}

func validate(params map[string]string) bool {
	return true
}

func build(params map[string]string) string {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		params["host"],
		params["port"],
		params["user"],
		params["dbname"],
		params["password"],
	)
	fmt.Println(dns)
	return dns
}
