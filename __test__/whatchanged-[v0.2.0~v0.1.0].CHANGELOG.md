## v0.2.0 (2020-11-27)

### New feature:

- refactor multiple version generation([`65c50c6`](https://github.com/release-lab/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf)) (@axetroy)
- unified usage when invalid options([`dde4ece`](https://github.com/release-lab/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec)) (@axetroy)
- print help/version information to stderr([`ddc712a`](https://github.com/release-lab/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9)) (@axetroy)
- version range of hash support 8-length short hash([`d9d3be8`](https://github.com/release-lab/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b)) (@axetroy)
- support to generate multiple specified versions ref: #4([`4a9e7ee`](https://github.com/release-lab/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec)) (@axetroy)
- rename tag:N to @N([`76eb777`](https://github.com/release-lab/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8)) (@axetroy)
- **cli**: rename '--dir' to '--project' and '--file' to '--output' #4([`189186a`](https://github.com/release-lab/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017)) (@axetroy)
- add whatchanged package for Go([`0bb75dd`](https://github.com/release-lab/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c)) (@axetroy)
- rename 'changelog' to 'whatchanged'([`ee18634`](https://github.com/release-lab/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d)) (@axetroy), Closes: #6
- implement revert parser([`2c800f2`](https://github.com/release-lab/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c)) (@axetroy)
- add chore block([`418b8f6`](https://github.com/release-lab/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f)) (@axetroy)
- add changelog repl. generate changelog online almose done.([`c7127b1`](https://github.com/release-lab/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c)) (@axetroy)
- enabled cors for changelog-remote([`8b5f7fb`](https://github.com/release-lab/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f)) (@axetroy)

### Bugs fixed:

- the breaking change block is generate in incorrect format([`5bf4bee`](https://github.com/release-lab/whatchanged/commit/5bf4beea7124cac872598c5487657548e7a826c9)) (@axetroy)
- the incorrect range generated when there are no parameters([`bbf648e`](https://github.com/release-lab/whatchanged/commit/bbf648e756ab74abb25764ee9ead032343832b3b)) (@axetroy)
- parse single version tag:N incorrect([`aa62d6b`](https://github.com/release-lab/whatchanged/commit/aa62d6be3294619c81159a39208c9f7bba07630f)) (@axetroy)
- **repl**: server not response markdown([`45b6ad2`](https://github.com/release-lab/whatchanged/commit/45b6ad20ec4a50dc7661bf575fa408ef6383c46b)) (@axetroy)
- **repl**: generate error([`8a92fd7`](https://github.com/release-lab/whatchanged/commit/8a92fd7693568683beba2431b0e0659fc99e3c82)) (@axetroy)
- **repl**: no-cors mode([`6410c1b`](https://github.com/release-lab/whatchanged/commit/6410c1be6cc35e3165172f99738add18ef4d5beb)) (@axetroy)
- **repl**: source block do not render after request success([`d4af36b`](https://github.com/release-lab/whatchanged/commit/d4af36be80ca60f4bbbcb96603b070883ac44a6a)) (@axetroy)
- **repl**: improve error handler([`b2bbbbd`](https://github.com/release-lab/whatchanged/commit/b2bbbbd7608501813986d74f6e44c233719246eb)) (@axetroy)
- **repl**: markdown is not rendered correctly([`4cf9f6e`](https://github.com/release-lab/whatchanged/commit/4cf9f6ee53d19f67380144030a38ede88cb1a59b)) (@axetroy)
- **deps**: update dependency ant-design-vue to v2.0.0-rc.2 (#8)([`a340efc`](https://github.com/release-lab/whatchanged/commit/a340efc8b86b1728eb1dcaedc9c101767582e811)) (@renovate[bot])
- **repl**: production assets should be set publicPath([`89ea856`](https://github.com/release-lab/whatchanged/commit/89ea856f4f2046f7347a5ebd2c9d60e3a3650595)) (@axetroy)
- **repl**: markdown render style([`acebfd0`](https://github.com/release-lab/whatchanged/commit/acebfd0bd736dac9c811186c82ba241d7b1e05e1)) (@axetroy)

### BREAKING CHANGES:

- rename

```diff
- tag:0~tag:1
+ @0~@1
```

- new flags

```diff
- --dir=/path/to/dir
+ --project=/path/to/project
```

```diff
- --file=CHANGELOG.md
+ --output=CHANGELOG.md
```

- repo and binary rename to whatchanged

```diff
- $ changelog --help
+ $ whatchanged --help
```

## v0.1.0 (2020-11-24)

### New feature:

- link commit for generation([`b9432db`](https://github.com/release-lab/whatchanged/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb)) (@axetroy)
- add full preset template([`7553570`](https://github.com/release-lab/whatchanged/commit/7553570590b571bd33e10a4f80ec5639d0613042)) (@axetroy)
- support changelog for git submodule([`ec6a957`](https://github.com/release-lab/whatchanged/commit/ec6a957752fbca9faa261d8694826779e2cbec1f)) (@axetroy)
- add writer step([`441ad13`](https://github.com/release-lab/whatchanged/commit/441ad1322b1fecaca89a170ecebaf2955a77d630)) (@axetroy)
- add formatter for markdown output([`8c177e0`](https://github.com/release-lab/whatchanged/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34)) (@axetroy)
- add --file flag to generate file([`0e4fb09`](https://github.com/release-lab/whatchanged/commit/0e4fb09789732fec5b09b247e208d61794c3da0d)) (@axetroy)
- support custom tmeplate and preset([`3aa0aee`](https://github.com/release-lab/whatchanged/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db)) (@axetroy)
- support tag ranges([`3d14c9c`](https://github.com/release-lab/whatchanged/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9)) (@axetroy)
- support version range. eg v2.0.0~v1.0.0([`a65a4a8`](https://github.com/release-lab/whatchanged/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3)) (@axetroy)
- parse commit message and generate to template([`7f67e78`](https://github.com/release-lab/whatchanged/commit/7f67e783926fed647d2ad5414f31448eea106fc3)) (@axetroy)

### Bugs fixed:

- improve ssh git url parser([`9993ee6`](https://github.com/release-lab/whatchanged/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f)) (@axetroy)
- commit range not include commit of tag([`f8df0ba`](https://github.com/release-lab/whatchanged/commit/f8df0ba654c8faf67eccf98262cd55807e53e597)) (@axetroy)
- unescape template([`e118cbf`](https://github.com/release-lab/whatchanged/commit/e118cbfafd201b945848f15303fdb261e251f058)) (@axetroy)
- if empty argument for command line([`9c79fd9`](https://github.com/release-lab/whatchanged/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd)) (@axetroy)
- **ci**: remove unsued code([`66bcf8f`](https://github.com/release-lab/whatchanged/commit/66bcf8f43db85409e0392c93f2e347ed91699e81)) (@axetroy)
