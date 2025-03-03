# Go の Ticker と Cron を使ったサンプル

## 概要

このプロジェクトは、Go の `time.Ticker` と `github.com/robfig/cron/v3` を活用して定期的な処理を実行する方法を示します。

## `time.Ticker` とは？

`time.Ticker` は、指定した間隔ごとに時間を通知するチャンネル (`ticker.C`) を提供する Go の機能です。定期的な処理を実行する際に便利です。

## `cron` とは？

`github.com/robfig/cron/v3` は、Unix の `cron` のように定期的なタスクをスケジュールできる Go のライブラリです。

## ディレクトリ構成

```
cmd/
│── basic/
│   └── main.go   # 基本的な Ticker と Cron の組み合わせ
│── timeIntervalUpdate/
│   └── main.go   # 動的に Ticker の間隔を変更する Cron の活用
```

---

## `basic/main.go`

### 説明

このプログラムは、以下の 2 つのジョブを実行します。

1. **Ticker ジョブ** (2 秒ごとに 3 回実行)
   - `time.NewTicker` を使用し、2 秒間隔で合計 3 回ログを出力します。
2. **Cron ジョブ** (5 秒ごとに実行)
   - `cron.AddFunc("@every 5s", cronJob)` を使用し、5 秒ごとにログを出力します。

### 実行方法

```sh
go run basic/main.go
```

### 出力例

```
Tickerジョブ: 2025-03-04 12:00:00.123456 +0000 UTC
Cronジョブ実行: 2025-03-04 12:00:05.123456 +0000 UTC
Tickerジョブ: 2025-03-04 12:00:02.123456 +0000 UTC
Tickerジョブ: 2025-03-04 12:00:04.123456 +0000 UTC
Cronジョブ実行: 2025-03-04 12:00:10.123456 +0000 UTC
...
```

### 学んだポイント

- `Ticker` を用いた定期処理の実行
- `Cron` を用いたスケジュール実行
- `goroutine` を使った非同期処理
- `time.Sleep()` を使ったメインループの制御

---

## `timeIntervalUpdate/main.go`

### 説明

このプログラムでは、`Ticker` の間隔を動的に変更する方法を示します。

1. **Ticker ジョブ** (最初は 5 秒ごとに実行)
   - `time.Sleep(tickerInterval)` を使い、初期状態では 5 秒間隔でログを出力します。
2. **Cron ジョブ** (15 秒ごとに実行)
   - 15 秒ごとに `changeInterval` 関数を呼び出し、Ticker の間隔を 2 秒に変更します。

### 実行方法

```sh
go run timeIntervalUpdate/main.go
```

### 出力例

```
Tickerジョブ実行: 2025-03-04 12:00:00.123456 +0000 UTC
Tickerジョブ実行: 2025-03-04 12:00:05.123456 +0000 UTC
Tickerジョブ実行: 2025-03-04 12:00:10.123456 +0000 UTC
Tickerの間隔を 2秒 に変更
Tickerジョブ実行: 2025-03-04 12:00:12.123456 +0000 UTC
Tickerジョブ実行: 2025-03-04 12:00:14.123456 +0000 UTC
...
```

### 学んだポイント

- `Ticker` の間隔を変更する方法
- `Cron` を使って定期的に `Ticker` の間隔を更新
- `goroutine` を用いた非同期処理
- `time.Sleep()` を用いたシンプルな時間制御

---

## 使用ライブラリ

- [robfig/cron](https://github.com/robfig/cron): Go 言語用の Cron スケジューラ

## 実行環境

- Go 1.18 以上推奨
- `go get github.com/robfig/cron/v3` を事前に実行

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
