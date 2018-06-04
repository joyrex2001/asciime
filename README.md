# ASCIIme

ASCIIme is a small application that will convert uploaded images to an
ASCII-art version of the image. The project is intended to demonstrate
how to use minimal docker images in OpenShift using go.

## Install in OpenShift

To install the project on your OpenShift/minishift platform:

```bash
oc new-project asciime
oc process -f deploy-ocp.yaml | oc create -f -
oc start-build asciime --from-dir=.
```

## Just run...

If for some reason, you are actually interested in the application itself,
rather than the process building it, you can run it ```go run main.go``` and
visit http://localhost:8080 to have some fun.
