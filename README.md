# Healthchecker
A mini project that regularly performs API healthchecks on configured endpoints, created using create-react-app and golang.

## Setup
1. Install dependencies for go
```
go mod tidy
```

2. Run backend - http://localhost:8080.
```
make run
```
3. Install dependencies for react
```
cd frontend && npm install
```

4. Run frontend - http://localhost:3000
```
npm run start
```

5. Frontend should display endpoints configured in config.yaml. The endpoints return 200 or 404 status in a random manner.