v0.4.0 (2021-10-03)
-------------------

### 🔥 New feature:

-	**BREAKING**: remove hashURL function in template render([`35a348f`](https://github.com/whatchanged-community/whatchanged/commit/35a348f3bbcb16afa4d74356899a2c651905bf12)) (@Axetroy)
-	**BREAKING**: remove stringsJoin method in template renderer and make([`cf4cb17`](https://github.com/whatchanged-community/whatchanged/commit/cf4cb179c62063ee97b89c9526f0e2b20d048a39)) (@Axetroy)
-	**BREAKING**: remove npm publish([`307a4ad`](https://github.com/whatchanged-community/whatchanged/commit/307a4ad842182e0edc8cd1bc3c7c45fa16fd2a87)) (@Axetroy)

### 🐛 Bugs fixed:

-	lint([`2082a2d`](https://github.com/whatchanged-community/whatchanged/commit/2082a2d32f354b9d8ac3b3368e400a89d817ef80)) (@Axetroy)
-	**deps**: update dependency marked to v3.0.4 (#67)([`837dd74`](https://github.com/whatchanged-community/whatchanged/commit/837dd747f672bad269f4e433a2ba4960a33308b5)) (@renovate[bot])
-	**deps**: update github.com/shurcool/markdownfmt commit hash to 7513492 (#47)([`14084b0`](https://github.com/whatchanged-community/whatchanged/commit/14084b02e5db63457514c8ac6ffd75009ddc027a)) (@renovate[bot])
-	**deps**: update dependency ant-design-vue to v2.2.8 (#73)([`950f456`](https://github.com/whatchanged-community/whatchanged/commit/950f4563eefab1023f1d15c7319e6705b07c05df)) (@renovate[bot])
-	**deps**: update dependency marked to v3 (#66)([`a61e60a`](https://github.com/whatchanged-community/whatchanged/commit/a61e60aeadfde1835127a8dce583d5d6113b3c9b)) (@renovate[bot])
-	**deps**: update dependency ant-design-vue to v2.2.6 (#55)([`8cf6abf`](https://github.com/whatchanged-community/whatchanged/commit/8cf6abf7d36d8af7b8f1cc5c2c4b7eb92cba4c97)) (@renovate[bot])
-	**deps**: update dependency marked to v2.1.3 (#59)([`3ff2cb8`](https://github.com/whatchanged-community/whatchanged/commit/3ff2cb883ae4f8bf385d39b7e48cf0912dd47660)) (@renovate[bot])
-	**deps**: update module github.com/go-git/go-git/v5 to v5.4.2 (#53)([`221b029`](https://github.com/whatchanged-community/whatchanged/commit/221b0293bee5002a56d9366d2fff5a058afb75a0)) (@renovate[bot])
-	**deps**: update dependency marked to v2.0.3 (#51)([`7c8ab58`](https://github.com/whatchanged-community/whatchanged/commit/7c8ab58365356f9de2f02d6ffac3147f9669ced7)) (@renovate[bot])
-	**deps**: update dependency ant-design-vue to v2.1.3 (#52)([`3ae4ad6`](https://github.com/whatchanged-community/whatchanged/commit/3ae4ad603ee81a5b1045c248b080ee9c362637ea)) (@renovate[bot])
-	vite config([`c13faa1`](https://github.com/whatchanged-community/whatchanged/commit/c13faa11061e22dc359fd6fc040e48752b5cb83b)) (@axetroy)

### ❤️ BREAKING CHANGES:

-	feat(BREAKING): remove hashURL function in template render

before:

```bash
{{ hashURL .Hash}}
{{ hashURL .RevertCommitHash }}
```

after:

```bash
{{ .HashURL }}
{{ .RevertCommitHashURL }}
```

-	feat(BREAKING): remove stringsJoin method in template renderer and make

before:

```
Closes: {{ stringsJoin .Field.Footer.Closes "," }}
```

after:

```
Closes: {{ .Field.Footer.Closes }}
```

-	remove npm publish

	```diff
	- npm install @axetroy/whatchanged -g
	```

### 💪 Commits(68):

-	[`f8ac971`](https://github.com/whatchanged-community/whatchanged/commit/f8ac9715c0db2dc375b3b238744dce3d0ea3c8ee) - v0.4.0
-	[`51895bc`](https://github.com/whatchanged-community/whatchanged/commit/51895bc7e33f9566574ba0911a54591ae110e162) - update
-	[`f3da8fe`](https://github.com/whatchanged-community/whatchanged/commit/f3da8fe56e04939bafceda9b9ea6e7bfada49f96) - update
-	[`c4d4ab1`](https://github.com/whatchanged-community/whatchanged/commit/c4d4ab1d28f6137447ccbdd3eda2876f22cd5cfe) - add test
-	[`9e52464`](https://github.com/whatchanged-community/whatchanged/commit/9e52464106be5add22d543b17a2bda271e4c3f4c) - update changelog
-	[`35a348f`](https://github.com/whatchanged-community/whatchanged/commit/35a348f3bbcb16afa4d74356899a2c651905bf12) - feat(BREAKING): remove hashURL function in template render
-	[`ae56529`](https://github.com/whatchanged-community/whatchanged/commit/ae56529432ddbd415762c82196fa5a3cf87eea05) - fix test
-	[`47cea0c`](https://github.com/whatchanged-community/whatchanged/commit/47cea0c02f27514ac542004dbaafcfdba74217c0) - update template
-	[`2082a2d`](https://github.com/whatchanged-community/whatchanged/commit/2082a2d32f354b9d8ac3b3368e400a89d817ef80) - fix: lint
-	[`cf4cb17`](https://github.com/whatchanged-community/whatchanged/commit/cf4cb179c62063ee97b89c9526f0e2b20d048a39) - feat(BREAKING): remove stringsJoin method in template renderer and make
-	[`8c96222`](https://github.com/whatchanged-community/whatchanged/commit/8c96222f3a8f15cbb4319f971fc42b94568fb2d4) - update docs
-	[`8ecb7d7`](https://github.com/whatchanged-community/whatchanged/commit/8ecb7d7c269ee8f9a01d6f0c39c3cce7b95e751e) - remove playgloud
-	[`ff3f10d`](https://github.com/whatchanged-community/whatchanged/commit/ff3f10d6256c2b3407dea9d5ba8330f64b41e452) - update
-	[`5365a7f`](https://github.com/whatchanged-community/whatchanged/commit/5365a7f89b914eedbd882b05674697dc7186fa6a) - update
-	[`f080f73`](https://github.com/whatchanged-community/whatchanged/commit/f080f731c00a64d199e88f33c0afaca3eb79a89d) - update
-	[`422b05b`](https://github.com/whatchanged-community/whatchanged/commit/422b05bf114d6461dc19bc7bad7d9ec8231517b2) - add install from powershell
-	[`f33e72a`](https://github.com/whatchanged-community/whatchanged/commit/f33e72ab8a669c576f85dc5b14d091ce68d449ec) - fix
-	[`0622e61`](https://github.com/whatchanged-community/whatchanged/commit/0622e61a8b1e96b410ed5b3bf660376a1daa354b) - fix
-	[`11d466f`](https://github.com/whatchanged-community/whatchanged/commit/11d466fcaa6828ed96d99f0d79dab118e1929ab2) - improve breaking change matcher
-	[`307a4ad`](https://github.com/whatchanged-community/whatchanged/commit/307a4ad842182e0edc8cd1bc3c7c45fa16fd2a87) - feat(BREAKING): remove npm publish
-	[`4282414`](https://github.com/whatchanged-community/whatchanged/commit/4282414bc03125af388f6e54a63cca2922df4ce3) - update docs
-	[`004f3bb`](https://github.com/whatchanged-community/whatchanged/commit/004f3bbb101125ddb83fbc497d55ea8fb7c9f5cf) - fix test
-	[`1cd2954`](https://github.com/whatchanged-community/whatchanged/commit/1cd2954667c34be82cd34d23ddc8bebd63b8db4c) - update
-	[`900f93c`](https://github.com/whatchanged-community/whatchanged/commit/900f93c5fd7ca1685f7971c9f894e7d126c868b4) - add context
-	[`a31242d`](https://github.com/whatchanged-community/whatchanged/commit/a31242df9ee3e65c8be79fa43ee52e95e34255df) - chore: update
-	[`2f80d85`](https://github.com/whatchanged-community/whatchanged/commit/2f80d855856c8c0db63db82377d7f07edc37aedc) - Create codeql-analysis.yml
-	[`c2cacba`](https://github.com/whatchanged-community/whatchanged/commit/c2cacba586ce496a50cf137fbff6516b960a702f) - update
-	[`d5026b6`](https://github.com/whatchanged-community/whatchanged/commit/d5026b66f876ffdfa486f7373737fc6ad55196cb) - update
-	[`8dfd196`](https://github.com/whatchanged-community/whatchanged/commit/8dfd196b1e24c19609b2e23365d3c6f1c01cf503) - update ci
-	[`58c1174`](https://github.com/whatchanged-community/whatchanged/commit/58c117422d66c1aa08cb92e92107df4183df2990) - fix test
-	[`b951adc`](https://github.com/whatchanged-community/whatchanged/commit/b951adc35768d34c18552761368a3f85815c4bb1) - update
-	[`3cead03`](https://github.com/whatchanged-community/whatchanged/commit/3cead032d30c6d3c56afb31fe4791047dd43ff05) - update ci
-	[`94d6cb3`](https://github.com/whatchanged-community/whatchanged/commit/94d6cb3f769a090b115aae76d043c7f5f5e92c71) - update
-	[`c7910de`](https://github.com/whatchanged-community/whatchanged/commit/c7910de590b164f17e77a8e0a2e5afef109946c3) - update
-	[`ab1f3d8`](https://github.com/whatchanged-community/whatchanged/commit/ab1f3d83b010edb004587a3e66baf122cd4e199b) - update ci
-	[`b28eec1`](https://github.com/whatchanged-community/whatchanged/commit/b28eec1761e321cb48a50005364deee5b1526c70) - update ci
-	[`7390c46`](https://github.com/whatchanged-community/whatchanged/commit/7390c46341bd8d0eb4fc59e699602d3d26b30c8a) - update
-	[`415fcc3`](https://github.com/whatchanged-community/whatchanged/commit/415fcc3860ded0a59ea2889dea351621feba4b81) - update
-	[`7942ced`](https://github.com/whatchanged-community/whatchanged/commit/7942cedd8cf25a61cf4f87217a91cf3a373c9a7a) - chore(deps): update vue monorepo to v3.2.19 (#63)
-	[`deeeb32`](https://github.com/whatchanged-community/whatchanged/commit/deeeb323113e6473a53802bcb2a58b1cb9c7e355) - fix
-	[`837dd74`](https://github.com/whatchanged-community/whatchanged/commit/837dd747f672bad269f4e433a2ba4960a33308b5) - fix(deps): update dependency marked to v3.0.4 (#67)
-	[`e30c8f0`](https://github.com/whatchanged-community/whatchanged/commit/e30c8f03d8c5b296f2646e1f84da2b88eb14d897) - chore(deps): update dependency sass to v1.42.1 (#69)
-	[`afeb213`](https://github.com/whatchanged-community/whatchanged/commit/afeb2130ad5445f3bc2a956afd895e44673d37ec) - chore(deps): update dependency @types/node to v14.17.20 (#58)
-	[`14084b0`](https://github.com/whatchanged-community/whatchanged/commit/14084b02e5db63457514c8ac6ffd75009ddc027a) - fix(deps): update github.com/shurcool/markdownfmt commit hash to 7513492 (#47)
-	[`ccc54e9`](https://github.com/whatchanged-community/whatchanged/commit/ccc54e9a0fcca26ec5a8d6f69dc6f2d9ee825d3c) - chore(deps): update dependency @types/marked to v2.0.5 (#68)
-	[`5926e56`](https://github.com/whatchanged-community/whatchanged/commit/5926e56f74afc3cf61fa2979ca7423aa84af9924) - chore(deps): update dependency vite to v2.6.2 (#70)
-	[`44bd0d1`](https://github.com/whatchanged-community/whatchanged/commit/44bd0d14df1f7ddd68d88068f317d149aed31768) - chore(deps): update dependency @vitejs/plugin-vue to v1.9.2 (#71)
-	[`950f456`](https://github.com/whatchanged-community/whatchanged/commit/950f4563eefab1023f1d15c7319e6705b07c05df) - fix(deps): update dependency ant-design-vue to v2.2.8 (#73)
-	[`a61e60a`](https://github.com/whatchanged-community/whatchanged/commit/a61e60aeadfde1835127a8dce583d5d6113b3c9b) - fix(deps): update dependency marked to v3 (#66)
-	[`cce6d83`](https://github.com/whatchanged-community/whatchanged/commit/cce6d8348efcd6687841bd74344b560ada0d9a4c) - chore(deps): update dependency vite to v2.5.0 (#64)
-	[`8cf6abf`](https://github.com/whatchanged-community/whatchanged/commit/8cf6abf7d36d8af7b8f1cc5c2c4b7eb92cba4c97) - fix(deps): update dependency ant-design-vue to v2.2.6 (#55)
-	[`275c52c`](https://github.com/whatchanged-community/whatchanged/commit/275c52ccd9684b2176efa0d2f906a5e2abc003b7) - chore(deps): update dependency @vitejs/plugin-vue to v1.4.0 (#65)
-	[`3ff2cb8`](https://github.com/whatchanged-community/whatchanged/commit/3ff2cb883ae4f8bf385d39b7e48cf0912dd47660) - fix(deps): update dependency marked to v2.1.3 (#59)
-	[`223889d`](https://github.com/whatchanged-community/whatchanged/commit/223889d1d0cfb08b2a706c76577b3107b64270c0) - chore(deps): update dependency sass to v1.38.0 (#57)
-	[`f13e68f`](https://github.com/whatchanged-community/whatchanged/commit/f13e68fd59f0985f4a1997f6fd7816f20c002b07) - chore(deps): update dependency vite to v2.4.2 (#62)
-	[`716fb03`](https://github.com/whatchanged-community/whatchanged/commit/716fb03e9fe0b111737a6bc41c0eab7214a671ff) - chore(deps): update dependency @vitejs/plugin-vue to v1.2.5 (#60)
-	[`221b029`](https://github.com/whatchanged-community/whatchanged/commit/221b0293bee5002a56d9366d2fff5a058afb75a0) - fix(deps): update module github.com/go-git/go-git/v5 to v5.4.2 (#53)
-	[`6249a90`](https://github.com/whatchanged-community/whatchanged/commit/6249a90c873b1d98313d14b0f5427b6e88a6a1ba) - chore(deps): update vue monorepo to v3.1.4 (#61)
-	[`b39a017`](https://github.com/whatchanged-community/whatchanged/commit/b39a017739c66ac75426ef47d88e4365019745f1) - chore(deps): update dependency vite to v2.4.1 (#56)
-	[`7c8ab58`](https://github.com/whatchanged-community/whatchanged/commit/7c8ab58365356f9de2f02d6ffac3147f9669ced7) - fix(deps): update dependency marked to v2.0.3 (#51)
-	[`8576eac`](https://github.com/whatchanged-community/whatchanged/commit/8576eacc3442b9682c00377157f580776ceb3a92) - chore(deps): update dependency @types/node to v14.14.44 (#45)
-	[`46e06ea`](https://github.com/whatchanged-community/whatchanged/commit/46e06ea9b41753b01887e4fa82b6ffc6659d77e0) - chore(deps): update dependency vite to v2.2.4 (#54)
-	[`add6f52`](https://github.com/whatchanged-community/whatchanged/commit/add6f5234604fe213003b9db39f468c05ef5f127) - chore(deps): update vue monorepo to v3.0.11 (#49)
-	[`b31c3f4`](https://github.com/whatchanged-community/whatchanged/commit/b31c3f4f16c0eeb96c3c8e0052d4d15759c6561e) - chore(deps): update dependency sass to v1.32.12 (#46)
-	[`b2b7d91`](https://github.com/whatchanged-community/whatchanged/commit/b2b7d91cda57df2321a3ae7a9dec425fbc899780) - chore(deps): update dependency @vitejs/plugin-vue to v1.2.2 (#50)
-	[`16e6698`](https://github.com/whatchanged-community/whatchanged/commit/16e6698ff65049f6d39c13de3b9935893a84b7b8) - chore(deps): update dependency vite to v2.2.3 (#48)
-	[`3ae4ad6`](https://github.com/whatchanged-community/whatchanged/commit/3ae4ad603ee81a5b1045c248b080ee9c362637ea) - fix(deps): update dependency ant-design-vue to v2.1.3 (#52)
-	[`c13faa1`](https://github.com/whatchanged-community/whatchanged/commit/c13faa11061e22dc359fd6fc040e48752b5cb83b) - fix: vite config

v0.3.6 (2021-02-18)
-------------------

### 🐛 Bugs fixed:

-	**deps**: update dependency ant-design-vue to v2.0.0 (#39)([`3c827b3`](https://github.com/whatchanged-community/whatchanged/commit/3c827b326e2741bfc6e2b2354fec7387e15db97a)) (@renovate[bot])
-	**deps**: update dependency marked to v2 [security](#41)\([`e607ad2`](https://github.com/whatchanged-community/whatchanged/commit/e607ad207894f735ae1578a0558213c9086e1238)) (@renovate[bot])
-	**deps**: update dependency ant-design-vue to v2.0.0-rc.9 (#31)([`90903e2`](https://github.com/whatchanged-community/whatchanged/commit/90903e254d36da0dabfcd17af201e2aed884410f)) (@renovate[bot])

### 💪 Commits(20):

-	[`7d31bdc`](https://github.com/whatchanged-community/whatchanged/commit/7d31bdc155706647f668f9afc36aad846fb307a5) - v0.3.6
-	[`4453aca`](https://github.com/whatchanged-community/whatchanged/commit/4453aca655df5a0dea10abfc9cd2227ba1d0b846) - ci: update golangci-lint
-	[`0263f5a`](https://github.com/whatchanged-community/whatchanged/commit/0263f5a87a43d2ed43fefe76b09688bd4f9e7f81) - chore(deps): update dependency @vitejs/plugin-vue to v1.1.4 (#34)
-	[`fd753b8`](https://github.com/whatchanged-community/whatchanged/commit/fd753b8c2323422162ee003dc46f182aa2c6983a) - chore(deps): update dependency @types/node to v14.14.28 (#36)
-	[`2c39bd0`](https://github.com/whatchanged-community/whatchanged/commit/2c39bd0abcefb5f72061fcd6c4278727adb4aa15) - chore(deps): update dependency sass to v1.32.7 (#35)
-	[`7a8e849`](https://github.com/whatchanged-community/whatchanged/commit/7a8e8494bfc51d387dacea6d9869054526c3a2b9) - chore(deps): update module stretchr/testify to v1.7.0 (#38)
-	[`9d97671`](https://github.com/whatchanged-community/whatchanged/commit/9d97671f76e3d93772c828bd0d7c3bc4a3d6c7c4) - chore(deps): update dependency @types/marked to v1.2.2 (#33)
-	[`1a627b6`](https://github.com/whatchanged-community/whatchanged/commit/1a627b6b3bbb016ec3a2bf08fe0eb00904f5fe7d) - chore(deps): update dependency vite to v2.0.1 (#43)
-	[`3c827b3`](https://github.com/whatchanged-community/whatchanged/commit/3c827b326e2741bfc6e2b2354fec7387e15db97a) - fix(deps): update dependency ant-design-vue to v2.0.0 (#39)
-	[`e607ad2`](https://github.com/whatchanged-community/whatchanged/commit/e607ad207894f735ae1578a0558213c9086e1238) - fix(deps): update dependency marked to v2 [security](#41)
-	[`13586c5`](https://github.com/whatchanged-community/whatchanged/commit/13586c5d8f800f6ff617d4edda594464e5f3a78c) - chore(deps): update dependency vite to v2.0.0 (#32)
-	[`5386786`](https://github.com/whatchanged-community/whatchanged/commit/5386786b4e135bca9b266f27b0efe86540bf8155) - ci: update
-	[`90903e2`](https://github.com/whatchanged-community/whatchanged/commit/90903e254d36da0dabfcd17af201e2aed884410f) - fix(deps): update dependency ant-design-vue to v2.0.0-rc.9 (#31)
-	[`01e3b1b`](https://github.com/whatchanged-community/whatchanged/commit/01e3b1ba15b559b753f69cad40495556e3c794c2) - chore(deps): update dependency vite to v2.0.0-beta.48 (#30)
-	[`4d0582c`](https://github.com/whatchanged-community/whatchanged/commit/4d0582c1f44a71d45a55677a656bf7c7b32bf93b) - chore(deps): update dependency @vitejs/plugin-vue to v1.1.2 (#29)
-	[`c042279`](https://github.com/whatchanged-community/whatchanged/commit/c042279233c1ea9c53245477c9d61cc23cb1a57c) - refactor(playgroud): update template
-	[`b2b697e`](https://github.com/whatchanged-community/whatchanged/commit/b2b697ecf5eeb283773fe47926c1b70da633dc30) - chore(deps): update dependency sass to v1.32.5 (#26)
-	[`27150fd`](https://github.com/whatchanged-community/whatchanged/commit/27150fd3e97f58571cbcf0454ba3b25144a71755) - chore(deps): update dependency vite to v2.0.0-beta.37 (#27)
-	[`dfb1da0`](https://github.com/whatchanged-community/whatchanged/commit/dfb1da0dd6d206e65778542f4cfe849a8b1729f2) - chore(deps): update dependency @types/node to v14.14.22 (#24)
-	[`d501a57`](https://github.com/whatchanged-community/whatchanged/commit/d501a57253ac6d6c311d9e679493692ad2a69665) - chore(deps): update dependency @vitejs/plugin-vue to v1.1.0 (#25)

v0.3.5 (2021-01-20)
-------------------

### 🐛 Bugs fixed:

-	npm install in windows([`fb012b7`](https://github.com/whatchanged-community/whatchanged/commit/fb012b75380178856f6eb1a822d04f5078e2c2ac)) (@axetroy)
-	**playground**: build([`b2fc9b7`](https://github.com/whatchanged-community/whatchanged/commit/b2fc9b7ed85e1fb711b54268148f5675640ad742)) (@axetroy)

### 💪 Commits(7):

-	[`70d1924`](https://github.com/whatchanged-community/whatchanged/commit/70d1924f37b5bba5a48c3243c7b620f66e9dbe80) - v0.3.5
-	[`fb012b7`](https://github.com/whatchanged-community/whatchanged/commit/fb012b75380178856f6eb1a822d04f5078e2c2ac) - fix: npm install in windows
-	[`5b301aa`](https://github.com/whatchanged-community/whatchanged/commit/5b301aae694185e30c48903b12537d61f97755a1) - docs: fix link
-	[`ad68d8d`](https://github.com/whatchanged-community/whatchanged/commit/ad68d8df900c1fc00809cf2d1c58b2fd2ddf3fd2) - update LICENSE to anti-996
-	[`9e4cd75`](https://github.com/whatchanged-community/whatchanged/commit/9e4cd750c6fee9b5d2e2a5c26550bfd65169696b) - chore(deps): pin dependencies (#21)
-	[`b2fc9b7`](https://github.com/whatchanged-community/whatchanged/commit/b2fc9b7ed85e1fb711b54268148f5675640ad742) - fix(playground): build
-	[`1cc5ab8`](https://github.com/whatchanged-community/whatchanged/commit/1cc5ab8a3332cfb0494db3180e391cee4809310a) - refactor(playgroud): update vite

v0.3.4 (2021-01-04)
-------------------

### 🔥 New feature:

-	add simple preset template([`4e54f2e`](https://github.com/whatchanged-community/whatchanged/commit/4e54f2e43ca49d52d3c3cad409402dcfc6e5ab26)) (@axetroy)

### 💪 Commits(11):

-	[`574d4f4`](https://github.com/whatchanged-community/whatchanged/commit/574d4f434e94948820c4635107b69f6a35d1a458) - 0.3.4
-	[`4e54f2e`](https://github.com/whatchanged-community/whatchanged/commit/4e54f2e43ca49d52d3c3cad409402dcfc6e5ab26) - feat: add simple preset template
-	[`2a2766e`](https://github.com/whatchanged-community/whatchanged/commit/2a2766e4d8fb27909c814f95c496be16b0a3cf41) - docs: add zh-CN readme
-	[`0582024`](https://github.com/whatchanged-community/whatchanged/commit/0582024b8d67cfecfd9ddee3b007a15c3e763135) - docs: update changelog
-	[`53e9214`](https://github.com/whatchanged-community/whatchanged/commit/53e9214e2030eb5e830726bc5a6ba26968ef5725) - ci: fix lint
-	[`2bbc17d`](https://github.com/whatchanged-community/whatchanged/commit/2bbc17d5dfb0184c5772f36a80b320ba71c83d73) - ci: fix lint
-	[`3f32b0b`](https://github.com/whatchanged-community/whatchanged/commit/3f32b0bfd50e65ebd320282c40c6e9841b485751) - docs: add custom template docs. close #17
-	[`257f6a4`](https://github.com/whatchanged-community/whatchanged/commit/257f6a4bb1c3dd8a7edb195c3cbbe818a1fad5ec) - test: add more test
-	[`9b8a6dd`](https://github.com/whatchanged-community/whatchanged/commit/9b8a6dd8b5554cd8296c74455aee4240abaa42cd) - update changelog
-	[`23448a5`](https://github.com/whatchanged-community/whatchanged/commit/23448a5482359f28a0089b17280dd2a0a0eaef26) - chore: skip install golang in golanglint
-	[`a3dfef7`](https://github.com/whatchanged-community/whatchanged/commit/a3dfef78486dfa612db49ae67ed155962fecefa3) - chore: build with go1.16-beta1

v0.3.3 (2020-12-31)
-------------------

### 🐛 Bugs fixed:

-	generate url for commit([`bba03df`](https://github.com/whatchanged-community/whatchanged/commit/bba03dfa1ed948d3f5309b67cf5a357c978fd2ed)) (@axetroy)
-	generate version range([`b0dca5d`](https://github.com/whatchanged-community/whatchanged/commit/b0dca5dd226046c131165a8f211b8f190c54d04b)) (@axetroy)

### 💪 Commits(7):

-	[`3bdb4cd`](https://github.com/whatchanged-community/whatchanged/commit/3bdb4cd35e59865a9ee3f3889e13ddfb3e2d4601) - test: update test snapshot
-	[`bba03df`](https://github.com/whatchanged-community/whatchanged/commit/bba03dfa1ed948d3f5309b67cf5a357c978fd2ed) - fix: generate url for commit
-	[`f5f8bef`](https://github.com/whatchanged-community/whatchanged/commit/f5f8bef47670d7fece0c9dd7025f2d5b5b7b0143) - test: fix generation
-	[`2ffd630`](https://github.com/whatchanged-community/whatchanged/commit/2ffd6308ebed82db3367e78fd022e920b3c4e0e8) - add .gitattributes
-	[`3b18728`](https://github.com/whatchanged-community/whatchanged/commit/3b18728898dd840b043ed64818846ad039a335b1) - chore: update ci
-	[`413abed`](https://github.com/whatchanged-community/whatchanged/commit/413abedf4d058fcb7f9aca18eaa99a330a87169d) - v0.3.3
-	[`b0dca5d`](https://github.com/whatchanged-community/whatchanged/commit/b0dca5dd226046c131165a8f211b8f190c54d04b) - fix: generate version range

v0.3.2 (2020-12-30)
-------------------

### 🐛 Bugs fixed:

-	When generating an uncertain version, the generated range is incorrect([`7e1db8d`](https://github.com/whatchanged-community/whatchanged/commit/7e1db8dd3a4e3a6970672139552c9bf42b10edba)) (@axetroy)

### 💪 Commits(8):

-	[`2db6d7c`](https://github.com/whatchanged-community/whatchanged/commit/2db6d7c984d042e00493bbc8621d39c680f6d708) - v0.3.2
-	[`17d4565`](https://github.com/whatchanged-community/whatchanged/commit/17d4565edd6f0334fc01b08b33a77110fa1f6d24) - chore: update makefile
-	[`7e1db8d`](https://github.com/whatchanged-community/whatchanged/commit/7e1db8dd3a4e3a6970672139552c9bf42b10edba) - fix: When generating an uncertain version, the generated range is incorrect
-	[`e177518`](https://github.com/whatchanged-community/whatchanged/commit/e17751864ceb762d598c44e2d861be00b26847c5) - chore: support build armv5
-	[`97ccdec`](https://github.com/whatchanged-community/whatchanged/commit/97ccdecde838f5c03516fde5dcb23fcf26d0c928) - chore: update ci
-	[`44f28da`](https://github.com/whatchanged-community/whatchanged/commit/44f28da30b079c97444b4868cf67c044c4bb9525) - docs: update changelog
-	[`7c3db91`](https://github.com/whatchanged-community/whatchanged/commit/7c3db918b94aa3027b9656d6b584a854129f1424) - docs: update
-	[`1b863fa`](https://github.com/whatchanged-community/whatchanged/commit/1b863fa926c473436e2b07b5b724b30af32bd221) - docs: update

v0.3.1 (2020-12-30)
-------------------

### 🐛 Bugs fixed:

-	Generate duplicate commit([`b7fa03c`](https://github.com/whatchanged-community/whatchanged/commit/b7fa03c8e60d82fb4d06956a8f9c79c174bb227f)) (@axetroy)

### 💪 Commits(2):

-	[`7c09a29`](https://github.com/whatchanged-community/whatchanged/commit/7c09a2934698c1842f9d93d0c8ed414a34eed6bb) - v0.3.1
-	[`b7fa03c`](https://github.com/whatchanged-community/whatchanged/commit/b7fa03c8e60d82fb4d06956a8f9c79c174bb227f) - fix: Generate duplicate commit

v0.3.0 (2020-12-30)
-------------------

### 🔥 New feature:

-	**BREAKING**: update generation template.([`20dfe73`](https://github.com/whatchanged-community/whatchanged/commit/20dfe7361a679e64abd9ec2ba8d59f935627ebfe)) (@axetroy)
-	support release date for generation. close #16([`4e9b59a`](https://github.com/whatchanged-community/whatchanged/commit/4e9b59a5fbbe8c62d782e0350273fb5980e52b48)) (@axetroy)

### 💪 Commits(4):

-	[`63d1118`](https://github.com/whatchanged-community/whatchanged/commit/63d1118a4ce3051b91059be2dee52d595690628f) - v0.3.0
-	[`20dfe73`](https://github.com/whatchanged-community/whatchanged/commit/20dfe7361a679e64abd9ec2ba8d59f935627ebfe) - feat(BREAKING): update generation template.
-	[`4e9b59a`](https://github.com/whatchanged-community/whatchanged/commit/4e9b59a5fbbe8c62d782e0350273fb5980e52b48) - feat: support release date for generation. close #16
-	[`21817fa`](https://github.com/whatchanged-community/whatchanged/commit/21817fa1aba8d1e71b5fe73ba4047dc08eae197b) - chore: update sum file

v0.2.7 (2020-12-29)
-------------------

### 🔥 New feature:

-	add --skip-format and flag options for generation([`a03b7cb`](https://github.com/whatchanged-community/whatchanged/commit/a03b7cb49fa23436de43e6d3436cfd123d831ede)) (@axetroy)

### 🐛 Bugs fixed:

-	generate incorrect in mutiple version for same commit. close #15([`8fd5818`](https://github.com/whatchanged-community/whatchanged/commit/8fd58182b02ea657c37c09e4a734a48215eaffb6)) (@axetroy)
-	**deps**: update dependency ant-design-vue to v2.0.0-rc.3 (#14)([`e6275b7`](https://github.com/whatchanged-community/whatchanged/commit/e6275b7d7f7b009e79a70f6f869efebb9dc7a866)) (@renovate[bot])

### 💪 Commits(6):

-	[`aa904db`](https://github.com/whatchanged-community/whatchanged/commit/aa904db8a94b965365aa796518002c92e1158f4e) - v0.2.7
-	[`dd4c86f`](https://github.com/whatchanged-community/whatchanged/commit/dd4c86f61ff127b74b82d48d0919fc603651163e) - test: update testcase
-	[`8fd5818`](https://github.com/whatchanged-community/whatchanged/commit/8fd58182b02ea657c37c09e4a734a48215eaffb6) - fix: generate incorrect in mutiple version for same commit. close #15
-	[`a03b7cb`](https://github.com/whatchanged-community/whatchanged/commit/a03b7cb49fa23436de43e6d3436cfd123d831ede) - feat: add --skip-format and flag options for generation
-	[`e6275b7`](https://github.com/whatchanged-community/whatchanged/commit/e6275b7d7f7b009e79a70f6f869efebb9dc7a866) - fix(deps): update dependency ant-design-vue to v2.0.0-rc.3 (#14)
-	[`6ab9c0b`](https://github.com/whatchanged-community/whatchanged/commit/6ab9c0ba085699192d64e856dae5bd9367296a5d) - chore: update golanglint

v0.2.6 (2020-12-07)
-------------------

### 🐛 Bugs fixed:

-	npm binary([`8384bf7`](https://github.com/whatchanged-community/whatchanged/commit/8384bf782d8adf1627082f3e9030ed4a88c0fa5a)) (@axetroy)

### 💪 Commits(3):

-	[`f055c6f`](https://github.com/whatchanged-community/whatchanged/commit/f055c6f1033d3b559007f4e3449227067d31fdfb) - v0.2.6
-	[`46944c6`](https://github.com/whatchanged-community/whatchanged/commit/46944c641ff26417b77c8210b89c5ca09a1d2480) - chore(npm): auto bump version
-	[`8384bf7`](https://github.com/whatchanged-community/whatchanged/commit/8384bf782d8adf1627082f3e9030ed4a88c0fa5a) - fix: npm binary

v0.2.5 (2020-12-05)
-------------------

### 🐛 Bugs fixed:

-	**npm**: missing postinstall script([`6209122`](https://github.com/whatchanged-community/whatchanged/commit/6209122eacda6d86421a6955dad86785d0206b4b)) (@axetroy)

### 💪 Commits(3):

-	[`91f3fa9`](https://github.com/whatchanged-community/whatchanged/commit/91f3fa9c27b6421694be6b67308533049d8adbf4) - v0.2.5
-	[`6209122`](https://github.com/whatchanged-community/whatchanged/commit/6209122eacda6d86421a6955dad86785d0206b4b) - fix(npm): missing postinstall script
-	[`9650595`](https://github.com/whatchanged-community/whatchanged/commit/9650595bba4aea0c11fbd1ec82eb6a676f7daa75) - chore(deps): update dependency sass to v1.30.0 (#13)

v0.2.4 (2020-12-05)
-------------------

### 💪 Commits(3):

-	[`2872a05`](https://github.com/whatchanged-community/whatchanged/commit/2872a056e2772feffbdcd0adc73bd1c0c64c6127) - v0.2.4
-	[`9dff4fc`](https://github.com/whatchanged-community/whatchanged/commit/9dff4fc6a9d746ffd9dd10215cf04d2fec2edd2a) - update
-	[`c7bc358`](https://github.com/whatchanged-community/whatchanged/commit/c7bc35835361afd5ddffaa8a157513cf83b4b24b) - chore: update

v0.2.3 (2020-12-05)
-------------------

### 🔥 New feature:

-	add npm install([`89c46fe`](https://github.com/whatchanged-community/whatchanged/commit/89c46feebf8a467567d192cf40324b50f3080437)) (@axetroy)

### 🐛 Bugs fixed:

-	**deps**: pin dependencies (#11)([`bfc347c`](https://github.com/whatchanged-community/whatchanged/commit/bfc347c7d945e6ba47787758ba88bc6940d1341a)) (@renovate[bot])

### 💪 Commits(6):

-	[`4474fb8`](https://github.com/whatchanged-community/whatchanged/commit/4474fb8b40e54b6cd6e776e53dc0df895379c048) - v0.2.3
-	[`09cb54d`](https://github.com/whatchanged-community/whatchanged/commit/09cb54d8772bbecf03f98d5559bb02973423d14a) - ci: add npm publish
-	[`bfc347c`](https://github.com/whatchanged-community/whatchanged/commit/bfc347c7d945e6ba47787758ba88bc6940d1341a) - fix(deps): pin dependencies (#11)
-	[`a349f8b`](https://github.com/whatchanged-community/whatchanged/commit/a349f8b4418cec0230aea8e8be87d9bc9abe1ed3) - chore(deps): update dependency @types/marked to v1.2.1 (#12)
-	[`68df2e8`](https://github.com/whatchanged-community/whatchanged/commit/68df2e8545288a36bc946b615bdb7eb1f115c6e9) - docs: update package.json
-	[`89c46fe`](https://github.com/whatchanged-community/whatchanged/commit/89c46feebf8a467567d192cf40324b50f3080437) - feat: add npm install

v0.2.2 (2020-12-04)
-------------------

### 🐛 Bugs fixed:

-	path handler([`58d7a2d`](https://github.com/whatchanged-community/whatchanged/commit/58d7a2d8d04bc99adf0b6a19e0a74261ab48f477)) (@axetroy)

### 💪 Commits(1):

-	[`58d7a2d`](https://github.com/whatchanged-community/whatchanged/commit/58d7a2d8d04bc99adf0b6a19e0a74261ab48f477) - fix: path handler

v0.2.1 (2020-12-04)
-------------------

### 🔥 New feature:

-	add share button for playground([`e46ccfd`](https://github.com/whatchanged-community/whatchanged/commit/e46ccfdf4075210534651209b60858103850a49f)) (@axetroy)

### 🐛 Bugs fixed:

-	regexp not correct([`5a8d338`](https://github.com/whatchanged-community/whatchanged/commit/5a8d338b5c1b9f4abce79e9c248f160907da89b0)) (@axetroy)
-	try fix path handler in windows([`a9c32a6`](https://github.com/whatchanged-community/whatchanged/commit/a9c32a6e9ef7a629d2d13bc341e06f896a8abe15)) (@axetroy)
-	update playground([`3a5437a`](https://github.com/whatchanged-community/whatchanged/commit/3a5437acbd9b711e807a7f8019a6e6302dfe0646)) (@axetroy)
-	playground([`a867126`](https://github.com/whatchanged-community/whatchanged/commit/a8671265554c3ded3d20c8e1fbdb6e55ebc133f2)) (@axetroy)
-	**playground**: error logo([`3fa90ed`](https://github.com/whatchanged-community/whatchanged/commit/3fa90eda1cc98378f4d5a8197682405ebbe92532)) (@axetroy)

### 💪 Commits(22):

-	[`a94f18f`](https://github.com/whatchanged-community/whatchanged/commit/a94f18fbd2561d9b12735544e9e9b632c4ab15b0) - docs: update readme
-	[`952ac0e`](https://github.com/whatchanged-community/whatchanged/commit/952ac0ed4c34dacf55ae399b3b5d8a7f5cac0cfa) - refactor: remove console
-	[`2b7c5d4`](https://github.com/whatchanged-community/whatchanged/commit/2b7c5d4dc4809ce00f5dc48bd15445fd10e8fcbd) - refactor: improve dir detect
-	[`5a8d338`](https://github.com/whatchanged-community/whatchanged/commit/5a8d338b5c1b9f4abce79e9c248f160907da89b0) - fix: regexp not correct
-	[`a9c32a6`](https://github.com/whatchanged-community/whatchanged/commit/a9c32a6e9ef7a629d2d13bc341e06f896a8abe15) - fix: try fix path handler in windows
-	[`5050ea4`](https://github.com/whatchanged-community/whatchanged/commit/5050ea45deb73d861df1b2e62f945502a2c78c7d) - ci: update ci
-	[`505b899`](https://github.com/whatchanged-community/whatchanged/commit/505b899c9f280ae89884ee171f4afb7cf460e3aa) - update
-	[`78dd79a`](https://github.com/whatchanged-community/whatchanged/commit/78dd79a45745d6785856bc32c104fe436660bdb5) - chore(deps): pin dependencies (#10)
-	[`d06c3a6`](https://github.com/whatchanged-community/whatchanged/commit/d06c3a617de0d768d345d5cda7912927c35b844f) - update playground
-	[`1db175c`](https://github.com/whatchanged-community/whatchanged/commit/1db175cba538af330345ac53f4b4106e635ba882) - test: add more test for generation
-	[`e898be4`](https://github.com/whatchanged-community/whatchanged/commit/e898be4385c738737c67e54eb2a2cd490eafd155) - test: add test
-	[`8642531`](https://github.com/whatchanged-community/whatchanged/commit/8642531883c864bdd87a8c2cc9026f9614ca9d85) - docs: update readme
-	[`3a5437a`](https://github.com/whatchanged-community/whatchanged/commit/3a5437acbd9b711e807a7f8019a6e6302dfe0646) - fix: update playground
-	[`a867126`](https://github.com/whatchanged-community/whatchanged/commit/a8671265554c3ded3d20c8e1fbdb6e55ebc133f2) - fix: playground
-	[`3fa90ed`](https://github.com/whatchanged-community/whatchanged/commit/3fa90eda1cc98378f4d5a8197682405ebbe92532) - fix(playground): error logo
-	[`e46ccfd`](https://github.com/whatchanged-community/whatchanged/commit/e46ccfdf4075210534651209b60858103850a49f) - feat: add share button for playground
-	[`9b308bb`](https://github.com/whatchanged-community/whatchanged/commit/9b308bb334f7be38403097c5b5b0e0e4d9e1c523) - refactor: rename repl to playground 2
-	[`d96bc10`](https://github.com/whatchanged-community/whatchanged/commit/d96bc1032d854a02626695442518155970d84a53) - refactor: rename repl to playground
-	[`a9fa2e8`](https://github.com/whatchanged-community/whatchanged/commit/a9fa2e843aa06927f37a629592d20dc06f8efcaa) - refactor: update preset=full template
-	[`69cb636`](https://github.com/whatchanged-community/whatchanged/commit/69cb636c3a69e7f5e86d29125dde8a926f6f63c5) - refactor(repl): update default template and add warning on mounted
-	[`ce67088`](https://github.com/whatchanged-community/whatchanged/commit/ce670883aa5a27b010d07cadec8aa1754a23f0bc) - ci: update release template
-	[`c489594`](https://github.com/whatchanged-community/whatchanged/commit/c489594a58fcf9b1aa2cad887574712546001163) - docs: update changelog

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

v0.1.0 (2020-11-24)
-------------------

### 🔥 New feature:

-	link commit for generation([`b9432db`](https://github.com/whatchanged-community/whatchanged/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb)) (@axetroy)
-	add full preset template([`7553570`](https://github.com/whatchanged-community/whatchanged/commit/7553570590b571bd33e10a4f80ec5639d0613042)) (@axetroy)
-	support changelog for git submodule([`ec6a957`](https://github.com/whatchanged-community/whatchanged/commit/ec6a957752fbca9faa261d8694826779e2cbec1f)) (@axetroy)
-	add writer step([`441ad13`](https://github.com/whatchanged-community/whatchanged/commit/441ad1322b1fecaca89a170ecebaf2955a77d630)) (@axetroy)
-	add formatter for markdown output([`8c177e0`](https://github.com/whatchanged-community/whatchanged/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34)) (@axetroy)
-	add --file flag to generate file([`0e4fb09`](https://github.com/whatchanged-community/whatchanged/commit/0e4fb09789732fec5b09b247e208d61794c3da0d)) (@axetroy)
-	support custom tmeplate and preset([`3aa0aee`](https://github.com/whatchanged-community/whatchanged/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db)) (@axetroy)
-	support tag ranges([`3d14c9c`](https://github.com/whatchanged-community/whatchanged/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9)) (@axetroy)
-	support version range. eg v2.0.0~v1.0.0([`a65a4a8`](https://github.com/whatchanged-community/whatchanged/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3)) (@axetroy)
-	parse commit message and generate to template([`7f67e78`](https://github.com/whatchanged-community/whatchanged/commit/7f67e783926fed647d2ad5414f31448eea106fc3)) (@axetroy)

### 🐛 Bugs fixed:

-	improve ssh git url parser([`9993ee6`](https://github.com/whatchanged-community/whatchanged/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f)) (@axetroy)
-	commit range not include commit of tag([`f8df0ba`](https://github.com/whatchanged-community/whatchanged/commit/f8df0ba654c8faf67eccf98262cd55807e53e597)) (@axetroy)
-	unescape template([`e118cbf`](https://github.com/whatchanged-community/whatchanged/commit/e118cbfafd201b945848f15303fdb261e251f058)) (@axetroy)
-	if empty argument for command line([`9c79fd9`](https://github.com/whatchanged-community/whatchanged/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd)) (@axetroy)
-	**ci**: remove unsued code([`66bcf8f`](https://github.com/whatchanged-community/whatchanged/commit/66bcf8f43db85409e0392c93f2e347ed91699e81)) (@axetroy)

### 💪 Commits(43):

-	[`7a3fd6d`](https://github.com/whatchanged-community/whatchanged/commit/7a3fd6de68462fb64266bf3f606e957d32cdc2ae) - ci: add release template
-	[`00d6cc7`](https://github.com/whatchanged-community/whatchanged/commit/00d6cc7bdf42d6a0b58cf15e0aea65710f1efbf1) - ci: debug
-	[`e6f97da`](https://github.com/whatchanged-community/whatchanged/commit/e6f97daba33e8c912ee738e4fceb2bf3da877cee) - ci: debug
-	[`62e3ae1`](https://github.com/whatchanged-community/whatchanged/commit/62e3ae1045c800c49743374cf83705ba03ce5bd7) - ci: debug
-	[`10b0cf3`](https://github.com/whatchanged-community/whatchanged/commit/10b0cf3b6f91085627f8c59351570abdeb01d354) - ci: debug
-	[`cc5dd66`](https://github.com/whatchanged-community/whatchanged/commit/cc5dd665e055c066738dc2fcdb26aec61d4e61b0) - ci: update gorelease
-	[`343e4a8`](https://github.com/whatchanged-community/whatchanged/commit/343e4a8f3eaf8a715030652cfd3b7e696e40422e) - chore: add release.md to gitignore
-	[`6891359`](https://github.com/whatchanged-community/whatchanged/commit/68913598822611d64a60211918b94a462822e734) - ci: fix ci
-	[`34f51dc`](https://github.com/whatchanged-community/whatchanged/commit/34f51dc2c90e4c8ba822666dd5dcbabe312e2841) - build: rename binary name
-	[`fecc276`](https://github.com/whatchanged-community/whatchanged/commit/fecc27679d6431f840eb57c3359aa84bec1607ef) - ci: update gorelease
-	[`7c687f7`](https://github.com/whatchanged-community/whatchanged/commit/7c687f7d729f199ee295109012cc5a2661f96c54) - ci: increase linter timeout
-	[`b61c491`](https://github.com/whatchanged-community/whatchanged/commit/b61c49171a776045d89e6cbaad82326070e2db78) - refactor: improve default preset
-	[`1292dfc`](https://github.com/whatchanged-community/whatchanged/commit/1292dfc0e39abf24262e94cd9076d91808dd0cc4) - refactor: update template
-	[`9993ee6`](https://github.com/whatchanged-community/whatchanged/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f) - fix: improve ssh git url parser
-	[`b9432db`](https://github.com/whatchanged-community/whatchanged/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb) - feat: link commit for generation
-	[`7553570`](https://github.com/whatchanged-community/whatchanged/commit/7553570590b571bd33e10a4f80ec5639d0613042) - feat: add full preset template
-	[`7068672`](https://github.com/whatchanged-community/whatchanged/commit/706867220fa9ca537855f359d3e04d0c3762b793) - update readme
-	[`ec6a957`](https://github.com/whatchanged-community/whatchanged/commit/ec6a957752fbca9faa261d8694826779e2cbec1f) - feat: support changelog for git submodule
-	[`f540d6f`](https://github.com/whatchanged-community/whatchanged/commit/f540d6f7123334dac558a37c6ac056fed1021cda) - refactor: rename transform to transformer
-	[`441ad13`](https://github.com/whatchanged-community/whatchanged/commit/441ad1322b1fecaca89a170ecebaf2955a77d630) - feat: add writer step
-	[`f8df0ba`](https://github.com/whatchanged-community/whatchanged/commit/f8df0ba654c8faf67eccf98262cd55807e53e597) - fix: commit range not include commit of tag
-	[`8c177e0`](https://github.com/whatchanged-community/whatchanged/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34) - feat: add formatter for markdown output
-	[`a47fe84`](https://github.com/whatchanged-community/whatchanged/commit/a47fe84d2141635d82c2dc49500bfc2a81c03535) - docs: update how it works
-	[`0e4fb09`](https://github.com/whatchanged-community/whatchanged/commit/0e4fb09789732fec5b09b247e208d61794c3da0d) - feat: add --file flag to generate file
-	[`3aa0aee`](https://github.com/whatchanged-community/whatchanged/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db) - feat: support custom tmeplate and preset
-	[`e118cbf`](https://github.com/whatchanged-community/whatchanged/commit/e118cbfafd201b945848f15303fdb261e251f058) - fix: unescape template
-	[`3d14c9c`](https://github.com/whatchanged-community/whatchanged/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9) - feat: support tag ranges
-	[`44ce3cd`](https://github.com/whatchanged-community/whatchanged/commit/44ce3cd8d68b786a3aac6b5daba90e7f12b75200) - docs: update readme
-	[`9c79fd9`](https://github.com/whatchanged-community/whatchanged/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd) - fix: if empty argument for command line
-	[`0b6ba0c`](https://github.com/whatchanged-community/whatchanged/commit/0b6ba0c3fc49139467025eaebbe16c158a0cce65) - refactor: Refactor the software architecture to make it clearer and simpler
-	[`66bcf8f`](https://github.com/whatchanged-community/whatchanged/commit/66bcf8f43db85409e0392c93f2e347ed91699e81) - fix(ci): remove unsued code
-	[`89fecf5`](https://github.com/whatchanged-community/whatchanged/commit/89fecf5588f1133f70a8aaf6db3488184ba2ceb2) - refactor: remove unsued code
-	[`a65a4a8`](https://github.com/whatchanged-community/whatchanged/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3) - feat: support version range. eg v2.0.0~v1.0.0
-	[`6301c2f`](https://github.com/whatchanged-community/whatchanged/commit/6301c2f01d881ae861e6c4459a411a5f48c74ba0) - update help information
-	[`770ed02`](https://github.com/whatchanged-community/whatchanged/commit/770ed02d43c4593ab9db8e71a9f93987812d97bc) - update
-	[`585445d`](https://github.com/whatchanged-community/whatchanged/commit/585445d917d4cb74d80b1c385b446f7e09a1606c) - cache tags
-	[`b47e49c`](https://github.com/whatchanged-community/whatchanged/commit/b47e49cf8efac3f5aba9df0149295694424beaf1) - filter tag with semver version
-	[`7f67e78`](https://github.com/whatchanged-community/whatchanged/commit/7f67e783926fed647d2ad5414f31448eea106fc3) - feat: parse commit message and generate to template
-	[`b907117`](https://github.com/whatchanged-community/whatchanged/commit/b907117bca3d6306955070b933361ecb0da0627e) - update ci
-	[`ef5b38a`](https://github.com/whatchanged-community/whatchanged/commit/ef5b38ad50dd150cbdbeff031f6898d9d0aff35a) - update ci linter version
-	[`dd55cc8`](https://github.com/whatchanged-community/whatchanged/commit/dd55cc85d6a3b9c482ba376fa862f20b6de11d5b) - fix lint
-	[`26408cc`](https://github.com/whatchanged-community/whatchanged/commit/26408ccef0f6256ca70edda59e4ab1d1c15fca72) - init
-	[`824d351`](https://github.com/whatchanged-community/whatchanged/commit/824d3511bf90e8fda3d1c7be679274d71ce73f52) - Initial commit
