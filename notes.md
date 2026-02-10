


// https://www.perplexity.ai/search/php-deki-array-chunk-fonksiyon-4QFCCf_nQOSPSn3qkuoOKQ
// https://www.perplexity.ai/search/bana-golang-slice-icine-veri-e-OWvaQBJqRiWTrIqXJYNxag
// https://www.perplexity.ai/search/bana-golang-struct-slice-icine-9DOZYn4MRFqz8bpg7pU.PQ
// https://www.perplexity.ai/search/php-deki-array-chunk-fonksiyon-4QFCCf_nQOSPSn3qkuoOKQ
// https://www.perplexity.ai/search/golang-da-var-newdata-person-b-zqYQaxLhQvWcIPP33Kakog
//https://www.perplexity.ai/search/bana-golag-da-bir-metin-icinde-2hfskKVpQ6eDJCzaGzaLtA#da17ec90-a569-43a7-8324-2cda1b368bb4
/*
https://www.perplexity.ai/search/golang-da-var-newdata-person-b-zqYQaxLhQvWcIPP33Kakog?sm=d

Performans İpucu (Pre-allocation): Eğer ekleyeceğiniz verinin sayısı önceden belliyse (örneğin 100 kişi ekleyecekseniz), slice'ı make ile oluşturmak performansı artırır.

Yani var newData []Person yerine newData := make([]Person, 0, len(rawNames)) kullanırsanız, Go bellekte yer ayırır ve döngü sırasında sürekli yeni bellek tahsisi yapmaz.


```
learnGoWithTests
    |
    |-> helloworld
    |    |- hello.go
    |    |- hello_test.go    
    |
    |-> integers
    |    |- adder_test.go
    |
    |- go.mod
    |- README.md

```


```
$ go mod init example
go: creating new go.mod: module example
$ mkdir -p external
$ touch external/{external.go,external_test.go}
$ tree .
.
├── external
│   ├── external.go
│   └── external_test.go
└── go.mod

1 directory, 3 files
```
