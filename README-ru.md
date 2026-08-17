# Клиент CoinMarketCap для Go

![CoinMarketCap Go SDK](https://i.postimg.cc/tTfkMhvb/coinmarketcap-golang-api.jpg)

**Язык:** [English](README.md) · Русский
**Пакет:** [pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go](https://pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go)

Небольшой Go-клиент для [CoinMarketCap API](https://coinmarketcap.com/api/documentation/v1/). Через один настроенный клиент можно получать котировки и листинги криптовалют, данные о биржах, общие метрики рынка и результаты конвертации цен.

Библиотека намеренно возвращает ответ API в виде `map[string]interface{}`. Так он остаётся близким к исходному JSON CoinMarketCap, а новые query-параметры можно использовать без ожидания обновления моделей. Если в приложении удобнее работать со своими структурами, преобразуйте поле `data` после запроса.

## Что понадобится

- Go 1.21 или новее
- API-ключ CoinMarketCap — [получить ключ](https://coinmarketcap.com/api/)

## Установка

```bash
go get github.com/tigusigalpa/coinmarketcap-go
```

## Быстрый старт

```go
package main

import (
	"fmt"
	"log"

	coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
	client, err := coinmarketcap.NewClientBuilder().
		SetAPIKey("ваш-api-ключ").
		Build()
	if err != nil {
		log.Fatal(err)
	}

	quotes, err := client.Cryptocurrency().QuotesLatest(map[string]string{
		"symbol": "BTC,ETH", "convert": "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", quotes["data"])
}
```

Каждый метод принимает `map[string]string` с query-параметрами. Набор параметров, формат ответа и доступность по тарифу смотрите в [официальной документации CoinMarketCap](https://coinmarketcap.com/api/documentation/v1/).

## Настройка

По умолчанию клиент обращается к production API и ждёт ответ до 30 секунд.

```go
client, err := coinmarketcap.NewClientBuilder().
	SetAPIKey("ваш-api-ключ").
	SetTimeout(60 * time.Second).
	Build()
```

Для песочницы передайте sandbox API-ключ и вызовите `UseSandbox()`:

```go
client, err := coinmarketcap.NewClientBuilder().
	SetAPIKey("ваш-sandbox-api-ключ").
	UseSandbox().
	Build()
```

`SetBaseURL()` пригодится, если нужно указать другой совместимый адрес — например, локальный сервер в тестах.

## Поддерживаемые эндпоинты

| Сервис | Методы |
| --- | --- |
| `Cryptocurrency()` | `ListingsLatest`, `ListingsHistorical`, `QuotesLatest`, `QuotesHistorical`, `Info`, `Map`, `OHLCVLatest`, `OHLCVHistorical`, `Categories`, `TrendingLatest`, `TrendingGainersLosers` |
| `Exchange()` | `ListingsLatest`, `QuotesLatest`, `QuotesHistorical`, `Info`, `Map`, `MarketPairsLatest` |
| `GlobalMetrics()` | `QuotesLatest`, `QuotesHistorical` |
| `Tools()` | `PriceConversion` |

Например, первые десять криптовалют из листинга:

```go
listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
	"start": "1", "limit": "10", "convert": "USD",
})
```

## Ошибки

Для обработки конкретных ошибок используйте `errors.As`:

| Тип | HTTP-статус |
| --- | --- |
| `*errors.AuthenticationError` | 401 |
| `*errors.InvalidRequestError` | 400 |
| `*errors.NotFoundError` | 404 |
| `*errors.RateLimitError` | 429 |
| `*errors.APIError` | Другие ошибки API и запросов |

```go
var rateLimitErr *cmcerrors.RateLimitError
if errors.As(err, &rateLimitErr) {
	if retryAfter := rateLimitErr.GetRetryAfter(); retryAfter != nil {
		fmt.Printf("Повторите через %d секунд\n", *retryAfter)
	}
}
```

`APIError` содержит HTTP-статус, распарсенный ответ (если он есть) и исходную ошибку сети или декодирования.

## Тесты

```bash
go test ./...
go test -race ./...
```

Будем рады issue и pull request. В вопросах по параметрам, лимитам и доступности данных ориентируйтесь прежде всего на [официальную документацию CoinMarketCap](https://coinmarketcap.com/api/documentation/v1/).

## Лицензия

[MIT](LICENSE)
