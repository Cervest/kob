# kob
A tool for starting Kubernetes jobs easily from the command line.

For now, it only handles job arguments, as we commonly want to manually run a job but with specific inputs. A templating engine would be overkill here, and would produce unwanted artefacts. `kob` applies changes in memory, and then immediately creates the job.

To create a single new job from a K8s spec-file, with the args `["one", "two", "three"]`:
```
kob with-args -f path/to/spec.yml arg1 arg2 arg3
```
Args are passed as arguments to the Docker command, they don’t override the command itself.

## Building
`go build -o kob`

## Troubleshooting
If K8s is complaining that one of your arguments is not executable, make sure you have `ENTRYPOINT` in your Dockerfile, not `CMD`, so you can pass args to it.
