#!/bin/bash

# 1. Remove type: "tech" and type: "idea" from all articles
find ./articles -type f -name "*.md" -exec sed -i -E '/^type: "(tech|idea)"/d' {} \;

# 2. Convert to Hugo's shortcode
find ./articles -type f -name "*.md" -exec sed -i -E 's|^https://www\.youtube\.com/watch\?v=([a-zA-Z0-9_-]+)|{{< youtube \1 >}}|g' {} \;

find ./articles -type f -name "*.md" -exec sed -i -E 's|^https://twitter\.com/([^/]+)/status/([0-9]+)|{{< tweet user="\1" id="\2" >}}|g' {} \;
find ./articles -type f -name "*.md" -exec sed -i -E 's|^https://x\.com/([^/]+)/status/([0-9]+)|{{< tweet user="\1" id="\2" >}}|g' {} \;
