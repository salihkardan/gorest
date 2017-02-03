#!/bin/bash
mkdir logs
./restful 1>logs/out.log 2>logs/err.log &
