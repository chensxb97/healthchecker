## Healthchecker
A mini project created using create-react-app and golang.

### Summary
A dashboard that regularly performs API healthchecks on configured endpoints.

### Run Backend
1. Install dependencies
```
go mod tidy
```

2. Run web server - http://localhost:8080
```
make run
```

3. Validate endpoints' statuses - http://localhost:8080/status

### Run Frontend
1. Install dependencies
```
npm install
```

2. Run frontend - http://localhost:3000

```
npm run start
```
