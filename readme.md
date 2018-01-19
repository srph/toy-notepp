# Note++
A toy note-taking web app in Go and React.

## Requirements
- Go v1.8
- Node.js 8 & npm 5
- MySQL 5.6

## Installing
If you'd like to fiddle or run the code, feel free to do so.

- Install the dependencies
```bash
make install
```
- Setup environment config. After running the command below, properly assign the values to your respective configuration.
```bash
cp .env.example .env
```
- Migrate the database
```bash
make db.migrate
make db.seed
```
- Running the application
```bash
make start
```

Finally, you're good to go.