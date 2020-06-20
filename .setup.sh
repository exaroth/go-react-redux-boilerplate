#!/bin/bash

# Destination directory of modifications
DEST="."

pkgOrigN="github.com/exaroth/go-react-redux-boilerplate"

function prompt() {
    echo -n -e "\033[1;32m?\033[0m \033[1m$1\033[0m ($2) "
}

function replace() {
    if [[ ! -z "${TEST}" ]]; then
        dest=$(echo $2 | sed "s|^${DEST}/||")
        mkdir -p $(dirname "${DEST}/${dest}")
        if [[ "$2" == "${DEST}/${dest}" ]]; then
            sed -E -e "$1" $2 > ${DEST}/${dest}.new
            mv -f ${DEST}/${dest}.new ${DEST}/${dest}
        else
            sed -E -e "$1" $2 > ${DEST}/${dest}
        fi
    else
        sed -E -e "$1" $2 > $2.new
        mv -f $2.new $2
    fi
}

pkgDefName=${PWD##*src/}
prompt "Enter package name to use" ${pkgDefName}
read pkgN
pkgN=$(echo "${pkgN:-${pkgDefName}}" | sed 's/[[:space:]]//g')

find pkg -type f | while read file; do replace "s|${pkgOrigN}|${pkgN}|" "$file"; done
find cmd -type f | while read file; do replace "s|${pkgOrigN}|${pkgN}|" "$file"; done

# Other project files
declare -a files=("CHANGELOG.md" "go.mod")
for file in "${files[@]}"; do
    if [[ -f "${file}" ]]; then
        replace "s|${pkgOrigN}|${pkgN}|" ${file}
    fi
done

