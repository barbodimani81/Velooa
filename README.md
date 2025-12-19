```
velooa-backend/
├── cmd/
│   └── api/
│       └── main.go           # Entry point (Wires everything together)
├── internal/                 # Code private to this app (cannot be imported by others)
│   ├── config/               # Load .env variables
│   ├── database/             # Connect to Postgres (pgxpool)
│   ├── models/               # The Go Structs matching your DB tables
│   ├── store/                # REPOSITORY LAYER: Raw SQL queries
│   │   ├── users.go
│   │   ├── products.go
│   │   └── orders.go
│   ├── service/              # BUSINESS LOGIC LAYER: The "Brain"
│   │   ├── checkout.go       # Transaction logic, stock checking
│   │   ├── pricing.go        # Calculate discounts/totals
│   │   └── auth.go           # Hashing passwords, token generation
│   ├── api/                  # TRANSPORT LAYER: HTTP Handlers
│   │   ├── handlers/         # Parse JSON, validate input, call Service
│   │   ├── middleware/       # Auth guards, logging, CORS
│   │   └── router.go         # Setup Gin/Echo/Chi routes
│   └── utils/                # Shared helpers (e.g., Validator, Random String)
├── migrations/               # Your .sql schema files
├── go.mod
└── go.sum
```
