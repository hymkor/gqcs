![](GQuaccccccesS.png)

# GQuaccccccesS

- A spreadsheet-like editor for your database, right in the terminal.
- A standalone tool based on [SQL-Bless]&#x2019;s `edit` command.
- Unlike SQL-Bless, GQuaccccccesS applies all edits automatically when you finish editing.

![](demo.gif)

```console
Usage: gqcs [DRIVE-NAME] DATASOURCE
```

For example:

```console
$ gqcs sqlite3 :memory:
$ gqcs oracle://USERNAME:PASSWORD@HOSTNAME:PORT/SERVICE
$ gqcs postgres://USERNAME:PASSWORD@127.0.0.1:5555/DBNAME?sslmode=verify-full
$ gqcs sqlserver "Server=localhost\SQLEXPRESS;Database=master;Trusted_Connection=True;protocol=lpc"
$ gqcs mysql user:password@/database
```

- The title logo was generated using [GQtitle].
- The demo data comes from the [Japanese postal code CSV dataset][JapanPost], imported into SQLite3.
- I support Mobile Suit Gundam [GQuuuuuuX] and Microsoft [Access].

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/gqcs/releases) and extract the executable.

### via Scoop-installer (Windows)

```
scoop install https://raw.githubusercontent.com/hymkor/gqcs/master/gqcs.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install gqcs
```

### Installing via Go

```
go install github.com/hymkor/gqcs@latest
```

Author
------

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

[SQL-Bless]: https://github.com/hymkor/sqlbless
[GQtitle]: https://bills-appworks.github.io/GQtitle/
[JapanPost]: https://www.post.japanpost.jp/zipcode/download.html
[GQuuuuuuX]: https://www.gundam.info/feature/gquuuuuux/
[Access]: https://www.microsoft.com/ja-jp/microsoft-365/access
