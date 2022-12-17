package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func setup() {
	redisServer = mockRedis()
	conn = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
}

func teardown() {
	redisServer.Close()
}

// 正常系（セッション生成）
func TestNewSession_Normal(t *testing.T) {
	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)
	w := httptest.NewRecorder()

	err := NewSession(w, r, "test", "test")
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("HTTP status code:%d", w.Code)
	}
	assert.Nil(t, err)
}

// 正常系（セッション取得)
func TestGetSession_Normal(t *testing.T) {
	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// rediskeyを生成
	b := make([]byte, 64)
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// Cookieを登録
	r.AddCookie(&http.Cookie{
		Name:  "cookieKey",
		Value: newRedisKey,
	})

	// セッションを登録
	conn.Set(r.Context(), newRedisKey, "redisValueTest", 0)

	rtn, err := GetSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "redisValueTest", rtn)
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
