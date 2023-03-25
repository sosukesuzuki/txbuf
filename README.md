# txbuf

## Overview

txbuf は CLI ベースのメモアプリケーションです。
`~/.txbuf/` ディレクトリに、`20230301-1.txt` のような名前のファイルが作成されていきます。
ファイルの編集は Vim で行われます。
一部のコマンドは `~/.txbuf` ディレクトリが Git で管理されていることを前提とします。そのようなコマンドを使う場合には、あらかじめ `~/.txbuf` ディレクトリ上で `git init` をしておいてください。

## Installation

```
go install github.com/sosukesuzuki/txbuf@latest
```

## Requirements

- vim
- peco
- bash
- git

## Commands

### `txbuf`

`~/.txbuf/` に存在する最新のファイルを Vim で開きます。
まだファイルが存在しない場合は、`20230301-1.txt` のような名前のファイルを新しく作成し、Vim で開きます。

### `txbuf new`

`20230301-1.txt` のような名前の新しくファイルを作成し、Vim で開きます。
同じ日にすでにファイルを作成していた場合は、`20230301-2.txt` のような名前のファイルを新しく作成し、Vim で開きます。

### `txbuf query`

`~/.txbuf/` に存在するファイルに対してインクリメンタルサーチを行い、選択されたファイルを Vim で開きます。
インクリメンタルサーチのために peco を使うので、事前に peco がインストールされている必要があります。

### `txbuf clone`

最新のファイルをコピーして、`20230301-2.txt` のような名前のファイルを新しく作成します。最新のファイルが存在しない場合は、何もしません。

### `txbuf git`

`~/.txbuf/` ディレクトリ上で `git` コマンドを実行します。たとえば `txbuf git push origin HEAD` や `txbuf git pull` のように使います。

### `txbuf push`

`~/.txbuf` ディレクトリ上で `git add . && git commit -m 'Update' && git push origin HEAD` を実行します。
