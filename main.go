package main

import (
	"fmt"
	"log"

	"github.com/contest/db"
)

func main() {
	/*
		b, err := ioutil.ReadFile("e:\\2.json")
		if err != nil {
			log.Fatal("read json file failed ", err)
		}

		nc := db.Memberconfig{}
		//nc1 := db.Testconfig{}

		err = json.Unmarshal(b, &nc)
		if err != nil {
			log.Fatal("json unmarshal failed ", err)
		}

		fmt.Printf("%+v\n", nc)


			if nc.Rules != nil {
				rules := nc.Rules.([]interface{})

				for _, u := range rules {
					fmt.Println(u)
				}

			}*/

	fdb, err := db.NewFileDB("E:\\go\\test\\root\\controller.d")
	if err != nil {
		log.Fatal("New filedb failed", err)
	}

	fmt.Printf("%+v", fdb)

}
