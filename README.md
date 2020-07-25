# Auth Cookie System

Basic handlers and functions structure to create, valid and remove cookie using jwt-go lib (github.com/dgrijalva/jwt-go). 
It creates JWT valid cookies signed by secret key. The cookie store basic personal user data in claims as userID.

## Utility

This initial code can be used directly or adapted in bigger projects.
The project doesn't cover login auth system. It's a basic structure to help developers in first steps.

## Env file

The software loads secrete key from `.env` file in root directory using Viper lib (github.com/spf13/viper). Below a `.env` example:

```
JWTKEY="secret_key"
```
## How to use

1. Server web app `go run main.go`
2. `localhost:800/login/` to log in and create a cookie called "auth-cookie"
3. `localhost:800/welcome/` can only be accessed with valid cookie
4. `localhost:800/logout/` to remove the cookie
5. `localhost:800/refresh/` web app can use this endpoint to do background requests, using JavaScript for example, to refresh user cookie
