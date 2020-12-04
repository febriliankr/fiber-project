module go-mysql-react

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/gofiber/fiber/v2 v2.2.3 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/klauspost/compress v1.11.3 // indirect
	// github.com/joho/godotenv v1.3.0
	go-mysql-react/helpers v0.1.0
	go-mysql-react/models v0.1.0
	golang.org/x/sys v0.0.0-20201202213521-69691e467435 // indirect
	github.com/heavykenny/currency
)

replace (
	go-mysql-react/helpers => ./helpers
	go-mysql-react/models => ./models
)
