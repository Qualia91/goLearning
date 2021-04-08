Installation:

K3D: brew install k3d
Kubectl: snap install kubectl --classic

Build cluster:
k3d cluster create
(confirm with
	kubectl config get-clusters
	kubectl get node
)

Arkade (Go cli similar to brew but for k8s applications):
curl -sLS https://dl.get-arkade.dev | sudo sh

Postgres in Kubernetes
arkade install postgresql

to find info about postgress: arkade info postgresql

Use the commands to open up postgres cli and create table:

CREATE TABLE todo (
  id              INT GENERATED ALWAYS AS IDENTITY,
  description     text NOT NULL,
  created_date    timestamp NOT NULL,
  completed_date  timestamp
);

check table with: \dt


Installing the PLONK stack:
Prometheus: Metrics, auto-scaling and observability (health checking)
Linux
OpenFaas: Framework for building serverless functionswith docker and k8s
NATS: Cloud native computing foundation project for messaging and Pub/Sub
K8s

install openfaas:
arkade install openfaas

to verify:
kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"

Then read message to:
1) Install faas-cli: curl -SLsf https://cli.openfaas.com | sudo sh
2) Get password:
	PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
	echo $PASSWORD
3) Port forward the OpenFaas gateway UI via port 8080:
	kubectl rollout status -n openfaas deploy/gateway
	kubectl port-forward -n openfaas svc/gateway 8080:8080 &
4) Login via cli
	echo -n $PASSWORD | faas-cli login --username admin --password-stdin
	faas-cli store deploy figlet
	faas-cli list

default username is admin. localhost:8080 is ui and will ask for username and password

Creating a go api with plonk: Use faas templates:
find with:
faas-cli template store list | grep go

For our example, we want basic http style, so use golang-middleware
faas-cli template store pull golang-middleware

The scaffold an API with golang-middleware and you docker hub username
export PREFIX=<USER_NAME>
export LANG=golang-middleware
export API_NAME=todo

faas-cli new --lang $LANG --prefix $PREFIX $API_NAME


build and deploy code with:
faas-cli up -f todo.yaml

Go to function site:
curl http://127.0.0.1:8080/function/todo

Create K8s secrets file:
we need the postgresql information to set in secrest.
Make sure the export POSTGRES_PASSWORD has been run (found via arkade info postgresql)

The run:
export USERNAME="postgres"
export PASSWORD=$POSTGRES_PASSWORD
export HOST="postgresql.default"

to save the info in export variables. Then save these to secrets:
faas-cli secret create username --from-literal $USERNAME
faas-cli secret create password --from-literal $PASSWORD
faas-cli secret create host --from-literal $HOST


Check this has worked via either of the following:
faas-cli secret ls
kubectl get secret -n openfaas-fn

These secret names need to be added to the yaml file.

Run build again:
faas-cli build -f todo.yml --build-arg GO111MODULE=on

Invoke the endpoint:
echo | faas-cli invoke todo -f todo.yml

I had a not null constraint on as field that i wanted to be null.
To resolve this i ran:
alter table todo alter column completed_date drop not null;

To insert data, either do it on postman or run curl:
curl http://127.0.0.1:8080/function/todo/create --data "User data"

To Get a list of data in json format:
curl http://127.0.0.1:8080/function/todo/list