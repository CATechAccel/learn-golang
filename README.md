# learn-golang

STEP1: プログラミング言語 Go 完全入門をやる！

## TODO

1. Go に触れる

   [Go に触れる](https://docs.google.com/presentation/d/1a1w9c7-tArPrCXWMZxIJ0Ozexp0szehY0G1P_1ldhd0/edit#slide=id.g4e29971f9a_0_0)

2. 基本構文

   [基本構文](https://docs.google.com/presentation/d/1DtK5HyT8zV7weuHeaoR6Wg88lT3lmCE8De4GuO5sB74/edit#slide=id.g4e29971f9a_0_0)

3. 関数と型

   [関数と型](https://docs.google.com/presentation/d/1VHRNg-qZH-3ngwHnK8tMIDYyTKn6pA3yp8dRLLNQeLA/edit#slide=id.g4cbe4d134e_0_0)

4. パッケージ

   [パッケージ](https://docs.google.com/presentation/d/13ykFLzbQXgiHebsPtmvsov5Kqi7R0OnNxUxY6BZZ_dA/edit#slide=id.g4f1426e3ae_0_0)

5. 抽象化

   [抽象化](https://docs.google.com/presentation/d/1vKdJHHx4A_sP_Ft-3cakQqfJRbKxWSI9axzQ-GHHOoY/edit#slide=id.g4f15a7d687_0_0)

6. エラー処理

   [エラー処理](https://docs.google.com/presentation/d/1QekRR0VE_5kEwUW7OatAx7OJmr65BHD5vkfFZ1lY13Y/edit?usp=sharing)

7. HTTP サーバーとクライアント

   [HTTP サーバとクライアント](https://docs.google.com/presentation/d/1enuUT-sQcIA8vFLQtc-Hthj_xsbRx88mG1BMrnRSMhk/edit?usp=sharing)

8. データベース

   [データベース](https://docs.google.com/presentation/d/1IyuZbug63J3JbqdE5qt3Rtv_D6J1oP185C8lWpjS5jA/edit?usp=sharing)

## github の運用ルール

このプロジェクトは、`IDD(Issue Driven Develop)`で行います！

### IDD(Issue Driven Develop)とは？

issue に対する開発を行い、必ず PullRequest も issue に向けた変更になっているように設計する開発手法です。

### ルール

1. master, develop への直接 Push は 🆖
2. ブランチルールに則った Branch 名で作業ブランチを切ること
3. PullRequest は必ず develop に向けて作ること（develop→master = デプロイという位置付けになります）

### ブランチルール

```
feat/[issue番号]/[issueの内容（英語で）]

[例]
feat/1/createBaseWebApplication
```

### 作業の流れ

1. [Project](https://github.com/CATechAccel/learn-golang/projects/1)を確認して自分がアサインされている issue がないか確認
2. アサインされている issue の中から実装しようと決めた issue を「In Progress」に移動する
3. develop ブランチをチェックアウトして、`git pull`する
4. ブランチルールに従い、作業用のブランチを新規作成
5. issue 内容を満たすようにローカル環境で開発を進める
   1. commit はできるだけ作業内容が後から追ったときにわかりやすい単位で細かく切ること！
   2. コミットメッセージは日本語で良いので丁寧に書くこと！
6. 開発が完了したら、`develop`←`作業ブランチ`の PullRequest を作成する
7. Pull Request に`close #1`のように close コマンドを入力して、Pull Request をマージしたら紐づく issue が close されるようにしておく
8. Reviewers, Assignees, Labels の項目を issue と合わせて設定しておく
9. 「Files changed」で自分の書いたコードにバグなどがないか一通りチェックする
10. Reviewer に設定した人に Slack の`#notify-github`でレビュー依頼をする（レビュー依頼時は Pull Request のリンクも添付しましょう）
