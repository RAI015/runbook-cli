# runbook CLI

Go製CLI `runbook` は `runbook.yaml` を読み込み、Markdown形式のRunbookを生成します。

## 導入目的

- RunbookをYAMLで構造化して管理し、読みやすいMarkdownへ自動変換する
- 手順書フォーマットの揺れを減らし、レビューしやすい形に統一する
- 運用ナレッジを明文化して、担当者依存（属人化）を減らす
- 障害対応や運用手順の更新を、PRベースで継続的に改善しやすくする

## Usage

```bash
go run ./cmd/runbook generate -i runbook.yaml -o runbook.md
```

ビルド後は次の形で実行できます。

```bash
./runbook generate -i runbook.yaml -o runbook.md
```

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

## Example

サンプル入力は `examples/runbook.yaml` を参照してください。
