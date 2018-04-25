# et

Provides simple CLI to Go's elastictab formatting.

## Building

Once you have golang installed, just run:

```sh
go build
```

The executable binary `et` is now ready to be used.

## Usage

Just pipe or redirect any (bounded) input to `et` to reformat. For example:

```
cat data.tsv | et
```

```
et < data.tsv
```
