## Event Store Data Publisher

This package provides a general event publishing mechanism for
[oraeventstore](https://github.com/xtracdev/oraeventstore). It essentially
uses [orapub](https://github.com/xtracdev/orapub) for the actual publishing,
providing a process wrapper that first calls Initialize on all the 
registered event processors, then calls ProcessEvent for each event,
which distribute the event to each event processor.

This event publisher connects to Oracle using the [oraconn](https://github.com/xtracdev/oraconn)
package, taking connection parameters from the environment.

