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

## Semantic commits

Read about semantic commit messages [here](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716).


## Semantic branches

```
igit b
```

![image](https://user-images.githubusercontent.com/9773803/87811434-fa4c3200-c85e-11ea-936b-249426df0253.png)



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
