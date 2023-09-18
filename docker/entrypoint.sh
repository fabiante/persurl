#!/usr/bin/env sh

echo "running migrations before actually starting..."
/persurl migrate

echo "running application..."
/persurl run
