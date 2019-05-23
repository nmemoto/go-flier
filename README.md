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
[要約リスト](https://www.flierinc.com/summary/list) に含まれる要約ページへのリンクに使用されているIDのリストを標準出力する

### 使い方

#### 実行例
```
# go run cmd/id-list/main.go -page=3 -headless=false 
1818
1821
1820
1819
1816
1810
1817
1815
1814
1813
1812
1809
1805
1811
1808
1804
1803
1807
1802
1801
1806
1796
1800
1799
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