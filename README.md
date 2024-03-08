# An example of building a docker image


Build a docker image with overflow that uses the new onflow/crypto 0.25.0 that requires CGO


This uses goreleaser to build the app and the ko plugin to build the image with ko. 

Github actions will trigger on a tag pushed and build a new image an publis it. 
