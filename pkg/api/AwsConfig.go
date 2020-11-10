package api

type AwsAssumeRole struct {
	RoleArn string
}

type AwsConfig struct {
	Profile    string
	Region     string
	AssumeRole *AwsAssumeRole
}
