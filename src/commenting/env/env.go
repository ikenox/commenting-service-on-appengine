package env

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"os"
	"path/filepath"
	"time"
)

// application globals
// アプリにとって普遍とみなせる値や関数、環境変数など？
// どこからでも使われるような値や概念はRepositoryにするよりグローバル変数にしてしまったほうが楽そう
var Namespace string
var CurrentTime func() time.Time
var IsProduction bool
var FirebaseApp *firebase.App
var GCPCredentialOption option.ClientOption
var ProjectID string

func init() {
	ctx := context.Background()

	ProjectID = os.Getenv("APP_ID")

	Namespace = os.Getenv("NAMESPACE")

	IsProduction = os.Getenv("IS_PRODUCTION") == "True"

	// time
	CurrentTime = time.Now


	path, err := filepath.Abs(os.Getenv("SERVICE_ACCOUNT_PATH"))
	if err != nil {
		panic(err.Error())
	}
	GCPCredentialOption = option.WithCredentialsFile(path)

	// FIXME クライアント郡がenvにいるのはおかしいので、optionだけenvに残して移動

	// firebase
	if FirebaseApp, err = firebase.NewApp(ctx, nil, GCPCredentialOption); err != nil {
		panic(err.Error())
	}
}
