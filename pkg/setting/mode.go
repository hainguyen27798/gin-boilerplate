package setting

// AppMode defines the mode in which the application is running,
// such as "dev" for development or "prod" for production.
type AppMode string

// ProdMode represents the application mode for production environments.
// DevMode represents the application mode for development environments.
const (
	ProdMode AppMode = "prod"
	DevMode  AppMode = "dev"
)
