package constants

// Constants for User model
const (
	USER_NEXT_ID string = "user:next-id"
	USER_BY_ID string = "user:%d"
	USER_BY_USERNAME string = "user:by-username"
)

// Constants for Updates model
const (
	UPDATES_NEXT_ID string = "updates:next-id"
	UPDATE_BY_ID string = "update:%d"
	BODY string = "body"
	TIME string = "time"
)

const (
	ID string = "id"
	USERNAME string = "username"
	PASSWORD_HASH string = "password-hash"
	SESSION string = "session"
	USER_ID string = "user-id"
	UPDATES string = "updates"
)

// Error Constants
const (
	INVALID_PASSWORD string = "Invalid Password"
	UNAUTHORIZED string = "Unauthorized"
	REDIS_ERROR string = "Redis Error"
)

const (
	AUTH_ERROR_PAGE string = "/auth-error"
)

// Action constants
const (
	LOGIN_AGAIN string = "Login again!"
	SIGNUP_AGAIN string = "Signup again!"
)

// HTTP Status codes
const (
	HTTP_UNAUTHORIZED int = 64
)
