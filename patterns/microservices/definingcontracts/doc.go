/*Defining Contracts.

Consumer Driven Contract Design: Design contracts using interface definition language, which you can generate tests from.

Interface Definition Languages:
- OpenAPI specification (impl: Swagger)

go-swagger docker command in the folder you are using:
alias swagger="docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"

Swagger demo: Todo list:
Initialize a swagger file with swagger init. The following will generate our base example:

swagger init spec \
  --title "A Todo list application" \
  --description "From the todo list tutorial on goswagger.io" \
  --version 1.0.0 \
  --scheme http \
  --consumes definingcontracts.v1+json \
  --produces definingcontracts.v1+json


validate:
swagger validate ./swagger.yml

Generate:
swagger generate server -A todo-list -f ./swagger.yml

Generated package contains large --help output
*/

package definingcontracts
