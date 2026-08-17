# CoinMarketCap API Client для Go

![CoinMarketCap Go SDK](https://i.postimg.cc/tTfkMhvb/coinmarketcap-golang-api.jpg)

**🌐 Язык:** [English](README.md) | Русский

**📦 Пакет:** [coinmarketcap package](https://pkg.go.dev/github.com/tigusigalpa/coinmarketcap-go)

**Профессиональная интеграция CoinMarketCap API v1 для Go приложений**

Современный, готовый к продакшену Go пакет, который обеспечивает бесшовную интеграцию с API криптовалют CoinMarketCap.
Создан для Golang разработчиков, которым нужен надежный доступ к данным о криптовалютах в реальном времени, рыночным
данным, информации о биржах и блокчейн аналитике.

## 🌟 Почему выбрать этот CoinMarketCap Go клиент?

Независимо от того, создаете ли вы трекер криптовалютного портфеля, криптоторгового бота, дашборд блокчейн аналитики или
интегрируете цены Bitcoin и Ethereum в реальном времени в ваше Go приложение, этот пакет предоставляет все необходимое:

- **Готов к продакшену**: Протестированный код с комплексной обработкой ошибок
- **Удобен для разработчиков**: Интуитивный fluent интерфейс, разработанный специально для Go
- **Типобезопасный**: Полная типобезопасность Go со структурированной обработкой ошибок
- **Хорошо документирован**: Обширная документация с примерами реальной интеграции
- **Оптимизирован по производительности**: Эффективная реализация HTTP клиента
- **Гибкая архитектура**: Легко интегрируется в любое Go приложение

## 📦 Установка

### Шаг 1: Установка через Go Get

```bash
go get github.com/tigusigalpa/coinmarketcap-go
```

### Шаг 2: Получите API ключ CoinMarketCap

1. Посетите [CoinMarketCap API Portal](https://coinmarketcap.com/api/)
2. Зарегистрируйте бесплатный аккаунт (базовый план включает 10,000 вызовов API/месяц)
3. Перейдите в панель API
4. Скопируйте ваш API ключ

### Шаг 3: Настройте переменные окружения (опционально)

Создайте файл `.env` в корне вашего проекта:

```env
COINMARKETCAP_API_KEY=ваш-api-ключ
COINMARKETCAP_BASE_URL=https://pro-api.coinmarketcap.com
COINMARKETCAP_TIMEOUT=30
COINMARKETCAP_USE_SANDBOX=false
```

## Быстрый старт

### Базовое использование

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
        "symbol":  "BTC,ETH",
        "convert": "USD",
    })
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Результат: %v\n", quotes)
}
```

## Доступные методы API

### Cryptocurrency API

| Метод                     | Эндпоинт                                     | Описание                                |
|---------------------------|----------------------------------------------|-----------------------------------------|
| `ListingsLatest()`        | `/v1/cryptocurrency/listings/latest`         | Список криптовалют по капитализации     |
| `ListingsHistorical()`    | `/v1/cryptocurrency/listings/historical`     | Исторические листинги                   |
| `QuotesLatest()`          | `/v2/cryptocurrency/quotes/latest`           | Текущие котировки криптовалют           |
| `QuotesHistorical()`      | `/v2/cryptocurrency/quotes/historical`       | Исторические котировки                  |
| `Info()`                  | `/v2/cryptocurrency/info`                    | Метаданные (логотипы, ссылки, описание) |
| `Map()`                   | `/v1/cryptocurrency/map`                     | Маппинг CoinMarketCap ID                |
| `OHLCVLatest()`           | `/v2/cryptocurrency/ohlcv/latest`            | Последние OHLCV данные                  |
| `OHLCVHistorical()`       | `/v2/cryptocurrency/ohlcv/historical`        | Исторические OHLCV данные               |
| `Categories()`            | `/v1/cryptocurrency/categories`              | Категории криптовалют                   |
| `TrendingLatest()`        | `/v1/cryptocurrency/trending/latest`         | Трендовые криптовалюты                  |
| `TrendingGainersLosers()` | `/v1/cryptocurrency/trending/gainers-losers` | Топ растущих и падающих                 |

### Exchange API

| Метод                 | Эндпоинт                           | Описание                     |
|-----------------------|------------------------------------|------------------------------|
| `ListingsLatest()`    | `/v1/exchange/listings/latest`     | Список бирж по объему торгов |
| `QuotesLatest()`      | `/v1/exchange/quotes/latest`       | Текущие данные бирж          |
| `QuotesHistorical()`  | `/v1/exchange/quotes/historical`   | Исторические данные          |
| `Info()`              | `/v1/exchange/info`                | Метаданные бирж              |
| `Map()`               | `/v1/exchange/map`                 | Маппинг ID бирж              |
| `MarketPairsLatest()` | `/v1/exchange/market-pairs/latest` | Торговые пары бирж           |

### Global Metrics API

| Метод                | Эндпоинт                               | Описание                   |
|----------------------|----------------------------------------|----------------------------|
| `QuotesLatest()`     | `/v1/global-metrics/quotes/latest`     | Текущие глобальные метрики |
| `QuotesHistorical()` | `/v1/global-metrics/quotes/historical` | Исторические метрики       |

### Tools API

| Метод               | Эндпоинт                     | Описание                       |
|---------------------|------------------------------|--------------------------------|
| `PriceConversion()` | `/v2/tools/price-conversion` | Конвертация цен между валютами |

## 🛡️ Обработка ошибок

Пакет предоставляет комплексную обработку ошибок для всех ошибок CoinMarketCap API:

### Типы ошибок

| Тип ошибки            | HTTP код  | Описание                            |
|-----------------------|-----------|-------------------------------------|
| `AuthenticationError` | 401       | Неверный или отсутствующий API ключ |
| `RateLimitError`      | 429       | Превышен лимит запросов API         |
| `InvalidRequestError` | 400       | Неверные параметры запроса          |
| `NotFoundError`       | 404       | Ресурс не найден                    |
| `APIError`            | Различные | Общая ошибка API                    |

### Базовая обработка ошибок

```go
package main

import (
    "errors"
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
    cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("ваш-api-ключ").
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
        "limit":   "100",
        "convert": "USD",
    })
    
    if err != nil {
        var authErr *cmcerrors.AuthenticationError
        var rateLimitErr *cmcerrors.RateLimitError
        
        switch {
        case errors.As(err, &authErr):
            fmt.Printf("Ошибка аутентификации: %v\n", authErr)
        case errors.As(err, &rateLimitErr):
            fmt.Printf("Превышен лимит. Повторить через: %v\n", rateLimitErr.GetRetryAfter())
        default:
            fmt.Printf("Ошибка: %v\n", err)
        }
        return
    }
    
    fmt.Printf("Листинги: %v\n", listings)
}
```

## 📝 Лицензия

MIT License - см. файл LICENSE для деталей

## 🔗 Ссылки

- [Документация CoinMarketCap API](https://coinmarketcap.com/api/documentation/v1/)
- [Получить API ключ](https://coinmarketcap.com/api/)
- [GitHub репозиторий](https://github.com/tigusigalpa/coinmarketcap-go)
