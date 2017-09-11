export projectname=recipesvr

#run mongo docker container 
docker run --name mongo-rs -p 27017:27017 -d mongo


go build
go install && $projectname
