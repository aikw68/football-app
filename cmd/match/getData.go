package match

import (
	"encoding/json"
	"football/cmd/users"
	"football/cmd/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title         string
	SubTitle      string
	LoginFlg      bool
	MatchList     MatchLists
	ScoreList     ScoreLists
	StandingsList StandingsLists
	Message       string
}

var Title string = "DUELSCORE"
var Blunk string = ""

func (p Page) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// 試合データ取得
func GetMatchData(r *http.Request, fstLoginFlg bool) (Page, error) {

	// ログインチェック
	var loginFlg bool
	if fstLoginFlg {
		// 初回ログイン・サインイン時はログイン中にする
		loginFlg = true
	} else {
		//　ログインチェック
		_, err := users.CheckLogin(r)
		if err != nil {
			return Page{}, err
		}
	}

	// キャッシュ有無チェック
	cacheKey := os.Getenv("FOOTBALL_REDIS_CACHE_KEY")
	cache := util.GetMatchDataCacher(r, cacheKey)

	if cache != nil {

		var cachePage Page
		if err := json.Unmarshal(cache, &cachePage); err != nil {
			panic(err)
		}

		// キャッシュが存在する場合、キャッシュから試合データを取得する
		c1 := cachePage.MatchList
		c2 := cachePage.ScoreList
		c3 := cachePage.StandingsList

		p := Page{Title, Blunk, loginFlg, c1, c2, c3, ""}

		return p, nil

	} else {

		// キャッシュが存在しない場合、新たに試合データを取得する
		c1 := make(chan MatchLists)
		c2 := make(chan ScoreLists)
		c3 := make(chan StandingsLists)

		// 試合日程取得
		go Schedule(c1)
		// 得点ランキング取得
		go Score(c2)
		// 順位表取得
		go Standings(c3)

		//　取得結果を一つにまとめる
		p := Page{Title, Blunk, loginFlg, <-c1, <-c2, <-c3, ""}

		// キャッシュ保存
		cacheKey := os.Getenv("FOOTBALL_REDIS_CACHE_KEY")
		util.NewMatchDataCache(r, cacheKey, p)

		return p, nil
	}
}

// 試合データ取得API呼出
func ApiCall(param string) []byte {

	res, err := util.GetSecret("football-data_auth_APIkey")
	if err != nil {
		log.Println(err.Error())
	}

	secretValue := res["FOOTBALL_DATA_AUTH_HEADER_VALUE"].(string)

	url := "https://api.football-data.org/v4/competitions/PL/" + param
	authHeaderName := "X-Auth-Token"
	authHeaderValue := secretValue

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(authHeaderName, authHeaderValue)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error Request:", err.Error())
	}

	if resp.StatusCode != 200 {
		log.Println("Error Response:", resp.Status)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)

	return jsonBytes
}
