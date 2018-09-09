package usecase

type Code string

// UseCaseの処理結果を表すコード
// ユーザーに伝えたい単位で分割
// 結果を抽象化しておくことでcontrollerではhttp status code以外への変換にも対応できる
// この例ではhttp status codeと同じくらいの粒度だが、アプリの要件によってはもっと細分化されるかも
// なんのフィールドがバリデーションエラーになったかとかまでクライアントに知らせる必要あるならそのレベルでエラーコード定義
// 成功も失敗も本質的には変わらないのでは？
// Resultに結果のデータも持たせようと思ったがgolangだとジェネリクスがなくて型が失われるのでやめた
// ジェネリクスある言語なら持たせていい気がする
// TODO: 処理は失敗したけどなんらかのデータを返したい場合ってある？返さないほうがいい？
// TODO: 同じエラーがいろんなとこでinstanciate
// TODO: 階層構造にすればon demandで細分化できて良さそう？？
// Result
//   OK
//     CREATED
//
//   Error
//     INVALID
//       PAGEID
const (
	OK         Code = "ok"         // success
	CREATED    Code = "created"    // create success
	INVALID    Code = "invalid"    // validation failed
	NOTFOUND   Code = "not found"  // resource not found
	UNEXPECTED Code = "unexpected" // unknown error
)

type Result struct {
	code    Code
	message string
}

func NewResult(code Code, message string) *Result {
	return &Result{
		code:    code,
		message: message,
	}
}

func (e *Result) Code() Code {
	return e.code
}

func (e *Result) Message() string {
	return e.message
}