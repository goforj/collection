#!/usr/bin/env bash

go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@v1.1.0

gomarkdoc \
  --template-file file=docs/templates/file.gotxt \
  --template-file package=docs/templates/package.gotxt \
  --template-file type=docs/templates/type.gotxt \
  --template-file func=docs/templates/func.gotxt \
  --template-file value=docs/templates/value.gotxt \
  --template-file index=docs/templates/index.gotxt \
  --template-file example=docs/templates/example.gotxt \
  --template-file doc=docs/templates/doc.gotxt \
  --template-file import=docs/templates/import.gotxt \
  --output README.md \
  --embed \
  .



