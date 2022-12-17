package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 正常系
// APIコールが問題なく実行できること
func TestCall_Normal(t *testing.T) {
	actual, err := GetSecret("football-data_auth_APIkey")
	assert.NotNil(t, actual)
	assert.Nil(t, err)
}

// 異常系
// APIコールで指定のエラーが発生すること。
// 「システムエラーが発生しました。時間を空けて再度入力してください。」
func TestCall_Abormal(t *testing.T) {
	_, err := GetSecret("")
	assert.Equal(t, ERR_USER_SYSTEM_ERROR.Error(), err.Error())
}
