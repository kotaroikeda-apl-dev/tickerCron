# Go の Ticker と Cron を使った定期実行サンプル

## 概要

このプロジェクトは、Go の `time.Ticker` と `github.com/robfig/cron/v3` を使用して、異なる間隔で定期的にジョブを実行する方法を学ぶためのものです。

- `time.Ticker` を使用して **2 秒ごとに 3 回実行されるジョブ (tickerJob)**
- `cron/v3` を使用して **5 秒ごとに実行されるジョブ (cronJob)**
- `tickerJob` は **ゴルーチン** を使って並行実行される

## コードの説明

### **Ticker を使ったジョブ**

```go
func tickerJob() {
    ticker := time.NewTicker(2 * time.Second) // 2秒ごとに Tick を送る Ticker を作成
    defer ticker.Stop() // 関数終了時に Ticker を停止

    for i := 0; i < 3; i++ { // 3回だけ実行
        fmt.Println("Tickerジョブ:", time.Now())
        <-ticker.C // Tick を待つ
    }
}
```

**2 秒ごとに 3 回実行されるジョブ**

- `ticker.C` で 2 秒ごとに tick を受信し、処理を実行
- 3 回実行後に終了

### **Cron を使ったジョブ**

```go
func cronJob() {
    fmt.Println("Cronジョブ実行:", time.Now())
}
```

**5 秒ごとに実行されるジョブ**

- 現在時刻を表示

### **メイン関数**

```go
func main() {
    c := cron.New() // Cron スケジューラを作成
    c.AddFunc("@every 5s", cronJob) // 5秒ごとに `cronJob` を実行
    c.Start() // Cron を開始

    go tickerJob() // Tickerジョブを別ゴルーチンで動かす

    time.Sleep(30 * time.Second) // 30 秒間待機
    c.Stop() // Cron を停止
}
```

- Cron のスケジューラ (`c`) を作成し、5 秒ごとに `cronJob()` を実行
- `tickerJob()` をゴルーチン (`go`) で並行実行
- 30 秒待機 (`time.Sleep(30 * time.Second)`) し、その間に Ticker と Cron の処理が動く
- 30 秒後に `c.Stop()` を実行し、Cron を停止

## 実行方法

```sh
go run main.go
```

## 実行結果 (出力例)

```sh
Tickerジョブ: 2025-03-01 12:00:00
Tickerジョブ: 2025-03-01 12:00:02
Tickerジョブ: 2025-03-01 12:00:04
Cronジョブ実行: 2025-03-01 12:00:05
Cronジョブ実行: 2025-03-01 12:00:10
Cronジョブ実行: 2025-03-01 12:00:15
Cronジョブ実行: 2025-03-01 12:00:20
Cronジョブ実行: 2025-03-01 12:00:25
Cronジョブ実行: 2025-03-01 12:00:30
```

**動作のポイント**

- `tickerJob()` は 2 秒ごとに 3 回実行
- `cronJob()` は 5 秒ごとに実行
- `main()` が 30 秒後に終了し、`cron` も停止

## 学習ポイント

1. `time.Ticker` を使って、指定間隔 (`2秒`) で処理を実行する方法
2. `cron/v3` を使って、指定間隔 (`5秒`) で関数をスケジュールする方法
3. `go` を使って `tickerJob()` をゴルーチンで並行実行する方法
4. `time.Sleep()` で `main()` の処理をブロックし、Ticker & Cron の動作を保証する
5. `c.Stop()` で `cron` の処理を適切に終了させる

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
