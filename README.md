# TRIVIA

Dumb trivia question source

## No Docker 

Create a new database user, and give it the password of `theansweris` when prompted

    createuser -P trivia

Create a new database
    
    createdb -O trivia trivia

Install goose to run the database migrations
    
    go get bitbucket.org/liamstask/goose/cmd/goose

Run it 

    ./lazy-run

## Docker stuff

    cd db && docker build -t trivia-db . && cd -
    cd migrations && docker build -t trivia-migrations . && cd -
    cd cmd/trivia && docker build -t trivia-app . && cd -

    docker run -d --name trivia-db trivia-db
    docker run --rm --link trivia-db:pg trivia-migrations
    docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock jwilder/nginx-proxy
    docker run -d --link trivia-db:pg --name trivia-app -e VIRTUAL_HOST=trivia.ralreegorganon.com trivia-app
