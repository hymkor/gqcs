package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-colorable"

	"github.com/nyaosorg/go-box/v2"

	"github.com/hymkor/sqlbless/dialect"
	_ "github.com/hymkor/sqlbless/dialect/mysql"
	_ "github.com/hymkor/sqlbless/dialect/oracle"
	_ "github.com/hymkor/sqlbless/dialect/postgresql"
	_ "github.com/hymkor/sqlbless/dialect/sqlite"
	_ "github.com/hymkor/sqlbless/dialect/sqlserver"
	"github.com/hymkor/sqlbless/spread"
)

var flagDebugLog = flag.String("D", os.DevNull, "file to write debug logs to")

func scanAllStrings(rows *sql.Rows, n int) ([]sql.NullString, error) {
	refs := make([]any, n)
	data := make([]sql.NullString, n)
	for i := 0; i < n; i++ {
		refs[i] = &data[i]
	}
	if err := rows.Scan(refs...); err != nil {
		return nil, err
	}
	return data, nil
}

func findColumn(target string, list []string) int {
	for i, x := range list {
		if strings.EqualFold(x, target) {
			return i
		}
	}
	return -1
}

func listTable(ctx context.Context, d *dialect.Entry, conn *sql.DB) ([]string, error) {
	rows, err := conn.QueryContext(ctx, d.SqlForTab)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	tableIndex := findColumn(d.TableField, columns)
	if tableIndex < 0 {
		return nil, fmt.Errorf("Application error: table field '%s' not found", d.TableField)
	}
	var tables []string
	for rows.Next() {
		data, err := scanAllStrings(rows, len(columns))
		if err != nil {
			return nil, err
		}
		if nameColumn := data[tableIndex]; nameColumn.Valid {
			tables = append(tables, nameColumn.String)
		}
	}
	return tables, nil
}

func mains(args []string) (lastErr error) {
	disabler := colorable.EnableColorsStdout(nil)
	defer disabler()
	terminal := colorable.NewColorableStdout()

	dbg, err := os.Create(*flagDebugLog)
	if err != nil {
		return err
	}
	defer func() {
		if lastErr != nil {
			fmt.Fprintln(dbg, lastErr.Error())
		}
	}()
	d, err := dialect.ReadDBInfoFromArgs(args)
	if err != nil {
		return fmt.Errorf("Usage: %s {DRIVER} DATASOURCE\n%w", os.Args[0], err)
	}
	conn, err := sql.Open(d.Driver, d.DataSource)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err = conn.Ping(); err != nil {
		return err
	}

	ctx := context.Background()
	tables, err := listTable(ctx, d.Dialect, conn)
	if err != nil {
		return err
	}
	var tx *sql.Tx
	editor := &spread.Editor{
		Viewer: &spread.Viewer{
			HeaderLines: 1,
			Comma:       ',',
			Null:        "\u2400",
		},
		Entry: d.Dialect,
		Query: conn.QueryContext,
		Exec: func(ctx context.Context, sql string, args ...any) (sql.Result, error) {
			if tx == nil {
				var err error
				if tx, err = conn.BeginTx(ctx, nil); err != nil {
					return nil, fmt.Errorf("conn.BeginTx: %w", err)
				}
			}
			fmt.Fprintln(dbg, sql)
			result, err := tx.ExecContext(ctx, sql, args...)
			if err != nil {
				fmt.Fprintln(dbg, err.Error())
				err = fmt.Errorf("%s\n%w", sql, err)
			}
			return result, err
		},
	}
	for {
		fmt.Fprintln(terminal, "Select a table:")
		table, err := box.SelectStringContext(ctx, tables, false, terminal)
		fmt.Println()
		if err != nil {
			return err
		}
		if len(table) < 1 {
			return nil
		}
		err = editor.Edit(ctx, `"`+table[0]+`"`, terminal)
		if tx == nil {
			continue
		}
		if err != nil {
			fmt.Fprintln(terminal, "Transaction rolled back.")
			fmt.Fprintln(dbg, "Transaction rolled back.")
			tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		tx = nil
	}
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
