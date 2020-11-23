v4.0.0
======

### New feature:

-	deprecated(7729788) (thanks @axetroy)

### Code Refactoring:

-	rename function name(a6cb403) (thanks @axetroy)
-	do not create json module(76b550c) (thanks @axetroy)
-	update(cde4dda) (thanks @axetroy)
-	remove unused extension(c20037d) (thanks @axetroy)
-	move configuration to Core module(8066121) (thanks @axetroy)

### Documentation:

-	update changelog(d30f167) (thanks @axetroy)

### Commits(15):

-	**d30f167** docs: update changelog
-	**7729788** feat: deprecated
-	**6d13713** Revert "deprecated"
-	**bf08694** deprecated
-	**a6cb403** refactor: rename function name
-	**76b550c** refactor: do not create json module
-	**5981863** Update dependency execa to v4.0.1 (#175)
-	**86f527c** Update dependency eslint to v7 (#176)
-	**2d16ba4** Update dependency @types/node to v12.12.38 (#173)
-	**777fe9e** Update dependency jest to v26.0.1 (#171)
-	**25b190b** Revert "refactor: remove unused extension"
-	**479ecd2** Revert "refactor: update"
-	**cde4dda** refactor: update
-	**c20037d** refactor: remove unused extension
-	**8066121** refactor: move configuration to Core module

	v3.7.0
	======

	### New feature:

-	remove JSON import statement. Now, you cannot import JSON module.(23ff6f3) (thanks @axetroy)

-	added unstable settings option (#167)(1a4a230) (thanks @Luca Casonato)

### Code Refactoring:

-	sort deps tree viewer(c7009d7) (thanks @axetroy)
-	improve code(6414b9a) (thanks @axetroy)

### Testing:

-	fix test case(14603bd) (thanks @axetroy)

### Documentation:

-	update changelog(4e36065) (thanks @axetroy)
-	update i18n(2f7845f) (thanks @axetroy)
-	update readme(8ea4cf2) (thanks @axetroy)
-	update i18n(9ab4dd4) (thanks @axetroy)

### Commits(15):

-	**4e36065** docs: update changelog
-	**2f7845f** docs: update i18n
-	**23ff6f3** feat: remove JSON import statement. Now, you cannot import JSON module.
-	**8ea4cf2** docs: update readme
-	**9ab4dd4** docs: update i18n
-	**c7009d7** refactor: sort deps tree viewer
-	**1a4a230** feat: added unstable settings option (#167)
-	**27e9871** chore(deps): update dependency jest to v26 (#170)
-	**859de1e** chore(deps): update typescript-eslint monorepo to v2.31.0 (#169)
-	**4279482** chore(deps): update dependency jest to v25.5.4 (#168)
-	**b71977e** chore(deps): update dependency jest to v25.5.3 (#165)
-	**502f5fe** chore(deps): update dependency jest to v25.5.2 (#163)
-	**62279e9** chore(deps): update dependency jest to v25.5.1 (#162)
-	**14603bd** test: fix test case
-	**6414b9a** refactor: improve code

	v3.6.2
	======

	### Bugs fixed:

-	file protocol import statement not work. close #146(67897bc) (thanks @axetroy)

-	not deno project also show deno deps tree view.(81303df) (thanks @axetroy)

### Code Refactoring:

-	improve code(2ba8078) (thanks @axetroy)
-	fix typo(85514ef) (thanks @axetroy)
-	update i18n(30b78ef) (thanks @axetroy)
-	remove unused i18n(5c785cc) (thanks @axetroy)
-	remove unused code(e29effc) (thanks @axetroy)
-	improve code(96ed751) (thanks @axetroy)

### Documentation:

-	update changelog(8a9ea10) (thanks @axetroy)
-	fix typo (#158)(eedf35d) (thanks @bouzuya)

### Commits(21):

-	**8a9ea10** docs: update changelog
-	**67897bc** fix: file protocol import statement not work. close #146
-	**2ba8078** refactor: improve code
-	**85514ef** refactor: fix typo
-	**30b78ef** refactor: update i18n
-	**5c785cc** refactor: remove unused i18n
-	**ea23992** chore(deps): update dependency jest to v25.5.0 (#161)
-	**81303df** fix: not deno project also show deno deps tree view.
-	**bcd1ace** chore(deps): update typescript-eslint monorepo to v2.30.0 (#160)
-	**e29effc** refactor: remove unused code
-	**eedf35d** docs: fix typo (#158)
-	**6e61355** chore(deps): update dependency coveralls to v3.1.0 (#159)
-	**8ac253e** chore(deps): update dependency vscode to v1.1.37 (#156)
-	**1e1ccae** chore(deps): update dependency coveralls to v3.0.14 (#157)
-	**e0faaa5** chore(deps): update typescript-eslint monorepo to v2.29.0 (#151)
-	**1b1cd98** chore(deps): update dependency @types/node to v12.12.37 (#154)
-	**d905710** chore(deps): update dependency prettier to v2.0.5 (#152)
-	**4a1630f** chore(deps): update dependency coveralls to v3.0.13 (#153)
-	**73fc185** chore(deps): update dependency jest to v25.4.0 (#150)
-	**94e38e9** chore(deps): update dependency @types/node to v12.12.36 (#149)
-	**96ed751** refactor: improve code

	v3.6.1
	======

	### Bugs fixed:

-	In Deno's cache module, `x-typescript-types` and redirects are not parsed correctly. close #147(a3a957f) (thanks @axetroy)

-	**deps**: update dependency semver to v7.3.2 (#144)(8496c84) (thanks @renovate[bot])

-	**deps**: update dependency semver to v7.3.1 (#143)(16a112b) (thanks @renovate[bot])

### Code Refactoring:

-	improve code(ffdf7ff) (thanks @axetroy)
-	improve file_walker(cc79850) (thanks @axetroy)

### Documentation:

-	update changelog(963e134) (thanks @axetroy)
-	update example(09a57a2) (thanks @axetroy)

### Commits(8):

-	**963e134** docs: update changelog
-	**a3a957f** fix: In Deno's cache module, `x-typescript-types` and redirects are not parsed correctly. close #147
-	**072573c** chore(deps): update dependency ts-jest to v25.4.0 (#148)
-	**ffdf7ff** refactor: improve code
-	**09a57a2** docs: update example
-	**8496c84** fix(deps): update dependency semver to v7.3.2 (#144)
-	**16a112b** fix(deps): update dependency semver to v7.3.1 (#143)
-	**cc79850** refactor: improve file_walker

	v3.6.0
	======

	### New feature:

-	ignore typescript compile options which ignore by Deno(b48fed0) (thanks @axetroy)

### Bugs fixed:

-	import map with trailing slash (#142)(d5ecc7e) (thanks @Chuah Chee Shian)
-	**deps**: update dependency semver to v7.3.0 (#140)(13e3947) (thanks @renovate[bot])
-	**deps**: update dependency semver to v7.2.2 (#137)(d9bd9c8) (thanks @renovate[bot])

### Code Refactoring:

-	improve module resolver(6881e33) (thanks @axetroy)

### Documentation:

-	update changelog(41fbefe) (thanks @axetroy)
-	update readme(85b68d7) (thanks @axetroy)

### Commits(11):

-	**41fbefe** docs: update changelog
-	**d5ecc7e** fix: import map with trailing slash (#142)
-	**ef5dadf** chore: update ci
-	**423231e** chore(deps): update typescript-eslint monorepo to v2.28.0 (#138)
-	**13e3947** fix(deps): update dependency semver to v7.3.0 (#140)
-	**85b68d7** docs: update readme
-	**d9bd9c8** fix(deps): update dependency semver to v7.2.2 (#137)
-	**6881e33** refactor: improve module resolver
-	**b48fed0** feat: ignore typescript compile options which ignore by Deno
-	**7072f56** chore(deps): update dependency husky to v4.2.5 (#135)
-	**c5cf550** chore(deps): update dependency husky to v4.2.4 (#133)

	v3.5.1
	======

	### Bugs fixed:

-	add more test case for import_map. ref #132(e4b1d6a) (thanks @axetroy)

### Documentation:

-	update changelog(b7e2c14) (thanks @axetroy)
-	update readme(8f5ddf1) (thanks @axetroy)

### Commits(3):

-	**b7e2c14** docs: update changelog
-	**e4b1d6a** fix: add more test case for import_map. ref #132
-	**8f5ddf1** docs: update readme

	v3.5.0
	======

	### New feature:

-	compatible deno cache command since Deno v0.40.0(7309b0c) (thanks @axetroy)

-	support deno-types compile hint (#129)(9be33a4) (thanks @艾斯特洛)

### Bugs fixed:

-	refresh diagnostic not work(f8e8e70) (thanks @axetroy)

### Documentation:

-	update changelog(d6f099e) (thanks @axetroy)
-	update changelog(b85da47) (thanks @axetroy)

### Commits(7):

-	**d6f099e** docs: update changelog
-	**f8e8e70** fix: refresh diagnostic not work
-	**7309b0c** feat: compatible deno cache command since Deno v0.40.0
-	**9be33a4** feat: support deno-types compile hint (#129)
-	**3270344** chore(deps): update dependency @types/node to v12.12.35 (#131)
-	**d2f6ff5** chore(deps): update dependency jest to v25.3.0 (#130)
-	**b85da47** docs: update changelog

	v3.4.2
	======

	### Bugs fixed:

-	if `x-typescript-types` do not exist. then fallback to origin file(05496e3) (thanks @axetroy)

-	**deps**: update dependency semver to v7.2.1 (#128)(90f4ad2) (thanks @renovate[bot])

### Code Refactoring:

-	add extension support for resolvedModule(5ea1d19) (thanks @axetroy)

### Commits(9):

-	**05496e3** fix: if `x-typescript-types` do not exist. then fallback to origin file
-	**5ea1d19** refactor: add extension support for resolvedModule
-	**90f4ad2** fix(deps): update dependency semver to v7.2.1 (#128)
-	**1f03dd7** chore(deps): update typescript-eslint monorepo to v2.27.0 (#127)
-	**d1b204a** chore(deps): update dependency prettier to v2.0.4 (#126)
-	**57d7d87** chore(deps): update dependency ts-jest to v25.3.1 (#123)
-	**6bc4aa0** chore(deps): update dependency jest to v25.2.7 (#121)
-	**0e7c226** chore(deps): update dependency @types/jest to v25.2.1 (#120)
-	**1c461d9** chore(deps): update dependency jest to v25.2.6 (#119)

	v3.4.1
	======

	### Bugs fixed:

-	Tsserver crashes in some cases(11563b4) (thanks @axetroy)

### Documentation:

-	update changelog(27e717e) (thanks @axetroy)
-	add example for react(078065c) (thanks @axetroy)

### Commits(8):

-	**27e717e** docs: update changelog
-	**a2d7da1** chore(deps): update dependency @types/node to v12.12.34 (#115)
-	**0f98dd8** chore(deps): update typescript-eslint monorepo to v2.26.0 (#114)
-	**c981837** chore(deps): update dependency jest to v25.2.4 (#112)
-	**63e4874** chore(deps): update dependency ts-jest to v25.3.0 (#113)
-	**078065c** docs: add example for react
-	**11563b4** fix: Tsserver crashes in some cases
-	**cfd9623** chore(deps): update dependency @types/node to v12.12.32 (#110)

	v3.4.0
	======

	### New feature:

-	add diagnostic for checking valid import statement(b2f070a) (thanks @axetroy)

### Bugs fixed:

-	importmap not work when set to a relative path. close #103(0e8398f) (thanks @axetroy)

### Documentation:

-	update changelog(8831fac) (thanks @axetroy)

### Commits(3):

-	**8831fac** docs: update changelog
-	**b2f070a** feat: add diagnostic for checking valid import statement
-	**0e8398f** fix: importmap not work when set to a relative path. close #103

	v3.3.1
	======

	### Bugs fixed:

-	Triple-Slash Directive does not work. ref #102(e381390) (thanks @axetroy)

-	output target of compileOption now is esnext(fcbd234) (thanks @axetroy)

-	Synchronization Deno error code(4b6efd4) (thanks @axetroy)

-	auto import module not working properly in some edge cases(71b518b) (thanks @axetroy)

-	**deps**: update dependency vscode-languageclient to v6.1.3 (#101)(f5a743c) (thanks @renovate[bot])

-	**deps**: update dependency vscode-languageclient to v6.1.2 (#99)(550f0bc) (thanks @renovate[bot])

### Code Refactoring:

-	merge(b3c0cec) (thanks @axetroy)
-	format code(cfd2147) (thanks @axetroy)

### Documentation:

-	update changelog(a7881b9) (thanks @axetroy)

### Commits(19):

-	**a7881b9** docs: update changelog
-	**e381390** fix: Triple-Slash Directive does not work. ref #102
-	**806c591** chore(deps): update dependency jest to v25.2.3 (#109)
-	**10c54f5** chore(deps): update dependency jest to v25.2.1 (#108)
-	**fcbd234** fix: output target of compileOption now is esnext
-	**b3c0cec** refactor: merge
-	**4b6efd4** fix: Synchronization Deno error code
-	**71b518b** fix: auto import module not working properly in some edge cases
-	**5278331** update
-	**3827164** update
-	**76f0ee6** update
-	**a3ba1cb** update
-	**cfd2147** refactor: format code
-	**b6f71b2** Merge branch 'master' into renovate/prettier-2.x
-	**ded9f89** chore(deps): update dependency prettier to v2
-	**f5a743c** fix(deps): update dependency vscode-languageclient to v6.1.3 (#101)
-	**23820f0** chore(deps): update dependency coveralls to v3.0.11 (#100)
-	**550f0bc** fix(deps): update dependency vscode-languageclient to v6.1.2 (#99)
-	**2cbae37** chore(deps): update typescript-eslint monorepo to v2.24.0 (#98)

	v3.3.0
	======

	### New feature:

-	remove `deno.dts_file` configuration (#94)(f06e852) (thanks @艾斯特洛)

-	add copy to clipboard message(4dd4b8d) (thanks @axetroy)

### Bugs fixed:

-	cannot resolve module if location headers is relative or absolute path. close #97(75d6027) (thanks @axetroy)

### Code Refactoring:

-	improve module resolver(011cbb3) (thanks @axetroy)
-	improve code(427a3ca) (thanks @axetroy)
-	bump to vscode^1.43.0(1967ae3) (thanks @axetroy)
-	improve code(033cdb4) (thanks @axetroy)
-	improve hash_meta(a6be8c5) (thanks @axetroy)

### Testing:

-	improve test case(f9879b7) (thanks @axetroy)

### Documentation:

-	update changelog(3a4eec6) (thanks @axetroy)
-	update readme(445ea65) (thanks @axetroy)

### BREAKING CHANGES:

-	no more use `deno.dts_file` anymore We think this configuration item is redundant

	You can configure external declaration files in `tsconfig.json`

	before:

	```json5
	//.vscode/settings.json
	{
	  "deno.enable": true,
	  "deno.dts_file": ["./path/to/demo.d.ts"]
	}

	```

	after:

	```json5
	//.vscode/settings.json
	{
	  "deno.enable": true
	}

	```

### Commits(13):

-	**f9879b7** test: improve test case
-	**3a4eec6** docs: update changelog
-	**011cbb3** refactor: improve module resolver
-	**75d6027** fix: cannot resolve module if location headers is relative or absolute path. close #97
-	**42e9fb5** chore(deps): update dependency @types/node to v12.12.30 (#96)
-	**f06e852** feat: remove `deno.dts_file` configuration (#94)
-	**445ea65** docs: update readme
-	**427a3ca** refactor: improve code
-	**4dd4b8d** feat: add copy to clipboard message
-	**1967ae3** refactor: bump to vscode^1.43.0
-	**033cdb4** refactor: improve code
-	**ccc6510** chore(deps): update typescript-eslint monorepo to v2.23.0 (#92)
-	**a6be8c5** refactor: improve hash_meta

	v3.2.1
	======

	### New feature:

-	copy url if click code_lens(e3829d8) (thanks @axetroy)

-	make Deno declaration file read-only(fba8d89) (thanks @axetroy)

### Bugs fixed:

-	auto-import not work for some modules. close #44(11d38b3) (thanks @axetroy)
-	If query exists in url, module will not be parsed correctly(a8965b5) (thanks @axetroy)

### Code Refactoring:

-	improve code(7cd9a55) (thanks @axetroy)
-	improve code(de9506c) (thanks @axetroy)
-	re-enable noUnusedLocals in tsconfig.json(8e489b5) (thanks @axetroy)
-	remove unused code(3d642bd) (thanks @axetroy)
-	improve code(1a2b7b6) (thanks @axetroy)
-	Improve code quality(5794907) (thanks @axetroy)
-	improve code(f974d97) (thanks @axetroy)
-	improve code(1c156b1) (thanks @axetroy)

### Testing:

-	add more test case for module_resolver(649b9dc) (thanks @axetroy)
-	improve test case(2ca1638) (thanks @axetroy)

### Documentation:

-	update changelog(5c7604f) (thanks @axetroy)

### Commits(17):

-	**5c7604f** docs: update changelog
-	**11d38b3** fix: auto-import not work for some modules. close #44
-	**7cd9a55** refactor: improve code
-	**de9506c** refactor: improve code
-	**8e489b5** refactor: re-enable noUnusedLocals in tsconfig.json
-	**3d642bd** refactor: remove unused code
-	**e3829d8** feat: copy url if click code_lens
-	**1a2b7b6** refactor: improve code
-	**a8965b5** fix: If query exists in url, module will not be parsed correctly
-	**e871c18** chore(deps): update dependency @types/jest to v25.1.4 (#90)
-	**fba8d89** feat: make Deno declaration file read-only
-	**5794907** refactor: Improve code quality
-	**f974d97** refactor: improve code
-	**649b9dc** test: add more test case for module_resolver
-	**01d58ea** chore: fix format
-	**2ca1638** test: improve test case
-	**1c156b1** refactor: improve code

	v3.2.0
	======

	### New feature:

-	improve fetching module message(3bb70e2) (thanks @axetroy)

-	support Deno Dependency Viewer. close #83(1b327b8) (thanks @axetroy)

### Code Refactoring:

-	improve diagnostics(fe1682f) (thanks @axetroy)
-	rename function name(3f8d578) (thanks @axetroy)
-	normalize file path(d8abdf3) (thanks @axetroy)
-	improve start server logic(c0e2300) (thanks @axetroy)

### Testing:

-	write test case for deno_deps(df34869) (thanks @axetroy)
-	improve test case(4cfe7b8) (thanks @axetroy)
-	improve test(633ae31) (thanks @axetroy)
-	improve test case(7483cb7) (thanks @axetroy)
-	improve test case(22c8bce) (thanks @axetroy)
-	improve test case(64bf615) (thanks @axetroy)
-	add test case(a19c04b) (thanks @axetroy)
-	update test case(ba54a38) (thanks @axetroy)
-	add more test for util(465fa90) (thanks @axetroy)
-	fix test case(5291704) (thanks @axetroy)
-	add test case for file_walker(98aed6a) (thanks @axetroy)

### Performance improves:

-	improve performance for file_walker(77ce898) (thanks @axetroy)

### Documentation:

-	update changelog(1772477) (thanks @axetroy)

### Commits(25):

-	**1772477** docs: update changelog
-	**fe1682f** refactor: improve diagnostics
-	**3f8d578** refactor: rename function name
-	**df34869** test: write test case for deno_deps
-	**4cfe7b8** test: improve test case
-	**633ae31** test: improve test
-	**7483cb7** test: improve test case
-	**7249b3c** chore: update ci
-	**22c8bce** test: improve test case
-	**64bf615** test: improve test case
-	**a19c04b** test: add test case
-	**ba54a38** test: update test case
-	**465fa90** test: add more test for util
-	**4c52827** update
-	**cca6093** chore: try fix ci
-	**d8abdf3** refactor: normalize file path
-	**64e6c6c** chore: rewrite file_walker
-	**5291704** test: fix test case
-	**b0321ff** chore: update ci
-	**98aed6a** test: add test case for file_walker
-	**77ce898** perf: improve performance for file_walker
-	**c0e2300** refactor: improve start server logic
-	**3bb70e2** feat: improve fetching module message
-	**5441f6e** [ImgBot] Optimize images (#88)
-	**1b327b8** feat: support Deno Dependency Viewer. close #83

	v3.1.1
	======

	### Bugs fixed:

-	auto-import rewrite not work on Windows(3602979) (thanks @axetroy)

### Documentation:

-	update changelog(d68d065) (thanks @axetroy)

### Commits(2):

-	**d68d065** docs: update changelog
-	**3602979** fix: auto-import rewrite not work on Windows

	v3.1.0
	======

	### New feature:

-	add CodeLens for deno cached module which will show current URL(44bc4a1) (thanks @axetroy)

-	improve auto-import completion detail(78fa0e8) (thanks @axetroy)

### Bugs fixed:

-	normalize filepath(f5ecd71) (thanks @axetroy)

### Code Refactoring:

-	extra common code to util(5316648) (thanks @axetroy)
-	extra common code to util(b97a2cf) (thanks @axetroy)
-	update(8f79348) (thanks @axetroy)

### Documentation:

-	update changelog(521ca52) (thanks @axetroy)

### Commits(8):

-	**521ca52** docs: update changelog
-	**d2e2642** chore: fix ci
-	**f5ecd71** fix: normalize filepath
-	**5316648** refactor: extra common code to util
-	**b97a2cf** refactor: extra common code to util
-	**44bc4a1** feat: add CodeLens for deno cached module which will show current URL
-	**8f79348** refactor: update
-	**78fa0e8** feat: improve auto-import completion detail

	v3.0.6
	======

	### Bugs fixed:

-	typescript server crash if create a new untitled typescript TextDocument ref: #86(e5643e1) (thanks @axetroy)

### Code Refactoring:

-	improve code(002da57) (thanks @axetroy)

### Documentation:

-	update changelog(74c1954) (thanks @axetroy)

### Commits(5):

-	**74c1954** docs: update changelog
-	**e5643e1** fix: typescript server crash if create a new untitled typescript TextDocument ref: #86
-	**002da57** refactor: improve code
-	**edc034b** chore: add more information when assert
-	**c7df99b** chore(deps): update typescript-eslint monorepo to v2.22.0 (#85)

	v3.0.5
	======

	### Bugs fixed:

-	extension not work when project has tsconfig.json at root dir(9ce2874) (thanks @axetroy)

### Documentation:

-	update readme and changelog(4517b41) (thanks @axetroy)

### Commits(3):

-	**4517b41** docs: update readme and changelog
-	**9ce2874** fix: extension not work when project has tsconfig.json at root dir
-	**0fa2794** chore: update ci for test

	v3.0.4
	======

	### Bugs fixed:

-	invalid http tester regular expression(3d51ab0) (thanks @axetroy)

-	'fetch module' on work correctly for importmap module(087d834) (thanks @axetroy)

### Code Refactoring:

-	improve get type of hash_meta(b3b860b) (thanks @axetroy)
-	improve code(180bde1) (thanks @axetroy)

### Documentation:

-	update readme(310f488) (thanks @axetroy)
-	update changelog(e6d6e0a) (thanks @axetroy)

### Commits(9):

-	**310f488** docs: update readme
-	**e6d6e0a** docs: update changelog
-	**7fe19fe** chore: Extract common code `hashURL` util function
-	**713fa4c** chore: Extract common code `isHttpURL` util function
-	**4a99bc5** chore: Extract common code `isHttpURL` util function
-	**3d51ab0** fix: invalid http tester regular expression
-	**087d834** fix: 'fetch module' on work correctly for importmap module
-	**b3b860b** refactor: improve get type of hash_meta
-	**180bde1** refactor: improve code

	v3.0.3
	======

	### Bugs fixed:

-	Path to module not resolved correctly in Windows(926896d) (thanks @axetroy)

-	can not set TextDocument's language mode correctly in Windows(83d6e34) (thanks @axetroy)

-	somethine server does not ready and send notify(271c9cd) (thanks @axetroy)

### Code Refactoring:

-	improve var name(9cb72a8) (thanks @axetroy)
-	update(5e54d4e) (thanks @axetroy)
-	remove unnecessary await(d3e5654) (thanks @axetroy)
-	remove unnecessary await(e1ab675) (thanks @axetroy)
-	typescript-deno-plugin(6170e21) (thanks @axetroy)

### Documentation:

-	update changelog(b50908b) (thanks @axetroy)
-	update readme(360c19b) (thanks @axetroy)
-	Adding note about tsconfig.json (#82)(5fd5806) (thanks @Yusuke Sakurai)

### Commits(19):

-	**b50908b** docs: update changelog
-	**2cc71f9** chore: re-enable format test in Windows
-	**9cb72a8** refactor: improve var name
-	**926896d** fix: Path to module not resolved correctly in Windows
-	**67a0ec3** chore: add logger fore core module
-	**bb62e84** chore: fix ci
-	**83d6e34** fix: can not set TextDocument's language mode correctly in Windows
-	**5e54d4e** refactor: update
-	**453a9b0** Revert "refactor: typescript-deno-plugin"
-	**d3e5654** refactor: remove unnecessary await
-	**e1ab675** refactor: remove unnecessary await
-	**360c19b** docs: update readme
-	**225790e** chore: update deno and node version in ci env
-	**5fd5806** docs: Adding note about tsconfig.json (#82)
-	**a859b3f** chore: update ci
-	**42ca04c** chore: update ci
-	**1c2665f** chore: remove unused scripts
-	**271c9cd** fix: somethine server does not ready and send notify
-	**6170e21** refactor: typescript-deno-plugin

	v3.0.2
	======

	### Bugs fixed:

-	parse .vscode/settings.json with json5(b3de3d3) (thanks @axetroy)

### Documentation:

-	update changelog(1ab3532) (thanks @axetroy)
-	add readme for typescript-deno-plugin(da68d42) (thanks @axetroy)

### Commits(5):

-	**1ab3532** docs: update changelog
-	**30247f0** chore(deps): pin dependency @types/json5 to 0.0.30 (#81)
-	**b3de3d3** fix: parse .vscode/settings.json with json5
-	**1be0c24** Revert "docs: add readme for typescript-deno-plugin"
-	**da68d42** docs: add readme for typescript-deno-plugin

	v3.0.1
	======

	### New feature:

-	re-enable typescript-deno-plugin with workspace's typescript version. close #78(7a53e70) (thanks @axetroy)

### Documentation:

-	update changelog(6d856ee) (thanks @axetroy)

### Commits(2):

-	**6d856ee** docs: update changelog
-	**7a53e70** feat: re-enable typescript-deno-plugin with workspace's typescript version. close #78

	v3.0.0
	======

	### New feature:

-	Resurrected in Deno v0.35.0 🚀(3aff7ed) (thanks @axetroy)

-	Requires minimum version of Deno 0.35.0(35b810c) (thanks @axetroy)

-	set the language of the document automatically(8839799) (thanks @axetroy)

-	adapt Deno new cache layout(1cba5df) (thanks @axetroy)

-	support format range code(498d37f) (thanks @axetroy)

-	add import map file jso validator(2ccfa02) (thanks @axetroy)

-	Add memory cache module(c6cd7e8) (thanks @axetroy)

-	auto detect ./vscode/settings.json in typescript plugin #60(95d73c6) (thanks @axetroy)

-	support Import-Maps in new cache layout.(0c4cccd) (thanks @axetroy)

-	make manifest be a iterator.(8c7b7ce) (thanks @axetroy)

-	add new Deno's cache layout parser and test(7308978) (thanks @axetroy)

-	support @deno-types hint definition. #68(758c5be) (thanks @axetroy)

### Bugs fixed:

-	module not follow redirect(dc97014) (thanks @axetroy)
-	Possible errors caused by optional parameters(f91085c) (thanks @axetroy)
-	auto-import in new cache layout(7726fde) (thanks @axetroy)
-	get file hash without query(d81850f) (thanks @axetroy)
-	fix windows url path handle(b50548e) (thanks @axetroy)
-	fix invalid regexp for Windows(d37c846) (thanks @axetroy)
-	fix invalid regexp for Windows(a96e93a) (thanks @axetroy)
-	path resolution of Windows(a9e3336) (thanks @axetroy)
-	**deps**: update dependency vscode-languageserver-textdocument to v1.0.1 (#66)(c49b0fa) (thanks @renovate[bot])

### Code Refactoring:

-	improve code(89b7d4c) (thanks @axetroy)
-	improve code(3e560f2) (thanks @axetroy)
-	improve code(80a47f6) (thanks @axetroy)
-	add normalizeImportStatement and test case(e53af86) (thanks @axetroy)
-	improve code(2dd0bdf) (thanks @axetroy)
-	make sure invalid url will not throw error(06cd9d1) (thanks @axetroy)
-	remove unused code(cc8318c) (thanks @axetroy)
-	move common code to core(ef0e24a) (thanks @axetroy)
-	update code struct(42c65ab) (thanks @axetroy)
-	Extract common code(9568757) (thanks @axetroy)
-	import map(8b21c1d) (thanks @axetroy)
-	rename function name(e4493f7) (thanks @axetroy)
-	improve Deno Language Server(a472547) (thanks @axetroy)
-	improve diagnostics(cc20091) (thanks @axetroy)
-	improve path handler(b2927a7) (thanks @axetroy)
-	improve filepath handler(2223c40) (thanks @axetroy)

### Testing:

-	add test for module_resolver if redirect(68026ad) (thanks @axetroy)
-	add more test for module_resolver(0e6fb19) (thanks @axetroy)
-	add more test for hash_meta(429e9b2) (thanks @axetroy)
-	make sure deno_dir(5553b78) (thanks @axetroy)
-	add cache_map and test case(9663a63) (thanks @axetroy)
-	add deno_compile_hint test case(03a808b) (thanks @axetroy)
-	add more test for import_map(57647d9) (thanks @axetroy)
-	generate coverage(6d842b6) (thanks @axetroy)
-	add test case for remote module with query(4c05138) (thanks @axetroy)
-	remove unused test case(f69a6d1) (thanks @axetroy)
-	fix test(6206477) (thanks @axetroy)
-	add module resolver unit test(97bbd37) (thanks @axetroy)
-	add testcase for deno module(3c50e4c) (thanks @axetroy)
-	add getDenoDeps testcase(af59922) (thanks @axetroy)
-	add test case for manifest module(27ca82c) (thanks @axetroy)
-	add unit test for client/utils.ts(96f5c0a) (thanks @axetroy)
-	disable format test(b71ba68) (thanks @axetroy)
-	update(3f31927) (thanks @axetroy)
-	add unit test for deno class. #39(b85baf9) (thanks @axetroy)

### Documentation:

-	update changelog(85b4c4f) (thanks @axetroy)
-	update readme(85ffa3c) (thanks @axetroy)
-	update readme(8e4322f) (thanks @axetroy)
-	update readme(10d4223) (thanks @axetroy)
-	update(73d4dc7) (thanks @axetroy)

### Commits(81):

-	**85b4c4f** docs: update changelog
-	**3aff7ed** feat: Resurrected in Deno v0.35.0 🚀
-	**85ffa3c** docs: update readme
-	**8e4322f** docs: update readme
-	**68026ad** test: add test for module_resolver if redirect
-	**0e6fb19** test: add more test for module_resolver
-	**429e9b2** test: add more test for hash_meta
-	**35b810c** feat: Requires minimum version of Deno 0.35.0
-	**dc97014** fix: module not follow redirect
-	**8839799** feat: set the language of the document automatically
-	**1cba5df** feat: adapt Deno new cache layout
-	**498d37f** feat: support format range code
-	**61bce12** chore(deps): update dependency @types/node to v12.12.29 (#75)
-	**0fe3ac0** chore(deps): update dependency typescript to v3.8.3 (#76)
-	**2ccfa02** feat: add import map file jso validator
-	**89b7d4c** refactor: improve code
-	**f91085c** fix: Possible errors caused by optional parameters
-	**c36d2b7** chore(deps): pin dependencies (#73)
-	**5553b78** test: make sure deno_dir
-	**9663a63** test: add cache_map and test case
-	**3e560f2** refactor: improve code
-	**f499a13** chore: lint before commit
-	**80a47f6** refactor: improve code
-	**fb16e83** chore(deps): pin dependencies (#72)
-	**d5cfbdd** chore: add eslint
-	**894e978** chore: update tsconfig.json
-	**6161bf1** refacot: update normalizeImportStatement
-	**e53af86** refactor: add normalizeImportStatement and test case
-	**7726fde** fix: auto-import in new cache layout
-	**6ec1cb4** pref: cache completion. it should improve performance.
-	**c6cd7e8** feat: Add memory cache module
-	**2dd0bdf** refactor: improve code
-	**03a808b** test: add deno_compile_hint test case
-	**95d73c6** feat: auto detect ./vscode/settings.json in typescript plugin #60
-	**57647d9** test: add more test for import_map
-	**f837565** chore: update .vscodeignore
-	**8299544** chore: add build script
-	**10d4223** docs: update readme
-	**7293927** chore: update ci
-	**b130d13** chore(deps): pin dependency coveralls to 3.0.9 (#71)
-	**22b06c3** chore: fix ci
-	**0e84d6f** chore: report test coverage to coveralls
-	**6d842b6** test: generate coverage
-	**4c05138** test: add test case for remote module with query
-	**f69a6d1** test: remove unused test case
-	**06cd9d1** refactor: make sure invalid url will not throw error
-	**d81850f** fix: get file hash without query
-	**cc8318c** refactor: remove unused code
-	**6206477** test: fix test
-	**05e02a2** chore: improve typescript config
-	**ef0e24a** refactor: move common code to core
-	**0c4cccd** feat: support Import-Maps in new cache layout.
-	**97bbd37** test: add module resolver unit test
-	**3c50e4c** test: add testcase for deno module
-	**af59922** test: add getDenoDeps testcase
-	**8c7b7ce** feat: make manifest be a iterator.
-	**27ca82c** test: add test case for manifest module
-	**b50548e** fix: fix windows url path handle
-	**d37c846** fix: fix invalid regexp for Windows
-	**a96e93a** fix: fix invalid regexp for Windows
-	**7308978** feat: add new Deno's cache layout parser and test
-	**42c65ab** refactor: update code struct
-	**27a01db** chore(deps): pin dependencies (#69)
-	**9568757** refactor: Extract common code
-	**8b21c1d** refactor: import map
-	**e4493f7** refactor: rename function name
-	**96f5c0a** test: add unit test for client/utils.ts
-	**a9e3336** fix: path resolution of Windows
-	**b71ba68** test: disable format test
-	**3f31927** test: update
-	**c4da618** chore: setup deno in ci
-	**b85baf9** test: add unit test for deno class. #39
-	**a472547** refactor: improve Deno Language Server
-	**758c5be** feat: support @deno-types hint definition. #68
-	**cc20091** refactor: improve diagnostics
-	**c49b0fa** fix(deps): update dependency vscode-languageserver-textdocument to v1.0.1 (#66)
-	**811497b** chore(deps): update dependency typescript to v3.8.2 (#64)
-	**73d4dc7** docs: update
-	**b2927a7** refactor: improve path handler
-	**2223c40** refactor: improve filepath handler
-	**e743f3b** chore(deps): update dependency @types/node to v12.12.28 (#63)

	v2.0.4
	======

	### Bugs fixed:

-	Try to fix the path processing under windows. ref: #61(e3d5bf2) (thanks @axetroy)

### Documentation:

-	update changelog(a650f19) (thanks @axetroy)

### Commits(2):

-	**a650f19** docs: update changelog
-	**e3d5bf2** fix: Try to fix the path processing under windows. ref: #61

	v2.0.3
	======

	### Bugs fixed:

-	Try to fix the path processing under windows. ref: #61(8c02221) (thanks @axetroy)

### Documentation:

-	update changelog(eeab8fb) (thanks @axetroy)

### Commits(2):

-	**eeab8fb** docs: update changelog
-	**8c02221** fix: Try to fix the path processing under windows. ref: #61

	v2.0.2
	======

	### Bugs fixed:

-	Auto-Import for Deno module incorrectly. now use http protocol modules instead of relative paths. close #44(df71fd1) (thanks @axetroy)

-	typescript-deno-plugin will be disable when open the file out of workspace.(b0f3aa6) (thanks @axetroy)

### Code Refactoring:

-	enable strict mode(a05a777) (thanks @axetroy)

### Documentation:

-	update changelog(5cb1f61) (thanks @axetroy)

### Commits(4):

-	**5cb1f61** docs: update changelog
-	**df71fd1** fix: Auto-Import for Deno module incorrectly. now use http protocol modules instead of relative paths. close #44
-	**b0f3aa6** fix: typescript-deno-plugin will be disable when open the file out of workspace.
-	**a05a777** refactor: enable strict mode

	v2.0.1
	======

	### New feature:

-	improve fetch module with progress(1eaeb41) (thanks @axetroy)

### Bugs fixed:

-	improve typescript-deno-plugin. Make it as unaffected as possible.(a6ad52f) (thanks @axetroy)
-	**deps**: update dependency vscode-languageclient to v6.1.1 (#58)(ae547f9) (thanks @renovate[bot])
-	**deps**: update dependency vscode-languageserver to v6.1.1 (#57)(e687f20) (thanks @renovate[bot])
-	import module from 'file:///path/to/module/mod.ts' not work(962411d) (thanks @axetroy)

### Code Refactoring:

-	improve code(5ad9851) (thanks @axetroy)
-	improve code(4ef9673) (thanks @axetroy)
-	improve code(57703f4) (thanks @axetroy)
-	update(3b58dbd) (thanks @axetroy)

### Documentation:

-	update changelog(9eeac21) (thanks @axetroy)
-	update readme(d48f7de) (thanks @axetroy)
-	update readme(eee5114) (thanks @axetroy)
-	update(41d6e73) (thanks @axetroy)
-	update readme(4592767) (thanks @axetroy)

### Commits(16):

-	**9eeac21** docs: update changelog
-	**4a58f4e** chore(deps): pin dependency @types/deep-equal to 1.0.1 (#59)
-	**f1ef40f** chore: fix ci
-	**a6ad52f** fix: improve typescript-deno-plugin. Make it as unaffected as possible.
-	**ae547f9** fix(deps): update dependency vscode-languageclient to v6.1.1 (#58)
-	**e687f20** fix(deps): update dependency vscode-languageserver to v6.1.1 (#57)
-	**1eaeb41** feat: improve fetch module with progress
-	**5ad9851** refactor: improve code
-	**d48f7de** docs: update readme
-	**eee5114** docs: update readme
-	**41d6e73** docs: update
-	**962411d** fix: import module from 'file:///path/to/module/mod.ts' not work
-	**4ef9673** refactor: improve code
-	**4592767** docs: update readme
-	**57703f4** refactor: improve code
-	**3b58dbd** refactor: update

	v2.0.0
	======

	### New feature:

-	Deno minimum required v0.33.0(014192a) (thanks @axetroy)

-	rename configuration `deno.dtsFilepaths` to `deno.dts_file` (#49)(555a230) (thanks @Axetroy)

-	remove `deno.enable` & `deno.disable` command (#48)(8ecae2c) (thanks @Axetroy)

-	upgrade Deno formatter (#50)(e872d1c) (thanks @Axetroy)

### Documentation:

-	update changelog(ee36bd6) (thanks @axetroy)
-	update readme(ff14377) (thanks @axetroy)
-	update readme(865503c) (thanks @axetroy)

### Commits(7):

-	**ee36bd6** docs: update changelog
-	**ff14377** docs: update readme
-	**865503c** docs: update readme
-	**014192a** feat: Deno minimum required v0.33.0
-	**555a230** feat: rename configuration `deno.dtsFilepaths` to `deno.dts_file` (#49)
-	**8ecae2c** feat: remove `deno.enable` & `deno.disable` command (#48)
-	**e872d1c** feat: upgrade Deno formatter (#50)

	v1.23.0
	=======

	### New feature:

-	add the tips for Deno's minimum version for this extension.(8b5c54b) (thanks @axetroy)

-	Now opening the js file will also launch the extension. the same with tsserver.(d4a9beb) (thanks @axetroy)

-	support external type definitions with `X-TypeScript-Types` headers. close #35(98253dd) (thanks @axetroy)

### Code Refactoring:

-	update DENO_DIR(541456f) (thanks @axetroy)

### Documentation:

-	update changelog(84ab4bd) (thanks @axetroy)
-	update README.md(46a4798) (thanks @axetroy)

### Commits(10):

-	**84ab4bd** docs: update changelog
-	**0194ff0** chore(deps): pin dependencies (#54)
-	**46a4798** docs: update README.md
-	**8b5c54b** feat: add the tips for Deno's minimum version for this extension.
-	**d4a9beb** feat: Now opening the js file will also launch the extension. the same with tsserver.
-	**98253dd** feat: support external type definitions with `X-TypeScript-Types` headers. close #35
-	**d9237c9** chore(deps): update dependency husky to v4.2.3 (#53)
-	**9ce54b8** chore(deps): update dependency husky to v4.2.2 (#52)
-	**c8cc1cb** chore(deps): update dependency @types/node to v12.12.27 (#51)
-	**541456f** refactor: update DENO_DIR

	v1.22.0
	=======

	### New feature:

-	Add translations for dutch and german (#42)(ed2b7a4) (thanks @Luca Casonato)

-	improve module import intelligent(faf76c9) (thanks @axetroy)

### Bugs fixed:

-	Module index is incorrect. close #47(d69e90a) (thanks @axetroy)
-	module import intelligent no work correctly when import from 'http/server.ts'(055d062) (thanks @axetroy)

### Code Refactoring:

-	update(4098c00) (thanks @axetroy)

### Documentation:

-	update changelog(b98fe61) (thanks @axetroy)
-	update changelog(5a31ba3) (thanks @axetroy)
-	Improve grammar and sentence structure (#46)(9d709e3) (thanks @David)

### Commits(8):

-	**b98fe61** docs: update changelog
-	**5a31ba3** docs: update changelog
-	**d69e90a** fix: Module index is incorrect. close #47
-	**ed2b7a4** feat: Add translations for dutch and german (#42)
-	**9d709e3** docs: Improve grammar and sentence structure (#46)
-	**055d062** fix: module import intelligent no work correctly when import from 'http/server.ts'
-	**4098c00** refactor: update
-	**faf76c9** feat: improve module import intelligent

	v1.21.0
	=======

	### New feature:

-	support external type definitions with '/// <reference types=https://raw.githubusercontent.com/date-fns/date-fns/master/typings.d.ts />'. ref: #35(f7affb2) (thanks @axetroy)

### Documentation:

-	update changelog(7f4509e) (thanks @axetroy)

### Commits(6):

-	**7f4509e** docs: update changelog
-	**d46a3c2** chore: try ignore format check in windows
-	**86ce29f** chore: update ci
-	**0d197e6** chore(deps): pin dependencies (#45)
-	**e9b08a9** chore: remove Deno formatter before commit. use prettier instead. close #43
-	**f7affb2** feat: support external type definitions with '/// <reference types=https://raw.githubusercontent.com/date-fns/date-fns/master/typings.d.ts />'. ref: #35

	v1.20.0
	=======

	### New feature:

-	remove `lock std version` and `prefer HTTPS` diagnostics. close #33(2480791) (thanks @axetroy)

### Bugs fixed:

-	update ignore diagnostics code. close #41(34e6c10) (thanks @axetroy)

### Code Refactoring:

-	rewrite typescript-deno-plugin(8c4cc59) (thanks @axetroy)
-	reformat code(499b119) (thanks @axetroy)

### Documentation:

-	update changelog(d054f27) (thanks @axetroy)
-	add Github issue template(77a6740) (thanks @axetroy)
-	Fix wrong configuration file path(07ef9ae) (thanks @axetroy)

### Commits(7):

-	**d054f27** docs: update changelog
-	**2480791** feat: remove `lock std version` and `prefer HTTPS` diagnostics. close #33
-	**34e6c10** fix: update ignore diagnostics code. close #41
-	**8c4cc59** refactor: rewrite typescript-deno-plugin
-	**499b119** refactor: reformat code
-	**77a6740** docs: add Github issue template
-	**07ef9ae** docs: Fix wrong configuration file path

	v1.19.0
	=======

	### New feature:

-	remove extension name diagnostic. close #12(892bb3f) (thanks @axetroy)

-	support import ECMA script module. close #37(1b68068) (thanks @axetroy)

### Bugs fixed:

-	esm module resolver(ffe30fb) (thanks @axetroy)

### Documentation:

-	update changelog(56f5096) (thanks @axetroy)
-	add screenshot(a7dffc8) (thanks @axetroy)

### Commits(5):

-	**56f5096** docs: update changelog
-	**a7dffc8** docs: add screenshot
-	**ffe30fb** fix: esm module resolver
-	**892bb3f** feat: remove extension name diagnostic. close #12
-	**1b68068** feat: support import ECMA script module. close #37

	v1.18.1
	=======

### Commits(1):

-	**341165e** Revert "feat: support top-level await with typescript 3.8"

	v1.18.0
	=======

	### New feature:

-	no more use workspace typescript version(2a6f9da) (thanks @axetroy)

-	require min vscode version 1.42.0(ab2cc6e) (thanks @axetroy)

-	support top-level await with typescript 3.8(cb0e592) (thanks @axetroy)

### Bugs fixed:

-	create local module no work(bcceff2) (thanks @axetroy)

### Code Refactoring:

-	improve typescript plugin(9b70328) (thanks @axetroy)
-	Remove redundant code(75fd7f2) (thanks @axetroy)
-	refactor typescript plugin(46d9450) (thanks @axetroy)

### Documentation:

-	update CHANGELOG.md(d35c922) (thanks @axetroy)

### Commits(8):

-	**d35c922** docs: update CHANGELOG.md
-	**2a6f9da** feat: no more use workspace typescript version
-	**ab2cc6e** feat: require min vscode version 1.42.0
-	**cb0e592** feat: support top-level await with typescript 3.8
-	**bcceff2** fix: create local module no work
-	**9b70328** refactor: improve typescript plugin
-	**75fd7f2** refactor: Remove redundant code
-	**46d9450** refactor: refactor typescript plugin

	v1.17.0
	=======

	### New feature:

-	fully i18n supported. #31(04e3938) (thanks @axetroy)

### Bugs fixed:

-	create a local module if is not relative or absolute path(21bacce) (thanks @axetroy)

### Documentation:

-	update changelog(327782b) (thanks @axetroy)
-	update configuration description(b67ff06) (thanks @axetroy)
-	update README.md(a2dfbdb) (thanks @axetroy)
-	add screenshot(92cac97) (thanks @axetroy)
-	update README.md(14cbcf8) (thanks @axetroy)
-	update README.md(7b9267c) (thanks @axetroy)

### Commits(9):

-	**327782b** docs: update changelog
-	**21bacce** fix: create a local module if is not relative or absolute path
-	**04e3938** feat: fully i18n supported. #31
-	**b67ff06** docs: update configuration description
-	**a2dfbdb** docs: update README.md
-	**7f04f64** [ImgBot] Optimize images (#32)
-	**92cac97** docs: add screenshot
-	**14cbcf8** docs: update README.md
-	**7b9267c** docs: update README.md

	v1.16.0
	=======

	### New feature:

-	support Import Maps for Deno. close #3(eb187af) (thanks @axetroy)

-	add lock deno_std version diagnostic(8d9097e) (thanks @axetroy)

-	add default content for creating a file when fix missing local module(1404f2f) (thanks @axetroy)

### Code Refactoring:

-	refactor code(a3c1828) (thanks @axetroy)

### Documentation:

-	update CHANGELOG.md(ef03c70) (thanks @axetroy)

### Commits(5):

-	**ef03c70** docs: update CHANGELOG.md
-	**eb187af** feat: support Import Maps for Deno. close #3
-	**8d9097e** feat: add lock deno_std version diagnostic
-	**1404f2f** feat: add default content for creating a file when fix missing local module
-	**a3c1828** refactor: refactor code

	v1.15.0
	=======

	### New feature:

-	support quickly fix for diagnostics. close #29(da85926) (thanks @axetroy)

### Bugs fixed:

-	**deps**: pin dependency execa to 4.0.0 (#30)(47ca6e4) (thanks @renovate[bot])
-	`typescript-deno-plugin` may not find modules and cause `typescript` to crash(8bdc5db) (thanks @axetroy)

### Documentation:

-	update CHANGELOG.md(f059bc3) (thanks @axetroy)
-	update README.md(b6e6d76) (thanks @axetroy)

### Commits(5):

-	**f059bc3** docs: update CHANGELOG.md
-	**b6e6d76** docs: update README.md
-	**47ca6e4** fix(deps): pin dependency execa to 4.0.0 (#30)
-	**da85926** feat: support quickly fix for diagnostics. close #29
-	**8bdc5db** fix: `typescript-deno-plugin` may not find modules and cause `typescript` to crash

	v1.14.0
	=======

	### New feature:

-	Added i18n support for Chinese Traditional(ca93cd2) (thanks @axetroy)

-	add `deno.restart_server` command to restart `Deno Language Server`. close #28(9a66f86) (thanks @axetroy)

-	improve status bar. show more information(6fb83c4) (thanks @axetroy)

### Bugs fixed:

-	lock prettier version to make sure formatter work on deno v0.32.0. We will switch to dprint in a future release and only suppport formatting typescript/javascipt code.(78b3266) (thanks @axetroy)

### Code Refactoring:

-	reformat code(1f1cd14) (thanks @axetroy)

### Documentation:

-	update changelog(4f14872) (thanks @axetroy)
-	update README.md(e8e13d9) (thanks @axetroy)
-	update screenshot.gif(7184c9b) (thanks @axetroy)
-	update README.md(aa9cb4a) (thanks @axetroy)

### Commits(9):

-	**4f14872** docs: update changelog
-	**ca93cd2** feat: Added i18n support for Chinese Traditional
-	**e8e13d9** docs: update README.md
-	**9a66f86** feat: add `deno.restart_server` command to restart `Deno Language Server`. close #28
-	**6fb83c4** feat: improve status bar. show more information
-	**1f1cd14** refactor: reformat code
-	**78b3266** fix: lock prettier version to make sure formatter work on deno v0.32.0. We will switch to dprint in a future release and only suppport formatting typescript/javascipt code.
-	**7184c9b** docs: update screenshot.gif
-	**aa9cb4a** docs: update README.md

	v1.13.1
	=======

	### Bugs fixed:

-	cannot find module if redirected. close #27(6fd7b13) (thanks @axetroy)

### Documentation:

-	update changelog(377ad30) (thanks @axetroy)

### Commits(2):

-	**377ad30** docs: update changelog
-	**6fd7b13** fix: cannot find module if redirected. close #27

	v1.13.0
	=======

	### New feature:

-	improve diagnostics(a5f029e) (thanks @axetroy)

### Bugs fixed:

-	can not import module which end with `.ts`(0169107) (thanks @axetroy)
-	**deps**: pin dependency vscode-uri to 2.1.1 (#26)(5cdf757) (thanks @renovate[bot])
-	improve import module position(8a999c6) (thanks @axetroy)

### Code Refactoring:

-	Remove redundant code(f188e99) (thanks @axetroy)
-	remove utils.ts(c9f51c1) (thanks @axetroy)
-	diagnostics(a74a9ec) (thanks @axetroy)

### Documentation:

-	update changelog(ad95ce5) (thanks @axetroy)

### Commits(9):

-	**ad95ce5** docs: update changelog
-	**0169107** fix: can not import module which end with `.ts`
-	**5cdf757** fix(deps): pin dependency vscode-uri to 2.1.1 (#26)
-	**f188e99** refactor: Remove redundant code
-	**a5f029e** feat: improve diagnostics
-	**c9f51c1** refactor: remove utils.ts
-	**a74a9ec** refactor: diagnostics
-	**ab9f6a9** chore: update Deno version for Github Action
-	**8a999c6** fix: improve import module position

	v1.12.0
	=======

	### New feature:

-	improve folder picker(71e9658) (thanks @axetroy)

-	Warning when import from http(72d9db3) (thanks @axetroy)

-	remove `deno.enable = true` by default(532cdf0) (thanks @axetroy)

### Code Refactoring:

-	diagnostic(45b4267) (thanks @axetroy)
-	server(8e7bf88) (thanks @axetroy)
-	rewrite extension(916e5ca) (thanks @axetroy)
-	improve multiple workspace. close #23(4e9ab53) (thanks @axetroy)

### Documentation:

-	add CHANGELOG.md(e8e5004) (thanks @axetroy)
-	fix typo(b6cd461) (thanks @axetroy)
-	update README.md(27c2440) (thanks @axetroy)
-	update README.md(e4fb383) (thanks @axetroy)
-	update README.md(356530c) (thanks @axetroy)

### Commits(13):

-	**e8e5004** docs: add CHANGELOG.md
-	**b6cd461** docs: fix typo
-	**27c2440** docs: update README.md
-	**71e9658** feat: improve folder picker
-	**e6b10bf** refactor
-	**45b4267** refactor: diagnostic
-	**72d9db3** feat: Warning when import from http
-	**8e7bf88** refactor: server
-	**e4fb383** docs: update README.md
-	**356530c** docs: update README.md
-	**532cdf0** feat: remove `deno.enable = true` by default
-	**916e5ca** refactor: rewrite extension
-	**4e9ab53** refactor: improve multiple workspace. close #23

	v1.11.0
	=======

	### New feature:

-	add diagnostics checking for disable non-extension name module import. close #12(8c1c244) (thanks @axetroy)

### Bugs fixed:

-	**deps**: pin dependency typescript to 3.7.5 (#21)(3b3049c) (thanks @renovate[bot])
-	add missing typescript deps(751261a) (thanks @axetroy)
-	**deps**: pin dependency get-port to 5.1.1 (#18)(d3cf219) (thanks @renovate[bot])

### Code Refactoring:

-	improve sync workspace forlder(8bb4703) (thanks @axetroy)

### Documentation:

-	update README.md(2d9a32d) (thanks @axetroy)

### Commits(6):

-	**3b3049c** fix(deps): pin dependency typescript to 3.7.5 (#21)
-	**751261a** fix: add missing typescript deps
-	**8c1c244** feat: add diagnostics checking for disable non-extension name module import. close #12
-	**8bb4703** refactor: improve sync workspace forlder
-	**d3cf219** fix(deps): pin dependency get-port to 5.1.1 (#18)
-	**2d9a32d** docs: update README.md

	v1.10.1
	=======

	### Bugs fixed:

-	formatter not run at workspace folder(bf6195a) (thanks @axetroy)

### Commits(1):

-	**bf6195a** fix: formatter not run at workspace folder

	v1.10.0
	=======

	### Bugs fixed:

-	completion show everywhere(21741a2) (thanks @axetroy)

### Code Refactoring:

-	refactor code(a9760fe) (thanks @axetroy)
-	add freeport(749e181) (thanks @axetroy)
-	rewrite to C/S model with LSP (#15)(8674e32) (thanks @Axetroy)

### Commits(5):

-	**21741a2** fix: completion show everywhere
-	**a9760fe** refactor: refactor code
-	**749e181** refactor: add freeport
-	**bd30dba** chore(deps): pin dependencies (#16)
-	**8674e32** refactor: rewrite to C/S model with LSP (#15)

	v1.9.2
	======

	### Bugs fixed:

-	resolve can not import module not end with .ts when module does not found. close #5(1143a97) (thanks @axetroy)

### Commits(1):

-	**1143a97** fix: resolve can not import module not end with .ts when module does not found. close #5

	v1.9.1
	======

	### New feature:

-	support top-level await. close #10(d1cd97c) (thanks @axetroy)

### Documentation:

-	add screenshot.gif(7872947) (thanks @axetroy)

### Commits(6):

-	**d1cd97c** feat: support top-level await. close #10
-	**525abb1** chore: add renovate.json
-	**d65aad8** [ImgBot] Optimize images (#9)
-	**6dc5002** chore: update vscodeignore
-	**74ba03c** chore: ignore image for vscode package
-	**7872947** docs: add screenshot.gif

	v1.9.0
	======

	### New feature:

-	enable jsx options by default for typescript-deno-plugin(b9c2fba) (thanks @axetroy)

-	support import installed module intelligent. close #4(6d9baaa) (thanks @axetroy)

### Commits(2):

-	**b9c2fba** feat: enable jsx options by default for typescript-deno-plugin
-	**6d9baaa** feat: support import installed module intelligent. close #4

	v1.8.0
	======

	### New feature:

-	add deno.dtsFilepaths configuration(458666e) (thanks @axetroy)

### Bugs fixed:

-	only allow .d.ts file for deno.dtsFilepaths(8916695) (thanks @axetroy)

### Documentation:

-	update README.md(9a58fc5) (thanks @axetroy)
-	update README.md(06d4110) (thanks @axetroy)

### Commits(5):

-	**8916695** fix: only allow .d.ts file for deno.dtsFilepaths
-	**458666e** feat: add deno.dtsFilepaths configuration
-	**9a58fc5** docs: update README.md
-	**a0f2c99** remove deno.alwaysShowStatus configuration. close #6
-	**06d4110** docs: update README.md

	v1.7.0
	======

	### New feature:

-	enable/disable typescript-deno-plugin in extension scope(fc2c197) (thanks @axetroy)

### Commits(1):

-	**fc2c197** feat: enable/disable typescript-deno-plugin in extension scope

	v1.6.1
	======

	### Bugs fixed:

-	support import.meta.url for Deno(3a26287) (thanks @axetroy)

### Commits(1):

-	**3a26287** fix: support import.meta.url for Deno

	v1.6.0
	======

	### Bugs fixed:

-	try fix ci(e21e3f9) (thanks @axetroy)

-	use yarn package for vsce(67b2efd) (thanks @axetroy)

### Documentation:

-	update README.md(356a86c) (thanks @axetroy)
-	update readme(44ba16b) (thanks @axetroy)

### Commits(7):

-	**e21e3f9** fix: try fix ci
-	**a4c9625** chore: update package.json
-	**f4a9ae2** chore: add postscript for typescript-deno-plugin
-	**356a86c** docs: update README.md
-	**67b2efd** fix: use yarn package for vsce
-	**44ba16b** docs: update readme
-	**44f32c2** transfer to Axetroy

	v1.5.0
	======

### Commits(0):

v1.4.1
======

### Commits(1):

-	**d537df8** rewrite vscode-deno

	v1.4.0
	======

### Commits(5):

-	**8debe3a** suport vscode >= 1.14.0
-	**f1980d4** upgrade pkgs
-	**2565d1f** fix deno finder
-	**5cef835** upgrade ts to v3.7.2
-	**91046b7** chore: upgrade deno version in ci and lock std version

	v1.3.3
	======

### Commits(10):

-	**8880916** fix deno version for ci (#52)
-	**74d4dd0** upgrade typescript-deno-plugin@1.2.7 (#51)
-	**4e7a884** upgrade deps (#50)
-	**36eabed** Bump lodash from 4.17.11 to 4.17.14 (#45)
-	**a567c2d** fix format incorrect when set vscode config autosave=off (#43)
-	**9c321b9** add open_collective
-	**9539695** resolve an spelling mistake
-	**9c27335** remove GitHub Sponsors
-	**07516aa** Yarn upgrade (#39)
-	**de80a8d** Create FUNDING.yml

	v1.3.2
	======

	### New feature:

-	add format provider (#32)(9636ee2) (thanks @Axetroy)

### Commits(5):

-	**23078d6** upgrade typescript-deno-plugin to 1.2.5 (#37)
-	**3df7c4e** format code
-	**6a3355b** use `allowSyntheticDefaultImports`
-	**9636ee2** feat: add format provider (#32)
-	**e9b6b7b** Fix typo in extension.ts. (#30)

	v1.3.0
	======

### Commits(2):

-	**6fe2402** Upgrade deno v0.3.9 check (#28)
-	**098939d** Update LICENSE

	v1.2.1
	======

### Commits(0):

v1.2.0
======

### Commits(13):

-	**cc6a95c** Generate Deno's .d.ts file (#21)
-	**7fa51ec** foramt package.json
-	**d8605f6** upgrade typescript-deno-plugin
-	**2479cbf** use deno v0.3.7 format
-	**8d29205** Revert "implement auto format on save (#13)" (#19)
-	**6927e6a** implement auto format on save (#13)
-	**485132f** Merge branch 'master' of github.com:justjavac/vscode-deno
-	**26c5993** change pr template
-	**cec97af** add i18n(zh-cn)
-	**3e205c3** format code
-	**8a9c3af** use deno formater
-	**7e83e8b** Merge pull request #14 from axetroy/configurationSection
-	**7e81cdb** add badge for downloads

	v1.0.7
	======

### Commits(3):

-	**e1e7c25** Merge pull request #8 from axetroy/snippets
-	**c3e5aa0** Merge pull request #7 from axetroy/disposables
-	**66986b0** Merge pull request #6 from axetroy/master

	v1.0.6
	======

### Commits(1):

-	**2b451ff** Revert "bundling extension using webpack"

	v1.0.5
	======

### Commits(3):

-	**61286ba** add schemas & snippets
-	**99b13e8** typo
-	**b227930** update readme.md

	v1.0.4
	======

### Commits(1):

-	**fbeb470** Merge pull request #3 from justjavac/fix-warning-on-found

	v1.0.3
	======

### Commits(5):

-	**e381cd7** release v1.0.3
-	**3883c63** add ISSUE_TEMPLATEs
-	**0954ca6** add 'PULL_REQUEST_TEMPLATE'
-	**7073153** add 'How to contribute to vscode-deno'
-	**07877f9** add 'Contributor Covenant Code of Conduct'

	v1.0.2
	======

### Commits(4):

-	**bf9c0b5** add screenshots
-	**1752c24** add some ignores
-	**a310fff** fix link
-	**6fda196** add title

	v1.0.1
	======

### Commits(0):
