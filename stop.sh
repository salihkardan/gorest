#!/bin/bash
kill -9 `ps aux | grep re[s]tful | awk {'print $2'} | head -1`