# HNG13 — Backend Stage 0: Dynamic Profile Endpoint.

This small Go project implements Stage 0 of the HNG Backend track: a single
GET endpoint that returns a profile object (your details) plus a dynamic cat
fact fetched from the Cat Facts API.

The project uses Gin for HTTP routing. The endpoint implemented:

- GET /me — returns a JSON response with your profile and a dynamic cat fact.

## JSON response format (required)

The .me endpoint returns JSON in this format:

{
	"status": "success",
	"user": {
		"email": "<your email>",
		"name": "<your full name>",
		"stack": "<your backend stack>"
	},
	"timestamp": "<current UTC time in ISO 8601 format>",
	"fact": "<random cat fact from Cat Facts API>"
}

Example:

{
	"status": "success",
	"user": {
		"email": "iJeddy5.0@gmail.com",
		"name": "Iwegbu Jedidiah",
		"stack": "Go/Gin"
	},
	"timestamp": "2025-10-15T12:34:56.789Z",
	"fact": "Cats have five toes on their front paws, but only four toes on their back paws."
}

## Requirements covered

- GET `/me` implemented using Gin
- Response Content-Type is `application/json`
- Dynamic cat fact fetched on every request from https://catfact.ninja/fact
- Timestamp uses UTC in ISO 8601 format (RFC3339/RFC3339Nano)
- Reasonable timeout for external API calls and graceful fallback when API is down

## How to run locally (Windows / PowerShell)

1. Install Go (1.20+ recommended). Verify `go` is in PATH.
2. From the project root:

```powershell
cd 'C:\Users\USER\Documents\Work\HNG13_Backend_Stage0'
go mod tidy
go run ./cmd
```

This starts the server (by default the code uses Gin — check `cmd/main.go` for the exact bind address).

## Quick test

With the server running, you can request the endpoint:

```powershell
Invoke-RestMethod -Uri http://localhost:8080/me -Method GET
```
[wahala for who dey use bash sha :)]

You should receive the JSON structure explained above.

## Troubleshooting

- If `go run` fails with `unexpected NUL in input` for a `.go` file, that file may be saved in UTF-16 or contains NUL bytes. Re-save the file as UTF-8 without BOM.
- If the Cat Facts API is unreachable, the server will return a fallback fact string and still respond with status 200 and the required JSON structure.

## Notes for submission
- Include link to this repository when submitting the Stage 0 form.
- Make sure the `/me` endpoint returns the exact JSON schema required by the task.

## Contact

Email: iJeddy5.0@gmail.com
Name: Iwegbu Jedidiah
Stack: Go/Gin


Good luck!
