# DBマイグレーション

## 調査時期

2025/1/3 now

## 前提

- GitサーバはGitHubのみ。
- 以下のデータベースに対応していること。
    - MySQL
    - PostgreSQL

## 定量的な比較表

`Star`および`Watch`はGitHub上のもの。

| 名称<br>(リンク)                                                   | 最終<br>更新日 | Star | Watch | 最新<br>Version | License | 開発元 | ドキュメント                                                         |
| :----------------------------------------------------------: |:------------------:| :----: |:-----:|:-------------:| :-----: | :----------------: | :-------------------------------------------------------------------: |
| [golang-migrate/migrate](https://github.com/golang-migrate/migrate) |     2024-09-24     | 15.8k  |  86   |    v4.18.1    | MIT     | コミュニティ        | [公式リポジトリ](https://github.com/golang-migrate/migrate/tree/master/GETTING_STARTED.md) |
| [pressly/goose](https://github.com/pressly/goose)            |     2024-12-20     | 7.4k   |  60   |    v3.24.0    | MIT     | Pressly           | [README](https://github.com/pressly/goose#readme)                   |
| [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)  |     2024-12-12     | 3.3k   |  36   |    v1.7.1     | MIT     | 個人/コミュニティ    | [README](https://github.com/rubenv/sql-migrate#readme)               |
| [amacneil/dbmate](https://github.com/amacneil/dbmate)        |     2024-12-20     | 5.6k   |  31   |    v2.24.2    | MIT     | 個人/コミュニティ    | [README](https://github.com/amacneil/dbmate#readme)                  |

### GitHub Star History

[startchart.cc](https://starchart.cc/)を使用。

#### migrate
![migrate](./GitHubStarHistory/migrate.png)

#### goose
![goose](./GitHubStarHistory/goose.png)

#### sql-migrate
![sql-migrate](./GitHubStarHistory/sql-migrate.png)

#### dbmate
![dbmate](./GitHubStarHistory/dbmate.png)

## 定性的な比較表

| 観点                                             |                         golang-migrate<br>/migrate                         |                        pressly<br>/goose                         |                rubenv<br>/sql-migrate                |                          amacneil<br>/dbmate                           |
| :------------------------------------------: |:--------------------------------------------------------------------------:|:----------------------------------------------------------------:|:----------------------------------------------------:|:----------------------------------------------------------------------:|
| **人気度／コミュニティ**                       |                        非常に活発。IssueやPRも多く、エコシステムが充実                         |                       比較的人気。シンプルな仕組みで使いやすい                       |              コミュニティ規模は中程度。Issueはそこそこ活発               |                      徐々にユーザが増加中。Docker環境などで使いやすい                       |
| **学習コスト**                                 |                            基本的なGoの知識があればすぐ利用可能                             |                 直感的なCLIと設定。マイグレーションファイル形式も分かりやすい                 |               SQLベースのマイグレーションで学習コストは低め               |                   基本的には簡単に使えるが、特定の要件に合わせた拡張時にやや工夫が必要                   |
| **パフォーマンス考慮**                         |                  大規模システムでの採用例も多く、マルチDB・並列実行など幅広いケースをサポート                   |                大量のマイグレーションにも対応可能だが、運用時に工夫が必要な場合あり                |        設定次第で性能面の調整が可能。Go以外のプロジェクトからの移行にも活用可能         |                      テーブルロックなど利用DBに応じて挙動が異なる点に注意                       |
| **マイグレーション方式**                       |               SQLファイルに加えてGoコードベースにも対応可能（Driverによっては複数方式を選択）                |              基本的にはSQLファイルをベースとしつつ、Goコードから呼び出す形にも対応               |       SQLファイルをベースにマイグレーションを行う。DB依存ロジックにも対応しやすい       |               SQLファイルを利用し、設定が比較的シンプル。DockerやCI環境での運用を想定                |
| **CLI機能／ツールの充実度**                    |             `migrate` CLIが充実。バージョン管理からロールバック、ステップ実行など多彩なオプション              |        `goose` CLIがわかりやすく、マイグレーション管理と簡単な初期化・状態確認などの操作性が良い        |    `sql-migrate` CLIでの操作が一通りそろっている。外部ツールとも連携しやすい     |               `dbmate` CLIが用意され、Dockerなどにも相性が良い。操作もシンプル                |
| **トランザクション対応／ロールバック**           |           多くのDBでのトランザクション対応を提供。ロールバック機能も充実。ただしDBによってサポート状況の差異はある           |        マイグレーション実行時にトランザクションを使用可能だが、特定のSQL(DDL)の扱いには注意が必要         |  トランザクション単位でのマイグレーション実行が可能。ロールバック定義もSQLファイルに書けば対応可能  |             DB種類によってはロックを取得する方法などに依存。ロールバックは手動で作成が必要なケースあり             |
| **拡張性／カスタマイズ性**                     |       ドライバやプラグインの追加、マルチ環境へのデプロイなど柔軟に対応可能。Goコードと組み合わせることで高度なカスタマイズが可能       | シンプルな構成のため、独自ロジックを組み込みたい場合は追加実装が必要となることも。ただしGoのコードから直接制御する運用にも合う |      SQL中心なので拡張は比較的しやすいが、Goコード内での高度な制御はやや工夫が必要       | マイグレーション周りの拡張は比較的容易。DockerやCI/CDに組み込みやすい構成だが、独自要件には追加スクリプトなどで対応する必要がある |
| **ユースケース**                               |                    多数のDBをサポートしており、マイクロサービスなど大規模環境でも広く利用                    |                  迅速に導入したい中小規模プロジェクトやPoCなどに向いている                  |         SQLを中心に扱うプロジェクトや、シンプルなマイグレーションを望むケース         |                    コンテナ化した環境で手軽にマイグレーション処理を組み込みたい場合                    |
