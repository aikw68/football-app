package util

type AppErr string

func (e AppErr) Error() string {
	msg, ok := AppErrCodeMessages[e]
	if !ok {
		return string(e)
	}
	return msg
}

// エラーコード→エラーメッセージに置換
var AppErrCodeMessages = map[AppErr]string{
	ERR_USER_EMAIL_REGISTERED:          "同一のメールアドレスが既に登録されています。",
	ERR_USER_EMAIL_NOT_EXIST:           "入力されたメールアドレスのユーザーが存在しません。",
	ERR_USER_EMAIL_INCORRECT_FORMAT:    "正しいメールアドレス形式で入力してください。",
	ERR_USER_PASSWORD_MISMATCH:         "入力されたパスワードが一致しません。",
	ERR_USER_PASSWORD_INCORRECT_FORMAT: "正しいパスワード形式で入力してください。",
	ERR_USER_SYSTEM_ERROR:              "システムエラーが発生しました。改善されない場合、お手数ですがサイト管理者までお問い合わせください。",
	ERR_PASSWORD_ENCRYPT_FAILED:        "パスワード暗号化中にエラーが発生しました。",
	ERR_SESSION_KEY_GENERATE_FAILED:    "Sessionキー生成時にエラーが発生しました。",
	ERR_SESSION_KEY_UNREGISTERED:       "Sessionキーが登録されていません。",
	ERR_SESSION_REGISTRATION_FAILED:    "Session登録時にエラーが発生しました。",
	ERR_SESSION_EXTEND_FAILED:          "Session有効期限延長時にエラーが発生しました。",
	ERR_SESSION_GET_FAILED:             "Session取得時にエラーが発生しました。",
	ERR_CACHE_REGISTRATION_FAILED:      "Cache登録時にエラーが発生しました。",
	ERR_CACHE_GET_FAILED:               "Cache取得時にエラーが発生しました。",
	ERR_API_CALL_FAILED:                "APIコール時にエラーが発生しました。",
	ERR_404_NOT_FOUND:                  "404エラーが発生しました。",
}

// エラーコード定義
const (
	ERR_USER_EMAIL_REGISTERED          AppErr = "ERR_USER_EMAIL_REGISTERED"
	ERR_USER_EMAIL_NOT_EXIST           AppErr = "ERR_USER_EMAIL_NOT_EXIST"
	ERR_USER_EMAIL_INCORRECT_FORMAT    AppErr = "ERR_USER_EMAIL_INCORRECT_FORMAT"
	ERR_USER_PASSWORD_MISMATCH         AppErr = "ERR_USER_PASSWORD_MISMATCH"
	ERR_USER_PASSWORD_INCORRECT_FORMAT AppErr = "ERR_USER_PASSWORD_INCORRECT_FORMAT"
	ERR_USER_SYSTEM_ERROR              AppErr = "ERR_USER_SYSTEM_ERROR"
	ERR_PASSWORD_ENCRYPT_FAILED        AppErr = "ERR_PASSWORD_ENCRYPT_FAILED"
	ERR_SESSION_KEY_GENERATE_FAILED    AppErr = "ERR_SESSION_KEY_GENERATE_FAILED"
	ERR_SESSION_KEY_UNREGISTERED       AppErr = "ERR_SESSION_UNREGISTERED"
	ERR_SESSION_REGISTRATION_FAILED    AppErr = "ERR_SESSION_REGISTRATION_FAILED"
	ERR_SESSION_EXTEND_FAILED          AppErr = "ERR_SESSION_EXTEND_FAILED"
	ERR_SESSION_GET_FAILED             AppErr = "ERR_SESSION_GET_FAILED"
	ERR_CACHE_REGISTRATION_FAILED      AppErr = "ERR_CACHE_REGISTRATION_FAILED"
	ERR_CACHE_GET_FAILED               AppErr = "ERR_CACHE_GET_FAILED"
	ERR_API_CALL_FAILED                AppErr = "ERR_API_CALL_FAILED"
	ERR_404_NOT_FOUND                  AppErr = "ERR_404_NOT_FOUND"
)
