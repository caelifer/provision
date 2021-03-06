drpcli workflows wait
=====================

Wait for a workflow's field to become a value within a number of seconds

Synopsis
--------

This function waits for the value to become the new value.

Timeout is optional, defaults to 1 day, and is measured in seconds.

Returns the following strings: complete - field is equal to value
interrupt - user interrupted the command timeout - timeout has exceeded

::

    drpcli workflows wait [id] [field] [value] [timeout] [flags]

Options
-------

::

      -h, --help   help for wait

Options inherited from parent commands
--------------------------------------

::

      -c, --catalog string      The catalog file to use to get product information (default "https://repo.rackn.io")
      -d, --debug               Whether the CLI should run in debug mode
      -E, --endpoint string     The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force               When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string       The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string     password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -r, --ref string          A reference object for update commands that can be a file name, yaml, or json blob
      -T, --token string        token of the Digital Rebar Provision access
      -t, --trace string        The log level API requests should be logged at on the server side
      -Z, --traceToken string   A token that individual traced requests should report in the server logs
      -U, --username string     Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli workflows <drpcli_workflows.html>`__ - Access CLI commands
   relating to workflows
