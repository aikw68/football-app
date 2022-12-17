package main

import (
	"football/cmd/match"
	"football/cmd/users"
	"football/cmd/util"
	"log"
	"net/http"
	"os"
	"syscall"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

const (
	SITE_TITLE string = "DUELSCORE"
)

// メイン画面表示
func index(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/" {
		notFoundHandler(w, r.URL.Path)
		return
	}
	p, err := match.GetMatchData(r, false)
	if err != nil {
		systemServerErrorHandler(w, err)
		return
	}

	// レンダリング実行
	renderTemplate(w, "index", p)
}

// 利用規約画面表示
func terms(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/terms" {
		notFoundHandler(w, r.URL.Path)
		return
	}

	// ログインチェック
	loginFlg, err := users.CheckLogin(r)
	if err != nil {
		systemServerErrorHandler(w, err)
		return
	}

	// レンダリング実行
	p := match.Page{Title: SITE_TITLE, SubTitle: "ユーザー登録", LoginFlg: loginFlg}
	renderTemplate(w, "terms", p)
}

// プライバシーポリシー画面表示
func privacy(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/privacy" {
		notFoundHandler(w, r.URL.Path)
		return
	}

	// ログインチェック
	loginFlg, err := users.CheckLogin(r)
	if err != nil {
		systemServerErrorHandler(w, err)
		return
	}

	// レンダリング実行
	p := match.Page{Title: SITE_TITLE, SubTitle: "プライバシーポリシー", LoginFlg: loginFlg}
	renderTemplate(w, "privacy", p)
}

// サインアップ画面表示
func getSignup(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/getSignup" {
		notFoundHandler(w, r.URL.Path)
		return
	}

	// レンダリング実行
	p := match.Page{Title: SITE_TITLE, SubTitle: "ユーザー登録"}
	renderTemplate(w, "signup", p)
}

// サインアップ処理
func postSignup(w http.ResponseWriter, r *http.Request) {

	//　サインアップ処理実行
	if _, err := users.Signup(r); err != nil {

		// サインアップに失敗した場合
		log.Println("サインアップに失敗")
		log.Println(err)
		// レンダリング実行（サインアップ画面にエラーメッセージを返す）
		p := match.Page{Title: SITE_TITLE, SubTitle: "ユーザー登録", Message: err.Error()}
		renderTemplate(w, "signup", p)

	} else {

		// サインアップに成功した場合
		// CookieKey取得
		cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE_KEY")
		// ログインセッション&Cookie生成
		util.NewSession(w, r, r.FormValue("email"), cookieKey)

		p, err := match.GetMatchData(r, true)
		if err != nil {
			systemServerErrorHandler(w, err)
			return
		}

		// トップページに遷移
		renderTemplate(w, "index", p)
	}
}

// ログイン画面表示
func getLogin(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/getLogin" {
		notFoundHandler(w, r.URL.Path)
		return
	}

	// レンダリング実行
	p := match.Page{Title: SITE_TITLE, SubTitle: "ログイン"}
	renderTemplate(w, "login", p)
}

// ログイン処理
func postLogin(w http.ResponseWriter, r *http.Request) {

	// POSTされたEメール、パスワードを取得
	email := r.FormValue("email")
	password := r.FormValue("password")

	// ログイン処理実行
	if _, err := users.Login(email, password); err != nil {

		// ログインに失敗した場合
		log.Println(err.Error())
		// レンダリング実行（ログイン画面にエラーメッセージを返す）
		p := match.Page{Title: SITE_TITLE, SubTitle: "ログイン", Message: err.Error()}
		renderTemplate(w, "login", p)

	} else {

		// ログインに成功した場合
		// CookieKey取得
		cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE_KEY")
		// ログインセッション&Cookie生成
		util.NewSession(w, r, email, cookieKey)

		p, err := match.GetMatchData(r, true)
		if err != nil {
			systemServerErrorHandler(w, err)
			return
		}

		// トップページに遷移
		renderTemplate(w, "index", p)
	}
}

// ログアウト処理
func logout(w http.ResponseWriter, r *http.Request) {
	users.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// パスワードリセット画面
func getReset(w http.ResponseWriter, r *http.Request) {

	// URLパスチェック
	if r.URL.Path != "/getReset" {
		notFoundHandler(w, r.URL.Path)
		return
	}

	// ログインチェック
	loginFlg, err := users.CheckLogin(r)
	if err != nil {
		systemServerErrorHandler(w, err)
		return
	}

	// レンダリング実行
	p := match.Page{Title: SITE_TITLE, SubTitle: "パスワードリセット", LoginFlg: loginFlg}
	renderTemplate(w, "reset", p)
}

// パスワードリセット処理（メール送信）
func postReset(w http.ResponseWriter, r *http.Request) {

	// POSTされたEメールを取得
	email := r.FormValue("email")

	// メールアドレスチェック
	if _, err := users.CheckMail(email); err != nil {

		// レンダリング実行（パスワードリセット画面にエラーメッセージを返す）
		p := match.Page{Title: SITE_TITLE, SubTitle: "パスワードリセット", Message: err.Error()}
		renderTemplate(w, "reset", p)

	} else {

		// パスワードリセットURLトークン生成
		token, err := users.UrlTokenGenerate(email, time.Now())
		if err != nil {
			return // エラー
		}

		// メール送信
		users.SendPasswordResetMail(email, token)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// GoogleChrome 二重送信防止用
func faviconHandler(w http.ResponseWriter, r *http.Request) {}

// レンダリング処理
func renderTemplate(w http.ResponseWriter, tmpl string, p match.Page) {
	t := template.Must(template.ParseFiles(
		"../assets/"+tmpl+".html",
		"../assets/header.html",
		"../assets/footer.html",
	))
	if err := t.ExecuteTemplate(w, tmpl+".html", p); err != nil {

		if errors.Is(err, syscall.EPIPE) {
			// "broken pipe"エラー対策
			// "broken pipe"エラーの場合、何もしない
			return

		} else {

			// "broken pipe"エラー以外の場合(500内部サーバーエラー)
			systemServerErrorHandler(w, err)
			return
		}
	}
}

// 404NotFound発生時
func notFoundHandler(w http.ResponseWriter, url string) {

	// レンダリング実行(404 Not Found)
	log.Printf("%s:入力されたURL=%s", util.ERR_404_NOT_FOUND, url)
	p := match.Page{Title: SITE_TITLE, SubTitle: "404 Not Found"}
	renderTemplate(w, "404", p)
}

// 500内部サーバーエラー発生時
func systemServerErrorHandler(w http.ResponseWriter, e error) {

	// レンダリング実行(500 Internal Server Error)
	log.Println(e.Error())
	p := match.Page{Title: SITE_TITLE, SubTitle: "500 Internal Server Error"}
	renderTemplate(w, "error", p)
}
