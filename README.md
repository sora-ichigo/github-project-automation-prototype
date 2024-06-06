# github-project-automation [prototype]


`POST: /run` を Cloud Run & Cloud Scheduler 定期実行している。ジョブでは GitHub Projects のアイテム追加やステータス更新が定期的に実行される。

GCP リソースの定義は別リポで terraform 管理しているのでここにはない。

