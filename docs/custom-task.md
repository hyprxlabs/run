# Custom Task

Custom tasks can be imported at the top level or as needed 
at the task level. 

## alias name:

The name of the task.

## alias description:

The description of the task, which becomes the default description of the task if its not
set in the runfile.yaml

## alias author:

The author of the task.

## alias inputs.{id}

The id of the input. Inputs will be passed in as environment variables and the
names will be converted from kebab case to screaming snake case with the prefix "INPUT_".

If the id of the input is **my-key**, the environment variable name will become
INPUT_MY_KEY.  Input values can use environment variable interpolation.

## alias input.{*}.uses

The built in task id that the step should use such as bash, pwsh, deno, node, shell, etc.

## alias input.{*}.run

The script or path to the file for steps that use shells or languages to write the task.

## alias input.{*}.env

Additional environment variables for the given step

## alias input.{*}.with

The input values if the built in tasks has parameters. If the built in task does not have input
parameters, this is ignored.

## alias input.{*}.if

Allows to step to only run if the condition is true e.g. `{{eq .os "windows"}}`

## alias input.{*}.cwd

Allows the step to set the current working directory.  Environment variables may be used to construct
the path.

## alias input.{*}.force

Forces the step to always run even if other steps have failed if it is set to true.


```yaml
tasks:
  - id: mkcert
    description: |
        Creates a mkcert self signed certificate

    inputs:
      domains:
        required: true
        desc: The domains created for the certificate
      name:
        required: false
        desc: The name of the public and private key that will get generated
      dest:
        required: false
        desc: The destination folder for the generated secrets

    outputs:
      domains:
        desc: The domains used to generate the cert
      public-key:
        desc: The full path to the public key
      private-key:
        desk: The full path to the private key

    steps:
      - uses: bash
        run: #my bash code here
```
