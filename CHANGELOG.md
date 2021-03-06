# CHANGELOG.md

## 2.2.2 (2021-02-25)

Features:

  - Trim invalid characters
  - Compatible with ArgoCD

## 2.2.1 (2020-11-01)

Features:

  - Docker Registry login
  - Optimize docker image size

## 2.1.2 (2020-10-29)

Improvements:

  - Refactor code
  - Add unit tests
  - Fix minor bug

## 2.1.1 (2020-10-25)

Features:

  - Move configs to ConfigMap
  - Log format can be set to plain text or json

## 2.0.1 (2020-10-24)

Features:

  - `Interval` unit will be set to `minutes` if missing

## 2.0.0 (2020-10-23)

Features:

  - Rewrite in Golang [cron](https://godoc.org/github.com/robfig/cron)
  - `Interval` is defined flexibly

## 1.16.0 (2020-10-23)

Features:

  - First release using Python [schedule](https://pypi.org/project/schedule/)
  - Interval is defined in minutes