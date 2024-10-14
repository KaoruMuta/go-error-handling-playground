# go_error_handling_playground

## Overview

This repository is a playground to confirm error handling strategy with Go and the following packages.

- [errors](https://pkg.go.dev/errors)
- [failure](https://github.com/morikuni/failure)

## Setup

1. Install Go 1.23 and configure GOPATH
2. Install deps

```bash
go install github.com/labstack/echo/v4
go install github.com/morikuni/failure/v2
go install github.com/air-verse/air@latest
```

3. Run server

```bash
air -c .air.toml
```
