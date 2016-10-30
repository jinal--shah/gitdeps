# gitdeps

[1]: https://github.com/hashicorp/go-getter "Hashicorp go-getter on github"
[2]: https://github.com/toml-lang/toml "toml on github"
[3]: https://github.com/jinal--shah/go-get-gitdeps/issues "go-get-gitdeps issues"

**WARNING: not yet usable ...**

_... buildable binary to recursively grab sources from git repos by branch or tag_

* Create a .gitdeps toml file with repo information in your current working directory.

* Run the gitdeps binary (or `go run gitdeps.go`) and it will fetch those sources from git.

* If any of the fetched sources contain a .gitdeps file, it'll fetch those too.

There is no circular dependency checking, just a default maximum recursion level of 3.

i.e. gitdeps can fetch repos in the current dir's .gitdeps (level 1) and fetch any gitdeps
listed in those repos (level 2). It can continue to fetch any gitdeps in the level 2 repos
(level 3) but after that it will fail.

## BUILD

This should create you a binary that works on any modern linux, 
though note the shelled out commands that you also need
available in your path ...

```bash

# ... assuming you've installed golang 1.6+

export GOBIN=/usr/local/bin # will install gitdeps binary to $GOBIN
export GOPATH=$HOME/work # amend to your workspace
[ ! -e $GOPATH ] && mkdir $GOPATH; cd $GOPATH

go get github.com/jinal--shah/gitdeps
go install github.com/jinal--shah/gitdeps/cmd/gitdeps.go

# Now go forth and pepper your git repos with .gitdeps files ...

```

## RUN THE BINARY

You will need the following cmds installed and available in your PATH:

* git
* grep
* awk

Build as above and just run `gitdeps` from the dir containing your
.gitdeps file. 

### ... or RUN WITHOUT BUILDING ...
Alternatively, install golang and create a GOPATH as above then:

```bash

go get github.com/jinal--shah/gitdeps
go run github.com/jinal--shah/gitdeps/cmd/gitdeps.go

```

## .gitdeps file

* [TOML] [2] format (cheers Tom, Open source thanks you)

* but all values must be quoted 

* opening section header must be: `[deps]`

* repo section headers must be `[deps.some_dir_for_clone]`

### .gitdeps repo section

```toml
[deps]
    [deps.<dir to clone in to>]
    src   = "<mandatory, any uri accepted by git clone cmd>"
    ref   = "<mandatory, any value accepted by git --branch option>"
    depth = "<optional, any value accepted by git --depth option>"
```

### example: .gitdeps file

```toml
[deps]
    [deps.build_alpine]
    src = "git@github.com:jinal--shah/build_ami"
    ref = "master"
    depth = "1"

    [deps.coreos]
    src = "https://github.com/jinal--shah/packer_includes"
    ref = "2.1.0"
```

## Why?

### tl;dr

I wanted an easy way to pull and assemble multiple versioned git sources required to
build a project.

### ... which is to say ...

Not every framework or tool supports modules out-of-the-box.

I have boilerplate makefile targets and bundles of shell scripts in tagged
git repos that I want available when building a project.

But I don't need to build and version packages of these flat file assets and stick
them in an asset repo when I've already version-tagged the git sources.

_Seriously, I just want to tie versions of those sources to a version of the project_
_I want to build ..._

Something like the mercurial Hashimoto's [go-getter] [1], but 
that could be built as a binary without bindings to os-specific c libs.

### ... um, why not just use a shell script?

Tried it. The shell script started off as a couple of for-loops but became quite
the monstrosity after adding config file parsing, error checking and validation.

It also ended up wanting some of those versioned reusable bundles of shell scripts sitting in
a git repo. _Chicken, meet egg._

But hey, if you find a way, please do get in touch.

## TODO

See [Issues] [3]

