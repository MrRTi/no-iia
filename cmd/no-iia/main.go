package main

import (
	"fmt"
	"log"

	"github.com/MrRTi/no-iia/pkg/files"
	"github.com/MrRTi/no-iia/pkg/translit"
)

func main() {
	filePath := "./ru-pnames-list/lists/female_names_rus.txt"
	rows, err := files.Read(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(rows); i++ {
		eng := translit.Transliterate(rows[i])
		fmt.Println(rows[i] + "\t" + eng)
	}
}
