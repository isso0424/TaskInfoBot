# 使い方

## コマンド
### 課題の追加
実行方法  
`!task add <name> <subject> <limit>`  
|名称|内容|
|---|---|
|name|課題の名前|
|subject|教科の名前(config.jsonで指定したもの)|
|limit|課題の締め切り(デフォルトで次の日)|

limitは省略が可能

### 課題の表示
実行方法  
`!task list <subject>`
|名称|内容|
|---|---|
|subject|フィルターする教科の名前|

### 課題の削除
実行方法  
`!task remove <name>`  
|名称|内容|
|---|---|
|name|課題の名前|

### help
実行方法
`!task help`

## 設定
config.jsonで行います  
##### courses
配列として教科の情報を渡します  
教科の情報は以下のとおりです  
|名称|内容|型|
|---|---|---|
|alias|各学科の名称|string|
|subjects|学科ごとの教科一覧|Subjects|

Subjects
|名称|内容|型|
|---|---|---|
|major|主専攻の教科一覧|array|
|minor|副専攻の教科一覧|array|

##### channels
通知や登録をするチャンネルのIDを指定する
|名称|内容|型|
|---|---|---|
|regist|登録を行うチャンネルのID|string|
|notify|通知/課題の表示を行うチャンネルの一覧|Notify|

|名称|内容|型|
|---|---|---|
|major|主専攻の通知を行うチャンネルの一覧|array|
|minor|副専攻の通知を行うチャンネルの一覧|array|

