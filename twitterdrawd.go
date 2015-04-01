package main

import (
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	debug = kingpin.Flag("debug", "Print debug information").Bool()
	campaign_file = kingpin.Arg("campaing-file", "Campaign file configuration").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
}
