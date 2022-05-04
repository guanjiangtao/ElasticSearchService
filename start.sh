#!/bin/bash
set -ex
make
./ElasticSearchService --config=config/config.conf
