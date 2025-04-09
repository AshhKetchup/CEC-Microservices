func main() {
	cfg := config.LoadConfig()

	// Create data directory if not exists
	if err := os.MkdirAll(cfg.DataDir, 0755); err != nil {
		log.Fatalf("failed to create data directory: %v", err)
	}

	// Initialize file storage
	store := storage.NewFileStore(cfg.DataDir)

	// Repository
	userRepo := repository.NewUserRepository(store)

	// Rest of the code remains the same...
}