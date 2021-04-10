/*Simple Go Docker project.

This project is the MVP for creating a Go docker image and running it in a container via docker.

To build and run:

docker image build -t simplegodocker .

docker container run -p 3000:3000 simplegodocker

-p info:
the first port is the port the internals will be mapped to external to the docker container
the second is the port internal to the docker for other images to access

to post to localhost you need to run with --net=host to give it host network access. The ports exposed are then ignored


*/

package main
