package util

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pkg/errors"
)

// 認証情報取得(AWS Secrets Managerへcall)
func GetSecret(param string) (map[string]interface{}, error) {

	secretName := param
	region := "ap-northeast-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, errors.WithStack(ERR_USER_SYSTEM_ERROR)
	}

	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, errors.WithStack(ERR_USER_SYSTEM_ERROR)
	}

	var secretString string = *result.SecretString

	res := make(map[string]interface{})
	json.Unmarshal([]byte(secretString), &res)

	return res, nil
}
