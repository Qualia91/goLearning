/*Context pattern.

When dealing with Go Servers, each request handles its own go routines, and need to access specific information
about who requested them. In order to manage this, the Context interface was created to pass request scoped
values, cancellation signals and deadlines across API boundaries.

More information can be found: https://blog.golang.org/context
*/

package main
