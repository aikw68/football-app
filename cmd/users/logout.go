package users

import (
	"football/cmd/util"
	"net/http"
	"os"
)

// ログアウト
func Logout(w http.ResponseWriter, r *http.Request) error {

	cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE")
	if err := util.DeleteSession(w, r, cookieKey); err != nil {
		return err
	}
	return nil
}
