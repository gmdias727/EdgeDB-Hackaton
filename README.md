# EdgeDB-Hackaton

### How to run this project?

install EdgeDB CLI globally: 
```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.edgedb.com | sh
```

healthcheck EdgeDB: 
```bash
edgedb --version
```

Clone the repo
```bash
git clone git@github.com:gmdias727/EdgeDB-Hackaton.git
```

Get inside the project directory
```bash
cd EdgeDB-Hackaton
```

how to install EdgeDB lib:
```bash
go install github.com/edgedb/edgedb-go/cmd/edgeql-go@latest
```



### Resources:
1. How to learn EdgeDB? [*here](https://docs.edgedb.com/tutorial).
2. How to setup environment variables? [Recommended](https://docs.edgedb.com/database/reference/projects#ref-guide-using-projects) or [also here](https://docs.edgedb.com/libraries/connection).

Note:
1. edgedb uses [edgeql](https://docs.edgedb.com/database/edgeql) as his main query language.
2. Consider installing EdgeDB extension for your favorite code editor, mine is [Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=magicstack.edgedb&ssr=false#review-details).