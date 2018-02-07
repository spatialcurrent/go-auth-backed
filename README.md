# go-auth-backend

# Description

**go-auth-backend** contains a simple interface for building authentication [plugins](https://golang.org/pkg/plugin/) for Go servers.

# Building

```
go get github.com/spatialcurrent/go-auth-backend
```

# Usage

**Server**

```
backend, err := LoadFromPlugin("/path/to/custom/auth/plugin.so", "NewAuthenticationBackend")
```

**Plugin**

```
package main

import (
  "net/http"
)

import (
  "github.com/spatialcurrent/go-auth-backend/authbackend"
)

func NewAuthenticationBackend() authbackend.Backend {
  return &MyCustomAuthBackend{}
}

func (a *AuthenticationBackend) Type() string {
  return "My Custom Auth Backend"
}

func (a *AuthenticationBackend) Connect(options map[string]string) (error) {
...
}

func (a *AuthenticationBackend) AuthenticateRequest(r *http.Request) (bool, authbackend.User, authbackend.Team, error) {
...
}
```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-auth-backend/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
