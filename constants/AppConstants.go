package constants

// Constants for models
const (
	USER_NEXT_ID string = "user:next-id"
	USER_BY_ID string = "user:%d"
	USER_BY_USERNAME string = "user:by-username"
)

const (
	ID string = "id"
	USERNAME string = "username"
	PASSWORD_HASH string = "password-hash"
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
