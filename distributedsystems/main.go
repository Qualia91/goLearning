/*Notes

4 of the characteristics
1) Service discovery
2) Load balancing
3) Distributed tracing and logging
4) Service monitoring

Types of DS:
1) Hub and Spoke (Centralized):
	Hub sets up coordination and spoke interaction is done through hub
	Good for load balancing and centralized tracing and logging
	Bad for single point of failure (hub) and the hub has multiple roles so can be complex
2) Peer to Peer (DIstributed):
	Gateway communicates to peers directly, and peers communicate to all peers
	Good for no single point of failure, and they are highly decoupled
	Bad for service discovery and load balancing
3) Message Queues (All talk to queues):
	Gateway interacts with queue, which is middleware to peers
	Good for easy to scale and message persistance (if system fails, messages can be stored)
	Bad for single failure point (message queue) and can be difficult to configure
4) Hybrid (Mix):
	Complicated mix of the first 2
	Good for load balancing and more robust against service failure
	Bad for very complex and central service is prone to scope creep

Architectural Elements:
1) Languages
2) Frameworks
3) Transports (How messages are sent like HTTP or RPC)
4) Protocols (Method like JSON or Protocol Buffers)

Frameworks:
1) Go-Kit.io
2) Go-Micro.dev

Ports for current app:
RegistryService: 		3000
LogService: 			4000
TeacherPortalService:	5000
GradingService: 		6000
*/

package main
