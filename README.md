[![CircleCI](https://circleci.com/gh/hirakiuc/alfred-github-workflow/tree/master.svg?style=svg&circle-token=fa58af4989cde7043d7b8ea72e53359355d7da9c)](https://circleci.com/gh/hirakiuc/alfred-github-workflow/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/hirakiuc/alfred-github-workflow/badge.svg?branch=master)](https://coveralls.io/github/hirakiuc/alfred-github-workflow?branch=master)

# alfred-github-workflow

Alfred workflow to view the github resources.

# Commands

## User resource

| status | command | description |
|:-------|:--------|:------------|
| x | gh user {query} | show repositories |

## Repository resource

| status | command | description |
|:-------|:--------|:------------|
| x | gh user/repo | Show sub commands |
| x | gh user/repo issues {query} | Show issues |
| x | gh user/repo branches {query} | Show branches |
| x | gh user/repo pulls {query} | Show pulls |
| x | gh user/repo milestones {query} | Show milestones |
| x | gh user/repo projects {query} | Show projects |
| x | gh user/repo releases {query} | Show releases |

| status | command | description |
|:-------|:--------|:------------|
| x | gh user/repo new issue | Create new issue |
| x | gh user/repo new pull | Create new pull request |

| status | command | description |
|:-------|:--------|:------------|
| x | gh user/repo wiki | Show the wiki page |
| x | gh user/repo security | Show the security page |
| x | gh user/repo insights | Show the insights page |
| x | gh user/repo settings | Show the settings page |
| x | gh user/repo clone | Show the clone page |

## My

### Pulls

| status | command | description |
|:-------|:--------|:------------|
| x | gh my pulls created | ... |
| x | gh my pulls assigned | ... |
| x | gh my pulls mentioned | ... |
| x | gh my pulls review-requests | ... |

### Issues

| status | command | description |
|:-------|:--------|:------------|
| x | gh my issues created | ... |
| x | gh my issues assigned | ... |
| x | gh my issues mentioned | ... |

## Configs

| status | command | description |
|:-------|:--------|:------------|
| x | gh > token | Configure the github api token |
| x | gh > clear-cache | clear the current caches |

# License

See LICENSE file.
