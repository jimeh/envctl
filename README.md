<h1 align="center">
  envctl
</h1>

<p align="center">
  <strong>
    Go package providing test helper functions to temporarily change and restore
    environment variables.
  </strong>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/jimeh/envctl">
    <img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white"
  alt="Go Reference">
  </a>
  <a href="https://github.com/jimeh/envctl/releases">
    <img src="https://img.shields.io/github/v/tag/jimeh/envctl?label=release" alt="GitHub tag (latest SemVer)">
  </a>
  <a href="https://github.com/jimeh/envctl/actions">
    <img src="https://img.shields.io/github/workflow/status/jimeh/envctl/CI.svg?logo=github" alt="Actions Status">
  </a>
  <a href="https://codeclimate.com/github/jimeh/envctl">
    <img src="https://img.shields.io/codeclimate/coverage/jimeh/envctl.svg?logo=code%20climate" alt="Coverage">
  </a>
  <a href="https://github.com/jimeh/envctl/issues">
    <img src="https://img.shields.io/github/issues-raw/jimeh/envctl.svg?style=flat&logo=github&logoColor=white"
alt="GitHub issues">
  </a>
  <a href="https://github.com/jimeh/envctl/pulls">
    <img src="https://img.shields.io/github/issues-pr-raw/jimeh/envctl.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests">
  </a>
  <a href="https://github.com/jimeh/envctl/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/jimeh/envctl.svg?style=flat" alt="License Status">
  </a>
</p>

```go
os.Setenv("PORT", "8080")

envctl.With(map[string]string{"BIND": "0.0.0.0", "PORT": "3000"}, func() {
	fmt.Println(os.Getenv("BIND") + ":" + os.Getenv("PORT"))
})

fmt.Println(os.Getenv("BIND") + ":" + os.Getenv("PORT"))
```

```
0.0.0.0:3000
:8080
```

## Documentation

Please see the
[Go Reference](https://pkg.go.dev/github.com/jimeh/envctl#section-documentation)
for documentation and examples.

## Benchmarks

Benchmark reports and graphs are available here:
https://jimeh.me/envctl/dev/bench/

## License

[MIT](https://github.com/jimeh/envctl/blob/master/LICENSE)
