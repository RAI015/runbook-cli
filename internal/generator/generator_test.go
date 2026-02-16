package generator

import (
	"strings"
	"testing"
)

func TestGenerateFromYAML_Success(t *testing.T) {
	input := []byte(`title: 障害一次対応（API 5xx増加）
owner: Backend
severity: SEV2
purpose: 顧客影響を最小化し、原因切り分けと一次復旧を行う
prechecks:
  - 影響範囲（ユーザー/機能）を確認
steps:
  - title: 状況把握
    items:
      - 監視ダッシュボードで5xx率/レイテンシを確認
rollback:
  criteria:
    - 新規デプロイ後にエラーが増加
  actions:
    - 直近リリースをロールバック
notes:
  - 対応後はポストモーテム作成
`)

	got, err := GenerateFromYAML(input)
	if err != nil {
		t.Fatalf("GenerateFromYAML returned error: %v", err)
	}

	checks := []string{
		"# 障害一次対応（API 5xx増加）",
		"## Purpose",
		"## Steps",
		"1. 状況把握",
		"監視ダッシュボードで5xx率/レイテンシを確認",
		"## Rollback",
		"### Criteria",
		"## Notes",
	}
	for _, c := range checks {
		if !strings.Contains(got, c) {
			t.Fatalf("output missing expected content %q\n%s", c, got)
		}
	}
}

func TestGenerateFromYAML_MissingRequiredField(t *testing.T) {
	input := []byte(`title: Deploy API
purpose: test
steps:
  - title: Step 1
`)

	_, err := GenerateFromYAML(input)
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "missing required field: steps[0].items") {
		t.Fatalf("unexpected error: %v", err)
	}
}
