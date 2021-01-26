module goserver

go 1.15

replace(
    goserver/thirdparty/logger => ./thirdparty/logger
    goserver/thirdparty/github.com/go-sql-driver/mysql => ./thirdparty/github.com/go-sql-driver/mysql
    goserver/thirdparty/github.com/mattn/go-sqlite3 => ./thirdparty/github.com/mattn/go-sqlite3
    goserver/thirdparty/github.com/dgrijalva/jwt-go => ./thirdparty/github.com/dgrijalva/jwt-go
)

require(
    goserver/thirdparty/logger v0.0.0
    goserver/thirdparty/github.com/go-sql-driver/mysql v0.0.0
    goserver/thirdparty/github.com/mattn/go-sqlite3 v0.0.0
    goserver/thirdparty/github.com/dgrijalva/jwt-go v0.0.0
)
