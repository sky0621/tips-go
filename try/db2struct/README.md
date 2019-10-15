# try xo

go get -u github.com/xo/xo
go get -tags postgres -u github.com/xo/xo

xo 'pgsql://localhost:5432/localdb?sslmode=disable' -o model

