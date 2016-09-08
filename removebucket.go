package main

import (
    "github.com/minio/minio-go"
    "log"
    "os"
)

func main() {
    endpoint := "192.168.10.129:9000"
    accessKeyID := "LJ93AK3HYCDAHV4OISOI"
    secretAccessKey := "Ws3+Gfyq7sNytuXzgvKPUGA0Uhev+ebAeIN8Wpoy"
    useSSL := false
  
    // Bucket name passed as a command line argument
    bucketName := os.Args[1]

    // Initialize a new connection
    minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
    if err != nil {
        log.Fatalln(err)
    }
    
    // Create a new bucket
    err = minioClient.RemoveBucket(bucketName)
    if err != nil {
        log.Fatalln(err)
    }
   
    log.Printf("Successfully deleted bucket %s", bucketName)
}
