install:
	@echo "[FAILBOOK] Installing dependencies..."
	go get ./...
.PHONY: install

start:
	@echo "[FAILBOOK] Starting..."
	go run main.go web
.PHONY: start

start.bcrypt:
	@echo "[FAILBOOK] Starting testbcrypt..."
	go run test/testbcrypt.go
.PHONY: start.bcrypt

db.migrate:
	@echo "[FAILBOOK] Running migrations"
	go run main.go db:migrate
.PHONY: db.migrate

db.clean:
	@echo "[FAILBOOK] Dropping migrations"
	go run main.go db:clean
.PHONY: db.clean

db.seed:
	@echo "[FAILBOOK] Seeding database data..."
	go run main.go db:seed
.PHONY: db.clean