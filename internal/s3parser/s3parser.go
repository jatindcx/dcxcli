package s3parser

import (
	"dcxcli/pkg/types"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"io"
)

var (
	bucketNames        string
	objectKey          string
	ignoreSuffixes     string
	ignoreSuffixesList []string
	bucketsList        map[string]string
)

func Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&bucketNames, "buckets", "b", "", "S3 bucket name to parse")
	cmd.Flags().StringVarP(&objectKey, "object", "o", "", "S3 object key to parse")
	cmd.Flags().StringVarP(&ignoreSuffixes, "ignore-suffixes", "i", "", "Ignore URLS with these suffixes (comma-separated)")
}

func S3ParserCommand(ctx *types.Context, cmd *cobra.Command, _ []string) {
	setDefaults()

	s3Client, err := getS3Client(ctx)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("Failed to load AWS config: %v", err))

		return
	}

	_, err = getData(ctx, s3Client)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("Failed to get object from S3: %v", err))

		return
	}
}

func getData(ctx *types.Context, s3Client *s3.Client) ([]byte, error) {
	var finalResponse []byte

	for bucket, prefix := range bucketsList {
		latestKey, err := getLatestObject(ctx, s3Client, bucket, prefix)
		if err != nil {
			return nil, fmt.Errorf("failed to get latest object from S3: %v", err)
		}

		resp, err := s3Client.GetObject(ctx.Ctx, &s3.GetObjectInput{
			Bucket: &bucket,
			Key:    latestKey.Key,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get object from S3 bucket %s with key %s: %w", bucketNames, objectKey, err)
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)

		if len(finalResponse) != 0 {
			finalResponse = append(finalResponse, '\n')
		}

		finalResponse = append(finalResponse, data...)
	}

	return finalResponse, nil
}

func getS3Client(ctx *types.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx.Ctx)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func parse(ctx *types.Context, body []byte) error {
	return nil
}
