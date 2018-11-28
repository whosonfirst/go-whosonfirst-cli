package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-cli/dsn"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"log"
)

func main() {

	str_dsn := flag.String("dsn", "", "...")

	var keys flags.MultiString
	flag.Var(&keys, "key", "...")

	flag.Parse()

	dsn_map, err := dsn.StringToDSNWithKeys(*str_dsn, keys...)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(dsn_map)
}
