package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/pacts", dir)


func TestProvider(t *testing.T) {
	// ローカルデーモン接続用にpactの作成
	pact := &dsl.Pact{
		Consumer: "MyConsumer",
		Provider: "MyProvider",
	}
	// サーバの起動
	go startServer()

	// ローカルpactファイルを用いた検証
	pactsDir := filepath.ToSlash(fmt.Sprintf("%s/myconsumer-myprovider.json", pactDir))
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:        "http://localhost:8000",
		PactURLs:               []string{pactsDir},
		ProviderStatesSetupURL: "http://localhost:8000/api/v1/animal/1",
	})
}

// プロダクトコード（になる予定）
func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/animal/1", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// contractの期待通りのレスポンスだった場合
		fmt.Fprintf(w, "{\"id\": 1, \"name\": \"サーバル\"}")

		// contract作成時からバージョンが上がって項目が追加された場合
		//fmt.Fprintf(w, "{\"id\": 1, \"name\": \"サーバル\", \"age\": 12}")

		// contract作成時からバージョンが上がって破壊的変更が含まれた場合
		// （ name → real_name に変更になりました ）
		//fmt.Fprintf(w, "{\"id\": 1, \"real_name\": \"サーバル\"}")
	})
	log.Fatal(http.ListenAndServe(":8000", mux))
}
