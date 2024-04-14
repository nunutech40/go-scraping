
## RUN PROJECT
### RUN the main / run the go project
```bash
go build main.go
go build -o myapp main.go
```
### and how to run the main.go
```bash
./myapp
```

### cara nge curl + query
```bash
curl "http://localhost:8080/converdolartorp?rupiah=30000.0"
```

### cara nge curl + query html
```bash
curl "http://localhost:8080/scraphtml?url=https://github.com/nunutech40/go-scraping"
```

## GO MOD
* Setelah ada file go.mod, build atau perbarui file executable myapp menggunakan
```bash
go build -o myapp
```

### Cara menambahkan Go MOD Init
```bash
go mod init github.com/nunutech40/go-scraping
```

### cara import path modulenya
```go
import "github.com/nunutech40/go-scraping/handlers"
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
