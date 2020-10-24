# go-osmand-tracker

A basic application for submitting and retrieving live location updates. Originally focused on OsmAnd, but basically any application that supports REST could use it.

[![Build Status](https://travis-ci.org/ricardobalk/go-osmand-tracker.svg?branch=master)](https://travis-ci.org/ricardobalk/go-osmand-tracker)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ricardobalk/go-osmand-tracker?color=%2300aa00)](./go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/ricardobalk/go-osmand-tracker)](https://goreportcard.com/report/github.com/ricardobalk/go-osmand-tracker)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/ricardobalk/go-osmand-tracker?color=%2300aa00)](https://github.com/ricardobalk/go-osmand-tracker/graphs/commit-activity)
[![GitHub open issues](https://img.shields.io/github/issues/ricardobalk/go-osmand-tracker)](https://github.com/ricardobalk/go-osmand-tracker/issues?q=is%3Aopen+is%3Aissue)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/ricardobalk/go-osmand-tracker?color=%2300aa00)](https://github.com/ricardobalk/go-osmand-tracker/issues?q=is%3Aissue+is%3Aclosed)
[![Requirements Status](https://requires.io/github/ricardobalk/go-osmand-tracker/requirements.svg?branch=master)](https://requires.io/github/ricardobalk/go-osmand-tracker/requirements/?branch=master)
[![License](https://img.shields.io/github/license/ricardobalk/go-osmand-tracker?color=%2300aa00)](./LICENSE.txt)
[![GitHub Hacktoberfest combined status](https://img.shields.io/github/hacktoberfest/2020/ricardobalk/go-osmand-tracker)](https://github.com/ricardobalk/go-osmand-tracker/issues?q=is%3Aopen+is%3Aissue+label%3Ahacktoberfest)


The back-end (REST API) of this application is made with Go, the front-end (map interface) is made with modern Vue 3 and TypeScript. Because of its modular setup, it is possible to run the back-end, front-end or a combination of both. :tada:

![OsmAnd with activated live tracking and the corresponding console output from go-osmand-tracker](./docs/tracking-example.png)

Image: OsmAnd app submitting location updates to the back-end, front-end retrieving the current location and showing it on a map. [Learn how to set up OsmAnd][OsmAnd documentation].

---

## Resources

### Getting started

Please read the [installation instructions][] to learn how to get started. It explains how to run the back-end, front-end or both, with and without Docker. :wink:

### API Documentation

To learn more about the API and see some examples of using it, take a look at the [API Documentation][].

### Contributing

The [contribution guidelines][] document briefly explains how to contribute to this repository. Also consider taking a look at the [open issues][] to know what needs to be done and the [kan-ban board][] to see who's currently working on what.

### License

The license this project uses is the EUPL v1.2 or later. See [LICENSE.txt](LICENSE.txt). Available in [other languages](./EUPL) as well.



[installation instructions]: ./docs/Installation/README.md "Installation Instructions"
[API Documentation]: ./api/README.md "API Documentation"
[contribution guidelines]: ./CONTRIBUTING.md	"Contribution guidelines"
[open issues]: https://github.com/ricardobalk/go-osmand-tracker/issues "Open issues of go-osmand-tracker"
[kan-ban board]: https://github.com/ricardobalk/go-osmand-tracker/projects/1 "Who's working on what?"

[OsmAnd documentation]: ./docs/OsmAnd/README.md "How to set up OsmAnd"