package setting

// AppMode represents the application mode, which can be one of "prod", "dev", or "test".
type AppMode string

const (
	ProdMode AppMode = "prod"
	DevMode  AppMode = "dev"
	TestMode AppMode = "test"
)
