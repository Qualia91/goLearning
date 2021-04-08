/* This is a doc block.
This file is used to provide documentation about the package as a whole.
The first line is shown with the name.

Some info:
1)  package names should be lower case and not have spaces or _.
2)  should match the parent folder.
3)  init() functions in package are called after package initialisation (including package variables), before any code.
4)  init() in main.go always fires last due to package loading order. order for others is not defined.
5)  cant call init() manually.
6)  Scope: Capitalise (Public), Lower (Package).
7)  Internal Package: Scoped to a package.
8)  go doc <package> <Optional function>.
9)  _ import: Importing for side effects runs all init() functions in package.
10) Can aliase package import by putting name before import.
11) Internal package: Call package and folder "internal".
*/

package packagesingo
