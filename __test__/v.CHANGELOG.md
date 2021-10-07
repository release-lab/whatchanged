## Unreleased (2020-12-03)

### Revert:

- revert [`e798326`](https://github.com/vlang/v/commit/e798326a1a410e79219056c3e42b7c6371974de8), gen: implement labelled break and continue (#6880)([`4328233`](https://github.com/vlang/v/commit/43282335044e7fc1c12185d8f5fd5739a73e2eb4))
- revert [`958577b`](https://github.com/vlang/v/commit/958577b98b71a1cc829fc3fbc021682bff6029b0), cgen: enable new if expression implementation everywhere([`30214a7`](https://github.com/vlang/v/commit/30214a7700172dfa2a06f8a627495e724509f2d9))
- revert [`b3f503e`](https://github.com/vlang/v/commit/b3f503e0ce3a8d4fa4c8d0671a11ad8bef72408d), ci: fix building v from vc([`7c66953`](https://github.com/vlang/v/commit/7c6695317c4196c951e90e718993c23e435afa7b))
- revert [`60fbcc3`](https://github.com/vlang/v/commit/60fbcc37fcfd09355f4eb7f9cf2861f878b181f8), gen: scape string function in gen/str.v (#6452)([`403cd0d`](https://github.com/vlang/v/commit/403cd0d915738d176d167ee7ce64c74c483c4155))
- revert [`818220d`](https://github.com/vlang/v/commit/818220de45782832d14a40642fc20a427a858240), checker: fix := test([`3956ea4`](https://github.com/vlang/v/commit/3956ea4665b43e195b2d42df0b0f28e0e2436d78))
- revert [`cc714ca`](https://github.com/vlang/v/commit/cc714ca5cc0e0500573a0aeae7636d1b77e23963), cgen: handle variables called "array([`333f355`](https://github.com/vlang/v/commit/333f355e23cd4a5a9c20a95699d09f86ecc8aa16))
- revert [`ae349ca`](https://github.com/vlang/v/commit/ae349ca6ba432e427eb2f81c8bf6d96a524fefa8), cgen: fix all -Wmissing-variable-declarations (#5802)([`2f6757a`](https://github.com/vlang/v/commit/2f6757a56af3d18eea662e3821f6910615c0000f))

## 0.1.29 (2020-08-06)

### Revert:

- revert [`19c226f`](https://github.com/vlang/v/commit/19c226fcf84aff292064a0764ff02c8a852d1421), parser: Support `unsafe(expr)` (#5973)([`f269cbd`](https://github.com/vlang/v/commit/f269cbdc94c5fb7173bcf8075386a016a574e888))
- revert [`38000f8`](https://github.com/vlang/v/commit/38000f862260fafa310f92c07f052ef0fc5e48fe), cgen: sort const inits/cleanups topologically too([`2425c05`](https://github.com/vlang/v/commit/2425c05c42725c7c2fcd4f2d277b726de8be66f1))
- revert [`f03688e`](https://github.com/vlang/v/commit/f03688e443b2519299706e7b3dc2ecfad07c412d), parser: advanced division by zero check (#5629)([`3d3549d`](https://github.com/vlang/v/commit/3d3549d65a0c2e0330c1dd53ff2791eddfb78bd5))
- revert [`6b2808a`](https://github.com/vlang/v/commit/6b2808a3f90b5664a78fe8593ee0a79ad7d891cd), fmt: re-format parser.v and cgen.v([`74af88b`](https://github.com/vlang/v/commit/74af88bc9261d615ea0a20d93bebcbccd9a17fb4))
- checker: error if variable used before decleration p1([`7bc9e23`](https://github.com/vlang/v/commit/7bc9e234a32f42ae9bbf05273e9cc5d6bf082205))

## 0.1.28 (2020-06-18)

### Revert:

- revert [`6b06184`](https://github.com/vlang/v/commit/6b06184ef4851ec689619a56f76843c86ecc4d62), vweb: @include tempaltes([`b13c95e`](https://github.com/vlang/v/commit/b13c95ea4894417b3db955c23182f7c0f5d4d466))
- revert [`b306c04`](https://github.com/vlang/v/commit/b306c04e99236dcbe61a1cf82a1f89b3a00f1bb5), _vinit: use static initialization of large const []number([`0a07dc5`](https://github.com/vlang/v/commit/0a07dc57623445365abdc5721d422f62cc8cfaa9))
- revert [`d7c6392`](https://github.com/vlang/v/commit/d7c63922d51355ea6caf13ae0b5efd9cf9dbb986), parser: allow void return type for C functions([`3bbda71`](https://github.com/vlang/v/commit/3bbda7103f9152e777a644dacbfcb1b033368b54))
- revert [`9a237c3`](https://github.com/vlang/v/commit/9a237c3e8249fec3a012417ca866d10080b3b615), all: C++ compiler support([`07a78b2`](https://github.com/vlang/v/commit/07a78b2843c81d316e56736775790e625893b9d3))
- revert [`c7e4f5e`](https://github.com/vlang/v/commit/c7e4f5eefbd82eb2dcdeb839f32385d6597a598d), makefile: simplify and speed up([`478ebed`](https://github.com/vlang/v/commit/478ebed0692359e889d96314f16706d6bad381a0))
- revert [`538662d`](https://github.com/vlang/v/commit/538662d99a7d70ffb5ba1d65bc06617ae41f9a14), tests: add more tests on interfaces ([`2618b4f`](https://github.com/vlang/v/commit/2618b4fbd376b0054e2d37b08adcbd7003dace9e))
- revert [`1ea13ac`](https://github.com/vlang/v/commit/1ea13ac7f399ba7026e8f6e86d46e6bbdaf2309f), tests: valgrind: enable string tests([`6ea741e`](https://github.com/vlang/v/commit/6ea741e26ec08a2ac7b75cc5f0f179c25b1e81b8))

## 0.1.27 (2020-05-06)

### Revert:

- parser: simplify array push detection([`41cc96a`](https://github.com/vlang/v/commit/41cc96aaec7d8cee578da075f10ec1c0726152a6))
- parser: allow deref assign without parens pt2 update vlib([`7177e71`](https://github.com/vlang/v/commit/7177e714442f1e7e5e818208034630e2b8664964))
- revert [`0f4c5fb`](https://github.com/vlang/v/commit/0f4c5fb1c9a9dbe80ff5361f48ee6995524dc195), comptime: enable again skipping parsing of other platform branches([`6b31ebe`](https://github.com/vlang/v/commit/6b31ebe456d516bfa0d8ac63f96c199eea3ef598))
- revert [`ba85b8d`](https://github.com/vlang/v/commit/ba85b8d02448f8057440b1298b842e333f33d1cd), Revert "compiler: support setting pref.output_cross_c with -os cross([`1d9d975`](https://github.com/vlang/v/commit/1d9d9757cf062c39c7ea6c08cc980c4b304cc7eb))
- revert [`878be4d`](https://github.com/vlang/v/commit/878be4d8861bccdfd275d26feb71a31ef869ad97), compiler: support setting pref.output_cross_c with -os cross([`ba85b8d`](https://github.com/vlang/v/commit/ba85b8d02448f8057440b1298b842e333f33d1cd))
- revert [`d42725a`](https://github.com/vlang/v/commit/d42725aafec2fd041b71a1c2e12f54f73957f5d8), tetris: part 1 of fixing building it with v2([`4b3c44c`](https://github.com/vlang/v/commit/4b3c44cfd757831403eace5abd865cfbfa302734))
- revert [`4b11075`](https://github.com/vlang/v/commit/4b110756e0bf00d2495840a75e30f7a5825b6d47), fmt: `(var f Foo)`([`da28bc7`](https://github.com/vlang/v/commit/da28bc7026f18e27779936a4b059c6e53c67f8e5))
- revert [`78440be`](https://github.com/vlang/v/commit/78440be2b2d54a56b6c6a97a0a0ce26ca99cfed3), cgen: print bool in struct correctly([`8909402`](https://github.com/vlang/v/commit/890940292bcf372958f95bb9d72ff6fc4cbafae6))
- revert [`cd6d175`](https://github.com/vlang/v/commit/cd6d17518096a53cc8dbcba276117ed47d95ea41), db: increase db module level in vlib ([`ec4be80`](https://github.com/vlang/v/commit/ec4be80bcc0a727f24dc1b80d70a36e675696881))
- term: make compilable with v2 on windows([`b4561fa`](https://github.com/vlang/v/commit/b4561fa8149ecc876427759a52693a0d5a46ab1d))
- revert [`34d9263`](https://github.com/vlang/v/commit/34d926350b30c4577fc0ab5cacc10114960e05ed), map: use hashmap instead of b-tree([`f413b2f`](https://github.com/vlang/v/commit/f413b2fa49045bac5195e9b7c12f2a43f0cab662))

## 0.1.24 (2020-01-08)

### Revert:

- revert [`0bd84e8`](https://github.com/vlang/v/commit/0bd84e80609184d3fbeecc47807cdd2840bc5362), ci: bring back x64 test([`c949e9e`](https://github.com/vlang/v/commit/c949e9e636c954deadf2dd2549ec55299cdf64a5))
- revert [`d226fa7`](https://github.com/vlang/v/commit/d226fa7b17f8d0e6a3c6f9792e0c6f8e8e38006c), enable macos syscalls([`78c706a`](https://github.com/vlang/v/commit/78c706ab71fd3be2743c0f65c8096d79c168606d))
- revert [`81ae54d`](https://github.com/vlang/v/commit/81ae54d9bd1ee7095bf01eeb1183e2a8bbac2c50), x64, v2 backends([`da5fb5d`](https://github.com/vlang/v/commit/da5fb5dcbdc561fafdf27bccaba67ce0af6666f1))
- revert [`8a4bce6`](https://github.com/vlang/v/commit/8a4bce667c40d2aece97d6f5d01532684b943026), B-tree map ([`907254b`](https://github.com/vlang/v/commit/907254b9e830defdd4a6468ecdb63556a7149644))
- revert [`507c71a`](https://github.com/vlang/v/commit/507c71ad80828a34f243df9d1203a3723ab73752), cgen.prepend_to_statement()([`47f9c02`](https://github.com/vlang/v/commit/47f9c02331f8d19bbf8ace7b4b14e80af4c644ef))
- revert [`9352903`](https://github.com/vlang/v/commit/93529031ded570fa0ff1d9b55bed2c70a041ec0a), use stdint.h with msvc([`b9a24e3`](https://github.com/vlang/v/commit/b9a24e3b51bd6dc993793f6753be3b4520974c85))

## 0.1.23 (2019-12-03)

### Revert:

- revert [`2c424c1`](https://github.com/vlang/v/commit/2c424c1aa6485f3b7b8addecb606657daa6b5a3d), update vcreate.v([`dcbb196`](https://github.com/vlang/v/commit/dcbb196e21d30673845546f8dbd6f3360a8e0557))
- revert [`d1e7a54`](https://github.com/vlang/v/commit/d1e7a54f3a2984c09e491d087de21af838b8faad), print_backtrace_skipping_top_frames:  Implementation for MSVC ([`ffa9646`](https://github.com/vlang/v/commit/ffa9646749ec8c8c39873b0183cd7cebca1c7319))
- revert [`e3ad367`](https://github.com/vlang/v/commit/e3ad367b805edb7d44da5bbb150821ee36034369), ci: disable vid build for now([`73bd82e`](https://github.com/vlang/v/commit/73bd82e70642eb6f6098342dea11454a1031a3b0))
- revert [`3b3f0eb`](https://github.com/vlang/v/commit/3b3f0eb50731382712800cd4dd6e7202342b5763), vtools: add an info message". This broke repl tests.([`56e1dac`](https://github.com/vlang/v/commit/56e1dac03a6551324298317eda28fcf1dc7e161f))
- revert [`b1eb9d6`](https://github.com/vlang/v/commit/b1eb9d6b15824c043503dd3f6de6a8421cb40a26), Revert "array: fix and document array functions([`ae696e7`](https://github.com/vlang/v/commit/ae696e7ccbd120e065412885461d942dc6d754a4))
- revert [`7fa33fc`](https://github.com/vlang/v/commit/7fa33fc250407b6e8d5eb4d09bebf94007e620a8), array: fix and document array functions([`b1eb9d6`](https://github.com/vlang/v/commit/b1eb9d6b15824c043503dd3f6de6a8421cb40a26))
- revert [`8373264`](https://github.com/vlang/v/commit/83732642ac979b92b64a9c8a9c072aa44cb3ff84), repl: add readline for user input ([`5faa7e7`](https://github.com/vlang/v/commit/5faa7e78616d80f606870167e8f88845d9dc5c6a))
- revert [`1956c6f`](https://github.com/vlang/v/commit/1956c6f906c25a07273f5de9d45cbf116a9b602e), repl: readline line editing ([`66f36be`](https://github.com/vlang/v/commit/66f36be7d0f08fcf815a8589c9571c36963fbd59))
- revert [`61af044`](https://github.com/vlang/v/commit/61af0443161e8f492d7e653905b5774f668397e5), Revert "parser: fix programs without fn main([`526f1a3`](https://github.com/vlang/v/commit/526f1a3172a603a0a9990315063b85d1d8d00a67))
- revert [`3748de8`](https://github.com/vlang/v/commit/3748de8736798888fdfa8fe501c3d72e3b4f8c01), move compiler/main.v to v.v([`b237ffc`](https://github.com/vlang/v/commit/b237ffcf094cd1963febd7757add6aee18bff101))
- revert [`4229892`](https://github.com/vlang/v/commit/4229892e29fa09c1cd4bec42151fc3b66ebd8ab4), update makefile([`5a80bf2`](https://github.com/vlang/v/commit/5a80bf27f4deb6cac00337403d91b12489dfb9ba))
- revert [`bf21108`](https://github.com/vlang/v/commit/bf21108fdb2fa20484108fe561020d73db9e00e4), parser: fix programs without fn main([`61af044`](https://github.com/vlang/v/commit/61af0443161e8f492d7e653905b5774f668397e5))
- revert [`67ae167`](https://github.com/vlang/v/commit/67ae1670131cd51d11bd82eb9df46efcff401cfc), compiler: remove math dependency([`c3787e1`](https://github.com/vlang/v/commit/c3787e17fd64929c82433703bd2b8fa5043f6473))
- revert [`7eaf289`](https://github.com/vlang/v/commit/7eaf289e33f2ba39269e7b94cf03ea569396893f), remove accidental duplicate code([`7454133`](https://github.com/vlang/v/commit/745413331b9a1259bdb768776991b0c78544b402))

## 0.1.21 (2019-09-30)

### Revert:

- revert [`8992707`](https://github.com/vlang/v/commit/8992707fbbcbbe5eeee8fbd2665715726b3b4b2c), parser: cache tokens (first step)([`0fcdd7d`](https://github.com/vlang/v/commit/0fcdd7db358bf5580ca7a073faf7c2d109654f60))
- revert [`9d1814a`](https://github.com/vlang/v/commit/9d1814ab81f4de5405cf4176dcbbad7192e58c64), remove Travis: GitHub CI is a lot faster and more reliable([`b10886b`](https://github.com/vlang/v/commit/b10886bc217c22d18f9256a6351208df478c49c5))

## 0.1.20 (2019-09-17)

### Revert:

- revert [`279f7d5`](https://github.com/vlang/v/commit/279f7d57cb53e10b7054fffb292de19e4d3b4133), os: fix get_line in windows([`57cfdee`](https://github.com/vlang/v/commit/57cfdeeaf01b48ad9b6d1f052d76567764239c78))
- revert [`982a162`](https://github.com/vlang/v/commit/982a162fbf1a0f166ea4e1c25ad11216b9cd9b8b), compiler: pass -l flags without space (needed for tcc)([`4a43c2f`](https://github.com/vlang/v/commit/4a43c2fa1a50ade5c7aa16c299c207b06934221f))

## 0.1.19 (2019-09-13)

### Revert:

- revert [`86d95fc`](https://github.com/vlang/v/commit/86d95fcd22b05aa1cd6fdba252da6b9c150277f1), travis: bring back msvs test([`3b4703e`](https://github.com/vlang/v/commit/3b4703e3b521af48bf098bb0d0c9211618213472))
- revert [`3bb559b`](https://github.com/vlang/v/commit/3bb559b4c83c8a7f04eeb5016ab0549624088390), Delete .gitattributes([`e38ee80`](https://github.com/vlang/v/commit/e38ee80c98fb6ccfc2c90d434d7d1ac73e438ac0))
- revert [`a43b831`](https://github.com/vlang/v/commit/a43b831965600bfc2502525b168784d94339f7dd), switch deprecation notice([`f6147b7`](https://github.com/vlang/v/commit/f6147b7ac7e06f1ab98138f62eb8eb8ef4f851db))

## 0.1.18 (2019-08-17)

### Revert:

- revert [`15c5f67`](https://github.com/vlang/v/commit/15c5f671f055c5c5ec89380a5187a7e8ba7a8e8f), remove pg from vlib([`9eb385d`](https://github.com/vlang/v/commit/9eb385d9ee732d89b08b9fea7f07e83d9d1cfc19))
- revert [`cdfc4c8`](https://github.com/vlang/v/commit/cdfc4c83725a171c87227d49b3aa7975169036ec), remove vlib/glfw (it's a vpm module now)([`8990eb0`](https://github.com/vlang/v/commit/8990eb06ec4a9f4b3a2771cc39e25f711131859e))
- revert [`acd28fa`](https://github.com/vlang/v/commit/acd28fa4952c00f1548041ede945493114969dcd), Support for the printf optimisation for windows and wide strings([`79be98d`](https://github.com/vlang/v/commit/79be98d2fb21a3f3b83e02a86a18dca38484bd5c))

## 0.1.17 (2019-07-29)

### Bugs fixed:

- segfault when using string.ustring_tmp()([`58577f5`](https://github.com/vlang/v/commit/58577f57c6bd23a42ae2acd02acc2f686c16ab6b)) (@Maciej Sopyło)

### Revert:

- revert [`d38940a`](https://github.com/vlang/v/commit/d38940ad576ade8aea17011c033bd0a591394c0b), modules: fix "is not a directory" error([`dbb64ec`](https://github.com/vlang/v/commit/dbb64ec14917e93019a6adc1080a6db0e4cdd1f3))

## 0.1.16 (2019-07-24)

### Revert:

- revert [`8462e99`](https://github.com/vlang/v/commit/8462e99bc5af9c28248225784204996e6afbd49c), Windows Unicode support([`2291e9f`](https://github.com/vlang/v/commit/2291e9fcfe88460f628f3adfe63e2f669a19bc16))
- revert [`3e00507`](https://github.com/vlang/v/commit/3e005074a34e3bf8c26a0aa1c16b5e95c375468e), Windows Unicode I/O ([`23c5f88`](https://github.com/vlang/v/commit/23c5f88f3e7dd5a6fc6f0e78dfe66a7f352f14ff))

## 0.1.15 (2019-07-16)

## v0.1.13 (2019-07-10)

### Revert:

- revert [`0f0ed8d`](https://github.com/vlang/v/commit/0f0ed8d716485b2a8bc8c5bbac9a01f9ba4bbdbc), make function arguments immutable([`d47e2f1`](https://github.com/vlang/v/commit/d47e2f113fb419e8a4309364d4c53e34f26045ed))
- revert [`adef37f`](https://github.com/vlang/v/commit/adef37f0f47ca0c67791a5af38a19e276767e06c), Added const INVALID_HANDLE_VALUE([`ebbea9f`](https://github.com/vlang/v/commit/ebbea9f560c70c94842f12772b483b3de2d8b957))

## v0.1.12 (2019-07-04)

### Revert:

- revert [`ae1313a`](https://github.com/vlang/v/commit/ae1313a35ccc76d72b87d3e5a4e6d36ae58c3af3), Added permission bits.([`95841a3`](https://github.com/vlang/v/commit/95841a31d40f62659b12ea068d3fb05a78ec0f2a))

## v0.1.11 (2019-07-01)

## v0.1.10 (2019-06-29)

## v0.1.9 (2019-06-29)

### Revert:

- compiler: allow mut passed as argument to be modified'. This broke([`249fa95`](https://github.com/vlang/v/commit/249fa95eab6c80267bb753a52481bc0315cb8737))

## v0.1.8 (2019-06-28)

## v0.1.7 (2019-06-27)

## v0.1.6 (2019-06-27)

## v0.1.5 (2019-06-27)

### Revert:

- revert [`223c35f`](https://github.com/vlang/v/commit/223c35ffb90240e5a17bd4e1ea280847922c44e4), compiler & builtin: bitshifts CAO fix and C code removal in utf8([`b61d2ac`](https://github.com/vlang/v/commit/b61d2ac346a16cc3c3a65225752e617833a85e19))
- revert [`6cb5eee`](https://github.com/vlang/v/commit/6cb5eee1b2e67f2e77e62ff9e7e845aa82a62183), add `unsetenv`, `setenv`, `clearenv` to os module (#583)([`cc1ecd1`](https://github.com/vlang/v/commit/cc1ecd1996eb25bc3abdc40093f0337911b74e61))
- revert [`bda50e9`](https://github.com/vlang/v/commit/bda50e96f047ea9ff447ea57df0b0126734827b1), Revert "getline: check if newline at the end to not cut end of line([`fc7ac25`](https://github.com/vlang/v/commit/fc7ac25a98aa1c4f0ac73d2e494d68ce9078acd1))
- revert [`f6a401a`](https://github.com/vlang/v/commit/f6a401aa025c603b58056a3d008fe03827796a34), removed bits C code inside utf8 builtin([`b901180`](https://github.com/vlang/v/commit/b9011804fc87b762ac99840e604f18121e9759fd))
- revert [`d24be8c`](https://github.com/vlang/v/commit/d24be8cf6c6b13ad36743b7f6436b7673424de04), getline: check if newline at the end to not cut end of line([`bda50e9`](https://github.com/vlang/v/commit/bda50e96f047ea9ff447ea57df0b0126734827b1))

## v0.1.4 (2019-06-26)

### Revert:

- revert [`8ef27f0`](https://github.com/vlang/v/commit/8ef27f0bb32f36770d2fff0668009a1b9f092086), add log replace color([`b00a47b`](https://github.com/vlang/v/commit/b00a47be66e868b087146feb7999fd72f2af3da0))

## v0.1.3 (2019-06-25)

### Bugs fixed:

- add CRLF handling to scanner.v([`461b78b`](https://github.com/vlang/v/commit/461b78bc778feca199c8290d9213a9e87c0a4e6d)) (@Oxylibrium)
- fix error messages([`cf6aa16`](https://github.com/vlang/v/commit/cf6aa16ee4a72c30e38d3c8f86e8817cd2a56936)) (@Yoshiya Hinosawa)

### Revert:

- revert [`8f6ca60`](https://github.com/vlang/v/commit/8f6ca60876b3be552899ccd3ecea6fe84d625722), fixed indentation on multiline condition([`056b415`](https://github.com/vlang/v/commit/056b41521abc1dd276afeeeea66fcaebd5752e04))
