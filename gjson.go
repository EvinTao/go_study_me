package main

import "github.com/tidwall/gjson"


func main() {
	value := gjson.Get(json, "data.proxy_pool.*.#(city_code=371600)")
	print(value.String())
}