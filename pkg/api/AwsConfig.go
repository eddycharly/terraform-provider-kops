package api

type AwsAssumeRole struct {
	RoleArn string
}

type AwsConfig struct {
	Profile    string
	AssumeRole *AwsAssumeRole
}
