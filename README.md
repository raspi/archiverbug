# archiverbug

See https://github.com/mholt/archiver/issues/341

```shell
% go build .
% ./archiverbug 
trying to open "empty.txt"
panic: file does not exist

goroutine 1 [running]:
main.main()
        main.go:32 +0x87
```

```shell
% go version
go version go1.18.3 linux/amd64
```
