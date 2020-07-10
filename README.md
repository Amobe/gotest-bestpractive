# Best Practice in Go Test

本專案將透過 go build tag 的機制，分離測試案例。藉此讓我們可以在不同的情境下執行特定的測試。

## File structure

go build tag 是檔案層級的選擇器，因此我們需要定義好每個檔案的用途。

```
exa
├── Makefile
├── exa.go                  // 程式邏輯
├── exa_test.go             // 公用程式邏輯測試
├── exa_internal_test.go    // 私用程式邏輯測試
└── exa_integration_test.go // 整合測試
```

## Setting Build Tag

在檔案的第一行設定了 unit 與 integration 兩種 build tag。並預留一個檔案沒有設定任何 tag。

File: exa_internal_test.go
```go
// +build unit

package exa
```

File: exa_integration_test.go
```go
// +build integration

package exa_test
```

File: exa_test.go
```go
package exa_test
```

## Execute Test

執行測試的指令中，可以加入 `tags` 參數，選擇我們要執行的測試。
實際指令如下。

```bash
# run without tag
go test -v ./...

# run with tag
go test -v -tags=unit ./...

# run with multiple tag
go test -v -tags=unit,integration ./...
go test -v -tags="unit integration" ./...
```

根據 tag 組合，就可以依照情境執行不同的測試案例。
配對表如下

| Tags | Include Test Case |
| --- | --- |
| none | exa_test.go |
| unit | exa_test.go<br>exa_internal_test.go |
| integration | exa_test.go<br>exa_integration_test.go |
| unit<br>integration | exa_test.go<br>exa_internal_test.go<br>exa_integration_test.go |

## Use Case

在 `exa_test.go` 沒有設定 build tag ，是為了呈現沒有指定 tag 的檔案，無論在什麼情況下都會被執行。

實際上，我們會將 `exa_test.go` 加上 `unit` 的 tag。

CI/CD 流程中，在 build 之後執行 unit test，並在 deploy 之後執行 integration test。 

## Bonus

* Goland 會忽略被加上 tag 的檔案，被忽略的檔案不會有 pre-compile 與 type check

    在設定頁中，加上 `unit integration`，即可避免檔案被忽視。

    Preferences > Go > Build Tags and Vendoring > Custom tags

* Go Test Sum

    https://github.com/gotestyourself/gotestsum
    
    非常便利的 golang 測試工具，可以簡要呈現測試的報告，支援不同型態的檔案輸出，可以匯入 Jenkins 等 CI/CD 平台。

