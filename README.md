# An example of building a docker image


Build a docker image with overflow that uses the new onflow/crypto 0.25.0 that requires CGO


This uses goreleaser to build the app and the ko plugin to build the image with ko. 

Github actions will trigger on a tag pushed and build a new image an publis it. 


## Error with ko

When using ko in goreleaser with an image built with CGO_ENABLED=1 (in main branch) i do not get a working image.

```
docker-compose up
Pulling example (ghcr.io/bjartek/overflow-example-image:0.2.3)...
0.2.3: Pulling from bjartek/overflow-example-image
0d64d3736fe6: Already exists
250c06f7c38e: Already exists
76fca8495c8e: Pull complete
Digest: sha256:f23cb9b574f90ee6123b24b45933cfca11764a600b44940e134d8862fdc8f55a
Status: Downloaded newer image for ghcr.io/bjartek/overflow-example-image:0.2.3
Recreating overflow-example-image_example_1 ... done
Attaching to overflow-example-image_example_1
example_1  | exec /ko-app/overflow-example-image: no such file or directory
overflow-example-image_example_1 exited with code 1
```

However if i build it using CGO_ENABLED=0 it works in the branch (with-old-crypto)

```
Starting overflow-example-image_example_1 ... done
Attaching to overflow-example-image_example_1
example_1  | foo
example_1  | ⭐ Script test run result:"test"
``
