package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	// データベース関連
	InsertDataFailed   ErrCode = "S001" // データ挿入失敗
	GetDataFailed      ErrCode = "S002" // データ取得失敗
	NAData             ErrCode = "S003" // データなし
	NoTargetData       ErrCode = "S004" // 対象データなし
	UpdateDataFailed   ErrCode = "S005" // 更新失敗
	DBConnectionFailed ErrCode = "S006" // データベース接続失敗
	TransactionFailed  ErrCode = "S007" // トランザクション失敗
	DeleteDataFailed   ErrCode = "S008" // データ削除失敗

	// リクエスト関連
	ReqBodyDecodeFailed ErrCode = "R001" // リクエストボディのデコード失敗
	BadParam            ErrCode = "R002" // 無効なリクエストパラメータ
)

type MyError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyError{ErrCode: code, Message: message, Err: err}
}
