package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		exitErrorf("Bucket name required\nUsage: %s bucket_name",
			os.Args[0])
	}

	bucket := os.Args[1]
	s3Client := s3Client()
	maxKeys := int64(1)
	var token *string = nil

	for {
		resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket), ContinuationToken: token, MaxKeys: &maxKeys})

		if err != nil {
			exitErrorf("Error listing bucket %q: %v", bucket, err)
		}
		for _, item := range resp.Contents {
			fmt.Println(*item.Key)
		}
		if *resp.IsTruncated {
			token = resp.NextContinuationToken
		} else {
			break
		}
	}
}

func s3Client() *s3.S3 {
	sess := session.New(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	return s3.New(sess)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
