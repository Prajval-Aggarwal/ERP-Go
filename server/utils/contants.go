package utils

const (
	HTTP_BAD_REQUEST                     int64  = 400
	HTTP_UNAUTHORIZED                    int64  = 401
	HTTP_PAYMENT_REQUIRED                int64  = 402
	HTTP_FORBIDDEN                       int64  = 403
	HTTP_NOT_FOUND                       int64  = 404
	HTTP_METHOD_NOT_ALLOWED              int64  = 405
	HTTP_NOT_ACCEPTABLE                  int64  = 406
	HTTP_PROXY_AUTHENTICATION_REQUIRED   int64  = 407
	HTTP_REQUEST_TIMEOUT                 int64  = 408
	HTTP_CONFLICT                        int64  = 409
	HTTP_GONE                            int64  = 410
	HTTP_LENGTH_REQUIRED                 int64  = 411
	HTTP_PRECONDITION_FAILED             int64  = 412
	HTTP_PAYLOAD_TOO_LARGE               int64  = 413
	HTTP_URI_TOO_LONG                    int64  = 414
	HTTP_UNSUPPORTED_MEDIA_TYPE          int64  = 415
	HTTP_RANGE_NOT_SATISFIABLE           int64  = 416
	HTTP_EXPECTATION_FAILED              int64  = 417
	HTTP_TEAPOT                          int64  = 418
	HTTP_MISDIRECTED_REQUEST             int64  = 421
	HTTP_UNPROCESSABLE_ENTITY            int64  = 422
	HTTP_LOCKED                          int64  = 423
	HTTP_FAILED_DEPENDENCY               int64  = 424
	HTTP_UPGRADE_REQUIRED                int64  = 426
	HTTP_PRECONDITION_REQUIRED           int64  = 428
	HTTP_TOO_MANY_REQUESTS               int64  = 429
	HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE int64  = 431
	HTTP_UNAVAILABLE_FOR_LEGAL_REASONS   int64  = 451
	HTTP_INTERNAL_SERVER_ERROR           int64  = 500
	HTTP_NOT_IMPLEMENTED                 int64  = 501
	HTTP_BAD_GATEWAY                     int64  = 502
	HTTP_SERVICE_UNAVAILABLE             int64  = 503
	HTTP_GATEWAY_TIMEOUT                 int64  = 504
	HTTP_HTTP_VERSION_NOT_SUPPORTED      int64  = 505
	HTTP_VARIANT_ALSO_NEGOTIATES         int64  = 506
	HTTP_INSUFFICIENT_STORAGE            int64  = 507
	HTTP_LOOP_DETECTED                   int64  = 508
	HTTP_NOT_EXTENDED                    int64  = 510
	HTTP_NETWORK_AUTHENTICATION_REQUIRED int64  = 511
	HTTP_OK                              int64  = 200
	HTTP_NO_CONTENT                      int64  = 204
	DB_MIGRATION_ERROR                   string = "Error while DB migration"
)

const (
	ERROR                       string = "Error"
	TOKEN                       string = "token"
	AUTHORIZATION_HEADER        string = "Authorization"
	FAILURE                     string = "Failure"
	SUCCESS                     string = "Success"
	TOKEN_NOT_FOUND             string = "Token not Found"
	ERROR_IN_HTTP_REQUEST       string = "Error in making HTTP request"
	SERVER_ERROR                string = "INTERNAL SERVER ERROR"
	RESPONSE_BODY_ERROR         string = "Error in reading response body"
	UNMARSHALLING_ERROR         string = "Error in unmarshalling the reponse body"
	DECODING_ERROR              string = "Error decoding request"
	ENCODING_ERROR              string = "Error encoding request"
	LOGIN_SUCCESSFULL           string = "User Login sucessfull"
	LOGOUT_SUCCESSFULL          string = "User Logout sucessfull"
	PASSWORD_UPDATE_SUCCESSFULL string = "Password update sucessfully"
	QUERYEXECUTOR_ERROR         string = "Error while executing query"
)

const (
	BASE_URL                        string = "https://mattermost.local.chicmic.co.in/"
	STAGING_USER_AUTHENTICATION_URL string = "https://timedragon.staging.frimustechnologies.com/v1/auth/check_authenticated"
	REQUEST_POST                    string = "POST"
	STAGING_USER_URL                string = "https://timedragon.staging.frimustechnologies.com/v1/user?_id="
	REQUEST_GET                     string = "GET"
	MATTERMOST_LOGIN_URL            string = BASE_URL + "api/v4/users/login"
	MATTERMOST_SIGNUP_URL           string = BASE_URL + "api/v4/users"
	MATTERMOST_LOGOUT_URL           string = BASE_URL + "api/v4/users/logout"
	CUSTOM_HEADER_KEY_1             string = "X-Requested-With"
	CUSTOM_HEADER_VALUE_1           string = "XMLHttpRequest"
	STAGING_DOMAIN                  string = ".chicmic.co.in"
	DATA_FETCH_SUCESS               string = "Data Fetched sucessfully"
)

const (
	EMAILID  string = "emailid"
	NAME     string = "name"
	PASSWORD string = "password"
	EMP_ID   string = "employeeId"
)
