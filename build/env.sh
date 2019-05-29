#!/bin/sh

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
srcdir="$workspace/src/github.com/wylzabc"
if [ ! -L "$srcdir/friday" ]; then
    mkdir -p "$srcdir"
    cd "$srcdir"
    ln -s ../../../../../. friday
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$srcdir/friday"
PWD="$srcdir/friday"

# Launch the arguments with the configured environment.
exec "$@"

