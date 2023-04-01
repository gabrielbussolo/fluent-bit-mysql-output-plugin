# Fluent Bit MySQL Output Plugin

This is plugin for Fluent Bit, an open-source data collector that can be used to collect, process, and forward logs and metrics data. This plugin allows you to write data to a MySQL database. It's influenced by the [PostgreSQL output plugin for Fluent Bit](https://docs.fluentbit.io/manual/pipeline/outputs/postgresql) and the [MySQL Fluentd plugin](https://github.com/fluent/fluent-plugin-sql).

## Usage

Download the last release and add to your plugin directory, example:

```bash
curl -o /fluent-bit/etc/mysql-output-plugin.so https://github.com/gabrielbussolo/fluent-bit-mysql-output-plugin/releases/download/v0.1.0-alpha/mysql-output-plugin.so
```

Create a file to register the custom plugin (if you don't have one already):

`plugins.conf`:
```
[PLUGINS]
    Path /fluent-bit/etc/mysql-output-plugin.so
```

Add the configs to your `fluent-bit.conf`:

| Key      | Description                                 | Default        |
|----------|---------------------------------------------|----------------|
| Address  | The address of the MySQL server             | localhost:3306 |
| User     | The user to connect to the MySQL server     | root           |
| Password | The password to connect to the MySQL server | -              |
| Database | The database to connect to                  | -              |
| Table    | The table to connect to                     | -              |

example:
```
# import the plugin
[SERVICE]
    Flush           1
    Log_Level       info
    plugins_file    /fluent-bit/etc/plugins.conf

# configure the plugin    
[OUTPUT]
    Name mysql
    Address localhost:3306
    User root
    Password my-secret-pw
    Database fluent_bit
    Table test
```

Is expected that the table already exists in the database with the following fields `data`, `datetime` and `tag`. The `data` field is a `json` type, the `datetime` field is a `datetime` and `tag` a `varchar`.

Example:
```sql
create table fluent_bit.test
(
    id       int auto_increment
        primary key,
    data     json         null,
    datetime datetime     null,
    tag      varchar(100) null
);
```

## Next steps:
- [ ] Add support to create a default table if it doesn't exist
- [ ] Support async writes
- [ ] Add options for the connection pool
- [ ] Implement CI/CD