package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 正常系
// パスワードの暗号化と比較処理に成功すること（nilを返す）
func TestCrypto_Normal(t *testing.T) {
	hash, _ := PasswordEncrypt("test")
	actual := CompareHashAndPassword(hash, "test")
	assert.Nil(t, actual)
}

// 異常系
// パスワードの暗号化に成功するが、比較処理でエラーになること。
func TestCrypto_Abnormal(t *testing.T) {
	hash, _ := PasswordEncrypt("test")
	actual := CompareHashAndPassword(hash, "test2")
	fmt.Println(actual)
	assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", actual.Error())
}
