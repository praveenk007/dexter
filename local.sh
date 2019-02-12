docker build -t dexter .
docker run --env-file=.env -p 4000:4000 dexter