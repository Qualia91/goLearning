/*Splitting up Monolithic Databases.

Idea: Database per service = Microdatabase

Approach to design:
- Data first vs Function first: Which is designed first.
- Data first approach leads to monolithic db (Anti-Pattern)
- Function first:
	- Start with bounded contexts models, which then define what should be in db for each microservice

Patterns:
- Event driven: Share data changes using events after saving it to local db

*/

package splittingmonolithicdatabases
