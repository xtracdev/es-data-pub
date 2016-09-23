# Event Store Data Publisher

This package provides a general event publishing mechanism for
[oraeventstore](https://github.com/xtracdev/oraeventstore). It essentially
uses [orapub](https://github.com/xtracdev/orapub) for the actual publishing,
providing a process wrapper that first calls Initialize on all the 
registered event processors, then calls ProcessEvent for each event,
which distribute the event to each event processor.

This event publisher connects to Oracle using the [oraconn](https://github.com/xtracdev/oraconn)
package, taking connection parameters from the environment.

## Contributing

To contribute, you must certify you agree with the [Developer Certificate of Origin](http://developercertificate.org/)
by signing your commits via `git -s`. To create a signature, configure your user name and email address in git.
Sign with your real name, do not use pseudonyms or submit anonymous commits.


In terms of workflow:

0. For significant changes or improvement, create an issue before commencing work.
1. Fork the respository, and create a branch for your edits.
2. Add tests that cover your changes, unit tests for smaller changes, acceptance test
for more significant functionality.
3. Run gofmt on each file you change before committing your changes.
4. Run golint on each file you change before committing your changes.
5. Make sure all the tests pass before committing your changes.
6. Commit your changes and issue a pull request.

## License

(c) 2016 Fidelity Investments
Licensed under the Apache License, Version 2.0