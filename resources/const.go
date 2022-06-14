package resources

const (
	DatabaseMySQL      = "DBMySQL"
	DatabasePostgreSQL = "DBPostgreSQL"
	RedisDefault       = "redis"

	NotFound = "not_found"

	//http status strings
	StatusUnauthorized        = "status_unauthorized"
	StatusBadRequest          = "bad_request"
	StatusNotFound            = "not_found"
	StatusRequestTimeout      = "timeout"
	StatusExpectationFailed   = "failed"
	StatusInternalServerError = "internalError"
	StatusTokenExpired        = "token_expired"
	StatusForbidden           = "forbidden"

	StatusUnauthorizedMessage = "status unauthorized"
	StatusBadRequestMessage   = "bad request"
	StatusNotFoundMessage     = "data not found"
	StatusOKMessage           = "ok"
	StatusTokenExpiredMessage = "token expired"
	StatusUnprocessableEntity = "status unprocessable entity"
	StatusOK                  = "status_ok"

	//GIN
	Release = "release"
)
