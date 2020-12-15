# Detecting Apple Silicon via Golang

## Build

1. Install `gotip` or Go 1.16

2. Run `make detectapplesilicon`

## Running

On an Intel mac:

```
$ ./detectapplesilicon
Running on intel mac
```

Natively on an Apple Silicon mac:

```
$ ./detectapplesilicon
Running on apple silicon natively
```

Running via Rosetta 2

```
$ arch -arch x86_64 ./detectapplesilicon
Running on apple silicon under Rosetta 2
```

Based on https://steipete.com/posts/apple-silicon-mac-mini-for-ci/#detecting-apple-silicon-via-scripts.