package s3

import (
	"os"
	"testing"

	"github.com/crowdmob/goamz/aws"
	"github.com/docker/docker-registry/storagedriver"
	"github.com/docker/docker-registry/storagedriver/testsuites"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("S3_BUCKET")
	encrypt := os.Getenv("S3_ENCRYPT")

	s3DriverConstructor := func() (storagedriver.StorageDriver, error) {
		return NewDriver(accessKey, secretKey, aws.GetRegion(region), true, bucket)
	}

	skipCheck := func() string {
		if accessKey == "" || secretKey == "" || region == "" || bucket == "" || encrypt == "" {
			return "Must set ACCESS_KEY, SECRET_KEY, AWS_REGION, S3_BUCKET, and S3_ENCRYPT to run S3 tests"
		}
		return ""
	}

	testsuites.RegisterInProcessSuite(s3DriverConstructor, skipCheck)
	testsuites.RegisterIPCSuite("s3", map[string]string{"accessKey": accessKey, "secretKey": secretKey, "region": region, "bucket": bucket, "encrypt": encrypt}, skipCheck)
}
