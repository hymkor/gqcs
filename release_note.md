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
