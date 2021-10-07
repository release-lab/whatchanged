v4.0.0 (2020-05-27)
-------------------

### 🔥 New feature:

-	deprecated([`7729788`](https://github.com/axetroy/vscode-deno/commit/7729788c37ae23fe206ca218a0bbf8253063077c)) (@axetroy)

### 🔙 Revert:

-	revert [`bf08694`](https://github.com/axetroy/vscode-deno/commit/bf0869485d8e072fc6a6ebaeffbad67b5a0c1824), deprecated([`6d13713`](https://github.com/axetroy/vscode-deno/commit/6d1371328a843e0c937ea80ae81fb07400a9826a)\)
-	revert [`c20037d`](https://github.com/axetroy/vscode-deno/commit/c20037dd95a828230277f12ff2e8c4b73bd05425), refactor: remove unused extension([`25b190b`](https://github.com/axetroy/vscode-deno/commit/25b190bf790d930962ab63faa285b0b51cbca9b3)\)
-	revert [`cde4dda`](https://github.com/axetroy/vscode-deno/commit/cde4dda61a565370285cce8f14a3962c8b5d3f63), refactor: update([`479ecd2`](https://github.com/axetroy/vscode-deno/commit/479ecd2268b9b3edc087043cca6b4fce954959c0)\)

v3.7.0 (2020-05-05)
-------------------

### 🔥 New feature:

-	remove JSON import statement. Now, you cannot import JSON module.([`23ff6f3`](https://github.com/axetroy/vscode-deno/commit/23ff6f3fa6a5388273f2aeaedd8029b5d18b3e7d)) (@axetroy)
-	added unstable settings option (#167)([`1a4a230`](https://github.com/axetroy/vscode-deno/commit/1a4a230c57a670632c43f947305e7f5fa11d1531)) (@Luca Casonato)

v3.6.2 (2020-04-29)
-------------------

### 🐛 Bugs fixed:

-	file protocol import statement not work. close #146([`67897bc`](https://github.com/axetroy/vscode-deno/commit/67897bc762b1c1b2997b36292ca91eb777a3e9bf)) (@axetroy)
-	not deno project also show deno deps tree view.([`81303df`](https://github.com/axetroy/vscode-deno/commit/81303df0e2cc8a942d4a5cd4a7c157b7bebeac5e)) (@axetroy)

v3.6.1 (2020-04-17)
-------------------

### 🐛 Bugs fixed:

-	In Deno's cache module, `x-typescript-types` and redirects are not parsed correctly. close #147([`a3a957f`](https://github.com/axetroy/vscode-deno/commit/a3a957f617617e19ea69e0941f961b57589574fc)) (@axetroy)
-	**deps**: update dependency semver to v7.3.2 (#144)([`8496c84`](https://github.com/axetroy/vscode-deno/commit/8496c845e17d44e313b9e5139d7e993a865dd956)) (@renovate[bot])
-	**deps**: update dependency semver to v7.3.1 (#143)([`16a112b`](https://github.com/axetroy/vscode-deno/commit/16a112b69318b2ac1fd3cea6de8c4dcd75874e8e)) (@renovate[bot])

v3.6.0 (2020-04-14)
-------------------

### 🔥 New feature:

-	ignore typescript compile options which ignore by Deno([`b48fed0`](https://github.com/axetroy/vscode-deno/commit/b48fed01dcbd4d5d1bd08e6d060e5217c094b9e6)) (@axetroy)

### 🐛 Bugs fixed:

-	import map with trailing slash (#142)([`d5ecc7e`](https://github.com/axetroy/vscode-deno/commit/d5ecc7e2ead8df65fcbc30d1e1d08431b91b177e)) (@Chuah Chee Shian)
-	**deps**: update dependency semver to v7.3.0 (#140)([`13e3947`](https://github.com/axetroy/vscode-deno/commit/13e3947928f1999963c6e2166f5c3e57d61aa67d)) (@renovate[bot])
-	**deps**: update dependency semver to v7.2.2 (#137)([`d9bd9c8`](https://github.com/axetroy/vscode-deno/commit/d9bd9c8df4ffbe0c8446909717a3ce751c43f5eb)) (@renovate[bot])

v3.5.1 (2020-04-09)
-------------------

### 🐛 Bugs fixed:

-	add more test case for import_map. ref #132([`e4b1d6a`](https://github.com/axetroy/vscode-deno/commit/e4b1d6aee4f55f17f00ef204048e77285f5b5ee3)) (@axetroy)

v3.5.0 (2020-04-09)
-------------------

### 🔥 New feature:

-	compatible deno cache command since Deno v0.40.0([`7309b0c`](https://github.com/axetroy/vscode-deno/commit/7309b0c3b58bd1a2ed4fc5af65fa0bd3f4626336)) (@axetroy)
-	support deno-types compile hint (#129)([`9be33a4`](https://github.com/axetroy/vscode-deno/commit/9be33a41c8d94c49a1420d124dcdad74b7778ee9)) (@艾斯特洛)

### 🐛 Bugs fixed:

-	refresh diagnostic not work([`f8e8e70`](https://github.com/axetroy/vscode-deno/commit/f8e8e706f9ef2e45b78aafd5a7b28ddd03c80679)) (@axetroy)

v3.4.2 (2020-04-08)
-------------------

### 🐛 Bugs fixed:

-	if `x-typescript-types` do not exist. then fallback to origin file([`05496e3`](https://github.com/axetroy/vscode-deno/commit/05496e3022e371c5be7e821c9f778c8ed7cfbc77)) (@axetroy)
-	**deps**: update dependency semver to v7.2.1 (#128)([`90f4ad2`](https://github.com/axetroy/vscode-deno/commit/90f4ad2f94db9762f23fdd9b79d73572cf1db5e2)) (@renovate[bot])

v3.4.1 (2020-03-31)
-------------------

### 🐛 Bugs fixed:

-	Tsserver crashes in some cases([`11563b4`](https://github.com/axetroy/vscode-deno/commit/11563b4df4e84b0c92bc256f2be6bf0dd6fc9954)) (@axetroy)

v3.4.0 (2020-03-27)
-------------------

### 🔥 New feature:

-	add diagnostic for checking valid import statement([`b2f070a`](https://github.com/axetroy/vscode-deno/commit/b2f070a85058432542fe9721be2d3aa66e6901b3)) (@axetroy)

### 🐛 Bugs fixed:

-	importmap not work when set to a relative path. close #103([`0e8398f`](https://github.com/axetroy/vscode-deno/commit/0e8398f0892f1a705cff6c4f8721e7b3f71dc948)) (@axetroy)

v3.3.1 (2020-03-27)
-------------------

### 🐛 Bugs fixed:

-	Triple-Slash Directive does not work. ref #102([`e381390`](https://github.com/axetroy/vscode-deno/commit/e381390ac9461f9bdc7b67a9228e071d64d2ac3b)) (@axetroy)
-	output target of compileOption now is esnext([`fcbd234`](https://github.com/axetroy/vscode-deno/commit/fcbd234c5a426efac91fc7b4456b1ea63828fc3c)) (@axetroy)
-	Synchronization Deno error code([`4b6efd4`](https://github.com/axetroy/vscode-deno/commit/4b6efd41a0d5be3d2b681e0be8e20398a66cae92)) (@axetroy)
-	auto import module not working properly in some edge cases([`71b518b`](https://github.com/axetroy/vscode-deno/commit/71b518bf229546c4015763cfea3a30ab7739f7cc)) (@axetroy)
-	**deps**: update dependency vscode-languageclient to v6.1.3 (#101)([`f5a743c`](https://github.com/axetroy/vscode-deno/commit/f5a743c51e789f79540244fc5599291fca9741ab)) (@renovate[bot])
-	**deps**: update dependency vscode-languageclient to v6.1.2 (#99)([`550f0bc`](https://github.com/axetroy/vscode-deno/commit/550f0bcf21b9d9a83107c65367a1bd375ccbdb1e)) (@renovate[bot])

v3.3.0 (2020-03-16)
-------------------

### 🔥 New feature:

-	remove `deno.dts_file` configuration (#94)([`f06e852`](https://github.com/axetroy/vscode-deno/commit/f06e852b60883e7499a5c7976d026511bd6dd0ad)) (@艾斯特洛)
-	add copy to clipboard message([`4dd4b8d`](https://github.com/axetroy/vscode-deno/commit/4dd4b8dc2be31a4d0b628fac638981013284b36a)) (@axetroy)

### 🐛 Bugs fixed:

-	cannot resolve module if location headers is relative or absolute path. close #97([`75d6027`](https://github.com/axetroy/vscode-deno/commit/75d602755f81a0c22331e08b1c97fdf6782fd9ac)) (@axetroy)

v3.2.1 (2020-03-08)
-------------------

### 🔥 New feature:

-	copy url if click code_lens([`e3829d8`](https://github.com/axetroy/vscode-deno/commit/e3829d85c355866097f1b24f845e9accc1fd8904)) (@axetroy)
-	make Deno declaration file read-only([`fba8d89`](https://github.com/axetroy/vscode-deno/commit/fba8d89f70bbf42f0ea34eff4173ed0de00ef271)) (@axetroy)

### 🐛 Bugs fixed:

-	auto-import not work for some modules. close #44([`11d38b3`](https://github.com/axetroy/vscode-deno/commit/11d38b3f46b8521ec1ea89565f11ec65a9d2cb1d)) (@axetroy)
-	If query exists in url, module will not be parsed correctly([`a8965b5`](https://github.com/axetroy/vscode-deno/commit/a8965b5a5742b8ae473d3da85f8fe4829aee82ba)) (@axetroy)

v3.2.0 (2020-03-05)
-------------------

### 🔥 New feature:

-	improve fetching module message([`3bb70e2`](https://github.com/axetroy/vscode-deno/commit/3bb70e23f51750a3eab053289772f154740c3045)) (@axetroy)
-	support Deno Dependency Viewer. close #83([`1b327b8`](https://github.com/axetroy/vscode-deno/commit/1b327b869cf74d2022dab71451c92ec3c4b5c5ea)) (@axetroy)

### ⚡️ Performance improves:

-	improve performance for file_walker([`77ce898`](https://github.com/axetroy/vscode-deno/commit/77ce898b12c1e9ec9f995c1ae2b14342547f7cbd)) (@axetroy)

v3.1.1 (2020-03-03)
-------------------

### 🐛 Bugs fixed:

-	auto-import rewrite not work on Windows([`3602979`](https://github.com/axetroy/vscode-deno/commit/3602979225fac8a5ffb53e362ab473ea741a3bbf)) (@axetroy)

v3.1.0 (2020-03-03)
-------------------

### 🔥 New feature:

-	add CodeLens for deno cached module which will show current URL([`44bc4a1`](https://github.com/axetroy/vscode-deno/commit/44bc4a13638822e0ee482a2fba0f23c463e16820)) (@axetroy)
-	improve auto-import completion detail([`78fa0e8`](https://github.com/axetroy/vscode-deno/commit/78fa0e86ef017c30349f96794ff3b6eb53cd5879)) (@axetroy)

### 🐛 Bugs fixed:

-	normalize filepath([`f5ecd71`](https://github.com/axetroy/vscode-deno/commit/f5ecd71867f5ed6bf055d7606610c0079dc80065)) (@axetroy)

v3.0.6 (2020-03-03)
-------------------

### 🐛 Bugs fixed:

-	typescript server crash if create a new untitled typescript TextDocument ref: #86([`e5643e1`](https://github.com/axetroy/vscode-deno/commit/e5643e159042a72ce1871061ff5038be7b6cebb5)) (@axetroy)

v3.0.5 (2020-03-03)
-------------------

### 🐛 Bugs fixed:

-	extension not work when project has tsconfig.json at root dir([`9ce2874`](https://github.com/axetroy/vscode-deno/commit/9ce2874230d4c66ea657f5d2de19c38eb8719df6)) (@axetroy)

v3.0.4 (2020-03-02)
-------------------

### 🐛 Bugs fixed:

-	invalid http tester regular expression([`3d51ab0`](https://github.com/axetroy/vscode-deno/commit/3d51ab04359ad8bd83b5564c144759d08f9d0237)) (@axetroy)
-	'fetch module' on work correctly for importmap module([`087d834`](https://github.com/axetroy/vscode-deno/commit/087d8345ca3a717d55822dce6ab64c6d9385e790)) (@axetroy)

v3.0.3 (2020-03-02)
-------------------

### 🐛 Bugs fixed:

-	Path to module not resolved correctly in Windows([`926896d`](https://github.com/axetroy/vscode-deno/commit/926896d33b88971d66f910cf12f881ae5f38405d)) (@axetroy)
-	can not set TextDocument's language mode correctly in Windows([`83d6e34`](https://github.com/axetroy/vscode-deno/commit/83d6e342d1795c15e1fdba9218bbcddae910e02e)) (@axetroy)
-	somethine server does not ready and send notify([`271c9cd`](https://github.com/axetroy/vscode-deno/commit/271c9cdda40556f8b8efe8280a7a323074d62544)) (@axetroy)

### 🔙 Revert:

-	revert [`6170e21`](https://github.com/axetroy/vscode-deno/commit/6170e21417754dcca33bf1f9376dec37251ddc06), refactor: typescript-deno-plugin([`453a9b0`](https://github.com/axetroy/vscode-deno/commit/453a9b0240bda3d2521defa5f4dbf74faeeca7a4)\)

v3.0.2 (2020-03-01)
-------------------

### 🐛 Bugs fixed:

-	parse .vscode/settings.json with json5([`b3de3d3`](https://github.com/axetroy/vscode-deno/commit/b3de3d352784c6dcbcbaf4cfb1aa4550b249b618)) (@axetroy)

### 🔙 Revert:

-	revert [`da68d42`](https://github.com/axetroy/vscode-deno/commit/da68d4289b1713432a5b0203fa8f46f4008cab44), docs: add readme for typescript-deno-plugin([`1be0c24`](https://github.com/axetroy/vscode-deno/commit/1be0c24d49d811fb7a5499671d01fe5bcafae6f5)\)

v3.0.1 (2020-03-01)
-------------------

### 🔥 New feature:

-	re-enable typescript-deno-plugin with workspace's typescript version. close #78([`7a53e70`](https://github.com/axetroy/vscode-deno/commit/7a53e7019a9aeba64e0494fe7d9f666540f8a6ce)) (@axetroy)

v3.0.0 (2020-02-29)
-------------------

### 🔥 New feature:

-	Resurrected in Deno v0.35.0 🚀([`3aff7ed`](https://github.com/axetroy/vscode-deno/commit/3aff7edf1481a0a234a1b994b9f6cf692e444beb)) (@axetroy)
-	Requires minimum version of Deno 0.35.0([`35b810c`](https://github.com/axetroy/vscode-deno/commit/35b810ca4f5bbb50796d0df4d3efe27908e604a1)) (@axetroy)
-	set the language of the document automatically([`8839799`](https://github.com/axetroy/vscode-deno/commit/88397994925c42647c77df6e3626922ea3b9d953)) (@axetroy)
-	adapt Deno new cache layout([`1cba5df`](https://github.com/axetroy/vscode-deno/commit/1cba5df49c10d7d35316949a0e635262d7ee1b30)) (@axetroy)
-	support format range code([`498d37f`](https://github.com/axetroy/vscode-deno/commit/498d37ffa7b4e85efd53c8c34d96b48b313effa1)) (@axetroy)
-	add import map file jso validator([`2ccfa02`](https://github.com/axetroy/vscode-deno/commit/2ccfa026f28c2285cd94d8e066a96c37f57eef9a)) (@axetroy)
-	Add memory cache module([`c6cd7e8`](https://github.com/axetroy/vscode-deno/commit/c6cd7e8b766398d413c406258d286b7380f5087e)) (@axetroy)
-	auto detect ./vscode/settings.json in typescript plugin #60([`95d73c6`](https://github.com/axetroy/vscode-deno/commit/95d73c6b5efa7ba5951cc3c0dd150ce925a429b7)) (@axetroy)
-	support Import-Maps in new cache layout.([`0c4cccd`](https://github.com/axetroy/vscode-deno/commit/0c4cccda2c109963c28064b1f19c58508a16bae9)) (@axetroy)
-	make manifest be a iterator.([`8c7b7ce`](https://github.com/axetroy/vscode-deno/commit/8c7b7ce81d9eb2110a59b14d7e41d6c08b00c8d3)) (@axetroy)
-	add new Deno's cache layout parser and test([`7308978`](https://github.com/axetroy/vscode-deno/commit/730897813b2066533c39945bb7eae4af1b261033)) (@axetroy)
-	support @deno-types hint definition. #68([`758c5be`](https://github.com/axetroy/vscode-deno/commit/758c5be8256efa263a53c8df1bf68ce6beeaed78)) (@axetroy)

### 🐛 Bugs fixed:

-	module not follow redirect([`dc97014`](https://github.com/axetroy/vscode-deno/commit/dc97014e31db7bba9859c5d2ba4ef77ddb67c9f3)) (@axetroy)
-	Possible errors caused by optional parameters([`f91085c`](https://github.com/axetroy/vscode-deno/commit/f91085ce5b27dde09e22b0e52975f893c111cb6b)) (@axetroy)
-	auto-import in new cache layout([`7726fde`](https://github.com/axetroy/vscode-deno/commit/7726fde3ee0e9b1caf7a04bd167fac725600febe)) (@axetroy)
-	get file hash without query([`d81850f`](https://github.com/axetroy/vscode-deno/commit/d81850fde5197f39bc04adf28de84429b2d0d5a8)) (@axetroy)
-	fix windows url path handle([`b50548e`](https://github.com/axetroy/vscode-deno/commit/b50548e4bb87baa62c83fda420d7994ca0e9dba8)) (@axetroy)
-	fix invalid regexp for Windows([`d37c846`](https://github.com/axetroy/vscode-deno/commit/d37c846121d3e3a3decfdcf821205c2f5e40683b)) (@axetroy)
-	fix invalid regexp for Windows([`a96e93a`](https://github.com/axetroy/vscode-deno/commit/a96e93abeedc30d018c5c06d60c9e26909fd78f5)) (@axetroy)
-	path resolution of Windows([`a9e3336`](https://github.com/axetroy/vscode-deno/commit/a9e33363da91ae45f4178db001a9deaedc858429)) (@axetroy)
-	**deps**: update dependency vscode-languageserver-textdocument to v1.0.1 (#66)([`c49b0fa`](https://github.com/axetroy/vscode-deno/commit/c49b0fac07f1a4d4cdc9a425ef666d9741e22a97)) (@renovate[bot])

v2.0.4 (2020-02-20)
-------------------

### 🐛 Bugs fixed:

-	Try to fix the path processing under windows. ref: #61([`e3d5bf2`](https://github.com/axetroy/vscode-deno/commit/e3d5bf27fc0b678b0928caeb19a3735774179a36)) (@axetroy)

v2.0.3 (2020-02-20)
-------------------

### 🐛 Bugs fixed:

-	Try to fix the path processing under windows. ref: #61([`8c02221`](https://github.com/axetroy/vscode-deno/commit/8c02221cb2a5abfcafc108ecf2ae88afc3e90f3b)) (@axetroy)

v2.0.2 (2020-02-19)
-------------------

### 🐛 Bugs fixed:

-	Auto-Import for Deno module incorrectly. now use http protocol modules instead of relative paths. close #44([`df71fd1`](https://github.com/axetroy/vscode-deno/commit/df71fd1d4fa5f47423f1c00b9b181e81f0435dd4)) (@axetroy)
-	typescript-deno-plugin will be disable when open the file out of workspace.([`b0f3aa6`](https://github.com/axetroy/vscode-deno/commit/b0f3aa6d6646adf81a9ac091c2d89e82eda35e94)) (@axetroy)

v2.0.1 (2020-02-18)
-------------------

### 🔥 New feature:

-	improve fetch module with progress([`1eaeb41`](https://github.com/axetroy/vscode-deno/commit/1eaeb4193299aa77e5a430fb0469c8d020851524)) (@axetroy)

### 🐛 Bugs fixed:

-	improve typescript-deno-plugin. Make it as unaffected as possible.([`a6ad52f`](https://github.com/axetroy/vscode-deno/commit/a6ad52f058860767310c7774ab8bbe34289064c3)) (@axetroy)
-	**deps**: update dependency vscode-languageclient to v6.1.1 (#58)([`ae547f9`](https://github.com/axetroy/vscode-deno/commit/ae547f90153b4519cc3748a79ec5176c16bed46e)) (@renovate[bot])
-	**deps**: update dependency vscode-languageserver to v6.1.1 (#57)([`e687f20`](https://github.com/axetroy/vscode-deno/commit/e687f207960568e37bc445d63cd133bda413acff)) (@renovate[bot])
-	import module from 'file:///path/to/module/mod.ts' not work([`962411d`](https://github.com/axetroy/vscode-deno/commit/962411de1e6aa15d6a1eb122a6f0b3035017cc03)) (@axetroy)

v2.0.0 (2020-02-14)
-------------------

### 🔥 New feature:

-	Deno minimum required v0.33.0([`014192a`](https://github.com/axetroy/vscode-deno/commit/014192a0d1ce80aac3adff5d120fda06c384d03d)) (@axetroy)
-	rename configuration `deno.dtsFilepaths` to `deno.dts_file` (#49)([`555a230`](https://github.com/axetroy/vscode-deno/commit/555a230a0476f101a295fd877608f6834d3d6a79)) (@Axetroy)
-	remove `deno.enable` & `deno.disable` command (#48)([`8ecae2c`](https://github.com/axetroy/vscode-deno/commit/8ecae2c86e28138ac21d12ea29aba34860c3bb95)) (@Axetroy)
-	upgrade Deno formatter (#50)([`e872d1c`](https://github.com/axetroy/vscode-deno/commit/e872d1cee1af7d9bdf1227165dfecf1c69df8fbe)) (@Axetroy)

v1.23.0 (2020-02-14)
--------------------

### 🔥 New feature:

-	add the tips for Deno's minimum version for this extension.([`8b5c54b`](https://github.com/axetroy/vscode-deno/commit/8b5c54b8e9fc9f19c47e2b60f36cc140c587f885)) (@axetroy)
-	Now opening the js file will also launch the extension. the same with tsserver.([`d4a9beb`](https://github.com/axetroy/vscode-deno/commit/d4a9beb911cec9f02be8ce1faffe5bb4a10ba836)) (@axetroy)
-	support external type definitions with `X-TypeScript-Types` headers. close #35([`98253dd`](https://github.com/axetroy/vscode-deno/commit/98253dd0bda546b6f11beb83926d972540133e33)) (@axetroy)

v1.22.0 (2020-02-11)
--------------------

### 🔥 New feature:

-	Add translations for dutch and german (#42)([`ed2b7a4`](https://github.com/axetroy/vscode-deno/commit/ed2b7a496d31356331cfb3dda44e06c8020a7476)) (@Luca Casonato)
-	improve module import intelligent([`faf76c9`](https://github.com/axetroy/vscode-deno/commit/faf76c9b015778ef7bcf3994a9708c81d8dbacb3)) (@axetroy)

### 🐛 Bugs fixed:

-	Module index is incorrect. close #47([`d69e90a`](https://github.com/axetroy/vscode-deno/commit/d69e90a90df3d7367eb9cb0bd10ec5f3ad21033a)) (@axetroy)
-	module import intelligent no work correctly when import from 'http/server.ts'([`055d062`](https://github.com/axetroy/vscode-deno/commit/055d062c26aff15c5336c45aa952a1d653ce9cbc)) (@axetroy)

v1.21.0 (2020-02-10)
--------------------

### 🔥 New feature:

-	support external type definitions with '/// <reference types=https://raw.githubusercontent.com/date-fns/date-fns/master/typings.d.ts />'. ref: #35([`f7affb2`](https://github.com/axetroy/vscode-deno/commit/f7affb27fb073f22437db227b2c576e9406d4784)) (@axetroy)

v1.20.0 (2020-02-10)
--------------------

### 🔥 New feature:

-	remove `lock std version` and `prefer HTTPS` diagnostics. close #33([`2480791`](https://github.com/axetroy/vscode-deno/commit/2480791f9c002b8d0706f2ffedb5b93ff3c3b407)) (@axetroy)

### 🐛 Bugs fixed:

-	update ignore diagnostics code. close #41([`34e6c10`](https://github.com/axetroy/vscode-deno/commit/34e6c1053c7c4c7928fd3e83a59fdd1e92a11f95)) (@axetroy)

v1.19.0 (2020-02-08)
--------------------

### 🔥 New feature:

-	remove extension name diagnostic. close #12([`892bb3f`](https://github.com/axetroy/vscode-deno/commit/892bb3fe8822500b48d9b1bfacffaa1d4a7c17ba)) (@axetroy)
-	support import ECMA script module. close #37([`1b68068`](https://github.com/axetroy/vscode-deno/commit/1b6806854581b9f0b9460526c730eb19dcc511d4)) (@axetroy)

### 🐛 Bugs fixed:

-	esm module resolver([`ffe30fb`](https://github.com/axetroy/vscode-deno/commit/ffe30fbbde5e65b9d0741020b820d5b323db5cd1)) (@axetroy)

v1.18.1 (2020-02-07)
--------------------

### 🔙 Revert:

-	revert [`cb0e592`](https://github.com/axetroy/vscode-deno/commit/cb0e592136f569e58daee56a7d2f46759b7ca946), feat: support top-level await with typescript 3.8([`341165e`](https://github.com/axetroy/vscode-deno/commit/341165e7d1c25e1a4f2d7aab8866f54fb9b8f110)\)

v1.18.0 (2020-02-07)
--------------------

### 🔥 New feature:

-	no more use workspace typescript version([`2a6f9da`](https://github.com/axetroy/vscode-deno/commit/2a6f9da82aac305431dccc6539b66eb66866155e)) (@axetroy)
-	require min vscode version 1.42.0([`ab2cc6e`](https://github.com/axetroy/vscode-deno/commit/ab2cc6e677f08e0392fc8b551fbae9e8e303bcee)) (@axetroy)
-	support top-level await with typescript 3.8([`cb0e592`](https://github.com/axetroy/vscode-deno/commit/cb0e592136f569e58daee56a7d2f46759b7ca946)) (@axetroy)

### 🐛 Bugs fixed:

-	create local module no work([`bcceff2`](https://github.com/axetroy/vscode-deno/commit/bcceff232ded01eb28575db7151b4116968945c1)) (@axetroy)

v1.17.0 (2020-02-06)
--------------------

### 🔥 New feature:

-	fully i18n supported. #31([`04e3938`](https://github.com/axetroy/vscode-deno/commit/04e3938197c6de200f79b6115c8ab3b9cff3651e)) (@axetroy)

### 🐛 Bugs fixed:

-	create a local module if is not relative or absolute path([`21bacce`](https://github.com/axetroy/vscode-deno/commit/21bacce8dbba3837a363aeb47ba8aefd262295a4)) (@axetroy)

v1.16.0 (2020-02-06)
--------------------

### 🔥 New feature:

-	support Import Maps for Deno. close #3([`eb187af`](https://github.com/axetroy/vscode-deno/commit/eb187afd06685c9462fcdace820f29754385f860)) (@axetroy)
-	add lock deno_std version diagnostic([`8d9097e`](https://github.com/axetroy/vscode-deno/commit/8d9097e3cb23925966e7339b344fa99cd6d6d491)) (@axetroy)
-	add default content for creating a file when fix missing local module([`1404f2f`](https://github.com/axetroy/vscode-deno/commit/1404f2f712867116801cd09a0f1122298218fd42)) (@axetroy)

v1.15.0 (2020-02-05)
--------------------

### 🔥 New feature:

-	support quickly fix for diagnostics. close #29([`da85926`](https://github.com/axetroy/vscode-deno/commit/da859261e33d86b22e01560557f71f4d76b087c2)) (@axetroy)

### 🐛 Bugs fixed:

-	**deps**: pin dependency execa to 4.0.0 (#30)([`47ca6e4`](https://github.com/axetroy/vscode-deno/commit/47ca6e47d3dc0e8dbb350225d269ccae7daca278)) (@renovate[bot])
-	`typescript-deno-plugin` may not find modules and cause `typescript` to crash([`8bdc5db`](https://github.com/axetroy/vscode-deno/commit/8bdc5db5863212efee62e51b9965c811c1cdeb34)) (@axetroy)

v1.14.0 (2020-02-05)
--------------------

### 🔥 New feature:

-	Added i18n support for Chinese Traditional([`ca93cd2`](https://github.com/axetroy/vscode-deno/commit/ca93cd24b28924fd065554f748eb653d23b3a449)) (@axetroy)
-	add `deno.restart_server` command to restart `Deno Language Server`. close #28([`9a66f86`](https://github.com/axetroy/vscode-deno/commit/9a66f867f93729b4b753abc4117fb65f3cba72a3)) (@axetroy)
-	improve status bar. show more information([`6fb83c4`](https://github.com/axetroy/vscode-deno/commit/6fb83c4a496a2c031247d4675c2838b073318911)) (@axetroy)

### 🐛 Bugs fixed:

-	lock prettier version to make sure formatter work on deno v0.32.0. We will switch to dprint in a future release and only suppport formatting typescript/javascipt code.([`78b3266`](https://github.com/axetroy/vscode-deno/commit/78b3266ab426b28e288ff02c677f44593647e2b9)) (@axetroy)

v1.13.1 (2020-02-04)
--------------------

### 🐛 Bugs fixed:

-	cannot find module if redirected. close #27([`6fd7b13`](https://github.com/axetroy/vscode-deno/commit/6fd7b13dc1394687dbae6a6a6e5f60d01f72cd64)) (@axetroy)

v1.13.0 (2020-02-04)
--------------------

### 🔥 New feature:

-	improve diagnostics([`a5f029e`](https://github.com/axetroy/vscode-deno/commit/a5f029e35f9af3692d8f7192fecd648237d00c1c)) (@axetroy)

### 🐛 Bugs fixed:

-	can not import module which end with `.ts`\([`0169107`](https://github.com/axetroy/vscode-deno/commit/01691075d9d236b6a0780f960f871206788fea44)) (@axetroy)
-	**deps**: pin dependency vscode-uri to 2.1.1 (#26)([`5cdf757`](https://github.com/axetroy/vscode-deno/commit/5cdf7571673a9c5fbfbfe8858488fbb7525e1027)) (@renovate[bot])
-	improve import module position([`8a999c6`](https://github.com/axetroy/vscode-deno/commit/8a999c667ea474ee769dbf72972a08f9d8f71465)) (@axetroy)

v1.12.0 (2020-02-03)
--------------------

### 🔥 New feature:

-	improve folder picker([`71e9658`](https://github.com/axetroy/vscode-deno/commit/71e9658fb4aea962941a0d4b7f03a6cbc80d50e1)) (@axetroy)
-	Warning when import from http([`72d9db3`](https://github.com/axetroy/vscode-deno/commit/72d9db3c7ce5b483ef0fb7d3e6310b7adf5974c2)) (@axetroy)
-	remove `deno.enable = true` by default([`532cdf0`](https://github.com/axetroy/vscode-deno/commit/532cdf0af76ac436243b20c885c406386a20f202)) (@axetroy)

v1.11.0 (2020-01-31)
--------------------

### 🔥 New feature:

-	add diagnostics checking for disable non-extension name module import. close #12([`8c1c244`](https://github.com/axetroy/vscode-deno/commit/8c1c24419fe07bd7f511605ce0b062b0ae16199a)) (@axetroy)

### 🐛 Bugs fixed:

-	**deps**: pin dependency typescript to 3.7.5 (#21)([`3b3049c`](https://github.com/axetroy/vscode-deno/commit/3b3049c19d1438b3357eef77f4b6241b535db3d2)) (@renovate[bot])
-	add missing typescript deps([`751261a`](https://github.com/axetroy/vscode-deno/commit/751261aaeb8a6a2931687a8082b5bbce591d7ba2)) (@axetroy)
-	**deps**: pin dependency get-port to 5.1.1 (#18)([`d3cf219`](https://github.com/axetroy/vscode-deno/commit/d3cf21902f0b6930640cd7b1f603649746833ac5)) (@renovate[bot])

v1.10.1 (2020-01-31)
--------------------

### 🐛 Bugs fixed:

-	formatter not run at workspace folder([`bf6195a`](https://github.com/axetroy/vscode-deno/commit/bf6195a1978787c53b5135a43245ee6295ca945f)) (@axetroy)

v1.10.0 (2020-01-31)
--------------------

### 🐛 Bugs fixed:

-	completion show everywhere([`21741a2`](https://github.com/axetroy/vscode-deno/commit/21741a265e38c1187c9e8a8cc71465489a250db1)) (@axetroy)

v1.9.2 (2020-01-29)
-------------------

### 🐛 Bugs fixed:

-	resolve can not import module not end with .ts when module does not found. close #5([`1143a97`](https://github.com/axetroy/vscode-deno/commit/1143a97d59672439bb5bf1e9b0fd5279df78d4eb)) (@axetroy)

v1.9.1 (2020-01-29)
-------------------

### 🔥 New feature:

-	support top-level await. close #10([`d1cd97c`](https://github.com/axetroy/vscode-deno/commit/d1cd97ce0748ff1b4df96726efb6cda308197dd8)) (@axetroy)

v1.9.0 (2020-01-26)
-------------------

### 🔥 New feature:

-	enable jsx options by default for typescript-deno-plugin([`b9c2fba`](https://github.com/axetroy/vscode-deno/commit/b9c2fbaefe733288ec62e46e8798d6b634f7eea9)) (@axetroy)
-	support import installed module intelligent. close #4([`6d9baaa`](https://github.com/axetroy/vscode-deno/commit/6d9baaadf3ae1eeb75e9fd46e1e567c9c8c66086)) (@axetroy)

v1.8.0 (2020-01-26)
-------------------

### 🔥 New feature:

-	add deno.dtsFilepaths configuration([`458666e`](https://github.com/axetroy/vscode-deno/commit/458666eba1c649a673c56e8d28179e4cd9860d6a)) (@axetroy)

### 🐛 Bugs fixed:

-	only allow .d.ts file for deno.dtsFilepaths([`8916695`](https://github.com/axetroy/vscode-deno/commit/8916695fe86b2281c66f2ca75642b6511e6a744c)) (@axetroy)

v1.7.0 (2020-01-25)
-------------------

### 🔥 New feature:

-	enable/disable typescript-deno-plugin in extension scope([`fc2c197`](https://github.com/axetroy/vscode-deno/commit/fc2c1977fc320b0b4609ca50bb02466fdbc7cc23)) (@axetroy)

v1.6.1 (2020-01-25)
-------------------

### 🐛 Bugs fixed:

-	support import.meta.url for Deno([`3a26287`](https://github.com/axetroy/vscode-deno/commit/3a26287d3d38a6aa9a55d87c2652c3839e91793c)) (@axetroy)

v1.6.0 (2020-01-24)
-------------------

### 🐛 Bugs fixed:

-	try fix ci([`e21e3f9`](https://github.com/axetroy/vscode-deno/commit/e21e3f9a89bd3be7706a323af0e4e6b21450c77e)) (@axetroy)
-	use yarn package for vsce([`67b2efd`](https://github.com/axetroy/vscode-deno/commit/67b2efdc7363206efda14dd7c40c57db6b4162a2)) (@axetroy)

v1.5.0 (2020-01-23)
-------------------

v1.4.1 (2020-01-23)
-------------------

v1.4.0 (2019-12-07)
-------------------

v1.3.3 (2019-08-29)
-------------------

v1.3.2 (2019-06-06)
-------------------

### 🔥 New feature:

-	add format provider (#32)([`9636ee2`](https://github.com/axetroy/vscode-deno/commit/9636ee26359339ffd2149557af2e76a8f29b6e29)) (@Axetroy)

v1.3.0 (2019-04-28)
-------------------

v1.2.1 (2019-04-19)
-------------------

v1.2.0 (2019-04-19)
-------------------

### 🔙 Revert:

-	revert [`6927e6a`](https://github.com/axetroy/vscode-deno/commit/6927e6a2d0dd7ac88dcc82c4dba388ef6aae394c), implement auto format on save (#13)" (#19)([`8d29205`](https://github.com/axetroy/vscode-deno/commit/8d29205104fc6ca7e35de01a24ceb240f2ae7d77)\)

v1.0.7 (2019-03-11)
-------------------

v1.0.6 (2019-03-09)
-------------------

### 🔙 Revert:

-	revert [`20b683a`](https://github.com/axetroy/vscode-deno/commit/20b683ab0c7ebd0cdee71bebff58456e62df6aaa), bundling extension using webpack([`2b451ff`](https://github.com/axetroy/vscode-deno/commit/2b451ff9b6f50ccc7edab19d28f376f5ee19923d)\)

v1.0.5 (2019-03-08)
-------------------

v1.0.4 (2019-03-07)
-------------------

v1.0.3 (2019-03-07)
-------------------

v1.0.2 (2019-03-07)
-------------------

v1.0.1 (2019-03-07)
-------------------
