# gramework  [![codecov](https://codecov.io/gh/waveteam/gramework/branch/master/graph/badge.svg)](https://codecov.io/gh/waveteam/gramework) [![Build Status](https://travis-ci.org/waveteam/gramework.svg?branch=master)](https://travis-ci.org/waveteam/gramework) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1203/badge)](https://bestpractices.coreinfrastructure.org/projects/1203) [![Backers on Open Collective](https://opencollective.com/gramework/backers/badge.svg)](#backers) [![Sponsors on Open Collective](https://opencollective.com/gramework/sponsors/badge.svg)](#sponsors) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgramework%2Fgramework.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgramework%2Fgramework?ref=badge_shield)

The Good Framework

[![Gramework Stats Screenshot](https://10357-5.s.cdn13.com/docs/gramework_stats_screenshot.png)](https://grafana.com/dashboards/3422)

_Gramework long-term testing stand metrics screenshot made with [Gramework Stats Dashboard](https://grafana.com/dashboards/3422) and [metrics middleware](https://github.com/waveteam/gramework/tree/dev/metrics)_

### What is it?
Fast, highly effective, reliable SPA-first go-way web framework made by fasthttp maintainer. You get the simple yet powerful API, we handle optimizations internally.
We glad to see your feature requests and PRs.

-----

**Reasons to use Gramework**

- Gramework has stable API
- Gramework is battle-tested
- Gramework made by fasthttp maintainer
- Gramework is one of the rare frameworks that can help you to use server's resources effectively
- Gramework make your projects' infrastructure costs more effective by using as less memory as possible
- Gramework helps you serve requests faster, and so it helps you increase conversions ([source 1](https://blog.kissmetrics.com/speed-is-a-killer/), [source 2](https://blog.hubspot.com/marketing/page-load-time-conversion-rates))
- With Gramework you can build software faster using simple yet optimized API
- With Gramework you can get agile support and get answers to your questions
- At Gramework team we respect our users
- You can contact me and donate for high priority feature
- You can be sure that all license questions are OK with gramework
- You can buy a corporate-grade support

**Go >= 1.9.6 is the oldest continously tested and supported version.**


### Useful links and info
If you find it, you can submit vulnerability via k@gramework.win.

| Name  | Link/Badge  	|
|---	|---		|
| Docs  | [GoDoc](https://godoc.org/github.com/waveteam/gramework) |
| Our Jira | [Jira](https://gramework.atlassian.net) |
| License Report | [Report](https://github.com/waveteam/gramework/tree/dev/third_party_licenses/REPORT.md) |
| Changelog | [Changelog](https://github.com/waveteam/gramework/tree/dev/docs/CHANGELOG.md) |
| Support us with a donation or become a sponsor | [OpenCollective](https://opencollective.com/gramework) |
| Our Telegram chat | [@gramework](https://t.me/gramework) |
| We have #gramework channel in the Gophers Slack | https://gophers.slack.com |
| Our Discord Server | https://discord.gg/HkW8DsD |
| Master branch coverage | [![codecov](https://codecov.io/gh/waveteam/gramework/branch/master/graph/badge.svg)](https://codecov.io/gh/waveteam/gramework) |
| Master branch status | [![Build Status](https://travis-ci.org/waveteam/gramework.svg?branch=master)](https://travis-ci.org/waveteam/gramework) |
| Dev branch coverage | [![codecov](https://codecov.io/gh/waveteam/gramework/branch/dev/graph/badge.svg)](https://codecov.io/gh/waveteam/gramework) |
| Dev branch status | [![Build Status](https://travis-ci.org/waveteam/gramework.svg?branch=dev)](https://travis-ci.org/waveteam/gramework) |
| CII Best Practices | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1203/badge)](https://bestpractices.coreinfrastructure.org/projects/1203) |
| Gramework Stats Dashboard for Grafana | https://grafana.com/dashboards/3422 |

# TOC
- [Benchmarks](#benchmarks)
- [3rd-party license info](#3rd-party-license-info)
- [Basic usage](#basic-usage)
  - [Hello world](#hello-world)
  - [Serving a dir](#serving-a-dir)
  - [Serving prepared bytes](#serving-prepared-bytes)
  - [Using dynamic handlers, example 1](#using-dynamic-handlers-example-1)
  - [Using dynamic handlers, example 2](#using-dynamic-handlers-example-2)

# Benchmarks
[![benchmark](https://raw.githubusercontent.com/smallnest/go-web-framework-benchmark/master/benchmark.png)](https://github.com/smallnest/go-web-framework-benchmark)

## Contributors
This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="graphs/contributors"><img src="https://opencollective.com/gramework/contributors.svg?width=890&button=false" /></a>

## Backers
Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/gramework#backer)]

<a href="https://opencollective.com/gramework#backers" target="_blank"><img src="https://opencollective.com/gramework/backers.svg?width=890"></a>

## Sponsors
Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/gramework)]

<a href="https://opencollective.com/gramework/sponsor/0/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/1/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/2/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/3/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/4/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/5/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/6/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/7/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/8/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/gramework/sponsor/9/website" target="_blank"><img src="https://opencollective.com/gramework/sponsor/9/avatar.svg"></a>

# 3rd-party license info
- Gramework is now powered by [fasthttp](https://github.com/valyala/fasthttp) and custom fasthttprouter, that is embedded now.
  You can find licenses in `/third_party_licenses/fasthttp` and `/third_party_licenses/fasthttprouter`.
- The 3rd autoTLS implementation, placed in `nettls_*.go`, is an integrated version of
  [caddytls](https://github.com/mholt/caddy/tree/d85e90a7b4c06d1698d0b96b695b05d41833fcd3/caddytls), because using it by simple import isn't an option:
  gramework based on `fasthttp`, that is incompatible with `net/http`.
  In [the commit I based on](https://github.com/mholt/caddy/tree/d85e90a7b4c06d1698d0b96b695b05d41833fcd3), caddy is `Apache-2.0` licensed.
  It's license placed in `/third_party_licenses/caddy`. @mholt [allow us](https://github.com/mholt/caddy/issues/1520#issuecomment-286907851) to copy the code in this repo.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgramework%2Fgramework.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgramework%2Fgramework?ref=badge_large)

# Basic usage
### Hello world
The example below will serve "hello, grameworld". Gramework will register flag "bind" for you, that allows you to choose another ip/port that gramework should listen:

```go
package main

import (
	"github.com/waveteam/gramework"
)

func main() {
	app := gramework.New()

	app.GET("/", "hello, grameworld")

	app.ListenAndServe()
}
```

If you don't want support `bind` flag, pass the optional address argument to `ListenAndServe`.

**NOTE**: all examples below will register `bind` flag.

### JSON world ;) Part 1
From version: 1.1.0-rc1

The example below will serve `{"hello":"grameworld"}` from the map. Gramework will register flag "bind" for you, that allows you to choose another ip/port that gramework should listen:

```go
package main

import (
	"github.com/waveteam/gramework"
)

func main() {
	app := gramework.New()

	app.GET("/", func() map[string]interface{} {
		return map[string]interface{}{
			"hello": "gramework",
		}
	})

	app.ListenAndServe()
}
```

### JSON world. Part 2
From version: 1.1.0-rc1

The example below will serve `{"hello":"grameworld"}` from the struct. Gramework will register flag "bind" for you, that allows you to choose another ip/port that gramework should listen:

```go
package main

import (
	"github.com/waveteam/gramework"
)

type SomeResponse struct {
	hello string
}

func main() {
	app := gramework.New()

	app.GET("/", func() interface{} {
		return SomeResponse{
			hello: "gramework",
		}
	})

	app.ListenAndServe()
}
```

### Serving a dir
The example below will serve static files from ./files:

```go
package main

import (
	"github.com/waveteam/gramework"
)

func main() {
	app := gramework.New()

	app.GET("/*any", app.ServeDir("./files"))

	app.ListenAndServe()
}
```

### Serving prepared bytes
The example below will serve byte slice:

```go
package main

import (
	"github.com/waveteam/gramework"
)

func main() {
	app := gramework.New()

	app.GET("/*any", []byte("some data"))

	app.ListenAndServe()
}
```

### Using dynamic handlers, example 1.
This example show you some ways to send response, how to use preconfigured logger etc.

```go
package main

import (
	"github.com/waveteam/gramework"
)

type SomeData struct {
	Name string
	Age  uint8
}

func main() {
	app := gramework.New()

	d := SomeData{
		Name: "Grame",
		Age: 20,
	}

	// service-wide CORS. you can also instead of using middleware
	// call ctx.CORS() manually
	app.Use(app.CORSMiddleware())

	app.GET("/someJSON", func(ctx *gramework.Context) {
		// send json, no metter if user asked for json, xml or anything else.
		if err := ctx.JSON(d); err != nil {
			// you can return err instead of manual checks and Err500() call.
			// See next handler for example.
			ctx.Err500()
		}
	})

	app.GET("/simpleJSON", func(ctx *gramework.Context) error {
		return ctx.JSON(d)
	})

	app.GET("/someData", func(ctx *gramework.Context) error {
		// send data in one of supported encodings user asked for.
		// Now we support json, xml and csv. More coming soon.
		sentType, err := ctx.Encode(d)
		if err != nil {
			ctx.Logger.WithError(err).Error("could not process request")
			return err
		}
		ctx.Logger.WithField("sentType", sentType).Debug("some request-related message")
		return nil
	})

	// you can omit context if you want, return `interface{}`, `error` or both.
	app.GET("/simplestJSON", func() interface{} {
		return d
	})

	// you can also use one of built-in types as a handler, we got you covered too
	app.GET("/hostnameJSON", fmt.Sprintf(`{"hostname": %q}`, os.Hostname()))

	wait := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Minute)
		app.Shutdown()
		wait <- struct{}{}
	}()

	app.ListenAndServe()

	// allow Shutdown() to stop the app properly.
	// ListenAndServe will return before Shutdown(), so we should wait.
	<- wait
}
```

### Using dynamic handlers, example 2. Simple FastHTTP-compatible handlers.
This example shows you how to migrate from fasthttp to gramework
without rewriting all your handlers.

```go
package main

import (
	"github.com/waveteam/gramework"
)

func main() {
	app := gramework.New()

	app.GET("/someJSON", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("another data")
	})

	app.ListenAndServe()
}
```
