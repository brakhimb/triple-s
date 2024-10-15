package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"tripleS/pkg/handler"
)

func main() {
	help := flag.Bool("help", false, "help")
	port := flag.String("port", "8080", "port")
	directory := flag.String("directory", "repository", "directory")

	flag.Parse()
	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", *directory)
	}
	if *help {
		fmt.Print("Simple Storage Service.\n\n" +
			"**Usage:**\n" +
			"    triple-s [-port <N>] [-dir <S>]  \n" +
			"    triple-s --help\n\n" +
			"**Options:**\n" +
			"- --help     Show this screen.\n" +
			"- --port N   Port number\n" +
			"- --dir S    Path to the directory")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			handler.CreatBucketHandler(w, r, *directory)
		case "GET":
			handler.ListBucketHandler(w, r)
		case "DELETE":
			handler.DeleteBucketHandler(w, r)
		}
	})
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
