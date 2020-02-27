#!/bin/bash

# Skipping test for generated models by sqlboiler
find . -type f -wholename 'app/models/*test.go'\
    -exec sed -i '/func.*testing.T.*{$/a\
                  \ if testing.Short() {\
                  \   t.Skip("skipping testing in short mode")\
                  \ }' \;
