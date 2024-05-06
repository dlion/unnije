# Unni jè 

_(Pronunced `/ˈʊn.ni ʤe/` from Sicilian language)_

Simple DNS Resolver written in Go

## Author
Domenico Luciani    
https://domenicoluciani.com

## Tests

The project has been implemented following a TDD approach as much as possible -due of the time I had available :)-, to run tests: `go test ./...`

## Run it
```sh
> ./unnije <domain> [<domain> ...]
```

## Output example
```sh
dlion@darkness unnije % ./unnije domenicoluciani.com
Querying 198.41.0.4 for domenicoluciani.com
Querying 192.41.162.30 for domenicoluciani.com
Querying 108.162.192.65 for domenicoluciani.com
104.21.47.30
```
or

```sh
dlion@darkness unnije % ./unnije domenicoluciani.com twitter.com
Querying 198.41.0.4 for domenicoluciani.com
Querying 192.41.162.30 for domenicoluciani.com
Querying 108.162.192.65 for domenicoluciani.com
172.67.144.42
Querying 198.41.0.4 for twitter.com
Querying 192.41.162.30 for twitter.com
Querying 198.41.0.4 for a.r06.twtrdns.net
Querying 192.55.83.30 for a.r06.twtrdns.net
Querying 205.251.195.207 for a.r06.twtrdns.net
Querying 205.251.192.179 for twitter.com
104.244.42.129
```


## Idea
This is the implementation of the [Build Your Own DNS Resolver](https://codingchallenges.fyi/challenges/challenge-dns-resolver/) with Go, following those requirements step-by-step.
Any idea or feedback to make it better are always welcomed!