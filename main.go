package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	// list out all of my buckets
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []types.Bucket
	if err != nil {
		log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	} else {
		buckets = result.Buckets
	}

	for _, bucket := range buckets {
		// BUCKET NAME IS A PTR STRING
		fmt.Println("Bucket:", *bucket.Name)

		// get the bucket's objects
		objects, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: bucket.Name,
		})
		if err != nil {
			log.Printf("Couldn't list objects for your account. Here's why: %v\n", err)
		} else {
			for _, object := range objects.Contents {
				// OBJECT NAME IS A PTR STRING
				fmt.Println("Object:", *object.Key)
			}
		}
	}
}
