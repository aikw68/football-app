package users

import (
	"football/cmd/util"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"primaryKey"`
	Password string
}

// サインアップ処理
func Signup(r *http.Request) (*User, error) {

	user := User{}

	if rtn := Validation(r); rtn != "" {
		return nil, rtn
	}

	// ユーザー入力パラメータ（メールアドレス、パスワード）
	email := r.FormValue("email")
	password := r.FormValue("password")

	// DB接続
	db, err := util.DbConnect()
	if err != nil {
		return nil, err
	}

	// DBマイグレーション
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// メールアドレス二重登録チェック
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return nil, util.ERR_USER_EMAIL_REGISTERED
	}

	// パスワード暗号化
	encryptPw, err := util.PasswordEncrypt(password)
	if err != nil {
		return nil, util.ERR_USER_SYSTEM_ERROR
	}

	// ユーザー登録
	user = User{Email: email, Password: encryptPw}
	db.Create(&user)

	return &user, nil
}

// ログイン
func Login(email, password string) (*User, error) {

	// DB接続
	db, err := util.DbConnect()
	if err != nil {
		return nil, err
	}

	// メールアドレスチェック
	user := User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, util.ERR_USER_EMAIL_NOT_EXIST
	}

	// パスワードチェック
	err = util.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, util.ERR_USER_PASSWORD_MISMATCH
	}
	return &user, nil
}

// ログインチェック
func CheckLogin(r *http.Request) bool {

	// CookieKey取得
	cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE")
	// ログインチェック
	id := util.GetSession(r, cookieKey)
	if id == nil {
		return false // 未ログイン
	} else {
		return true // ログイン中
	}
}
