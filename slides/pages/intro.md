---
layout: section
---

# Introduction to Go

---

# History

- Created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson
- First public release in 2009
- Version 1.0 released in March 2012
- Go 1.0 code still works with the latest Go version, over a decade later
- Allegedly created in response to frustrations with C++ build times and complexity

---

## Where is Go used?

- Docker
- Kubernetes
- Istio
- Helm

---

## Where is Go used?

- Terraform
- Prometheus
- Grafana
- Loki
- HashiCorp tools (Consul, Vault, Nomad)

---

## Where is Go used?

- Kong
- Traefik
- Caddy

---

## Where is Go used?

- Google
- CloudFlare
- Just Eat
- Uber
- Twitch
- Dropbox
- Monzo
- ITV
- Aviva
- Thought Machine

---

## Where is Go used?

- Ollama
- TypeScript compiler

---
layout: two-cols-header
---

# Build simple, secure, scalable systems with Go

::left::

- An open-source programming language supported by Google
- Easy to learn and great for teams
- Built-in concurrency and a robust standard library
- Large ecosystem of partners, communities, and tools

::right::

<img src="/Go_Logo_Blue.svg" alt="Go logo" style="width: 200px; margin-left: 80px;">

---

# An open source programming language

- Download and use for free
- Compiler and tools are open source
- Developed in the open on GitHub
  - Anyone can contribute

---

# Supported by Google

- Uses Go
- Employs Go team members
- Security team
- Patent protection
- Go project infrastructure
  - Package cache
  - Vulnerability database

---
layout: two-cols-header
---

# Easy to learn

::left::

- Simple, familiar syntax (looks like C / Java / JavaScript / C#)
- Can learn the basics in a few hours
- Built-in tooling

::right::

```bash
go run
```

```bash
go build
```

```bash
go test
```

```bash
go get github.com/3rdparty/package
```

```bash
gofmt
```

```bash
gopls
```

<!--

- Don't need to choose between `pytest` or `unittest` - there's just `go test`
- No need to choose eslint or prettier, and choose rules - there's just `gofmt`
- No need to choose gradle or maven - there's just `go build` and `go get`

-->

---
layout: two-cols-header
---

# Great for teams

::left::

- Simple language, hard to mess up
- Go code tends to look similar, regardless of who wrote it
- Everyone uses the same tools
- Helpful warnings and error messages
- Use any editor you like

::right::

<img src="/ide_help.png"/>

---
layout: two-cols-header
---

# Built-in concurrency

::left::

- Language-level support for concurrent programming
- No async/await, promises, callbacks, or event loops needed
- goroutines: lightweight threads managed by the Go runtime
- Channels: typed conduits for communication between goroutines

::right::

```go {|1-10|2,9|3|4-8|5-6|7|12-19|13|14-16|15|18}
func backgroundTask(ctx context.Context, i int) {
  for {
    fmt.Printf("Hello from background task %d!\n", i)
    select {
      case <-ctx.Done():
          return
      case <-time.After(1 * time.Second):
    }
  }
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  for i := range 10000 {
    go backgroundTask(ctx, i)
  }
  time.Sleep(5 * time.Second)
  cancel()
}
```

---
layout: two-cols-header
---

# Robust standard library

<v-switch tag="ul" childTag="li" unmount>
  <template #1><code>bufio</code> buffered I/O</template>
  <template #2><code>bytes</code> / <code>strings</code> byte and string manipulation</template>
  <template #3><code>compress</code> zlib, gzip, bzip2, lzw, flate etc.</template>
  <template #4><code>context</code> cancellation, timeouts, request-scoped values</template>
  <template #5><code>crypto</code> TLS, SHA, AES, RSA, ECDSA etc.</template>
  <template #6><code>database/sql</code> SQL database access</template>
  <template #7><code>embed</code> embed files in the binary</template>
  <template #8><code>encoding</code> JSON, XML, CSV, Gob etc.</template>
  <template #9><code>errors</code> error handling</template>
  <template #10><code>flag</code> command-line flag parsing</template>
  <template #11><code>fmt</code> formatted I/O</template>
  <template #12><code>go</code> Go source code parsing and analysis</template>
  <template #13><code>hash</code> hash functions, e.g. MD5, SHA1, SHA256</template>
  <template #14><code>html</code> HTML encoding, decoding, templating</template>
  <template #15><code>image</code> image manipulation, encoding, decoding</template>
  <template #16><code>io</code> standard interfaces, copy operations etc.</template>
  <template #17><code>iter</code> iterator patterns</template>
  <template #18><code>log</code> / <code>slog</code> logging / structured logging</template>
  <template #19><code>math</code> / <code>math/rand</code> mathematics, random numbers</template>
  <template #20><code>net</code> networking, TCP, UDP, IP etc.</template>
  <template #21><code>net/http</code> HTTP client and server</template>
  <template #22><code>os</code> OS functions, file system access</template>
  <template #23><code>path</code> / <code>path/filepath</code> path manipulation</template>
  <template #24><code>regexp</code> regular expressions</template>
  <template #25><code>sort</code> sorting slices and user-defined collections</template>
  <template #26><code>strconv</code> string conversions</template>
  <template #27><code>sync</code> concurrency primitives</template>
  <template #28><code>testing</code> unit testing</template>
  <template #29><code>time</code> time and duration</template>
  <template #30><code>unicode</code> Unicode handling</template>
</v-switch>

---

# Robust standard library

- High quality, well tested, actively maintained over time
- Consistent API design
- Regularly updated with new features and improvements:
  - `net/http`
  - `context`
  - `slog`
  - `embed`
  - `iter`
- Crytographic audit and FIPS 140-2 compliance

---
layout: section
---

# Ecosystem of partners, communities, and tools

---
layout: two-cols-header
---

## Partners

::left::

- Cloud providers
- Databases
- Service providers
- etc.

<img src="/companies.png">

::right::

<img src="/aws-sdk.png" style="height: 200px; margin-left: 60px; margin-bottom: 20px;">

<img src="/stripe-sdk.png" style="height: 200px; margin-left: 60px;">

---
layout: two-cols-header
---

## Communities

::left::

- Slack
- Reddit
- Meetups
- Gophercon UK

<img src="/slack.png" style="height: 200px; margin-top: 50px; margin-bottom: 20px;">

::right::

<img src="/reddit.png" style="height: 200px; margin-left: 60px; margin-bottom: 20px;">

<img src="/meetup.png" style="height: 200px; margin-left: 60px;">

---
layout: two-cols-header
---

::left::

## Tools

- Autocompletion - `gopls`
- Linting - `golangci-lint`
- Security - `gosec`
- Dependency management - `dependabot`, `renovate`
- LLM assistants - `copilot`, `Claude`...
- Profiling - `pprof`
- Debugging - `delve`

::right::

## Editors

<img src="/editor.svg">

https://go.dev/blog/survey2024-h2-results#editor-awareness-and-preferences

<Tip>
The Go team maintains the VS Code Go extension, and maintains `gopls`, the language server that provides IDE features like auto-completion and go-to-definition, so the experience across editors is quite consistent.
</Tip>
