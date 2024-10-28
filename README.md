# go-dla

Draw images using [diffusion-limited aggregation](https://en.wikipedia.org/wiki/Diffusion-limited_aggregation)

## How to run

```
go run .
```

## Versions

I ran those on my MBP M1 Max.  There are 10 versions in branches `v1` to `v10`,
each with a different optimisation.

| Branch | Optimisation | benchmark time |
|-- |-- |-- |
| [v1](https://github.com/arnodel/go-dla/tree/v1) | None at all| 8min16s |
| [v1..v2](https://github.com/arnodel/go-dla/compare/v1...v2) | Better data structure for world map | 1min35s |
| [v2..v3](https://github.com/arnodel/go-dla/compare/v2...v3) | Use all CPU cores | 12.7s |
| [v3..v4](https://github.com/arnodel/go-dla/compare/v3...v4) | Remove branching on hot path | 7.86s |
| [v4..v5](https://github.com/arnodel/go-dla/compare/v4...v5) | Do not waste a random bit | 4.08s |
| [v5..v6](https://github.com/arnodel/go-dla/compare/v5...v6) | Precompute neighbours | 2.50s |
| [v6..v7](https://github.com/arnodel/go-dla/compare/v6...v7) | Speed up hot path by splitting a function to inline it | 2.10s |
| [v7..v8](https://github.com/arnodel/go-dla/compare/v7...v8) | Speed up hot path by merging two functions | 1.92s |
| [v8..v9](https://github.com/arnodel/go-dla/compare/v8...v9) | Use the power of Maths! | 0.873s |
| [v9..v10](https://github.com/arnodel/go-dla/compare/v9...v10) | Reduce contention on communication channel | 0.687s |
