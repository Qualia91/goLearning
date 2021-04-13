/*Data consistency

Options:
- Two phase commit pattern: Chooses consistency over availability
	ACID: Atomicity, Consitency, Isolation, Durability
- Saga Pattern: Trades atomicity for availability
- Eventual consitency pattern: Chooses availability over consistency.
	Uses data replication or event based.
	BASE: Basic availability, Soft state, Eventual consitency.
*/

package dataconsistency
