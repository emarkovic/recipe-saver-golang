export projectname=recipesvr
export PORT="8080"
export HOST="localhost"
export DBADDR="localhost:27017"
export REDISADDR="localhost:6379"

#run mongo docker container 
docker run --name mongo-rs -p 27017:27017 -d mongo
#run redis docker container
docker run --name redis-rs -p 6379:6379 -d redis

go build
go install && $projectname
