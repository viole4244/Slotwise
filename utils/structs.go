package utils

// ---------- Config ----------

// Config holds every setting read from the environment (.env) at startup.
type Config struct {
	Port        string
	DatabaseURL string
	RedisURL    string
}

// ---------- API responses ----------

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Redis    string `json:"redis"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
