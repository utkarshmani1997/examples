#!/bin/bash
set -eou pipefail


# one input, one template
geb multi -O output-1-1-A \
  -I input-1 \
  -T template-1

# one input, one template, again
geb multi -O output-1-1-B \
  -I input-3 \
  -T template-1

# multiple in, one template
geb multi -O output-M-1 \
  -I input-1 -I input-2 \
  -T template-2

# multiple in, multiple template
geb multi -O output-M-M \
  -I input-1 -I input-2 -I input-3 \
  -T template-2 -T template-3

# using repeats
geb multi -O output-R \
  -I input-1 -I input-2 -I input-3 \
  -T template-2 \
  -R "users:emails/version-1.txt:emails/{{name}}-v1.txt" \
  -R "users:emails/version-2.txt:emails/{{name}}-v2.txt"

