package constants

import "os"

var AUTH_JWT_SECRET = os.Getenv("JWT_SECRET")
