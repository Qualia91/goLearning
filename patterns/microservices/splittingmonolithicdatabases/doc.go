/*Splitting up Monolithic Databases.

Idea: Database per service = Microdatabase

Approach to design:
- Data first vs Function first: Which is designed first.
- Data first approach leads to monolithic db (Anti-Pattern)
- Function first:
	- Start with bounded contexts models, which then define what should be in db for each microservice

Patterns:
- Event driven: Share data changes using events after saving it to local db
- Event Sourcing: State is a series of changes, not actual state. To get current state, you need to replay all events in correct order.
- Command Query Response Segregation (C.Q.R.S): Separate command modules from query models (R/W separation)
	- Databases synced via low level database technologies

*/

package splittingmonolithicdatabases
