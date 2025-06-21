#!/bin/bash

# Delete all *_author.log files in ./storage/logs/ older than 3 days
find ./storage/logs/ -type f -name '*_author.log' -mtime +3 -exec rm -f {} +

echo "Deleted all *_author.log files in ./storage/logs/ older than 3 days"