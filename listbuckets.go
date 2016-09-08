package main

import (
    "github.com/minio/minio-go"
    "log"
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

    // Dump some info about the successfully established connection
    log.Println("%v", minioClient)

    // Simply list buckets
    buckets, err := minioClient.ListBuckets()
    if err != nil {
        log.Fatalln(err)
    }

    for _, bucket := range buckets {
        log.Println(bucket)
    }
}

