# Grafana v12 Dashboards

## pgbackrest_dashboard.json

### Description
This dashboard allows to see how old a daily incremental backup is. It is a simple example dashboard for using the "Backup_age_pbackrest" metrics.

There are two axes:
- y-axis --> contains the database name of the backup
- x-axis --> contains the age of the backup

The bars are coloured accordingly:
- green: The backup doesn't need to be renewed and it is not overdue (age: 0 hours < 24 hours).
- yellow: This is the time when the backup should normally be renewed (age: 24 hours < 25 hours).
- red: The backup is older than 25 hours, something went wrong when trying to take a new backup. Check it out or data can be lost!

### Required metrics
For the dashboard you need the metric "backup_age_pgbackrest".
You can find it on the main repo https://github.com/cybertec-postgresql/pgwatch/tree/master in the file /internal/metrics/metrics.yaml.

### Required extensions
The "backup_age_pgbackrest" metrics needs the Postgres extension "plpython3u" to work.
