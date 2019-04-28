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
_| - | gh user/repo milestones {query} | Show milestones |
| - | gh user/repo projects {query} | Show projects |
| - | gh user/repo releases {query} | Show releases |

| status | command | description |
|:-------|:--------|:------------|
| - | gh user/repo new issue | Create new issue |
| - | gh user/repo new pull | Create new pull request |

| status | command | description |
|:-------|:--------|:------------|
| - | gh user/repo admin | Show the admin page |
| - | gh user/repo clone | Show the clone page |
| - | gh user/repo graphs | Show the graphs page |
| - | gh user/repo network | Show the network page |
| - | gh user/repo pulse | Show the pulse page |
| - | gh user/repo wiki | Show the wiki page |

# License

See LICENSE file.
