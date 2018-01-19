install:
	@echo "[NOTE++] Installing dependencies..."
	go get ./...
.PHONY: install

start:
	@echo "[NOTE++] Starting..."
	go run main.go web
.PHONY: start

db.migrate:
	@echo "[NOTE++] Running migrations"
	go run main.go db:migrate
.PHONY: db.migrate

db.clean:
	@echo "[NOTE++] Dropping migrations"
	go run main.go db:clean
.PHONY: db.clean

db.seed:
	@echo "[NOTE++] Seeding database data..."
	go run main.go db:seed
.PHONY: db.seed