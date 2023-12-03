#!/bin/bash

function main {
    local session_file
    local date

    declare -i date
    date="$1"

    session_file="aoc-session"

    if [[ $date -gt 25 ]]; then
        echo "End of AOC-2023"
        exit 0
    fi

    if [[ ! -f $session_file ]]; then
      echo "[-] Missing session"
      exit 1
    fi

    cookie="$(cat ${session_file})"

    pushd $PWD >&/dev/null
    dir=$(printf "day%.2d" "${date}")
    if [[ -d $dir ]]; then
        echo "Directory already exists"
        exit 0
    fi

    mkdir -p $dir && cd $dir

    # input data
    mkdir -p data && cd data
    curl -s -O "https://adventofcode.com/2023/day/$date/input" -b "$cookie"
    cd ..

    # Makefile
    cat <<-EOF | tee Makefile
solution.out: src/*.go
  go build -o $< $*
EOF

    # go.mod
    go mod init $dir

    # src dir
    mkdir -p src && cd src

    cat <<-EOF | tee main.go
package main

func main() {
  fmt.Printf("Hello World!")
}
  EOF

  cat <<-EOF | tee utils.go
package main

import (
        "bufio"
        "log"
        "os"
        "strings"
)

func readFile(filename string) []string {
        input_file, err := os.Open(filename)

        if err != nil {
                log.Fatal("Error opening input file!")
        }

        defer input_file.Close()

        fileScanner := bufio.NewScanner(input_file)
        fileScanner.Split(bufio.ScanLines)

        buf := make([]string, 0)
        for fileScanner.Scan() {
                buf = append(buf, fileScanner.Text())
        }

        return buf
}
EOF

    cat <<-EOF | tee a.go | tee b.go
package main

EOF
    cd ..

    popd >&/dev/null
}

date="$1"
if [[ $# -lt 1 ]]; then
  date=$(date +%d)
fi
main "$(echo $date | sed 's/^0*//')"

