# github-actions-book-go-excellent

[GitHub CI/CD 実践ガイド | 技術評論社](https://gihyo.jp/book/2024/978-4-297-14173-8)
の "4.2 自動テスト" のところ

## プルリクを CLI で作る

```sh
git switch feature-branch # or `git switch -c feature-branch`
# do something
git commit -am 'foobar'
git push origen feature-branch
gh pr create -t "新機能追加" -b "詳細説明" -B main -H feature-branch
```

## プルリクを CLI で確認

```sh
gh pr list
# 自動テストの結果を確認
gh pr check <プルリクIDまたはURL>
# 拒否する
gh pr close <プルリクIDまたはURL>
# マージする
gh pr merge <プルリクIDまたはURL>
```

## タスクランナー

GoLang らしく
[Task](https://taskfile.dev/)
にした。

actionlint などと一緒に `aqua i` でインストール

```console
$ task -l
task: Available tasks for this project:
* hello:                       Print a greeting message
* lint:                        Lint Go code
* lint-workflows:              Lint GitHub Workflows
* lint-workflows-docker:       Lint GitHub Workflows with Docker
* run:                         Run excellent Go program
* test:                        Test Go code
```
