To create a single new job from a K8s spec-file, with the args `["one", "two", "three"]`:
```
kob with-args path/to/spec.yml arg1 arg2 arg3
```
Args are passed as arguments to the Docker command, they don’t override the command itself.
