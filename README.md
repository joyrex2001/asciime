# ASCIIme

Small demo project to show how to use minimal docker images in OpenShift using go.

To install the project on your OpenShift/minishift platform:

```bash
oc new-project asciime
oc process -f deploy-ocp.yaml | oc create -f -
oc start-build asciime --from-dir=.
```
