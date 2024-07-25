# Go Templates

Some templates for building Go projects.

## Templates

- [cli](cli/): CLI starter template

## Usage

Install `gonew` using `go install`:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Execute the command gonew in the parent directory of your new project.
Provide two arguments: first, the path to the template you want to copy,
and second, the module name for your new project.

For example:

```bash
gonew github.com/ioboi/go-templates/cli github.com/ioboi/my-cli
cd my-cli
```
