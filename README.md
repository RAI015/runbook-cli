# runbook CLI

[![CI](https://github.com/RAI015/runbook-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/RAI015/runbook-cli/actions/workflows/ci.yml)

Go製CLI `runbook` は `runbook.yaml` を読み込み、Markdown形式のRunbookを生成します。
Runbookを「PRで更新できる運用資産」にするための最小CLIです。

## Demo

入力 (`examples/runbook.yaml` の一部):

```yaml
title: "障害一次対応（API 5xx増加）"
purpose: "顧客影響を最小化し、原因切り分けと一次復旧を行う"
steps:
  - title: "状況把握"
    items:
      - "監視ダッシュボードで5xx率/レイテンシを確認"
```

出力 (`runbook.md` の一部):

```md
# 障害一次対応（API 5xx増加）

## 概要
- 目的: 顧客影響を最小化し、原因切り分けと一次復旧を行う
- 担当: Backend
- 重大度: SEV2

## 事前確認
- [ ] 影響範囲（ユーザー/機能）を確認

## 手順

### 1. 状況把握
- [ ] 監視ダッシュボードで5xx率/レイテンシを確認
```

## Quickstart

```bash
# テスト
go test ./...

# 生成（最短）
go run ./cmd/runbook generate -i examples/runbook.yaml -o runbook.md
```

`make` が使える場合:

```bash
make test
make build
make demo
```

インストール:

```bash
go install github.com/RAI015/runbook-cli/cmd/runbook@latest
```

## 導入目的

- RunbookをYAMLで構造化して管理し、読みやすいMarkdownへ自動変換する
- 手順書フォーマットの揺れを減らし、レビューしやすい形に統一する
- 運用ナレッジを明文化して、担当者依存（属人化）を減らす
- 障害対応や運用手順の更新を、PRベースで継続的に改善しやすくする

## 適用シナリオ

- リリース手順の標準化
- 障害一次対応手順の整備（オンコール向け）
- ロールバック手順の標準化
- 新メンバー向けの運用手順テンプレート作成
- 監査や引き継ぎで、手順の最新版をMarkdownとして配布する運用

## YAML Schema

必須項目:
- `title` (string)
- `purpose` (string)
- `steps` (array, 1件以上)
- `steps[].title` (string)
- `steps[].items` (string array, 1件以上)

任意項目:
- `owner` (string)
- `severity` (string)
- `prechecks` (string array)
- `rollback.criteria` (string array)
- `rollback.actions` (string array)
- `notes` (string array)

```yaml
title: string
purpose: string
owner: string # optional
severity: string # optional
prechecks: # optional
  - string
steps:
  - title: string
    items:
      - string
rollback: # optional
  criteria:
    - string
  actions:
    - string
notes: # optional
  - string
```

## License

MIT License
