## ポートフォリオ概要
- DUELSCORE
    - 海外サッカー試合情報サイト（試合日程、試合スコア、得点ランキング、順位表）
    - 試合情報はAPI（[football-data.org](https://www.football-data.org/)）を使い取得

## URL
[https://duelscore.net/](https://duelscore.net/)

## テストアカウント
ヘッダー内「ログイン」から、下記情報でログイン可
- メールアドレス：test@test.com
- パスワード：zzzzAAAA11

※ログイン後のマイページ機能等は、今後追加予定

## 主要技術
- 言語：Golang 1.17
- ORM：GORM 1.23.10
- JSONパーサー：GJSON 1.14.1
- DB：MySQL 8.0
- セッション管理：Redis 7.0.5
- バージョン管理：Github
- コンテナ管理：Docker/Docker-compose
- インフラ：AWS
    - VPC
    - ECS(Fargate)
    - ECR
    - ELB
    - S3
    - IAM
    - Route53
    - CloudWatch
    - Secrets Manager
- IaC：Terraform

## 実装機能
- 試合日程取得機能
- 得点ランキング取得機能
- 順位表取得機能
- ユーザー管理機能(ユーザー登録/ログイン/ログアウト)

## 工夫した点
- 不要なAPIコールを防ぐ手段として、APIから取得した試合情報は30秒間Redisで一時保持し、使い回すようにしている。
    
## 苦労した点
- エラーハンドリングの構成（開発者によって様々あり、都度見直し中）。
- Redisの操作方法（セッション生成、破棄など）について掲載されているサイトが少ない。
- インフラ(AWS)やIaC（Terraform）の仕様理解。現在も試行錯誤中。

## 今後実装したい機能/実装中機能
- ユーザー管理機能（登録情報変更）
- チャット機能
- オッズ投票機能
- 取得情報追加（各国リーグ情報、選手情報など）

## 今後実施したい取り組み
- CI/CDパイプライン構築（CircleCI、またはGithubActionsによる、自動デプロイ＆テスト）

## （参考）ポートフォリオのスクリーンショット


<img width="1440" alt="portfolio_sample1" src="https://user-images.githubusercontent.com/77528519/211614837-3ed1f0f0-7d6c-4af4-880d-c94077936d21.png">

<img width="1440" alt="portfolio_sample2" src="https://user-images.githubusercontent.com/77528519/211615026-f62b72ca-1134-43ea-ab18-ead0f4d3592c.png">

<img width="1437" alt="portfolio_sample3" src="https://user-images.githubusercontent.com/77528519/211615091-ffcfb0aa-5934-43cc-8ed3-6b9c8387b38a.png">

<img width="1440" alt="portfolio_sample4" src="https://user-images.githubusercontent.com/77528519/211615072-2b1bc199-cadf-4830-9466-4d5cec06971e.png">