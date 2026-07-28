# Blockchain Explorer CLI

A lightweight blockchain explorer written in Go.

The application works entirely with local JSON data and simulates common blockchain explorer commands.

## Features

- Address lookup
- Transaction lookup
- Block lookup
- Token information
- Pretty console output
- Offline mode

## Run

```bash
go run ./cmd address 0x123
go run ./cmd tx 0xabc
go run ./cmd block 125
go run ./cmd token USDC
```

## Future Improvements

- SQLite backend
- CSV export
- Interactive shell
- Search history
- Color themes
