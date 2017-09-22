# Example Blog site

This demo shows how to build a blogging site with Hofstadter.

Take a look around the design folders and files.

### Building

To generate the implementation:

```
make gen
```

To build the blog API server:

```
make blog
```

To build the tools: (database, data, cli)

```
make tools
```


### Running

First start the Postgres database:

```
./scripts/postgres.sh start
```

Next, migrate the tables: (only needed the first time the DB starts)

```
./blog-tool-db migrate-tables
``` 

Then, start the blog server:

```
./blog-server
```

