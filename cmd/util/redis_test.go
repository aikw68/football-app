package util

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

// 初期処理_テスト用redisサーバ作成
func NewMockRedis(t *testing.T) *redis.Client {
	t.Helper()

	// redisサーバ作成
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("unexpected error while createing test redis server '%#v'", err)
	}

	// *redis.Clientを用意
	conn := redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "",
		DB:       0,
	})
	return conn
}

// 正常系
func TestNewSession_Normal(t *testing.T) {

}

// 異常系
func TestNewSession_Abnormal(t *testing.T) {

}

// 正常系
func TestGetSession_Normal(t *testing.T) {

}

// 異常系
func TestGetSession_Abnormal(t *testing.T) {

}

// 正常系
func TestDeleteSession_Normal(t *testing.T) {

}

// 異常系
func TestDeleteSession_Abnormal(t *testing.T) {

}

// 正常系
func TestNewMatchDataCache_Normal(t *testing.T) {

}

// 異常系
func TestNewMatchDataCache_Abnormal(t *testing.T) {

}

// 正常系
func TestGetMatchDataCacher_Normal(t *testing.T) {

}

// 異常系
func TestGetMatchDataCacher_Abnormal(t *testing.T) {

}
