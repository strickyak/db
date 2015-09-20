set -ex

cd $(dirname $0)
go build

T=/tmp/test--github.com--strickyak--db--$$
trap "rm -rf $T" 0 1 2 3

function db() {
  ./db --db="$T" "$@"
}

db put color red
test '"red"' = $(db get color)
db put color purple
test '"purple"' = $(db get color)

db put gooselevel 42
test '"42"' = $(db get gooselevel)

set / $(db scan)
test '/ "color" :: "purple" "gooselevel" :: "42"' = "$*"

db del gooselevel

set / $(db scan)
test '/ "color" :: "purple"' = "$*"

db del color

set / $(db scan)
test '/' = "$*"

echo $0 : OKAY >&2
