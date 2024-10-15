package handler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func isValid(name string) bool {
	if len(name) < 3 || len(name) > 63 {
		return false
	}
	if name[0] == '.' || name[0] == '-' || name[len(name)-1] == '.' || name[len(name)-1] == '-' {
		return false
	}
	return true
}

func saveBucket(bucketName string) {
	metadata := filepath.Join(bucketName + "buckets.csv")

	file, err := os.OpenFile(metadata, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		bucketName,
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		"active",
	}
	writer.Write(record)
}

func CreatBucketHandler(w http.ResponseWriter, r *http.Request, directory string) {
	bucketName := r.URL.Path[1:]
	if r.Method != http.MethodPut {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if !isValid(bucketName) {
		http.Error(w, "Invalid bucket name", http.StatusMethodNotAllowed)
		return
	}
	dir := filepath.Join(directory, bucketName)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		http.Error(w, "Bucket already exists", http.StatusConflict)
		return
	}
	err := os.Mkdir(directory, 0o775)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	saveBucket(bucketName)
	w.WriteHeader(http.StatusOK)
}

func ListBucketHandler(w http.ResponseWriter, r *http.Request) {
}

func DeleteBucketHandler(w http.ResponseWriter, r *http.Request) {
}
