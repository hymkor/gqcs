- Changed log output to JSON format
- Stopped enclosing table names in double quotes in SQL when not necessary.

v0.1.0
======
Oct 14, 2025

- Update based SQL-Bless to v0.23.0
    - Unified the exit operation to the `ESC` key
        - `ESC` + `y`: Apply changes and exit
        - `ESC` + `n`: Discard changes and exit
        - `c`: Still supported but deprecated
        - `q`: Now equivalent to `ESC`
    - Changed the brackets around the table name display from `【】` to `[]`
- Record timestamp and parameters on the debug log now
- Update csvi package to v1.15.0 + snapshot:
    - Added key bindings `]` and `[` to adjust the width of the current column (widen and narrow, respectively).
    - Added `-rv` option to prevent unnatural colors on terminals with a white background
    - At startup, the width of ambiguous-width Unicode characters was being measured, but on terminals that do not support the cursor position query sequence `ESC[6n`, this could cause a hang followed by an error. To address this:
        - If `ESC[6n` is not supported, the program now continues without aborting.
        - Skipped the measurement of ambiguous-width Unicode characters when the environment variable `RUNEWIDTH_EASTASIAN` is defined.
    - Suppress color output if the `NO_COLOR` environment variable is set (following https://no-color.org/ )
    - When the environment variable `COLORFGBG` is defined in the form `(FG);(BG)` and `(FG)` is less than `(BG)`, the program now uses color settings designed for light backgrounds (equivalent to `-rv`).

v0.0.3
======
Sep 28, 2025

- Update based SQL-Bless to v0.21.0
    - Changed to use placeholders for value specification
    - Modified SQLite3 datetime column updates to normalize values in `WHERE` clauses according to column type:
        - `DATETIME` / `TIMESTAMP` columns → `datetime()`
        - `DATE` columns → `date()`
        - `TIME` columns → `time()`
        - This ensures updates work regardless of whether ISO8601 strings contain `T` or `Z`
- Display an error message if the number of updated rows is not the expected single row.
- When an error occurs during update, return to table selection mode instead of terminating immediately.

v0.0.2
======
Sep 15, 2025

- Fix: Escape sequence was not enabled on Windows.

v0.0.1
======
Sep 14, 2025

- Prototype
