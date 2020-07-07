# echo-unique
A utility to output unique strings passed into stdin, written in Golang

### Sample stdin

```
> cat sample_stdin.txt
staging.mbe.domain.com
domain.com
domain.co.uk
domain.ai
billing-api-staging.domain.com
mbe-staging.domain.com
mbe-staging.domain.com
staging.mbe.domain.com
duplicate
unique
duplicate
randomtext
random
```
### Usage
```
> cat sample_stdin.txt | echo-unique
staging.mbe.domain.com
domain.com
domain.co.uk
domain.ai
billing-api-staging.domain.com
mbe-staging.domain.com
duplicate
unique
randomtext
random
```

### Install with Homebrew

First install gomod package.
```
brew install filosottile/gomod/brew-gomod
brew gomod github.com/icheko/echo-unique
```
