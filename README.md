# new

*new* is an opinionated project generator for Go.

Run it from the command line to generate a new project structure. It currently takes the module name as a parameter, and `-l {license name}` as an optional flag to specify which license to use. If `-l {license name}` is omitted, it defaults to the [MIT license](https://opensource.org/licenses/MIT). 

Licenses are loaded from the [GitHub License API](https://developer.github.com/v3/licenses/), and the complete list of licenses can be obtained by hitting the url at [https://api.github.com/licenses](https://api.github.com/licenses). *new* uses the 'key' field from the list of licenses to identify a license.

To install *new*, run the following, which will install *new* as a binary

```bash
go get github.com/tophatsteve/new
```

### Example

Run the following to create a new project called *mynewproject*

```bash
new mynewproject -l gpl-3.0
```

It will generate a the following folder structure for the project, and will also initialise a git repository.

```
mynewproject
│   README.md
│   LICENSE
│   .gitignore
│   .travis.yml   
│   Gopkg.toml
│   mynewproject.go
│   mynewproject_test.go
│
└───cmd
│   │   main.go
│   │   main_test.go
```

### Future Plans

This is very much an alpha (or even pre-alpha) project, and I will be adding more facilities, such as the ability to base the project layout on a git repository, in the future.

