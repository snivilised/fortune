# 🌟 honour: ___Go template for library modules___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![A B](https://img.shields.io/badge/branch%20history-linear-blue?style=flat)](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/managing-a-branch-protection-rule)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/honour.svg)](https://pkg.go.dev/github.com/snivilised/honour)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/honour)](https://goreportcard.com/report/github.com/snivilised/honour)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/honour/badge.svg?branch=main)](https://coveralls.io/github/snivilised/honour?branch=main&kill_cache=1)
[![Astrolib Continuous Integration](https://github.com/snivilised/honour/actions/workflows/ci-workflow.yml/badge.svg)](https://github.com/snivilised/honour/actions/workflows/ci-workflow.yml)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![A B](https://img.shields.io/badge/commit-conventional-commits?style=flat)](https://www.conventionalcommits.org/)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD014/commands-show-output: Dollar signs used before commands without showing output mark down lint -->
<!-- MarkDownLint-disable MD014 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<!-- MD028/no-blanks-blockquote: Blank line inside blockquote -->
<!-- MarkDownLint-disable MD028 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

This project is a template to aid in the startup of Go library module projects.

## 📚 Usage

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" /></a>
</p>

+ unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
+ implemented with [🐍 Cobra](https://cobra.dev/) cli framework, assisted by [🐲 Cobrass](https://github.com/snivilised/cobrass)
+ i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)
+ linting configuration and pre-commit hooks, (see: [linting-golang](https://freshman.tech/linting-golang/)).

## 🔨 Developer Info

By using this template, there is no need to use the cobra-cli to scaffold your application as this has been done already. It should be noted that the structure that is generated by the cobra-cli has been significantly changed in this template, mainly to remove use of the __init()__ function and to minimise use of package level global variables. For a rationale, see [go-without-package-scoped-variables](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables).

### 📝 Checklist of required changes

The following is list of actions that must be performed before using this template. Most of the changes concern changing the name `astrolib` to the name of the new application. As the template is instantiated from github, the new name will automatically replace the top level directory name, that being ___astrolib___.

➕ The following descriptions use owner name ___pandora___ and repo name ___maestro___ as an example. That is to say the client has instantiated ___astrolib___ template into github at url _github.com/pandora/maestro_

#### 🤖 Automated changes

Automated via `automate-checklist.sh` script. When the user instantiates the repo, a github actions workflow is executed which applies changes to the clients repo automatically. The following description describes the changes that are applied on the user's behalf and the workflow is automatically deleted. However, there are other changes that should be made. These compose the manual checklist and should be heeded by the user.

##### ✅ Rename import statements

+ `rename import paths`: global search and replace ___snivilised/honour___ to ___pandora/maestro___

##### ✅ Identifiers

+ `change astrolibTemplData`: perform a refactor rename (_Rename Symbol_) to ___maestroTemplData___

##### ✅ Global search replace astrolib to maestro

Will take care of the following required changes:

+ `change module name`: update the module name inside the .mod file in the root directory
+ `change ApplicationName`: modify to reflect the new application name. This application name is incorporated into the name of any translation files to be loaded.
+ `update BINARY_NAME`: Inside _Taskfile.yml_, change the value of ___BINARY_NAME___ to the name of the client application.
+ `update github action workflows`: change the name of the workflows in the .yaml files to replace ___astrolib___ to ___Maestro___ (note the change of case, if this is important).

##### ✅ Localisation/Internationalisation

+ `change the names of the translation files`: eg change ___astrolib.active.en-GB.json___ to ___maestro.active.en-GB.json___

##### ✅ Miscellaneous automated changes

+ `reset version files`: this is optional because the release process automatically updates the version number according to the tag specified by the user, but will initially contain the version number which reflects the current value of astrolib at the time the client project is instantiated.
+ `change SOURCE_ID`: to "github.com/pandora/maestro"

#### 🖐 Manual changes

The following documents manual changes required. Manual checklist:

##### ☑️ Structural changes

+ `github actions workflow`: If the client does not to use github actions workflow automation, then these files ([ci-workflow](.github/workflows/ci-workflow.yml), [release-workflow](.github/workflows/release-workflow.yml), [.goreleaser.yaml](./.goreleaser.yaml)), should be deleted.

+ `rename the widget command`: rename __widget-cmd.go__ and its associated test __widget_test.go__ to whatever is the first command to be implemented in the application. The widget command can serve as a template as to how to define a new command, without having to start from scratch. It will be easier for the user to modify an existing command, so just perform a case sensitive search and replace for ___widget/Widget___ and replace with ___Foo/foo___ where foo represents the new command to be created.

+ `review bootstrap.go`: this will need to be modified to invoke creation of any custom commands. The `execute` method of __bootstrap__ should be modified to invoke command builder. Refer to the `widget` command to see how this is done.

#### ☑️ Github changes

Unfortunately, github doesn't copy over the template project's settings to the client project, so these changes must be made manually:

Under `Protect matching branches`

+ `Require a pull request before merging` ✅ _ENABLE_
+ `Require linear history` ✅ _ENABLE_
+ `Do not allow bypassing the above settings` ✅ _ENABLE_

Of course, its up to the user what settings they use in their repo, these are just recommended as a matter of good practice.

#### ☑️ Code coverage

+ `coveralls.io`: add maestro project

#### ☑️ Miscellaneous changes

+ `replace README content`
+ `update email address in copyright statement`: The __root.go__ file contains a placeholder for an email address, update this comment accordingly.
+ `create .env file`: Add any appropriate secrets to a newly created .env in the root directory and to enable the __deploy__ task to work, define a __DEPLOY_TO__ entry that defines where builds should be deployed to for testing
+ `install pre-commit hooks`: just run ___pre-commit install___
+ `update translation file`: Inside _Taskfile.yml_, add support for loading any translations that the app will support. By default, it deploys a translation file for __en-US__ so this needs to be updated as appropriate.

### 🌐 l10n Translations

This template has been setup to support localisation. The default language is `en-GB` with support for `en-US`. There is a translation file for `en-US` defined as __src/i18n/deploy/astrolib.active.en-US.json__. This is the initial translation for `en-US` that should be deployed with the app.

Make sure that the go-i18n package has been installed so that it can be invoked as cli, see [go-i18n](https://github.com/nicksnyder/go-i18n) for installation instructions.

To maintain localisation of the application, the user must take care to implement all steps to ensure translate-ability of all user facing messages. Whenever there is a need to add/change user facing messages including error messages, to maintain this state, the user must:

+ define template struct (__xxxTemplData__) in __src/i18n/messages.go__ and corresponding __Message()__ method. All messages are defined here in the same location, simplifying the message extraction process as all extractable strings occur at the same place. Please see [go-i18n](https://github.com/nicksnyder/go-i18n) for all translation/pluralisation options and other regional sensitive content.

For more detailed workflow instructions relating to i18n, please see [i18n README](./resources/doc/i18n-README.md)

### 🧪 Quick Test

To check the app is working (as opposed to running the unit tests), build and deploy:

> task tbd

(which performs a test, build then deploy)

NB: the `deploy` task has been set up for windows by default, but can be changed at will.

Check that the executable and the US language file __maestro.active.en-US.json__ have both been deployed. Then invoke the widget command with something like

> maestro widget -p "P?\<date\>" -t 30

Optionally, the user can also specify the ___directory___ flag:

> maestro widget -p "P?\<date\>" -t 30 -d foo-bar.txt

... where ___foo-bar.txt___ should be replaced with a file that actually exists.

This assumes that the the project name is `maestro`, change as appropriate.

Since the `widget` command uses `Cobrass` option validation to check that the file specified exists, the app will fail if the file does not exist. This serves as an example of how to implement option validation with `Cobrass`.
