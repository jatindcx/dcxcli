package s3parser

import (
	"dcxcli/pkg/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"strings"

	"sort"
)

func getLatestObject(ctx *types.Context, s3Client *s3.Client, bucketName string) (*s3Types.Object, error) {
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

	latestObject := listObjectsOutput.Contents[0]

	return &latestObject, nil
}

func setDefaults() {
	bucketsList = []string{
		"vpc-resolver-query-logs-ap-northeast-1-024221098369",
		"vpc-resolver-query-logs-ap-south-1-024221098369",
		"vpc-resolver-query-logs-ap-southeast-1-024221098369",
		"vpc-resolver-query-logs-me-south-1-024221098369",
		"vpc-resolver-query-logs-us-east-1-024221098369",
		"vpc-resolver-query-logs-ap-northeast-1-106328716869",
		"vpc-resolver-query-logs-ap-south-1-106328716869",
		"vpc-resolver-query-logs-ap-south-2-106328716869",
		"vpc-resolver-query-logs-ap-southeast-1-106328716869",
		"vpc-resolver-query-logs-us-east-1-106328716869",
		"vpc-resolver-query-logs-ap-northeast-1-410639636265",
		"vpc-resolver-query-logs-ap-south-1-410639636265",
		"vpc-resolver-query-logs-ap-south-2-410639636265",
		"vpc-resolver-query-logs-ap-southeast-1-410639636265",
		"vpc-resolver-query-logs-us-east-1-410639636265",
		"vpc-resolver-query-logs-us-east-1-821727966112",
		"vpc-resolver-query-logs-ap-northeast-1-693001631719",
		"vpc-resolver-query-logs-ap-south-2-693001631719",
		"vpc-resolver-query-logs-ap-southeast-1-693001631719",
		"vpc-resolver-query-logs-me-south-1-693001631719",
		"vpc-resolver-query-logs-us-east-1-693001631719",
		"vpc-resolver-query-logs-ap-northeast-1-147836456796",
		"vpc-resolver-query-logs-ap-south-1-147836456796",
		"vpc-resolver-query-logs-ap-southeast-1-147836456796",
		"vpc-resolver-query-logs-us-east-1-147836456796",
		"vpc-resolver-query-logs-ap-northeast-1-688619497087",
		"vpc-resolver-query-logs-ap-south-1-688619497087",
		"vpc-resolver-query-logs-ap-southeast-1-688619497087",
		"vpc-resolver-query-logs-us-east-1-688619497087",
		"vpc-resolver-query-logs-ap-northeast-1-127902694105",
		"vpc-resolver-query-logs-ap-south-1-127902694105",
		"vpc-resolver-query-logs-ap-southeast-1-127902694105",
		"vpc-resolver-query-logs-us-east-1-127902694105",
		"vpc-resolver-query-logs-ap-northeast-1-318622992493",
		"vpc-resolver-query-logs-ap-south-1-318622992493",
		"vpc-resolver-query-logs-ap-southeast-1-318622992493",
		"vpc-resolver-query-logs-us-east-1-318622992493",
		"vpc-resolver-query-logs-ap-northeast-1-835838111781",
		"vpc-resolver-query-logs-ap-south-1-835838111781",
		"vpc-resolver-query-logs-ap-southeast-1-835838111781",
		"vpc-resolver-query-logs-us-east-1-835838111781",
		"vpc-resolver-query-logs-ap-northeast-1-323308110940",
		"vpc-resolver-query-logs-ap-south-1-323308110940",
		"vpc-resolver-query-logs-ap-southeast-1-323308110940",
		"vpc-resolver-query-logs-us-east-1-323308110940",
		"vpc-resolver-query-logs-ap-northeast-1-672019911484",
		"vpc-resolver-query-logs-ap-south-1-672019911484",
		"vpc-resolver-query-logs-ap-southeast-1-672019911484",
		"vpc-resolver-query-logs-us-east-1-672019911484",
		"vpc-resolver-query-logs-ap-northeast-1-200587739478",
		"vpc-resolver-query-logs-ap-south-1-200587739478",
		"vpc-resolver-query-logs-ap-southeast-1-200587739478",
		"vpc-resolver-query-logs-us-east-1-200587739478",
		"vpc-resolver-query-logs-ap-northeast-1-477252799259",
		"vpc-resolver-query-logs-ap-south-1-477252799259",
		"vpc-resolver-query-logs-ap-southeast-1-477252799259",
		"vpc-resolver-query-logs-us-east-1-477252799259",
		"vpc-resolver-query-logs-ap-northeast-1-287002870177",
		"vpc-resolver-query-logs-ap-south-1-287002870177",
		"vpc-resolver-query-logs-ap-southeast-1-287002870177",
		"vpc-resolver-query-logs-us-east-1-287002870177",
		"vpc-resolver-query-logs-ap-northeast-1-509536004829",
		"vpc-resolver-query-logs-ap-south-1-509536004829",
		"vpc-resolver-query-logs-ap-southeast-1-509536004829",
		"vpc-resolver-query-logs-us-east-1-509536004829",
		"vpc-resolver-query-logs-ap-northeast-1-550716552410",
		"vpc-resolver-query-logs-ap-south-1-550716552410",
		"vpc-resolver-query-logs-ap-southeast-1-550716552410",
		"vpc-resolver-query-logs-us-east-1-550716552410",
		"vpc-resolver-query-logs-ap-northeast-1-522003900580",
		"vpc-resolver-query-logs-ap-south-1-522003900580",
		"vpc-resolver-query-logs-ap-southeast-1-522003900580",
		"vpc-resolver-query-logs-us-east-1-522003900580",
		"vpc-resolver-query-logs-ap-northeast-1-761938460123",
		"vpc-resolver-query-logs-ap-south-1-761938460123",
		"vpc-resolver-query-logs-ap-southeast-1-761938460123",
		"vpc-resolver-query-logs-us-east-1-761938460123",
		"vpc-resolver-query-logs-ap-northeast-1-503831806232",
		"vpc-resolver-query-logs-ap-south-1-503831806232",
		"vpc-resolver-query-logs-ap-southeast-1-503831806232",
		"vpc-resolver-query-logs-me-south-1-503831806232",
		"vpc-resolver-query-logs-us-east-1-503831806232",
		"vpc-resolver-query-logs-ap-northeast-1-188187289566",
		"vpc-resolver-query-logs-ap-south-1-188187289566",
		"vpc-resolver-query-logs-ap-southeast-1-188187289566",
		"vpc-resolver-query-logs-us-east-1-188187289566",
		"vpc-resolver-query-logs-ap-northeast-1-212257683650",
		"vpc-resolver-query-logs-ap-south-1-212257683650",
		"vpc-resolver-query-logs-ap-southeast-1-212257683650",
		"vpc-resolver-query-logs-me-south-1-212257683650",
		"vpc-resolver-query-logs-us-east-1-212257683650",
		"vpc-resolver-query-logs-ap-northeast-1-006240882047",
		"vpc-resolver-query-logs-ap-south-1-006240882047",
		"vpc-resolver-query-logs-ap-southeast-1-006240882047",
		"vpc-resolver-query-logs-eu-central-1-006240882047",
		"vpc-resolver-query-logs-us-east-1-006240882047",
		"vpc-resolver-query-logs-ap-northeast-1-892930679478",
		"vpc-resolver-query-logs-ap-south-1-892930679478",
		"vpc-resolver-query-logs-ap-southeast-1-892930679478",
		"vpc-resolver-query-logs-eu-central-1-892930679478",
		"vpc-resolver-query-logs-us-east-1-892930679478",
		"vpc-resolver-query-logs-ap-northeast-1-952252268202",
		"vpc-resolver-query-logs-ap-south-1-952252268202",
		"vpc-resolver-query-logs-ap-southeast-1-952252268202",
		"vpc-resolver-query-logs-me-south-1-952252268202",
		"vpc-resolver-query-logs-us-east-1-952252268202",
		"vpc-resolver-query-logs-ap-south-1-202533531934",
		"vpc-resolver-query-logs-me-south-1-202533531934",
		"vpc-resolver-query-logs-us-east-1-202533531934",
		"vpc-resolver-query-logs-ap-south-1-699475919553",
		"vpc-resolver-query-logs-me-south-1-699475919553",
		"vpc-resolver-query-logs-us-east-1-699475919553",
		"vpc-resolver-query-logs-ap-south-1-051826733157",
		"vpc-resolver-query-logs-me-south-1-051826733157",
		"vpc-resolver-query-logs-us-east-1-051826733157",
		"vpc-resolver-query-logs-ap-northeast-1-650251690667",
		"vpc-resolver-query-logs-ap-northeast-2-650251690667",
		"vpc-resolver-query-logs-ap-northeast-3-650251690667",
		"vpc-resolver-query-logs-ap-south-1-650251690667",
		"vpc-resolver-query-logs-ap-south-2-650251690667",
		"vpc-resolver-query-logs-ap-southeast-1-650251690667",
		"vpc-resolver-query-logs-ap-southeast-2-650251690667",
		"vpc-resolver-query-logs-ca-central-1-650251690667",
		"vpc-resolver-query-logs-eu-central-1-650251690667",
		"vpc-resolver-query-logs-eu-north-1-650251690667",
		"vpc-resolver-query-logs-eu-west-1-650251690667",
		"vpc-resolver-query-logs-eu-west-2-650251690667",
		"vpc-resolver-query-logs-eu-west-3-650251690667",
		"vpc-resolver-query-logs-sa-east-1-650251690667",
		"vpc-resolver-query-logs-us-east-1-650251690667",
		"vpc-resolver-query-logs-us-east-2-650251690667",
		"vpc-resolver-query-logs-us-west-1-650251690667",
		"vpc-resolver-query-logs-us-west-2-650251690667",
		"vpc-resolver-query-logs-ap-northeast-1-490004645851",
		"vpc-resolver-query-logs-ap-northeast-2-490004645851",
		"vpc-resolver-query-logs-ap-northeast-3-490004645851",
		"vpc-resolver-query-logs-ap-south-1-490004645851",
		"vpc-resolver-query-logs-ap-south-2-490004645851",
		"vpc-resolver-query-logs-ap-southeast-1-490004645851",
		"vpc-resolver-query-logs-ap-southeast-2-490004645851",
		"vpc-resolver-query-logs-ca-central-1-490004645851",
		"vpc-resolver-query-logs-eu-central-1-490004645851",
		"vpc-resolver-query-logs-eu-north-1-490004645851",
		"vpc-resolver-query-logs-eu-west-1-490004645851",
		"vpc-resolver-query-logs-eu-west-2-490004645851",
		"vpc-resolver-query-logs-eu-west-3-490004645851",
		"vpc-resolver-query-logs-sa-east-1-490004645851",
		"vpc-resolver-query-logs-us-east-1-490004645851",
		"vpc-resolver-query-logs-us-east-2-490004645851",
		"vpc-resolver-query-logs-us-west-1-490004645851",
		"vpc-resolver-query-logs-us-west-2-490004645851",
		"vpc-resolver-query-logs-ap-south-1-785316066641",
		"vpc-resolver-query-logs-us-east-1-785316066641",
		"vpc-resolver-query-logs-ap-south-1-675613596333",
		"vpc-resolver-query-logs-us-east-1-675613596333",
		"vpc-resolver-query-logs-ap-south-1-731493185918",
		"vpc-resolver-query-logs-us-east-1-731493185918",
		"vpc-resolver-query-logs-ap-south-1-585620823823",
		"vpc-resolver-query-logs-us-east-1-585620823823",
		"vpc-resolver-query-logs-ap-south-1-011563687552",
		"vpc-resolver-query-logs-me-south-1-011563687552",
	}
	if len(bucketNames) != 0 {
		bucketsList = append(bucketsList, strings.Split(bucketNames, ",")...)
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
