package users

import (
	"football/cmd/util"
	"net/http"
	"os"
)

// ログアウト
func Logout(w http.ResponseWriter, r *http.Request) {

	cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE")
	util.DeleteSession(w, r, cookieKey)
}
