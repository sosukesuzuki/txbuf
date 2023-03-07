# txbuf

## Overview

txbuf は CLI ベースのメモアプリケーションです。
`~/.txbuf/` ディレクトリに、`20230301-1.txt` のような名前のファイルが作成されていきます。
ファイルの編集は Vim で行われます。

## Installation

```
go install github.com/sosukesuzuki/txbuf@latest
```

## Requirements

- vim
- peco
- bash

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
