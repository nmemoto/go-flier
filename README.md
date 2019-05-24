# go-flier
https://www.flierinc.com/ を操作する

## login
ログインを実行し、任意のディレクトリにCookie等を保存することによって、ログインが必要な操作を実行できるようにする

### 事前準備
以下のようにログインに必要な情報を設定する
```
export FLIER_USER=test@example.com
export FLIER_PASS=hogehero
```

### 使い方

#### 実行例
```
# go run cmd/login/main.go -headless=false
```

#### 引数
- -user-data-dir: ログイン時の情報を保存するディレクトリの指定(デフォルトは./tmp)
- -headless: headlessで実行するか否か(デフォルトはtrue)


## max-page-num
[要約リスト](https://www.flierinc.com/summary/list) の総ページ数を標準出力する

### 使い方
#### 実行例
```
# go run cmd/max-page-num/main.go
74
```

## id-list
[要約リスト](https://www.flierinc.com/summary/list) に含まれる要約ページへのリンクに使用されているID, 要約本のタイトル, 著者名(一人のみ), 要約ページへのリンクのリストを標準出力する

### 使い方

#### 実行例
```
# go run cmd/id-list/main.go -page=2 -headless=false 
1861,"アナログの逆襲","デイビッド・サックス","https://www.flierinc.com/summary/summary/1861"
1851,"無敵の筋トレ食","岡田隆","https://www.flierinc.com/summary/summary/1851"
1855,"ささいなことに動揺してしまう敏感すぎる人の「仕事の不安」がなくなる本","みさきじゅり","https://www.flierinc.com/summary/summary/1855"
1854,"しょぼい起業で生きていく","えらいてんちょう","https://www.flierinc.com/summary/summary/1854"
1860,"勉強大全","伊沢拓司","https://www.flierinc.com/summary/summary/1860"
1859,"サンデル教授、中国哲学に出会う","マイケル・サンデル","https://www.flierinc.com/summary/summary/1859"
1858,"The San Francisco Fallacy","ジョナサン・シーゲル","https://www.flierinc.com/summary/summary/1858"
1853,"勝間式超コントロール思考","勝間和代","https://www.flierinc.com/summary/summary/1853"
1857,"人生は攻略できる","橘玲","https://www.flierinc.com/summary/summary/1857"
1856,"ことばの「なまり」が強みになる!","吉村誠","https://www.flierinc.com/summary/summary/1856"
1848,"居酒屋へ行こう。","太田和彦","https://www.flierinc.com/summary/summary/1848"
1852,"直感と論理をつなぐ思考法","佐宗邦威","https://www.flierinc.com/summary/summary/1852"
1850,"即動力","田村淳","https://www.flierinc.com/summary/summary/1850"
1846,"「金融パーソン」はどう生きるか","窪田泰彦","https://www.flierinc.com/summary/summary/1846"
1843,"Move Fast and Break Things","Jonathan Taplin","https://www.flierinc.com/summary/summary/1843"
1849,"管理ゼロで成果はあがる","倉貫義人","https://www.flierinc.com/summary/summary/1849"
1845,"東京大田区・弁当屋のすごい経営","菅原勇一郎","https://www.flierinc.com/summary/summary/1845"
1842,"平成最後のアニメ論","町口哲生","https://www.flierinc.com/summary/summary/1842"
1844,"東京格差","中川寛子","https://www.flierinc.com/summary/summary/1844"
1847,"経営戦略としての異文化適応力","宮森千嘉子","https://www.flierinc.com/summary/summary/1847"
1841,"OODA LOOP(ウーダループ)","チェット・リチャーズ","https://www.flierinc.com/summary/summary/1841"
1840,"伝達の整理学","外山滋比古","https://www.flierinc.com/summary/summary/1840"
1832,"丹羽宇一郎\u3000習近平の大問題","丹羽宇一郎","https://www.flierinc.com/summary/summary/1832"
1838,"鎌倉資本主義","柳澤大輔","https://www.flierinc.com/summary/summary/1838"
```

```
for i in $(seq 73)
do
    go run cmd/id-list/main.go -page $i >> idlist
done
```

#### 引数
- -page: 要約リストで使用されているページ数のクエリを指定する(デフォルトは1)
    - https://www.flierinc.com/summary/list?page=3 のIDのリストを取得したい場合、3を指定する
- -headless: headlessで実行するか否か(デフォルトはtrue)


## pdf-print-page-transition
[要約リスト](https://www.flierinc.com/summary/list) に含まれる要約ページのPDFダウンロード画面に遷移する

### 使い方

#### 引数
- -user-data-dir: ログイン時の情報を保存するディレクトリの指定(デフォルトは./tmp)
- -id: 要約ページへのリンクに使用されているID(デフォルトは1)
    - https://www.flierinc.com/summary/10 の要約ページを取得したい場合、10を指定する

#### 実行例
```
$ go run cmd/pdf-print-page-transition/main.go -id=1825 
```