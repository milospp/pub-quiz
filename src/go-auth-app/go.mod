module github.com/milospp/pub-quiz/src/go-auth-app

go 1.19

replace github.com/milospp/pub-quiz/src/go-global => ../go-global

require (
	github.com/jackc/pgconn v1.13.0
	github.com/jackc/pgx/v4 v4.17.0
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8
	github.com/milospp/pub-quiz/src/go-global v1.0.0

)

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
)
