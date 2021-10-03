v0.2.0 (2020-11-27)
-------------------

### 🔥 New feature:

-	refactor multiple version generation([`65c50c6`](https://github.com/whatchanged-community/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf)) (@axetroy)
-	unified usage when invalid options([`dde4ece`](https://github.com/whatchanged-community/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec)) (@axetroy)
-	print help/version information to stderr([`ddc712a`](https://github.com/whatchanged-community/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9)) (@axetroy)
-	version range of hash support 8-length short hash([`d9d3be8`](https://github.com/whatchanged-community/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b)) (@axetroy)
-	support to generate multiple specified versions ref: #4([`4a9e7ee`](https://github.com/whatchanged-community/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec)) (@axetroy)
-	rename tag:N to @N([`76eb777`](https://github.com/whatchanged-community/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8)) (@axetroy)
-	**cli**: rename '--dir' to '--project' and '--file' to '--output' #4([`189186a`](https://github.com/whatchanged-community/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017)) (@axetroy)
-	add whatchanged package for Go([`0bb75dd`](https://github.com/whatchanged-community/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c)) (@axetroy)
-	rename 'changelog' to 'whatchanged'([`ee18634`](https://github.com/whatchanged-community/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d)) (@axetroy) , Closes: #6
-	implement revert parser([`2c800f2`](https://github.com/whatchanged-community/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c)) (@axetroy)
-	add chore block([`418b8f6`](https://github.com/whatchanged-community/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f)) (@axetroy)
-	add changelog repl. generate changelog online almose done.([`c7127b1`](https://github.com/whatchanged-community/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c)) (@axetroy)
-	enabled cors for changelog-remote([`8b5f7fb`](https://github.com/whatchanged-community/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f)) (@axetroy)

### 🐛 Bugs fixed:

-	the breaking change block is generate in incorrect format([`5bf4bee`](https://github.com/whatchanged-community/whatchanged/commit/5bf4beea7124cac872598c5487657548e7a826c9)) (@axetroy)
-	the incorrect range generated when there are no parameters([`bbf648e`](https://github.com/whatchanged-community/whatchanged/commit/bbf648e756ab74abb25764ee9ead032343832b3b)) (@axetroy)
-	parse single version tag:N incorrect([`aa62d6b`](https://github.com/whatchanged-community/whatchanged/commit/aa62d6be3294619c81159a39208c9f7bba07630f)) (@axetroy)
-	**repl**: server not response markdown([`45b6ad2`](https://github.com/whatchanged-community/whatchanged/commit/45b6ad20ec4a50dc7661bf575fa408ef6383c46b)) (@axetroy)
-	**repl**: generate error([`8a92fd7`](https://github.com/whatchanged-community/whatchanged/commit/8a92fd7693568683beba2431b0e0659fc99e3c82)) (@axetroy)
-	**repl**: no-cors mode([`6410c1b`](https://github.com/whatchanged-community/whatchanged/commit/6410c1be6cc35e3165172f99738add18ef4d5beb)) (@axetroy)
-	**repl**: source block do not render after request success([`d4af36b`](https://github.com/whatchanged-community/whatchanged/commit/d4af36be80ca60f4bbbcb96603b070883ac44a6a)) (@axetroy)
-	**repl**: improve error handler([`b2bbbbd`](https://github.com/whatchanged-community/whatchanged/commit/b2bbbbd7608501813986d74f6e44c233719246eb)) (@axetroy)
-	**repl**: markdown is not rendered correctly([`4cf9f6e`](https://github.com/whatchanged-community/whatchanged/commit/4cf9f6ee53d19f67380144030a38ede88cb1a59b)) (@axetroy)
-	**deps**: update dependency ant-design-vue to v2.0.0-rc.2 (#8)([`a340efc`](https://github.com/whatchanged-community/whatchanged/commit/a340efc8b86b1728eb1dcaedc9c101767582e811)) (@renovate[bot])
-	**repl**: production assets should be set publicPath([`89ea856`](https://github.com/whatchanged-community/whatchanged/commit/89ea856f4f2046f7347a5ebd2c9d60e3a3650595)) (@axetroy)
-	**repl**: markdown render style([`acebfd0`](https://github.com/whatchanged-community/whatchanged/commit/acebfd0bd736dac9c811186c82ba241d7b1e05e1)) (@axetroy)

### ❤️ BREAKING CHANGES:

-	rename

	```diff
	- tag:0~tag:1
	+ @0~@1
	```

-	new flags

	```diff
	- --dir=/path/to/dir
	+ --project=/path/to/project
	```

	```diff
	- --file=CHANGELOG.md
	+ --output=CHANGELOG.md
	```

-	repo and binary rename to whatchanged

	```diff
	- $ changelog --help
	+ $ whatchanged --help
	```

### 💪 Commits(58):

-	[`f55c367`](https://github.com/whatchanged-community/whatchanged/commit/f55c3677031bbcf295684d2f0510ef838ab10ad4) - ci: update whatchanged command options
-	[`749a021`](https://github.com/whatchanged-community/whatchanged/commit/749a021a56fb4467d2e2bc5e4416ba3d87754c8d) - docs: update changelog
-	[`8608c37`](https://github.com/whatchanged-community/whatchanged/commit/8608c37b4bffd7971f31333a13b06181daa5f9bd) - v0.2.0
-	[`02136b9`](https://github.com/whatchanged-community/whatchanged/commit/02136b9cc348d81ec572daaa8b46d8eaf98dbc52) - ci: fix generation in ci
-	[`65c50c6`](https://github.com/whatchanged-community/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf) - feat: refactor multiple version generation
-	[`dde4ece`](https://github.com/whatchanged-community/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec) - feat: unified usage when invalid options
-	[`ddc712a`](https://github.com/whatchanged-community/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9) - feat: print help/version information to stderr
-	[`d9d3be8`](https://github.com/whatchanged-community/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b) - feat: version range of hash support 8-length short hash
-	[`71a51e1`](https://github.com/whatchanged-community/whatchanged/commit/71a51e17c7011e337e5c4d308d229860c91fcb8f) - faet: add style block for generation
-	[`4a9e7ee`](https://github.com/whatchanged-community/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec) - feat: support to generate multiple specified versions ref: #4
-	[`5bf4bee`](https://github.com/whatchanged-community/whatchanged/commit/5bf4beea7124cac872598c5487657548e7a826c9) - fix: the breaking change block is generate in incorrect format
-	[`bbf648e`](https://github.com/whatchanged-community/whatchanged/commit/bbf648e756ab74abb25764ee9ead032343832b3b) - fix: the incorrect range generated when there are no parameters
-	[`76eb777`](https://github.com/whatchanged-community/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8) - feat: rename tag:N to @N
-	[`e930b89`](https://github.com/whatchanged-community/whatchanged/commit/e930b89df4916e6596f9073be8545e44f3e4abaa) - refactor: improve the extractor
-	[`aa62d6b`](https://github.com/whatchanged-community/whatchanged/commit/aa62d6be3294619c81159a39208c9f7bba07630f) - fix: parse single version tag:N incorrect
-	[`f83db49`](https://github.com/whatchanged-community/whatchanged/commit/f83db49db790c912895963acff61884d4ca847ca) - build(heroku): fix heroku deploy
-	[`45b6ad2`](https://github.com/whatchanged-community/whatchanged/commit/45b6ad20ec4a50dc7661bf575fa408ef6383c46b) - fix(repl): server not response markdown
-	[`4f27a9f`](https://github.com/whatchanged-community/whatchanged/commit/4f27a9f22d37f4ddde213e506d788a4a21c1ee64) - docs: update help information
-	[`bbce6a7`](https://github.com/whatchanged-community/whatchanged/commit/bbce6a7ba7fa6c6bc5878b8d2a8ed70ed474da82) - ci: fix script
-	[`189186a`](https://github.com/whatchanged-community/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017) - feat(cli): rename '--dir' to '--project' and '--file' to '--output' #4
-	[`0bb75dd`](https://github.com/whatchanged-community/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c) - feat: add whatchanged package for Go
-	[`46b9831`](https://github.com/whatchanged-community/whatchanged/commit/46b98319804dbcdb24a937d7560417ec1ccdec23) - ci: update gorelease's build target
-	[`b8bc674`](https://github.com/whatchanged-community/whatchanged/commit/b8bc6742496673a169d2a893aac53924ed317ddb) - refactor: restruct folder
-	[`ee18634`](https://github.com/whatchanged-community/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d) - feat: rename 'changelog' to 'whatchanged'
-	[`39ac73a`](https://github.com/whatchanged-community/whatchanged/commit/39ac73ae8d42533c42cbc93fbb3b48cc74d9a475) - refactor: update repl template
-	[`8a92fd7`](https://github.com/whatchanged-community/whatchanged/commit/8a92fd7693568683beba2431b0e0659fc99e3c82) - fix(repl): generate error
-	[`6007691`](https://github.com/whatchanged-community/whatchanged/commit/6007691ab2123cb56edd41df06bbc67a9eaeae7f) - refactor(repl): update repl default template
-	[`2c800f2`](https://github.com/whatchanged-community/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c) - feat: implement revert parser
-	[`418b8f6`](https://github.com/whatchanged-community/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f) - feat: add chore block
-	[`6410c1b`](https://github.com/whatchanged-community/whatchanged/commit/6410c1be6cc35e3165172f99738add18ef4d5beb) - fix(repl): no-cors mode
-	[`d4af36b`](https://github.com/whatchanged-community/whatchanged/commit/d4af36be80ca60f4bbbcb96603b070883ac44a6a) - fix(repl): source block do not render after request success
-	[`b2bbbbd`](https://github.com/whatchanged-community/whatchanged/commit/b2bbbbd7608501813986d74f6e44c233719246eb) - fix(repl): improve error handler
-	[`4cf9f6e`](https://github.com/whatchanged-community/whatchanged/commit/4cf9f6ee53d19f67380144030a38ede88cb1a59b) - fix(repl): markdown is not rendered correctly
-	[`afa1da6`](https://github.com/whatchanged-community/whatchanged/commit/afa1da6e96523b608ad6cac785d641e1e78ea570) - chore(deps): pin dependency vite to 1.0.0-rc.13 (#9)
-	[`5be056f`](https://github.com/whatchanged-community/whatchanged/commit/5be056fcb8364bc2be1ea09695284791048234fe) - refactor(repl): use new setup scripts
-	[`ec1471f`](https://github.com/whatchanged-community/whatchanged/commit/ec1471fe298e567ecb42adcd8201972fe7f0c6f2) - refactor(repl): update default repo in repl
-	[`a3f02fe`](https://github.com/whatchanged-community/whatchanged/commit/a3f02fe5fa2b134804584af39f851ad2d82c0544) - ci: skip deploy repl in not master branch
-	[`a340efc`](https://github.com/whatchanged-community/whatchanged/commit/a340efc8b86b1728eb1dcaedc9c101767582e811) - fix(deps): update dependency ant-design-vue to v2.0.0-rc.2 (#8)
-	[`5a3e750`](https://github.com/whatchanged-community/whatchanged/commit/5a3e750f71b1e96b8b35e98a3f3c04d21c0490da) - chore(deps): pin dependencies (#7)
-	[`04768c0`](https://github.com/whatchanged-community/whatchanged/commit/04768c0a7bb06c0e790dc16ce766b56ed01a0a14) - chore(repl): update vite config
-	[`516d167`](https://github.com/whatchanged-community/whatchanged/commit/516d16722baca03c20cb358c8dced574e97b7513) - docs: update readme
-	[`fa7e09a`](https://github.com/whatchanged-community/whatchanged/commit/fa7e09acfdedbba398562aa1430eac2ba478b8fe) - chore(deps): remove unused deps
-	[`89ea856`](https://github.com/whatchanged-community/whatchanged/commit/89ea856f4f2046f7347a5ebd2c9d60e3a3650595) - fix(repl): production assets should be set publicPath
-	[`4253a69`](https://github.com/whatchanged-community/whatchanged/commit/4253a69547a7cdb727ba8dfa7eb398858a595af6) - ci: update repl
-	[`4632d13`](https://github.com/whatchanged-community/whatchanged/commit/4632d13a9081f1353fe005a85be77906f20c0cfd) - ci: fix repl delopment
-	[`13e4f7a`](https://github.com/whatchanged-community/whatchanged/commit/13e4f7a64156b06b55edaa3433b6abd1ad9f1d14) - ci: add repl workflow
-	[`acebfd0`](https://github.com/whatchanged-community/whatchanged/commit/acebfd0bd736dac9c811186c82ba241d7b1e05e1) - fix(repl): markdown render style
-	[`c7127b1`](https://github.com/whatchanged-community/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c) - feat: add changelog repl. generate changelog online almose done.
-	[`8b5f7fb`](https://github.com/whatchanged-community/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f) - feat: enabled cors for changelog-remote
-	[`e064058`](https://github.com/whatchanged-community/whatchanged/commit/e064058604658f748dd26ad6bbbee049ea98321e) - ci: fix clone all commit for generate changelog
-	[`5c9b47c`](https://github.com/whatchanged-community/whatchanged/commit/5c9b47c4b6cafb73962c68c619708b484fd00a7e) - ci: add step to generate all changelog for testing it works
-	[`e420ed1`](https://github.com/whatchanged-community/whatchanged/commit/e420ed1d0ae0f1578f1171f1c5ce7d583291b84c) - build: add heroku online server
-	[`50dc305`](https://github.com/whatchanged-community/whatchanged/commit/50dc305fc812151a2f3ae07dd72d83cd7e042007) - docs: fix how it work link
-	[`8f7eb5b`](https://github.com/whatchanged-community/whatchanged/commit/8f7eb5b77317f62cc445bf60bf9fe572441d6c65) - docs: update help information
-	[`5c03aa7`](https://github.com/whatchanged-community/whatchanged/commit/5c03aa7d62c2d27e586b7690b0dc1dc83cebc38e) - chore(deps): update module stretchr/testify to v1.6.1 (#3)
-	[`e6af81e`](https://github.com/whatchanged-community/whatchanged/commit/e6af81e4084680cd7c5a2618ed1c90cbc08751d3) - chore(deps): update module pkg/errors to v0.9.1 (#2)
-	[`e51a5c4`](https://github.com/whatchanged-community/whatchanged/commit/e51a5c434629f49ca7b14787d7bcc28c16e23131) - Add renovate.json
-	[`f34fba7`](https://github.com/whatchanged-community/whatchanged/commit/f34fba7c77e664954e2a954b23270b8bff25c84e) - chore: update install.sh
