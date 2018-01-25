## riff create node

Create a node.js function

### Synopsis


Generate the function based on the function source code specified as the filename, using the name
and version specified for the function image repository and tag.
For example, if you have a directory named 'square' containing a function 'square.js', you can simply type :

riff init node -f square

or

riff init node

from the 'square' directory

to generate the required Dockerfile and resource definitions using sensible defaults.

```
riff create node [flags]
```

### Options

```
  -h, --help   help for node
```

### Options inherited from parent commands

```
  -a, --artifact string       path to the function artifact, source code or jar file
      --config string         config file (default is $HOME/.riff.yaml)
      --dry-run               print generated function artifacts content to stdout only
  -f, --filepath string       path or directory to be used for the function resources, if a file is specified then the file's directory will be used (defaults to the current directory)
      --force                 overwrite existing functions artifacts
  -i, --input string          the name of the input topic (defaults to function name)
  -n, --name string           the name of the function (defaults to the functionName of the current directory)
  -o, --output string         the name of the output topic (optional)
  -p, --protocol string       the protocol to use for function invocations (defaults to 'stdio' for shell and python, to 'http' for java and node)
      --push                  push the image to Docker registry
      --riff-version string   the version of riff to use when building containers (default "0.0.2")
  -u, --useraccount string    the Docker user account to be used for the image repository (defaults to current OS username) (default "dturanski")
  -v, --version string        the version of the function image (default "0.0.1")
```

### SEE ALSO
* [riff create](riff_create.md)	 - Create a function

###### Auto generated by spf13/cobra on 25-Jan-2018