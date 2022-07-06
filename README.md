## Description

**Polygo** is a polynomials solver written in [golang](https://golang.org)

## Status - `WIP`

- [x] parse simple equations and convert them back to string
  supported syntax is:
    - [x] `+`/`-`/`=`
    - [x] letter for unknown (`x` expected)
    - [ ] braces `(`/`)`
- [ ] solve equations
    - [ ] horner's method (for all kinds of equations)
    - [ ] *optionally* solving particular types of equations using diffrent methods for better performence
- [ ] GUI - probably write some gui using [giu](https://github.com/AllenDang/giu) framework.
