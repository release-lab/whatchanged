#!/bin/sh

cat >output/cjs/package.json <<!EOF
{
    "type": "commonjs"
}
!EOF

cat >output/mjs/package.json <<!EOF
{
    "type": "module"
}
!EOF