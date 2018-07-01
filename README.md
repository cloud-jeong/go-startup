# kcem - Kubernetes-cem

To provision `Kubernetes cluster` to enterprise environment and provide easiest tools for operators

## What is kcem?

`kcem` helps you to get status of kubernetes cluster using k8s client library 

## Features

* Get certification information to connect Kubernetes
* Get current status of Kubernetes
* Deploy application to cluster

## Documentation

Documentation is in the `/docs` directory, 
<!--
 
 [and the index is here.](docs/README.md)


## Kubernetes Release Compatibility


### Kubernetes Version Support

kops is intended to be backward compatible.  It is always recommended to use the
latest version of kops with whatever version of Kubernetes you are using.  Always
use the latest version of kops.

One exception, in regards to compatibility, kops supports the equivalent
Kubernetes minor release number.  A minor version is the second digit in the
release number.  kops version 1.8.0 has a minor version of 8. The numbering
follows the semantic versioning specification, MAJOR.MINOR.PATCH.

For example, kops 1.8.0 does not support Kubernetes 1.9.2, but kops 1.9.0
supports Kubernetes 1.9.2 and previous Kubernetes versions. Only when kops minor
version matches, the Kubernetes minor version does kops officially support the
Kubernetes release.  kops does not stop a user from installing mismatching
versions of K8s, but Kubernetes releases always require kops to install specific
versions of components like docker, that tested against the particular
Kubernetes version.

#### Compatibility Matrix

| kops version | k8s 1.5.x | k8s 1.6.x | k8s 1.7.x | k8s 1.8.x | k8s 1.9.x |
|--------------|-----------|-----------|-----------|-----------|-----------|
| 1.9.x        | Y         | Y         | Y         | Y         | Y         |
| 1.8.x        | Y         | Y         | Y         | Y         | N         |
| 1.7.x        | Y         | Y         | Y         | N         | N         |
| 1.6.x        | Y         | Y         | N         | N         | N         |

Use the latest version of kops for all releases of Kubernetes, with the caveat
that higher versions of Kubernetes are not _officially_ supported by kops.

### kops Release Schedule

This project does not follow the Kubernetes release schedule.  `kops` aims to
provide a reliable installation experience for kubernetes, and typically
releases about a month after the corresponding Kubernetes release. This time
allows for the Kubernetes project to resolve any issues introduced by the new
version and ensures that we can support the latest features. kops will release
alpha and beta pre-releases for people that are eager to try the latest
Kubernetes release.  Please only use pre-GA kops releases in environments that
can tolerate the quirks of new releases, and please do report any issues
encountered.


## Installing

### Prerequisite

`kubectl` is required, see [here](http://kubernetes.io/docs/user-guide/prereqs/).


### OSX From Homebrew

```console
brew update && brew install kops
```

The `kops` binary is also available via our [releases](https://github.com/kubernetes/kops/releases/latest).


### Linux

```console
curl -LO https://github.com/kubernetes/kops/releases/download/$(curl -s https://api.github.com/repos/kubernetes/kops/releases/latest | grep tag_name | cut -d '"' -f 4)/kops-linux-amd64
chmod +x kops-linux-amd64
sudo mv kops-linux-amd64 /usr/local/bin/kops
```


## Release History

See the [releases](https://github.com/kubernetes/kops/releases) for more
information on changes between releases.


## Getting Involved and Contributing

Are you interested in contributing to kops? We, the maintainers and community,
would love your suggestions, contributions, and help! We have a quick-start
guide on [adding a feature](/docs/development/adding_a_feature.md). Also, the
maintainers can be contacted at any time to learn more about how to get
involved.

In the interest of getting more new folks involved with kops, we are starting to
tag issues with `good-starter-issue`. These are typically issues that have
smaller scope but are good ways to start to get acquainted with the codebase.

We also encourage ALL active community participants to act as if they are
maintainers, even if you don't have "official" write permissions. This is a
community effort, we are here to serve the Kubernetes community. If you have an
active interest and you want to get involved, you have real power! Don't assume
that the only people who can get things done around here are the "maintainers".

We also would love to add more "official" maintainers, so show us what you can
do!


What this means:

__Issues__
* Help read and triage issues, assist when possible.
* Point out issues that are duplicates, out of date, etc.
  - Even if you don't have tagging permissions, make a note and tag maintainers (`/close`,`/dupe #127`).

__Pull Requests__
* Read and review the code. Leave comments, questions, and critiques (`/lgtm` ).
* Download, compile, and run the code and make sure the tests pass (make test).
  - Also verify that the new feature seems sane, follows best architectural patterns, and includes tests.

This repository uses the Kubernetes bots.  See a full list of the commands [here](
https://github.com/kubernetes/test-infra/blob/master/commands.md).


## Office Hours

Kops maintainers set aside one hour every other week for **public** office hours. Office hours are hosted on a [zoom video chat](https://zoom.us/my/k8ssigaws) on Fridays at [12 noon (Eastern Time)/9 am (Pacific Time)](http://www.worldtimebuddy.com/?pl=1&lid=100,5,8,12), on odd week numbered weeks. We strive to get to know and help developers either working on `kops` or interested in getting to know more about the project.


### Open Office Hours Topics

Include but not limited to:

- Help and guide to those who attend, who are interested in contributing.
- Discuss the current state of the kops project, including releases.
- Strategize about how to move `kops` forward.
- Collaborate about open and upcoming PRs.
- Present demos.

This time is focused on developers, although we will never turn a courteous participant away. Please swing by, even if you've never actually installed kops.

We encourage you to reach out **beforehand** if you plan on attending. You're welcome to join any session, and please feel free to add an item to the  [agenda](https://docs.google.com/document/d/12QkyL0FkNbWPcLFxxRGSPt_tNPBHbmni3YLY-lHny7E/edit) where we track notes from office hours.

Office hours are hosted on [Zoom](https://zoom.us/my/k8ssigaws) video conference, held on Fridays at [12 noon (Eastern Time)/9 am (Pacific Time)](http://www.worldtimebuddy.com/?pl=1&lid=100,5,8,12) every other odd numbered week.

You can check your week number using:

```bash
date +%V
```

The maintainers and other community members are generally available on the [kubernetes slack](https://github.com/kubernetes/community/blob/master/communication.md#social-media) in [#kops-users](https://kubernetes.slack.com/messages/kops-users/), so come find and chat with us about how kops can be better for you!


## GitHub Issues


### Bugs

If you think you have found a bug please follow the instructions below.

- Please spend a small amount of time giving due diligence to the issue tracker. Your issue might be a duplicate.
- Set `-v 10` command line option and save the log output. Please paste this into your issue.
- Note the version of kops you are running (from `kops version`), and the command line options you are using.
- Open a [new issue](https://github.com/kubernetes/kops/issues/new).
- Remember users might be searching for your issue in the future, so please give it a meaningful title to helps others.
- Feel free to reach out to the kops community on [kubernetes slack](https://github.com/kubernetes/community/blob/master/communication.md#social-media).


### Features

We also use the issue tracker to track features. If you have an idea for a feature, or think you can help kops become even more awesome follow the steps below.

- Open a [new issue](https://github.com/kubernetes/kops/issues/new).
- Remember users might be searching for your issue in the future, so please give it a meaningful title to helps others.
- Clearly define the use case, using concrete examples. EG: I type `this` and kops does `that`.
- Some of our larger features will require some design. If you would like to include a technical design for your feature please include it in the issue.
- After the new feature is well understood, and the design agreed upon we can start coding the feature. We would love for you to code it. So please open up a **WIP** *(work in progress)* pull request, and happy coding. 

-->