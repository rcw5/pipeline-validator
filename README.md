# Concourse Secrets Validator

[Concourse](http://concourse.ci) is awesome. However validating whether a pipeline and its vars are complete is sub-awesome.

Pipeline vars can be surrounded with either brackets or curly braces. If a curly-braces var is missing from the vars file then `fly set-pipeline` will throw an error:

```
could not resolve old-style template vars: 1 error occurred:

* unbound variable in template: 'FooBar'
```

However if a var surrounded by brackets is undeclared then it'll just set the pipeline with the raw (bracketed) value.

There may be times when this behaviour is what you want, but most often I want to know that my vars are missing. Checking the `set-pipeline` output manually is difficult for complex pipelines.

This tool validates a pipeline definition against a set of vars files and will output:
- Any vars present in the pipeline but not declared in the vars files
- Any extra vars present in the vars files but unused in the pipeline

## Usage

```NAME:
   vars-validator - Validate a Concourse pipeline and its vars

USAGE:
   vars-validator [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --pipeline value, -p value        Pipeline definition
   --load-vars-from value, -l value  Vars (secrets) file to load
   --help, -h                        show help
   --version, -v                     print the version
```
