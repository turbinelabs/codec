
[//]: # ( Copyright 2017 Turbine Labs, Inc.                                   )
[//]: # ( you may not use this file except in compliance with the License.    )
[//]: # ( You may obtain a copy of the License at                             )
[//]: # (                                                                     )
[//]: # (     http://www.apache.org/licenses/LICENSE-2.0                      )
[//]: # (                                                                     )
[//]: # ( Unless required by applicable law or agreed to in writing, software )
[//]: # ( distributed under the License is distributed on an "AS IS" BASIS,   )
[//]: # ( WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or     )
[//]: # ( implied. See the License for the specific language governing        )
[//]: # ( permissions and limitations under the License.                      )

# turbinelabs/codec

[![Apache 2.0](https://img.shields.io/hexpm/l/plug.svg)](LICENSE)
[![GoDoc](https://https://godoc.org/github.com/turbinelabs/codec?status.svg)](https://https://godoc.org/github.com/turbinelabs/codec)
[![CircleCI](https://circleci.com/gh/turbinelabs/codec.svg?style=shield)](https://circleci.com/gh/turbinelabs/codec)

The codec project provides a simple interface for encoding and decoding values
with JSON and YAML implementations, along with a means to configure them
with a flag.FlagSet.

## Requirements

- Go 1.7.4 or later (previous versions may work, but we don't build or test against them)

## Dependencies

The codec project has no external dependencies; the tests depend on our
[test package](https://github.com/turbinelabs/test).
It should always be safe to use HEAD of all master branches of Turbine Labs
open source projects together, or to vendor them with the same git tag.

A [gomock](https://github.com/golang/mock)-based MockCodec is provided.

Additionally, we vendor
[github.com/ghodss/yaml](https://github.com/ghodss/yaml). This should be
considered an opaque implementation detail, see
[Vendoring](http://github.com/turbinelabs/developer/blob/master/README.md#vendoring)
for more discussion.

## Install

```
go get -u github.com/turbinelabs/codec/...
```

## Clone/Test

```
mkdir -p $GOPATH/src/turbinelabs
git clone https://github.com/turbinelabs/codec.git > $GOPATH/src/turbinelabs/codec
go test github.com/turbinelabs/codec/...
```

## Godoc

[`codec`](https://godoc.org/github.com/turbinelabs/codec)

## Versioning

Please see [Versioning of Turbine Labs Open Source Projects](http://github.com/turbinelabs/developer/blob/master/README.md#versioning).

## Pull Requests

Patches accepted! Please see [Contributing to Turbine Labs Open Source Projects](http://github.com/turbinelabs/developer/blob/master/README.md#contributing).

## Code of Conduct

All Turbine Labs open-sourced projects are released with a
[Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in our
projects you agree to abide by its terms, which will be vigorously enforced.
