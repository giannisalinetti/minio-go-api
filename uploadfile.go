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

    // Connect to local minio instance
    minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
    if err != nil {
        log.Fatalln(err)
    }

    // Upload file to a choosen bucket
    bucketName := os.Args[1]
    fileName := os.Args[2]
    objectName := os.Args[3]
    contentType := os.Args[4]

    n, err := minioClient.FPutObject(bucketName, objectName, fileName, contentType)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Successfully uploaded file %s of size %d\n", fileName, n) 
}
