/*The protocolbuffer package.

to get tools:
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto


.proto files are schemes for buffer data]


compile proto files using:
protoc person.proto -I. --go_out=:.

Remember: Package must be define at top of proto file for go generation.
*/
package main
