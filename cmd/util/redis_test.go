package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis

// miniredis（redisのモック）
func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

type testPage struct {
	Title         string
	SubTitle      string
	LoginFlg      bool
	MatchList     interface{}
	ScoreList     interface{}
	StandingsList interface{}
	Message       string
}

func (p testPage) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
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

	//　セッション生成
	err := NewSession(w, r, "test", "test")
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("HTTP status code:%d", w.Code)
	}

	// エラーが発生しないことを確認
	assert.Nil(t, err)
}

// 正常系（セッション登録済で取得できる場合)
func TestGetSession_Normal(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// rediskey生成
	b := make([]byte, 64)
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// Cookieを登録
	r.AddCookie(&http.Cookie{
		Name:  "cookieKey",
		Value: newRedisKey,
	})

	// セッションを登録
	conn.Set(r.Context(), newRedisKey, "redisValueTest", 0)

	// セッションを取得
	rtn, err := GetSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	// 取得結果が"redisValueTest"であることを確認
	assert.Equal(t, "redisValueTest", rtn)
}

// 正常系（セッション未登録で取得できない場合)
func TestGetSession_Normal2(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// セッションを取得
	rtn, err := GetSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	// 取得結果がブランクであることを確認
	assert.Equal(t, "", rtn)
}

// 正常系（Cookie登録済、かつ、セッション未登録で取得できない場合)
func TestGetSession_Normal3(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// rediskey生成
	b := make([]byte, 64)
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// Cookieを登録
	r.AddCookie(&http.Cookie{
		Name:  "cookieKey",
		Value: newRedisKey,
	})

	// セッションは未登録

	// セッションを取得
	rtn, err := GetSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	// 取得結果がブランクであることを確認
	assert.Equal(t, "", rtn)
}

// 異常系（保留）
func TestGetSession_Abnormal(t *testing.T) {
}

// 正常系
func TestExtendSession_Normal(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// rediskey生成
	b := make([]byte, 64)
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// Cookieを登録
	r.AddCookie(&http.Cookie{
		Name:  "cookieKey",
		Value: newRedisKey,
	})

	// セッションを登録（30分有効)
	conn.Set(r.Context(), newRedisKey, "redisValueTest", 10*time.Minute)

	conn.TTL(r.Context(), newRedisKey)

	// セッションの有効期限延長
	err := ExtendSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	lifetime := conn.TTL(r.Context(), newRedisKey)

	assert.Equal(t, 30*time.Minute, lifetime.Val())
}

// 異常系
func TestExtendSession_Abnormal(t *testing.T) {

}

// 正常系
func TestDeleteSession_Normal(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)
	w := httptest.NewRecorder()

	// rediskey生成
	b := make([]byte, 64)
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// Cookieを登録
	r.AddCookie(&http.Cookie{
		Name:  "cookieKey",
		Value: newRedisKey,
	})

	// セッションを登録
	conn.Set(r.Context(), newRedisKey, "redisValueTest", 0)

	// セッションを取得
	err := DeleteSession(w, r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	// セッションを取得
	rtn, err := GetSession(r, "cookieKey")
	if err != nil {
		fmt.Println(err)
	}

	// 取得結果がブランクであることを確認
	assert.Equal(t, "", rtn)
}

// 異常系
func TestDeleteSession_Abnormal(t *testing.T) {

}

// 正常系
func TestNewMatchDataCache_Normal(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// キャッシュ元データ初期化
	var cachePage testPage
	cache := []byte("TEST")
	if err := json.Unmarshal(cache, &cachePage); err != nil {
		fmt.Println(err)
	}

	p := testPage{"test", "test", false, cache, cache, cache, ""}

	//　キャッシュ生成
	err := NewMatchDataCache(r, "cacheKeyTest", p)
	if err != nil {
		fmt.Println(err)
	}

	// エラーが発生しないことを確認
	assert.Nil(t, err)
}

// 異常系
func TestNewMatchDataCache_Abnormal(t *testing.T) {

}

// 正常系
func TestGetMatchDataCacher_Normal(t *testing.T) {

	// miniredis初期化
	setup()
	defer teardown()

	// httptest初期化
	reqBody := bytes.NewBufferString("request body")
	r := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// キャッシュ元データ初期化
	var cachePage testPage
	cache := []byte("TEST")
	if err := json.Unmarshal(cache, &cachePage); err != nil {
		fmt.Println(err)
	}

	p := testPage{"test", "test", false, cache, cache, cache, ""}

	//　キャッシュ生成
	err := NewMatchDataCache(r, "cacheKeyTest", p)
	if err != nil {
		fmt.Println(err)
	}

	// キャッシュ取得
	var getCache []byte
	getCache, err = GetMatchDataCache(r, "cacheKeyTest")
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, getCache)
}

// 異常系
func TestGetMatchDataCacher_Abnormal(t *testing.T) {

}
