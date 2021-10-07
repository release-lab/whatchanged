Commit: feat: rename tag:N to @N [BREAKING CHANGE: rename

```diff
- tag:0~tag:1
+ @0~@1
```]
Commit message start ===
feat: rename tag:N to @N

BREAKING CHANGE: rename

```diff
- tag:0~tag:1
+ @0~@1
```

Commit message end ===
Commit: feat(cli): rename '--dir' to '--project' and '--file' to '--output' #4 [BREAKING CHANGE: new flags

```diff
- --dir=/path/to/dir
+ --project=/path/to/project
```

```diff
- --file=CHANGELOG.md
+ --output=CHANGELOG.md
```]
Commit message start ===
feat(cli): rename '--dir' to '--project' and '--file' to '--output' #4

BREAKING CHANGE: new flags

```diff
- --dir=/path/to/dir
+ --project=/path/to/project
```

```diff
- --file=CHANGELOG.md
+ --output=CHANGELOG.md
```

Commit message end ===
Commit: feat: rename 'changelog' to 'whatchanged' [BREAKING CHANGE: repo and binary rename to whatchanged

```diff
- $ changelog --help
+ $ whatchanged --help
```

Closes #6]
Commit message start ===
feat: rename 'changelog' to 'whatchanged'

BREAKING CHANGE: repo and binary rename to whatchanged

```diff
- $ changelog --help
+ $ whatchanged --help
```

Closes #6

Commit message end ===
Commit: chore(deps): pin dependency vite to 1.0.0-rc.13 (#9) [Co-authored-by: Renovate Bot <bot@renovateapp.com>]
Commit message start ===
chore(deps): pin dependency vite to 1.0.0-rc.13 (#9)

Co-authored-by: Renovate Bot <bot@renovateapp.com>
Commit message end ===
Commit: fix(deps): update dependency ant-design-vue to v2.0.0-rc.2 (#8) [Co-authored-by: Renovate Bot <bot@renovateapp.com>]
Commit message start ===
fix(deps): update dependency ant-design-vue to v2.0.0-rc.2 (#8)

Co-authored-by: Renovate Bot <bot@renovateapp.com>
Commit message end ===
Commit: chore(deps): pin dependencies (#7) [Co-authored-by: Renovate Bot <bot@renovateapp.com>]
Commit message start ===
chore(deps): pin dependencies (#7)

Co-authored-by: Renovate Bot <bot@renovateapp.com>
Commit message end ===
Commit: chore(deps): update module stretchr/testify to v1.6.1 (#3) [Co-authored-by: Renovate Bot <bot@renovateapp.com>]
Commit message start ===
chore(deps): update module stretchr/testify to v1.6.1 (#3)

Co-authored-by: Renovate Bot <bot@renovateapp.com>
Commit message end ===
Commit: chore(deps): update module pkg/errors to v0.9.1 (#2) [Co-authored-by: Renovate Bot <bot@renovateapp.com>]
Commit message start ===
chore(deps): update module pkg/errors to v0.9.1 (#2)

Co-authored-by: Renovate Bot <bot@renovateapp.com>
Commit message end ===
## v0.2.0 (2020-11-27)

### 🔥  New feature:

- refactor multiple version generation([`65c50c6`](https://github.com/release-lab/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf)) (@axetroy)
- unified usage when invalid options([`dde4ece`](https://github.com/release-lab/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec)) (@axetroy)
- print help/version information to stderr([`ddc712a`](https://github.com/release-lab/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9)) (@axetroy)
- version range of hash support 8-length short hash([`d9d3be8`](https://github.com/release-lab/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b)) (@axetroy)
- support to generate multiple specified versions ref: #4([`4a9e7ee`](https://github.com/release-lab/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec)) (@axetroy)
- rename tag:N to @N([`76eb777`](https://github.com/release-lab/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8)) (@axetroy)
- **cli**: rename '--dir' to '--project' and '--file' to '--output' #4([`189186a`](https://github.com/release-lab/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017)) (@axetroy)
- add whatchanged package for Go([`0bb75dd`](https://github.com/release-lab/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c)) (@axetroy)
- rename 'changelog' to 'whatchanged'([`ee18634`](https://github.com/release-lab/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d)) (@axetroy)
- implement revert parser([`2c800f2`](https://github.com/release-lab/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c)) (@axetroy)
- add chore block([`418b8f6`](https://github.com/release-lab/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f)) (@axetroy)
- add changelog repl. generate changelog online almose done.([`c7127b1`](https://github.com/release-lab/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c)) (@axetroy)
- enabled cors for changelog-remote([`8b5f7fb`](https://github.com/release-lab/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f)) (@axetroy)

### 🐛  Bugs fixed:

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

### ❤️ BREAKING CHANGES:

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

Closes #6
