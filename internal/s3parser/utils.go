package s3parser

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"

	"dcxcli/pkg/types"
)

func getS3Client(ctx *types.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx.Ctx)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func getLatestObject(ctx *types.Context, s3Client *s3.Client, bucketName, prefix string) (*s3Types.Object, error) {
	listObjectsOutput, err := s3Client.ListObjectsV2(ctx.Ctx, &s3.ListObjectsV2Input{
		Bucket: &bucketName,
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(listObjectsOutput.Contents, func(i, j int) bool {
		return listObjectsOutput.Contents[i].LastModified.After(*listObjectsOutput.Contents[j].LastModified)
	})

	if len(listObjectsOutput.Contents) == 0 {
		return nil, fmt.Errorf("no objects found in bucket %s", bucketName)
	}

	latestObject := listObjectsOutput.Contents[0]

	return &latestObject, nil
}

func getObjects(ctx *types.Context, s3Client *s3.Client, bucketName, prefix string) ([]s3Types.Object, error) {
	listObjectsOutput, err := s3Client.ListObjectsV2(ctx.Ctx, &s3.ListObjectsV2Input{
		Bucket: &bucketName,
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(listObjectsOutput.Contents, func(i, j int) bool {
		return listObjectsOutput.Contents[i].LastModified.After(*listObjectsOutput.Contents[j].LastModified)
	})

	if len(listObjectsOutput.Contents) == 0 {
		return nil, fmt.Errorf("no objects found in bucket %s", bucketName)
	}

	return listObjectsOutput.Contents, nil
}

func setDefaults() {
	bucketsList = map[string]string{
		"vpc-resolver-query-logs-ap-northeast-1-106328716869": "resolver-logs/AWSLogs/106328716869/vpcdnsquerylogs/vpc-0d0eb0edc77e67efc/2025/07/23/",
	}

	ignoreSuffixesList = []string{
		"amazonaws.com",
		"compute.internal",
		"cloud.rlrcp.com",
	}
	if len(ignoreSuffixes) != 0 {
		ignoreSuffixesList = append(ignoreSuffixesList, strings.Split(ignoreSuffixes, ",")...)
	}

	//	todo set for objectKey
}
