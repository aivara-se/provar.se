#!/bin/bash

docker compose up -d
yarn workspaces foreach --all --parallel --interlaced run dev
