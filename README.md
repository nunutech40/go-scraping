

### RUN the main / run the go project
```
go build main.go
go build -o myapp main.go
```
### and how to run the main.go
```
./myapp
```

### cara nge curl + query
```
curl "http://localhost:8080/converdolartorp?rupiah=30000.0"
```

### cara nge curl + query html
```
curl "http://localhost:8080/scraphtml?url=https://github.com/nunutech40/go-scraping"
```



// catatan untuk new version scraping
req, _ := http.NewRequest("GET", "https://twitter.com/somepage", nil)
req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
// Lanjutkan dengan membaca response
