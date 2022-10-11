package e

const SUC = 0

// 400 系列
const (
	ERR_BAD_REQUEST = 400000 + iota
	ERR_INVALID_ARGS
	ERR_INVALID_CAPTCHA
	ERR_ADD_ITEM
	ERR_ADD_TEAM
	ERR_ADD_DOC
)

// 401 未授权
const (
	ERR_UNAUTHORIZED = 401000 + iota
	ERR_NOT_AUTH
)

// 403 禁止
const (
	ERR_FORBIDEN = 403000 + iota
)

// 404
const (
	ERR_NOTFOUND = 404000 + iota
	ERR_SOURCE_INVAILID
	ERR_SOURCE_EXPIRE
)

// 415 unsupport
const (
	ERR_UNSUPPORT_MEDIA = 415000 + iota
	ERR_UNSUPPORT_FILE
	ERR_UNSUPPORT_PLATFORM
)

// 429
const (
	ERR_TOO_MANY_REQUESTS = 429000 + iota
	ERR_TOO_MANY_REQUESTS_IP
	ERR_TOO_MANY_REQUESTS_PHONE
)

// 500
const (
	ERR_SERVER = 500000 + iota
	ERR_UNKNOWN
)

var MsgFlags = map[int]string{
	// 400
	ERR_BAD_REQUEST: `bad request`,

	// 401
	ERR_NOT_AUTH: `unauthorized`,

	// 403
	ERR_FORBIDEN: `forbidden`,

	// 404
	ERR_NOTFOUND: `source not found`,

	// 415 Unsupported Media Type
	ERR_UNSUPPORT_MEDIA: `Unsupported Media Type`,
	// 429
	ERR_TOO_MANY_REQUESTS: `request frequently`,

	// 500 内部错误
	ERR_SERVER: `internal server error`,
}
