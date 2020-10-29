# CHANGELOG.md

## 2.1.2 (2020-10-29)

Improvements:

  - Refactor code
  - More unit tests
  - Fix minor bug

Helm chart:

  - version `1.1.2`

## 2.1.1 (2020-10-25)

Features:

  - Move configs to ConfigMap
  - Log format can be set to plain text or json

Helm chart:

  - version `1.1.1`

## 2.0.1 (2020-10-24)

Features:

  - `Interval` unit will be set to `minutes` if missing

Helm chart:

  - version `1.0.1`

## 2.0.0 (2020-10-23)

Features:

  - Rewrite in Golang [cron](https://godoc.org/github.com/robfig/cron)
  - `Interval` is defined flexibly

Helm chart:

  - version `1.0.0`

## 1.16.0 (2020-10-23)

Features:

  - First release using Python [schedule](https://pypi.org/project/schedule/)
  - Interval is defined in minutes

Helm chart:

  - versions `0.1.1`