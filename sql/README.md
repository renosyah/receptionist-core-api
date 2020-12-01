## Dump data 

```

pg_dump -U postgres -h localhost -p 5432 --data-only receptionist_db > data.sql

```


## Import data


```

psql receptionist_db < data.sql

```
