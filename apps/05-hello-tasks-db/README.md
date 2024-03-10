## Debug

```sql
-- get last inserted row
SELECT * FROM ads ORDER BY scrape_time DESC LIMIT 1;
```


## References

- reference project https://github.com/charmbracelet/taskcli
- good info about sqlite3, how to start in go https://earthly.dev/blog/golang-sqlite/

## Sqlite3 commands

```bash
sqlite3 .conf/tasks.db
```

```sql
-- list tables
.tables
-- list table schema
.schema tasks
-- information about columns in table
PRAGMA table_info(tasks);
-- list all columns in table, limit 10
SELECT * FROM tasks LIMIT 10;
```

examples of insert, update, delete, select

```sql
-- insert
INSERT INTO tasks (Name, Project) VALUES ("TaskName", "TaskProject");
-- update
UPDATE tasks SET Name="TaskNameUpdated" WHERE ID=5;
-- delete
DELETE FROM tasks WHERE ID=5;
```

