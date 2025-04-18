#!/bin/bash

set -eu -o pipefail

changelog=$(cat CHANGELOG.md)

regex='
## ([0-9]+\.[0-9]+\.[0-9]+(-[^ ]+)?) \(([0-9]{4}-[0-9]{2}-[0-9]{2})\)

((.|
)*)
'

if [[ ! $changelog =~ $regex ]]; then
      echo "Could not find date line in change log!"
      exit 1
fi

version="${BASH_REMATCH[1]}"
date="${BASH_REMATCH[3]}"
notes="$(echo "${BASH_REMATCH[4]}" | sed -n -E '/^## [0-9]+\.[0-9]+\.[0-9]+(-[^ ]+)?/,$!p')"
tag="v$version"

if [[ "$date" != "$(date +"%Y-%m-%d")" ]]; then
    echo "$date is not today!"
    exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
    echo ". is not clean." >&2
    exit 1
fi

echo $'\nVersion:'
echo "$version"

echo $'\nRelease notes:'
echo "$notes"

read -e -p "Push to origin? " should_push

if [ "$should_push" != "y" ]; then
    echo "Aborting"
    exit 1
fi

git push

gh release create --target "$(git branch --show-current)" -t "$version" -n "$notes" "$tag"
