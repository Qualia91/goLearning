/*Backwards Compatibility.

Microservices USP:
- Independently changeable
- independently deployable


How to keep backwards compatibility:
- Shared models and seperation of concerns: Seperating client contract models from internal models, and using adapter services to convert between.
- Contract testing: Keep document describing scheme of contracts and auto test against.
- Versioning strategies: major.minor.patch and choose which one breaks contracts. Send version in HTTP headers or request path
*/

package backwardscompatibility
