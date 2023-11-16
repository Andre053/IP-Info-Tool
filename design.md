# DESIGN OVERVIEW
## Files
- main.go       Hosts main() and run(), the latter of which runs the CLI core
- commands.go   Functions called for each command available in the CLI
- handlers.go   Functions for handling repeatable actions, such as network requests
- helpers.go    Functions and structs that commands use