package util

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const SESSION_LIFETIME = 30 * time.Minute

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "football-redis:6379",
		Password: "",
		DB:       0,
	})
}

// ログインセッション&Cookie生成
func NewSession(w http.ResponseWriter, r *http.Request, redisValue string, cookieKey string) error {

	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return errors.WithStack(ERR_SESSION_KEY_GENERATE_FAILED)
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// セッション生成
	if err := conn.Set(r.Context(), newRedisKey, redisValue, SESSION_LIFETIME).Err(); err != nil {
		return errors.WithStack(ERR_SESSION_REGISTRATION_FAILED)
	}
	// Cookie生成
	cookie := http.Cookie{
		Name:     cookieKey,
		Value:    newRedisKey,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// ログインセッション取得
func GetSession(r *http.Request, cookieKey string) (string, error) {

	redisKey, _ := r.Cookie(cookieKey)
	if redisKey == nil {
		return "", nil
	} else {
		redisValue, err := conn.Get(r.Context(), redisKey.Value).Result()
		switch {
		case err == redis.Nil:
			return "", nil
		case err != nil:
			return "", errors.WithStack(ERR_SESSION_GET_FAILED)
		}
		return redisValue, nil
	}
}

// ログインセッションの有効期限延長
func ExtendSession(r *http.Request, cookieKey string) error {

	redisKey, _ := r.Cookie(cookieKey)
	if err := conn.Expire(r.Context(), redisKey.Value, SESSION_LIFETIME).Err(); err != nil {
		return errors.WithStack(ERR_SESSION_EXTEND_FAILED)
	}
	return nil
}

// ログインセッション&Cookie削除
func DeleteSession(w http.ResponseWriter, r *http.Request, cookieKey string) {

	redisKey, _ := r.Cookie(cookieKey)
	// セッション削除
	conn.Del(r.Context(), redisKey.Value)
	// Cookie削除
	redisKey.MaxAge = -1
	http.SetCookie(w, redisKey)
}

// 試合データのキャッシュ生成
func NewMatchDataCache(r *http.Request, cacheKey string, redisValue interface{}) {

	if err := conn.Set(r.Context(), cacheKey, redisValue, 60*time.Second).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}
}

// 試合データのキャッシュ取得
func GetMatchDataCacher(r *http.Request, cacheKey string) []byte {

	redisValue, err := conn.Get(r.Context(), cacheKey).Result()
	switch {
	case err == redis.Nil:
		// キャッシュ無し（生成する）
		return nil
	case err != nil:
		// キャッシュ取得時にエラー発生
		return nil //,errors.WithStack(ERR_CACHE_GET_FAILED)
	}
	byteConvValue := []byte(redisValue)
	return byteConvValue //,nil
}
