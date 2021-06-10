package config

type Aws struct {
	// Profile defines the AWS profile to load when calling aws services
	Profile string
	// Region defines the AWS region
	Region string
	// Region defines the AWS access key
	AccessKey string
	// Region defines the AWS secret key
	SecretKey string
	// AssumeRole defines the AWS IAM role to be assumed
	AssumeRole *AwsAssumeRole
	// S3Endpoint defines S3 compatible endpoint
	S3Endpoint string
	// S3Region defines S3 compatible endpoint region
	S3Region string
	// S3AccessKey defines S3 compatible endpoint access key
	S3AccessKey string
	// S3SecretKey defines S3 compatible endpoint secret key
	S3SecretKey string
	// SkipRegionCheck skips validating region check
	SkipRegionCheck bool
	// SkipRequestingAccountId   bool
	// SkipCredentialsValidation bool
}
