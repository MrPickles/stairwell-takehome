# Generator

This folder contains a simple Python server that generates a UUID.

Your goal is to deploy this NOT on Kubernetes.
In fact, feel free to run it on your local machine.
Pretty simple, right?

The catch here is that the other server (in the sibling folder) must be able to
reach this server, but this server cannot be exposed to the public.
(I'd might as well say it, but don't try anything cheeky like hosting
everything on the same machine.)

## Development

Running the server is straightforward.
There are no flags.

```sh
./generator.py
```

You're welcome to modify the code, but I doubt that will be necessary.
