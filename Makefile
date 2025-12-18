.PHONY: run-be run-fe setup

# Run Backend (Go)
run-be:
	cd backend && make run

# Run Frontend (Vue)
run-fe:
	cd frontend && npm run dev

# Install Dependencies for both
setup:
	cd backend && go mod tidy
	cd frontend && npm install