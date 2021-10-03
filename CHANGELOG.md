Unreleased (2021-10-03)
-----------------------

### 🔥 New feature:

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

### 💪 Commits(62):

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
