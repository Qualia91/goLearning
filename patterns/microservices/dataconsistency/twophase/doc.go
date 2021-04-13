/*Two Phase.

Transaction manager manages transactions with following procedure:

- Prepare phase: asks services if they can save data
- Voting phase: Services tell manager if they can save data
- Commit: If all services can save data, manager issues commit (tells services to save data)
- If one service says no, manager issues rollback command (undo save)

In the following example:
Transaction manager sends prepare message to hello and world services
They send messages back saying they can save data
Manager sends a commit message to get services to save data
Services save data

Bad for scaling, only use on small networks
*/

package twophase
