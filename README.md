# TRIVIA

Dumb trivia question source

## Use it

Create a new database user, and give it the password of `theansweris` when prompted

    createuser -P trivia

Create a new database
    
    createdb -O trivia trivia

Specify connection string

    export TRIVIA_CONNECTION_STRING="user=trivia password=theansweris dbname=trivia timezone=UTC sslmode=disable"

Install goose to run the database migrations
    
    go get bitbucket.org/liamstask/goose/cmd/goose

Run the migrations.
    
    goose up

Run it 

    go get ./...
    trivia
