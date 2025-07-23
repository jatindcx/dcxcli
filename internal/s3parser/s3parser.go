package s3parser

import (
	"dcxcli/pkg/types"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

var (
	bucketName     string
	objectKey      string
	ignoreSuffixes string
)

func Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "S3 bucket name to parse")
	cmd.Flags().StringVarP(&objectKey, "object", "o", "", "S3 object key to parse")
	cmd.Flags().StringVarP(&ignoreSuffixes, "ignore-suffixes", "i", "", "Ignore URLS with these suffixes (comma-separated)")
}

func S3ParserCommand(ctx *types.Context, cmd *cobra.Command, _ []string) {
	if bucketName == "" || objectKey == "" {
		cmd.Help()

		return
	}

	ctx.Ctx = cmd.Context()

	s3Client, err := getS3Client(ctx)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("Failed to load AWS config: %v", err))

		return
	}

	resp, err := s3Client.GetObject(cmd.Context(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("failed to get object '%s' from S3: %v", objectKey, err))
	}

	defer resp.Body.Close()
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
