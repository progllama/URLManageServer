package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// パスの区切り文字にはos関係なくスラッシュで渡す。
func ServeFavicon(path string) gin.HandlerFunc {
	// /区切りのパスをosの区切り文字でのパスに変換。
	path = filepath.FromSlash(path)
	// から文字でないかつOSの区切り文字でないなら
	if len(path) > 0 && !os.IsPathSeparator(path[0]) {
		// 作業ディレクトリを取得
		wd, err := os.Getwd()
		// 作業ディレクトリの取得に失敗したら終了
		if err != nil {
			panic(err)
		}
		// 作業ディレクトリと引数のパスを結合
		path = filepath.Join(wd, path)
	}
	fmt.Println(path)

	// ファイルの情報を取得
	info, err := os.Stat(path)
	// エラーでないもしくは情報がないもしくはディレクトリなら終了
	if err != nil || info == nil || info.IsDir() {
		if err != nil {
			panic(err)
		}
		panic("Invalid favicon path: " + path)
	}

	// faviconを読み込む
	file, err := ioutil.ReadFile(path)
	// エラーなら終了
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(file)
	log.Println("Success to register ServeFavicon middleware.")

	// ハンドラを返す
	return func(c *gin.Context) {
		// リクエストのURIがちがければ終了
		if c.Request.RequestURI != "/favicon.ico" {
			fmt.Println(c.Request.RequestURI)
			log.Println("wrong uri.")
			return
		}
		// GET、HEAD以外なら
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			fmt.Println("wrong method.")
			status := http.StatusOK
			// OPTIONじゃないならStatusOK -> MethodNotAllowdに変更
			if c.Request.Method != "OPTIONS" {
				status = http.StatusMethodNotAllowed
			}
			c.Header("Allow", "GET,HEAD,OPTIONS")
			//ここで終了
			c.AbortWithStatus(status)
			return
		}
		// GET、HEADならfaviconを返す。
		c.Header("Content-Type", "image/x-icon")
		http.ServeContent(c.Writer, c.Request, "favicon.ico", info.ModTime(), reader)
	}
}
