package s3parser

import (
	"bufio"
	"bytes"
	"dcxcli/pkg/types"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"io"
	"net/url"
	"os"
	"strings"
)

var (
	logFile            string
	vendorFile         string
	bucketNames        string
	objectKey          string
	ignoreSuffixes     string
	ignoreSuffixesList []string
	bucketsList        []string
	env                string
)

type DNSLog struct {
	Version        string `json:"version"`
	AccountID      string `json:"account_id"`
	Region         string `json:"region"`
	VPCID          string `json:"vpc_id"`
	QueryTimestamp string `json:"query_timestamp"`
	QueryName      string `json:"query_name"`
	QueryType      string `json:"query_type"`
	QueryClass     string `json:"query_class"`
	RCode          string `json:"rcode"`
	Answers        []struct {
		Rdata string `json:"Rdata"`
		Type  string `json:"Type"`
		Class string `json:"Class"`
	} `json:"answers"`
	SrcAddr   string `json:"srcaddr"`
	SrcPort   string `json:"srcport"`
	Transport string `json:"transport"`
	SrcIDs    struct {
		Instance string `json:"instance"`
	} `json:"srcids"`
}

func Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&bucketNames, "buckets", "b", "", "S3 bucket name to parse")
	cmd.Flags().StringVarP(&logFile, "logFile", "l", "", "CSV file in log")
	cmd.Flags().StringVarP(&vendorFile, "vendorFile", "v", "", "CSV file from Krishna's vendor file")
	cmd.Flags().StringVarP(&objectKey, "object", "o", "", "S3 object key to parse")
	cmd.Flags().StringVarP(&env, "env", "e", "Staging", "Environment to use (default: Staging)")
	//cmd.Flags().StringVarP(&ignoreSuffixes, "ignore-suffixes", "i", "", "Ignore URLS with these suffixes (comma-separated)")
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
		"amazonaws.com.",
		"compute.internal.",
		"cloud.rlrcp.com.",
	}
	if len(ignoreSuffixes) != 0 {
		ignoreSuffixesList = append(ignoreSuffixesList, strings.Split(ignoreSuffixes, ",")...)
	}

	//	todo set for objectKey
}

func S3ParserCommand(ctx *types.Context, cmd *cobra.Command, _ []string) {
	setDefaults()

	//s3Client, err := getS3Client(ctx)
	//if err != nil {
	//	ctx.Logger.Error(fmt.Sprintf("Failed to load AWS config: %v", err))
	//
	//	return
	//}

	//resp, err := getObject(ctx, s3Client)
	//if err != nil {
	//	ctx.Logger.Error(fmt.Sprintf("Failed to get object from S3: %v", err))
	//
	//	return
	//}

	// Parse the response
	//err = parse(ctx, resp)
	//if err != nil {
	//	ctx.Logger.Error(fmt.Sprintf("Failed to parse response: %v", err))
	//}
}

func getObject(ctx *types.Context, s3Client *s3.Client) ([]byte, error) {
	var finalResponse []byte

	for _, bucket := range bucketsList {
		resp, err := s3Client.GetObject(ctx.Ctx, &s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &objectKey,
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

func S3Parse(ctx *types.Context, body []byte) (domains []string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(body))

	var ignored []string
	ignored = append(ignored, "cloud.rlrcp.com.")
	ignored = append(ignored, "compute.internal.")
	ignored = append(ignored, "amazonaws.com.")

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var entry DNSLog
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			ctx.Logger.Warn(fmt.Sprintf("Failed to parse line: %v \n", err))
			continue
		}

		ok := false
		for _, suffix := range ignored {
			if strings.HasSuffix(entry.QueryName, suffix) {
				ok = true
			}
		}

		if ok == false {
			cleanDomain := strings.TrimSuffix(entry.QueryName, ".")
			domains = append(domains, cleanDomain)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return domains, nil
}

func ParseCommand(ctx *types.Context, _ *cobra.Command, _ []string) {
	data1, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	data2, err := os.ReadFile(vendorFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	file1, err := CSVParse(ctx, data1)
	if err != nil {
		return
	}

	s3OutFile, err := os.Create("s3_imported_external_file.csv")
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("Error creating s3_imported_external_file.csv: %v", err))
		return
	}
	defer s3OutFile.Close()

	s3Writer := csv.NewWriter(s3OutFile)
	defer s3Writer.Flush()

	// Header
	s3Writer.Write([]string{"Domain"})

	// Write each domain
	for _, domain := range file1 {
		s3Writer.Write([]string{domain})
	}

	println("length of file1 is ", len(file1))

	file2, err := CSVParse(ctx, data2)
	if err != nil {
		return
	}
	println("length of file2 is ", len(file2))

	set1 := make(map[string]bool)
	for _, domain := range file1 {
		set1[domain] = true
	}

	set2 := make(map[string]bool)
	for _, domain := range file2 {
		set2[domain] = true
	}

	result := make(map[string]string)

	// Mark all domains from file1
	for domain := range set1 {
		if set2[domain] {
			result[domain] = "in both files"
		} else {
			result[domain] = "from S3 bucket"
		}
	}

	// Add domains from file2 that weren't in file1
	for domain := range set2 {
		if _, exists := result[domain]; !exists {
			result[domain] = "from Krishna's file"
		}
	}

	// Write output CSV
	outFile, err := os.Create("output_comparison.csv")
	if err != nil {
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Header
	writer.Write([]string{"Domain", "Source"})

	// Write results
	for domain, source := range result {
		writer.Write([]string{domain, source})
	}

	println("final length of string is ", len(result))

	ctx.Logger.Info("output_comparison.csv generated.")
}

func CSVParse(ctx *types.Context, data []byte) ([]string, error) {
	reader := csv.NewReader(bytes.NewReader(data))

	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	ignored := []string{
		"cloud.rlrcp.com.",
		"compute.internal.",
		"amazonaws.com.",
	}

	prodIndex := -1
	for i, col := range headers {
		if strings.TrimSpace(col) == env || strings.TrimSpace(col) == "query_name" {
			prodIndex = i
			break
		}
	}

	if prodIndex == -1 {
		return nil, fmt.Errorf("Production column not found")
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var domains []string
	for _, row := range records {
		if prodIndex >= len(row) {
			continue
		}

		rawURL := strings.TrimSpace(row[prodIndex])
		if rawURL == "" {
			continue
		}

		// Normalize the URL
		if !strings.HasPrefix(rawURL, "http://") &&
			!strings.HasPrefix(rawURL, "https://") &&
			!strings.HasPrefix(rawURL, "wss://") &&
			!strings.HasPrefix(rawURL, "ws://") {
			rawURL = "https://" + rawURL
		}

		parsed, err := url.Parse(rawURL)
		if err != nil || parsed.Host == "" {
			continue
		}

		domain := parsed.Host

		// Check if domain ends with any ignored suffix
		shouldIgnore := false
		for _, suffix := range ignored {
			if strings.HasSuffix(domain, suffix) {
				shouldIgnore = true
				break
			}
		}
		if shouldIgnore {
			continue
		}

		domains = append(domains, domain)
	}

	return domains, nil
}
