package config

type Aws struct {
	Profile    string
	Region     string
	AssumeRole *AwsAssumeRole
}
