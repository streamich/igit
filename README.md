# igit

Interactive CLI for semantic Git commit messages.

```
igit c
```

![image](https://user-images.githubusercontent.com/9773803/87804358-0b437600-c854-11ea-8897-60c82dbf7484.png)


## Installation

Download from [releases page](https://github.com/streamich/igit/releases) or install using NPM.

```
npm i -g igit-cli
```

Or use without installation.

```
npx igit-cli
```


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
