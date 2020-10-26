# BindingData
> A helper application for testing KubePreset

This application provides a web API to access data bound through KubePreset.
The data could be either environment variables or files mounted at a certain
directory indicated by `SERVICE_BINDING_ROOT` environment variable.  The
`SERVICE_BINDING_ROOT` is an environment variable set by the KubePreset.

This latest version of this application will be available as a docker image at
`quay.io/kubepreset/bindingdata:latest`.

## API

```
/env/{VARIABLE_NAME} => {"value": "<VALUE>"}`
```

The key will be always literval `value` and `<VALUE>` is environment variable
value.

This data will be returned from `/files` end point like this:

```
/files => {"account-database": [{"type": "<TYPE"},
                                   {"provider": "<PROVIDER>"},
                                   {"uri": "<URI>"},
                                   {"username": "<USERNAME>"},
                                   {"password": "<PASSWORD>"}],
              "transaction-event-stream": [{"type": "<TYPE>"},
                                   {"connection-count": "<connection-count>"},
                                   {"uri": "<URI>"},
                                   {"certificates": "<CERTIFICATES>"},
                                   {"private-key": "<PRIVATE-KEY>"}]}
```
The above data is based on
https://github.com/k8s-service-bindings/spec#example-directory-structure
