package main

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

func main() {
	//file, _ := os.Open("go_ini.ini")
	//defer file.Close()

	cfg, err := ini.Load("D:\\go_study\\src\\study7_config_analysis\\go_ini.ini")
	//cfg, err := ini.Load("src/study7_config_analysis/go_ini.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	fmt.Println("cfg: ", cfg)

	for _, section := range cfg.Sections() {
		fmt.Println(section.Name())
	}
	fmt.Println("ini: ", cfg.Section("NAME").Key("a").String())

	fmt.Println(cfg.HasSection("NAME"))

	if cfg.HasSection("NAME") {
		fmt.Println("find")
	} else {
		fmt.Println("cannot find")
	}

	fmt.Println(cfg.GetSection("NAME"))
	s1, _ := cfg.GetSection("NAME")
	fmt.Println(s1.Key("a"))
}
