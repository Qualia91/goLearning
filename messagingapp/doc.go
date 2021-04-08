/*
Running:
docker image build -t messagingapp .
docker container run -p 8080:8080 messagingapp


When using a scratch build (simple small container with no shell), and you want to look inside, user the following commands

docker create --name="tmp_$$" messagingapp:latest
docker export tmp_$$ | tar t
docker rm tmp_$$
*/
package main
