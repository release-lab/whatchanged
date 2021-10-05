v0.2.0 (2020-11-27)
-------------------

### 🔥 New feature:

-	refactor multiple version generation([`65c50c6`](https://github.com/release-lab/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf)) (@axetroy)
-	unified usage when invalid options([`dde4ece`](https://github.com/release-lab/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec)) (@axetroy)
-	print help/version information to stderr([`ddc712a`](https://github.com/release-lab/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9)) (@axetroy)
-	version range of hash support 8-length short hash([`d9d3be8`](https://github.com/release-lab/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b)) (@axetroy)
-	support to generate multiple specified versions ref: #4([`4a9e7ee`](https://github.com/release-lab/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec)) (@axetroy)
-	rename tag:N to @N([`76eb777`](https://github.com/release-lab/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8)) (@axetroy)
-	**cli**: rename '--dir' to '--project' and '--file' to '--output' #4([`189186a`](https://github.com/release-lab/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017)) (@axetroy)
-	add whatchanged package for Go([`0bb75dd`](https://github.com/release-lab/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c)) (@axetroy)
-	rename 'changelog' to 'whatchanged'([`ee18634`](https://github.com/release-lab/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d)) (@axetroy) , Closes: #6
-	implement revert parser([`2c800f2`](https://github.com/release-lab/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c)) (@axetroy)
-	add chore block([`418b8f6`](https://github.com/release-lab/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f)) (@axetroy)
-	add changelog repl. generate changelog online almose done.([`c7127b1`](https://github.com/release-lab/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c)) (@axetroy)
-	enabled cors for changelog-remote([`8b5f7fb`](https://github.com/release-lab/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f)) (@axetroy)

### 🐛 Bugs fixed:

-	the breaking change block is generate in incorrect format([`5bf4bee`](https://github.com/release-lab/whatchanged/commit/5bf4beea7124cac872598c5487657548e7a826c9)) (@axetroy)
-	the incorrect range generated when there are no parameters([`bbf648e`](https://github.com/release-lab/whatchanged/commit/bbf648e756ab74abb25764ee9ead032343832b3b)) (@axetroy)
-	parse single version tag:N incorrect([`aa62d6b`](https://github.com/release-lab/whatchanged/commit/aa62d6be3294619c81159a39208c9f7bba07630f)) (@axetroy)
-	**repl**: server not response markdown([`45b6ad2`](https://github.com/release-lab/whatchanged/commit/45b6ad20ec4a50dc7661bf575fa408ef6383c46b)) (@axetroy)
-	**repl**: generate error([`8a92fd7`](https://github.com/release-lab/whatchanged/commit/8a92fd7693568683beba2431b0e0659fc99e3c82)) (@axetroy)
-	**repl**: no-cors mode([`6410c1b`](https://github.com/release-lab/whatchanged/commit/6410c1be6cc35e3165172f99738add18ef4d5beb)) (@axetroy)
-	**repl**: source block do not render after request success([`d4af36b`](https://github.com/release-lab/whatchanged/commit/d4af36be80ca60f4bbbcb96603b070883ac44a6a)) (@axetroy)
-	**repl**: improve error handler([`b2bbbbd`](https://github.com/release-lab/whatchanged/commit/b2bbbbd7608501813986d74f6e44c233719246eb)) (@axetroy)
-	**repl**: markdown is not rendered correctly([`4cf9f6e`](https://github.com/release-lab/whatchanged/commit/4cf9f6ee53d19f67380144030a38ede88cb1a59b)) (@axetroy)
-	**deps**: update dependency ant-design-vue to v2.0.0-rc.2 (#8)([`a340efc`](https://github.com/release-lab/whatchanged/commit/a340efc8b86b1728eb1dcaedc9c101767582e811)) (@renovate[bot])
-	**repl**: production assets should be set publicPath([`89ea856`](https://github.com/release-lab/whatchanged/commit/89ea856f4f2046f7347a5ebd2c9d60e3a3650595)) (@axetroy)
-	**repl**: markdown render style([`acebfd0`](https://github.com/release-lab/whatchanged/commit/acebfd0bd736dac9c811186c82ba241d7b1e05e1)) (@axetroy)

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

-	[`f55c367`](https://github.com/release-lab/whatchanged/commit/f55c3677031bbcf295684d2f0510ef838ab10ad4) - ci: update whatchanged command options
-	[`749a021`](https://github.com/release-lab/whatchanged/commit/749a021a56fb4467d2e2bc5e4416ba3d87754c8d) - docs: update changelog
-	[`8608c37`](https://github.com/release-lab/whatchanged/commit/8608c37b4bffd7971f31333a13b06181daa5f9bd) - v0.2.0
-	[`02136b9`](https://github.com/release-lab/whatchanged/commit/02136b9cc348d81ec572daaa8b46d8eaf98dbc52) - ci: fix generation in ci
-	[`65c50c6`](https://github.com/release-lab/whatchanged/commit/65c50c6b30f5dfc608c260c73d55cc8601041bdf) - feat: refactor multiple version generation
-	[`dde4ece`](https://github.com/release-lab/whatchanged/commit/dde4ecee3db8925c804db839bed098eb4a0f82ec) - feat: unified usage when invalid options
-	[`ddc712a`](https://github.com/release-lab/whatchanged/commit/ddc712a8e4ec502976deb7430a79532d902bcbf9) - feat: print help/version information to stderr
-	[`d9d3be8`](https://github.com/release-lab/whatchanged/commit/d9d3be819e1214139342e48cdaf866ae3b628f4b) - feat: version range of hash support 8-length short hash
-	[`71a51e1`](https://github.com/release-lab/whatchanged/commit/71a51e17c7011e337e5c4d308d229860c91fcb8f) - faet: add style block for generation
-	[`4a9e7ee`](https://github.com/release-lab/whatchanged/commit/4a9e7ee70a80104933e60f20db4784ea472ae2ec) - feat: support to generate multiple specified versions ref: #4
-	[`5bf4bee`](https://github.com/release-lab/whatchanged/commit/5bf4beea7124cac872598c5487657548e7a826c9) - fix: the breaking change block is generate in incorrect format
-	[`bbf648e`](https://github.com/release-lab/whatchanged/commit/bbf648e756ab74abb25764ee9ead032343832b3b) - fix: the incorrect range generated when there are no parameters
-	[`76eb777`](https://github.com/release-lab/whatchanged/commit/76eb7774ac0a1f44ca6b66b9322870cba24a50a8) - feat: rename tag:N to @N
-	[`e930b89`](https://github.com/release-lab/whatchanged/commit/e930b89df4916e6596f9073be8545e44f3e4abaa) - refactor: improve the extractor
-	[`aa62d6b`](https://github.com/release-lab/whatchanged/commit/aa62d6be3294619c81159a39208c9f7bba07630f) - fix: parse single version tag:N incorrect
-	[`f83db49`](https://github.com/release-lab/whatchanged/commit/f83db49db790c912895963acff61884d4ca847ca) - build(heroku): fix heroku deploy
-	[`45b6ad2`](https://github.com/release-lab/whatchanged/commit/45b6ad20ec4a50dc7661bf575fa408ef6383c46b) - fix(repl): server not response markdown
-	[`4f27a9f`](https://github.com/release-lab/whatchanged/commit/4f27a9f22d37f4ddde213e506d788a4a21c1ee64) - docs: update help information
-	[`bbce6a7`](https://github.com/release-lab/whatchanged/commit/bbce6a7ba7fa6c6bc5878b8d2a8ed70ed474da82) - ci: fix script
-	[`189186a`](https://github.com/release-lab/whatchanged/commit/189186a89693c724ac794c17f2c35781b2fdc017) - feat(cli): rename '--dir' to '--project' and '--file' to '--output' #4
-	[`0bb75dd`](https://github.com/release-lab/whatchanged/commit/0bb75dd41758d85c4608f010298f823346a68a7c) - feat: add whatchanged package for Go
-	[`46b9831`](https://github.com/release-lab/whatchanged/commit/46b98319804dbcdb24a937d7560417ec1ccdec23) - ci: update gorelease's build target
-	[`b8bc674`](https://github.com/release-lab/whatchanged/commit/b8bc6742496673a169d2a893aac53924ed317ddb) - refactor: restruct folder
-	[`ee18634`](https://github.com/release-lab/whatchanged/commit/ee1863487bd70a2664ff856c4aacfc34d3a5043d) - feat: rename 'changelog' to 'whatchanged'
-	[`39ac73a`](https://github.com/release-lab/whatchanged/commit/39ac73ae8d42533c42cbc93fbb3b48cc74d9a475) - refactor: update repl template
-	[`8a92fd7`](https://github.com/release-lab/whatchanged/commit/8a92fd7693568683beba2431b0e0659fc99e3c82) - fix(repl): generate error
-	[`6007691`](https://github.com/release-lab/whatchanged/commit/6007691ab2123cb56edd41df06bbc67a9eaeae7f) - refactor(repl): update repl default template
-	[`2c800f2`](https://github.com/release-lab/whatchanged/commit/2c800f24894c495761e715e3a3f81863e0b3b96c) - feat: implement revert parser
-	[`418b8f6`](https://github.com/release-lab/whatchanged/commit/418b8f6383b9d710c043655a5dd28fd6627bd85f) - feat: add chore block
-	[`6410c1b`](https://github.com/release-lab/whatchanged/commit/6410c1be6cc35e3165172f99738add18ef4d5beb) - fix(repl): no-cors mode
-	[`d4af36b`](https://github.com/release-lab/whatchanged/commit/d4af36be80ca60f4bbbcb96603b070883ac44a6a) - fix(repl): source block do not render after request success
-	[`b2bbbbd`](https://github.com/release-lab/whatchanged/commit/b2bbbbd7608501813986d74f6e44c233719246eb) - fix(repl): improve error handler
-	[`4cf9f6e`](https://github.com/release-lab/whatchanged/commit/4cf9f6ee53d19f67380144030a38ede88cb1a59b) - fix(repl): markdown is not rendered correctly
-	[`afa1da6`](https://github.com/release-lab/whatchanged/commit/afa1da6e96523b608ad6cac785d641e1e78ea570) - chore(deps): pin dependency vite to 1.0.0-rc.13 (#9)
-	[`5be056f`](https://github.com/release-lab/whatchanged/commit/5be056fcb8364bc2be1ea09695284791048234fe) - refactor(repl): use new setup scripts
-	[`ec1471f`](https://github.com/release-lab/whatchanged/commit/ec1471fe298e567ecb42adcd8201972fe7f0c6f2) - refactor(repl): update default repo in repl
-	[`a3f02fe`](https://github.com/release-lab/whatchanged/commit/a3f02fe5fa2b134804584af39f851ad2d82c0544) - ci: skip deploy repl in not master branch
-	[`a340efc`](https://github.com/release-lab/whatchanged/commit/a340efc8b86b1728eb1dcaedc9c101767582e811) - fix(deps): update dependency ant-design-vue to v2.0.0-rc.2 (#8)
-	[`5a3e750`](https://github.com/release-lab/whatchanged/commit/5a3e750f71b1e96b8b35e98a3f3c04d21c0490da) - chore(deps): pin dependencies (#7)
-	[`04768c0`](https://github.com/release-lab/whatchanged/commit/04768c0a7bb06c0e790dc16ce766b56ed01a0a14) - chore(repl): update vite config
-	[`516d167`](https://github.com/release-lab/whatchanged/commit/516d16722baca03c20cb358c8dced574e97b7513) - docs: update readme
-	[`fa7e09a`](https://github.com/release-lab/whatchanged/commit/fa7e09acfdedbba398562aa1430eac2ba478b8fe) - chore(deps): remove unused deps
-	[`89ea856`](https://github.com/release-lab/whatchanged/commit/89ea856f4f2046f7347a5ebd2c9d60e3a3650595) - fix(repl): production assets should be set publicPath
-	[`4253a69`](https://github.com/release-lab/whatchanged/commit/4253a69547a7cdb727ba8dfa7eb398858a595af6) - ci: update repl
-	[`4632d13`](https://github.com/release-lab/whatchanged/commit/4632d13a9081f1353fe005a85be77906f20c0cfd) - ci: fix repl delopment
-	[`13e4f7a`](https://github.com/release-lab/whatchanged/commit/13e4f7a64156b06b55edaa3433b6abd1ad9f1d14) - ci: add repl workflow
-	[`acebfd0`](https://github.com/release-lab/whatchanged/commit/acebfd0bd736dac9c811186c82ba241d7b1e05e1) - fix(repl): markdown render style
-	[`c7127b1`](https://github.com/release-lab/whatchanged/commit/c7127b1b0e3869854d293b536eb2f21c4e0c8e3c) - feat: add changelog repl. generate changelog online almose done.
-	[`8b5f7fb`](https://github.com/release-lab/whatchanged/commit/8b5f7fbda0f6aefbc933de757a13ed34d105990f) - feat: enabled cors for changelog-remote
-	[`e064058`](https://github.com/release-lab/whatchanged/commit/e064058604658f748dd26ad6bbbee049ea98321e) - ci: fix clone all commit for generate changelog
-	[`5c9b47c`](https://github.com/release-lab/whatchanged/commit/5c9b47c4b6cafb73962c68c619708b484fd00a7e) - ci: add step to generate all changelog for testing it works
-	[`e420ed1`](https://github.com/release-lab/whatchanged/commit/e420ed1d0ae0f1578f1171f1c5ce7d583291b84c) - build: add heroku online server
-	[`50dc305`](https://github.com/release-lab/whatchanged/commit/50dc305fc812151a2f3ae07dd72d83cd7e042007) - docs: fix how it work link
-	[`8f7eb5b`](https://github.com/release-lab/whatchanged/commit/8f7eb5b77317f62cc445bf60bf9fe572441d6c65) - docs: update help information
-	[`5c03aa7`](https://github.com/release-lab/whatchanged/commit/5c03aa7d62c2d27e586b7690b0dc1dc83cebc38e) - chore(deps): update module stretchr/testify to v1.6.1 (#3)
-	[`e6af81e`](https://github.com/release-lab/whatchanged/commit/e6af81e4084680cd7c5a2618ed1c90cbc08751d3) - chore(deps): update module pkg/errors to v0.9.1 (#2)
-	[`e51a5c4`](https://github.com/release-lab/whatchanged/commit/e51a5c434629f49ca7b14787d7bcc28c16e23131) - Add renovate.json
-	[`f34fba7`](https://github.com/release-lab/whatchanged/commit/f34fba7c77e664954e2a954b23270b8bff25c84e) - chore: update install.sh

v0.1.0 (2020-11-24)
-------------------

### 🔥 New feature:

-	link commit for generation([`b9432db`](https://github.com/release-lab/whatchanged/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb)) (@axetroy)
-	add full preset template([`7553570`](https://github.com/release-lab/whatchanged/commit/7553570590b571bd33e10a4f80ec5639d0613042)) (@axetroy)
-	support changelog for git submodule([`ec6a957`](https://github.com/release-lab/whatchanged/commit/ec6a957752fbca9faa261d8694826779e2cbec1f)) (@axetroy)
-	add writer step([`441ad13`](https://github.com/release-lab/whatchanged/commit/441ad1322b1fecaca89a170ecebaf2955a77d630)) (@axetroy)
-	add formatter for markdown output([`8c177e0`](https://github.com/release-lab/whatchanged/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34)) (@axetroy)
-	add --file flag to generate file([`0e4fb09`](https://github.com/release-lab/whatchanged/commit/0e4fb09789732fec5b09b247e208d61794c3da0d)) (@axetroy)
-	support custom tmeplate and preset([`3aa0aee`](https://github.com/release-lab/whatchanged/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db)) (@axetroy)
-	support tag ranges([`3d14c9c`](https://github.com/release-lab/whatchanged/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9)) (@axetroy)
-	support version range. eg v2.0.0~v1.0.0([`a65a4a8`](https://github.com/release-lab/whatchanged/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3)) (@axetroy)
-	parse commit message and generate to template([`7f67e78`](https://github.com/release-lab/whatchanged/commit/7f67e783926fed647d2ad5414f31448eea106fc3)) (@axetroy)

### 🐛 Bugs fixed:

-	improve ssh git url parser([`9993ee6`](https://github.com/release-lab/whatchanged/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f)) (@axetroy)
-	commit range not include commit of tag([`f8df0ba`](https://github.com/release-lab/whatchanged/commit/f8df0ba654c8faf67eccf98262cd55807e53e597)) (@axetroy)
-	unescape template([`e118cbf`](https://github.com/release-lab/whatchanged/commit/e118cbfafd201b945848f15303fdb261e251f058)) (@axetroy)
-	if empty argument for command line([`9c79fd9`](https://github.com/release-lab/whatchanged/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd)) (@axetroy)
-	**ci**: remove unsued code([`66bcf8f`](https://github.com/release-lab/whatchanged/commit/66bcf8f43db85409e0392c93f2e347ed91699e81)) (@axetroy)

### 💪 Commits(43):

-	[`7a3fd6d`](https://github.com/release-lab/whatchanged/commit/7a3fd6de68462fb64266bf3f606e957d32cdc2ae) - ci: add release template
-	[`00d6cc7`](https://github.com/release-lab/whatchanged/commit/00d6cc7bdf42d6a0b58cf15e0aea65710f1efbf1) - ci: debug
-	[`e6f97da`](https://github.com/release-lab/whatchanged/commit/e6f97daba33e8c912ee738e4fceb2bf3da877cee) - ci: debug
-	[`62e3ae1`](https://github.com/release-lab/whatchanged/commit/62e3ae1045c800c49743374cf83705ba03ce5bd7) - ci: debug
-	[`10b0cf3`](https://github.com/release-lab/whatchanged/commit/10b0cf3b6f91085627f8c59351570abdeb01d354) - ci: debug
-	[`cc5dd66`](https://github.com/release-lab/whatchanged/commit/cc5dd665e055c066738dc2fcdb26aec61d4e61b0) - ci: update gorelease
-	[`343e4a8`](https://github.com/release-lab/whatchanged/commit/343e4a8f3eaf8a715030652cfd3b7e696e40422e) - chore: add release.md to gitignore
-	[`6891359`](https://github.com/release-lab/whatchanged/commit/68913598822611d64a60211918b94a462822e734) - ci: fix ci
-	[`34f51dc`](https://github.com/release-lab/whatchanged/commit/34f51dc2c90e4c8ba822666dd5dcbabe312e2841) - build: rename binary name
-	[`fecc276`](https://github.com/release-lab/whatchanged/commit/fecc27679d6431f840eb57c3359aa84bec1607ef) - ci: update gorelease
-	[`7c687f7`](https://github.com/release-lab/whatchanged/commit/7c687f7d729f199ee295109012cc5a2661f96c54) - ci: increase linter timeout
-	[`b61c491`](https://github.com/release-lab/whatchanged/commit/b61c49171a776045d89e6cbaad82326070e2db78) - refactor: improve default preset
-	[`1292dfc`](https://github.com/release-lab/whatchanged/commit/1292dfc0e39abf24262e94cd9076d91808dd0cc4) - refactor: update template
-	[`9993ee6`](https://github.com/release-lab/whatchanged/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f) - fix: improve ssh git url parser
-	[`b9432db`](https://github.com/release-lab/whatchanged/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb) - feat: link commit for generation
-	[`7553570`](https://github.com/release-lab/whatchanged/commit/7553570590b571bd33e10a4f80ec5639d0613042) - feat: add full preset template
-	[`7068672`](https://github.com/release-lab/whatchanged/commit/706867220fa9ca537855f359d3e04d0c3762b793) - update readme
-	[`ec6a957`](https://github.com/release-lab/whatchanged/commit/ec6a957752fbca9faa261d8694826779e2cbec1f) - feat: support changelog for git submodule
-	[`f540d6f`](https://github.com/release-lab/whatchanged/commit/f540d6f7123334dac558a37c6ac056fed1021cda) - refactor: rename transform to transformer
-	[`441ad13`](https://github.com/release-lab/whatchanged/commit/441ad1322b1fecaca89a170ecebaf2955a77d630) - feat: add writer step
-	[`f8df0ba`](https://github.com/release-lab/whatchanged/commit/f8df0ba654c8faf67eccf98262cd55807e53e597) - fix: commit range not include commit of tag
-	[`8c177e0`](https://github.com/release-lab/whatchanged/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34) - feat: add formatter for markdown output
-	[`a47fe84`](https://github.com/release-lab/whatchanged/commit/a47fe84d2141635d82c2dc49500bfc2a81c03535) - docs: update how it works
-	[`0e4fb09`](https://github.com/release-lab/whatchanged/commit/0e4fb09789732fec5b09b247e208d61794c3da0d) - feat: add --file flag to generate file
-	[`3aa0aee`](https://github.com/release-lab/whatchanged/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db) - feat: support custom tmeplate and preset
-	[`e118cbf`](https://github.com/release-lab/whatchanged/commit/e118cbfafd201b945848f15303fdb261e251f058) - fix: unescape template
-	[`3d14c9c`](https://github.com/release-lab/whatchanged/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9) - feat: support tag ranges
-	[`44ce3cd`](https://github.com/release-lab/whatchanged/commit/44ce3cd8d68b786a3aac6b5daba90e7f12b75200) - docs: update readme
-	[`9c79fd9`](https://github.com/release-lab/whatchanged/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd) - fix: if empty argument for command line
-	[`0b6ba0c`](https://github.com/release-lab/whatchanged/commit/0b6ba0c3fc49139467025eaebbe16c158a0cce65) - refactor: Refactor the software architecture to make it clearer and simpler
-	[`66bcf8f`](https://github.com/release-lab/whatchanged/commit/66bcf8f43db85409e0392c93f2e347ed91699e81) - fix(ci): remove unsued code
-	[`89fecf5`](https://github.com/release-lab/whatchanged/commit/89fecf5588f1133f70a8aaf6db3488184ba2ceb2) - refactor: remove unsued code
-	[`a65a4a8`](https://github.com/release-lab/whatchanged/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3) - feat: support version range. eg v2.0.0~v1.0.0
-	[`6301c2f`](https://github.com/release-lab/whatchanged/commit/6301c2f01d881ae861e6c4459a411a5f48c74ba0) - update help information
-	[`770ed02`](https://github.com/release-lab/whatchanged/commit/770ed02d43c4593ab9db8e71a9f93987812d97bc) - update
-	[`585445d`](https://github.com/release-lab/whatchanged/commit/585445d917d4cb74d80b1c385b446f7e09a1606c) - cache tags
-	[`b47e49c`](https://github.com/release-lab/whatchanged/commit/b47e49cf8efac3f5aba9df0149295694424beaf1) - filter tag with semver version
-	[`7f67e78`](https://github.com/release-lab/whatchanged/commit/7f67e783926fed647d2ad5414f31448eea106fc3) - feat: parse commit message and generate to template
-	[`b907117`](https://github.com/release-lab/whatchanged/commit/b907117bca3d6306955070b933361ecb0da0627e) - update ci
-	[`ef5b38a`](https://github.com/release-lab/whatchanged/commit/ef5b38ad50dd150cbdbeff031f6898d9d0aff35a) - update ci linter version
-	[`dd55cc8`](https://github.com/release-lab/whatchanged/commit/dd55cc85d6a3b9c482ba376fa862f20b6de11d5b) - fix lint
-	[`26408cc`](https://github.com/release-lab/whatchanged/commit/26408ccef0f6256ca70edda59e4ab1d1c15fca72) - init
-	[`824d351`](https://github.com/release-lab/whatchanged/commit/824d3511bf90e8fda3d1c7be679274d71ce73f52) - Initial commit
