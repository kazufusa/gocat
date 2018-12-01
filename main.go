package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			fi, err := os.Open(fname)
			defer func() {
				if err := fi.Close(); err != nil {
					log.Fatal(err)
				}
			}()
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(os.Stdout, fi)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
