# runbook CLI

Go製CLI `runbook` は `runbook.yaml` を読み込み、Markdown形式のRunbookを生成します。

## Usage

```bash
go run ./cmd/runbook generate -i runbook.yaml -o runbook.md
```

ビルド後は次の形で実行できます。

```bash
./runbook generate -i runbook.yaml -o runbook.md
```

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

サンプル入力は `/Users/mirai/Works/runbook-gen/examples/runbook.yaml` を参照してください。
