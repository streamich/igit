# igit

Interactive CLI for semantic Git commit messages.


## Semantic branch names

```
<prefix>/<date>/<description>/<issue>
```

Where `<prefix>` is one of:

- `feature`
- `bug`
- `hotfix`
- `release`

`<date>` is in `YYYY-MM-DD` format. `<description>` is optional user provided description.
And `<issue>` is optional tracker issue number.


## Development

Run tests.

```
make test
```

Run your code.

```
ARGS="cz" make run
```

Set commit body editor.

```
EDITOR=nano make run
```
