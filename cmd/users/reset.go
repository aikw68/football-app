package users

import (
	"football/cmd/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 作成中
const (
	secret   = "b7OkYK/x3Ff4/H93AFEH61XsoGeYvITrouYdTjBw8RbNzAmEuzxlsQ=="
	emailKey = "email"
	iatKey   = "iat"
	expKey   = "exp"
	lifetime = 30 * time.Minute
)

// メールアドレスチェック
func CheckMail(email string) (*User, error) {

	// DB接続
	db, err := util.DbConnect()
	if err != nil {
		return nil, err
	}

	user := User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, util.ERR_USER_EMAIL_NOT_EXIST
	}
	return &user, nil
}

// パスワードリセットURLトークン生成
func UrlTokenGenerate(email string, now time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		emailKey: email,
		iatKey:   now.Unix(),
		expKey:   now.Add(lifetime).Unix(),
	})
	return token.SignedString([]byte(secret))
}

// メール送信
func SendPasswordResetMail(email, token string) {

	subject := "パスワードリセット"
	body := "○○様\r\rDUELSCOREをご利用いただきありがとうございます。\rパスワード再設定は以下のURLからお願い致します。\r\rhttp://localhost:8080/getResetPass/" + token + "\r\r今後ともDUELSCOREをよろしくお願い致します。"
	util.SendMail(email, subject, body)

}
