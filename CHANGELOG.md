Unreleased
==========

### New feature:

-	link commit for generation([`b9432db`](https://github.com/axetroy/changelog/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb)) (thanks @axetroy)
-	add full preset template([`7553570`](https://github.com/axetroy/changelog/commit/7553570590b571bd33e10a4f80ec5639d0613042)) (thanks @axetroy)
-	support changelog for git submodule([`ec6a957`](https://github.com/axetroy/changelog/commit/ec6a957752fbca9faa261d8694826779e2cbec1f)) (thanks @axetroy)
-	add writer step([`441ad13`](https://github.com/axetroy/changelog/commit/441ad1322b1fecaca89a170ecebaf2955a77d630)) (thanks @axetroy)
-	add formatter for markdown output([`8c177e0`](https://github.com/axetroy/changelog/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34)) (thanks @axetroy)
-	add --file flag to generate file([`0e4fb09`](https://github.com/axetroy/changelog/commit/0e4fb09789732fec5b09b247e208d61794c3da0d)) (thanks @axetroy)
-	support custom tmeplate and preset([`3aa0aee`](https://github.com/axetroy/changelog/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db)) (thanks @axetroy)
-	support tag ranges([`3d14c9c`](https://github.com/axetroy/changelog/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9)) (thanks @axetroy)
-	support version range. eg v2.0.0~v1.0.0([`a65a4a8`](https://github.com/axetroy/changelog/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3)) (thanks @axetroy)
-	parse commit message and generate to template([`7f67e78`](https://github.com/axetroy/changelog/commit/7f67e783926fed647d2ad5414f31448eea106fc3)) (thanks @axetroy)

### Bugs fixed:

-	improve ssh git url parser([`9993ee6`](https://github.com/axetroy/changelog/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f)) (thanks @axetroy)
-	commit range not include commit of tag([`f8df0ba`](https://github.com/axetroy/changelog/commit/f8df0ba654c8faf67eccf98262cd55807e53e597)) (thanks @axetroy)
-	unescape template([`e118cbf`](https://github.com/axetroy/changelog/commit/e118cbfafd201b945848f15303fdb261e251f058)) (thanks @axetroy)
-	if empty argument for command line([`9c79fd9`](https://github.com/axetroy/changelog/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd)) (thanks @axetroy)
-	**ci**: remove unsued code([`66bcf8f`](https://github.com/axetroy/changelog/commit/66bcf8f43db85409e0392c93f2e347ed91699e81)) (thanks @axetroy)

### Commits(30):

-	\[[`9993ee6`](https://github.com/axetroy/changelog/commit/9993ee600c84cf77d3a0c634e8fa83c2580e137f)] - fix: improve ssh git url parser
-	\[[`b9432db`](https://github.com/axetroy/changelog/commit/b9432db1d1f5afe170296b9e0bfebee1aa62fabb)] - feat: link commit for generation
-	\[[`7553570`](https://github.com/axetroy/changelog/commit/7553570590b571bd33e10a4f80ec5639d0613042)] - feat: add full preset template
-	\[[`7068672`](https://github.com/axetroy/changelog/commit/706867220fa9ca537855f359d3e04d0c3762b793)] - update readme
-	\[[`ec6a957`](https://github.com/axetroy/changelog/commit/ec6a957752fbca9faa261d8694826779e2cbec1f)] - feat: support changelog for git submodule
-	\[[`f540d6f`](https://github.com/axetroy/changelog/commit/f540d6f7123334dac558a37c6ac056fed1021cda)] - refactor: rename transform to transformer
-	\[[`441ad13`](https://github.com/axetroy/changelog/commit/441ad1322b1fecaca89a170ecebaf2955a77d630)] - feat: add writer step
-	\[[`f8df0ba`](https://github.com/axetroy/changelog/commit/f8df0ba654c8faf67eccf98262cd55807e53e597)] - fix: commit range not include commit of tag
-	\[[`8c177e0`](https://github.com/axetroy/changelog/commit/8c177e032e8bdb1b76d135981ea10e7053f3ef34)] - feat: add formatter for markdown output
-	\[[`a47fe84`](https://github.com/axetroy/changelog/commit/a47fe84d2141635d82c2dc49500bfc2a81c03535)] - docs: update how it works
-	\[[`0e4fb09`](https://github.com/axetroy/changelog/commit/0e4fb09789732fec5b09b247e208d61794c3da0d)] - feat: add --file flag to generate file
-	\[[`3aa0aee`](https://github.com/axetroy/changelog/commit/3aa0aee2584036da1c63dea9bb399cb83b48a8db)] - feat: support custom tmeplate and preset
-	\[[`e118cbf`](https://github.com/axetroy/changelog/commit/e118cbfafd201b945848f15303fdb261e251f058)] - fix: unescape template
-	\[[`3d14c9c`](https://github.com/axetroy/changelog/commit/3d14c9cf2dc7d51e348fddc7764d8aba1691fac9)] - feat: support tag ranges
-	\[[`44ce3cd`](https://github.com/axetroy/changelog/commit/44ce3cd8d68b786a3aac6b5daba90e7f12b75200)] - docs: update readme
-	\[[`9c79fd9`](https://github.com/axetroy/changelog/commit/9c79fd91bbf88f7861b4aca89ced8384cf2b9bcd)] - fix: if empty argument for command line
-	\[[`0b6ba0c`](https://github.com/axetroy/changelog/commit/0b6ba0c3fc49139467025eaebbe16c158a0cce65)] - refactor: Refactor the software architecture to make it clearer and simpler
-	\[[`66bcf8f`](https://github.com/axetroy/changelog/commit/66bcf8f43db85409e0392c93f2e347ed91699e81)] - fix(ci): remove unsued code
-	\[[`89fecf5`](https://github.com/axetroy/changelog/commit/89fecf5588f1133f70a8aaf6db3488184ba2ceb2)] - refactor: remove unsued code
-	\[[`a65a4a8`](https://github.com/axetroy/changelog/commit/a65a4a8bd0122e41c7b20c98676e9def76e786d3)] - feat: support version range. eg v2.0.0~v1.0.0
-	\[[`6301c2f`](https://github.com/axetroy/changelog/commit/6301c2f01d881ae861e6c4459a411a5f48c74ba0)] - update help information
-	\[[`770ed02`](https://github.com/axetroy/changelog/commit/770ed02d43c4593ab9db8e71a9f93987812d97bc)] - update
-	\[[`585445d`](https://github.com/axetroy/changelog/commit/585445d917d4cb74d80b1c385b446f7e09a1606c)] - cache tags
-	\[[`b47e49c`](https://github.com/axetroy/changelog/commit/b47e49cf8efac3f5aba9df0149295694424beaf1)] - filter tag with semver version
-	\[[`7f67e78`](https://github.com/axetroy/changelog/commit/7f67e783926fed647d2ad5414f31448eea106fc3)] - feat: parse commit message and generate to template
-	\[[`b907117`](https://github.com/axetroy/changelog/commit/b907117bca3d6306955070b933361ecb0da0627e)] - update ci
-	\[[`ef5b38a`](https://github.com/axetroy/changelog/commit/ef5b38ad50dd150cbdbeff031f6898d9d0aff35a)] - update ci linter version
-	\[[`dd55cc8`](https://github.com/axetroy/changelog/commit/dd55cc85d6a3b9c482ba376fa862f20b6de11d5b)] - fix lint
-	\[[`26408cc`](https://github.com/axetroy/changelog/commit/26408ccef0f6256ca70edda59e4ab1d1c15fca72)] - init
-	\[[`824d351`](https://github.com/axetroy/changelog/commit/824d3511bf90e8fda3d1c7be679274d71ce73f52)] - Initial commit
