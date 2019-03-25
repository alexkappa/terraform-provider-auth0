# Contributing to Terraform - Auth0 Provider

**First:** if you're unsure or afraid of _anything_, just ask
or submit the issue or pull request anyways. You won't be yelled at for
giving your best effort. The worst that can happen is that you'll be
politely asked to change something. We appreciate any sort of contributions,
and don't want a wall of rules to get in the way of that.

However, for those individuals who want a bit more guidance on the
best way to contribute to the project, read on. This document will cover
what we're looking for. By addressing all the points we're looking for,
it raises the chances we can quickly merge or address your contributions.

Specifically, we have provided checklists below for each type of issue and pull
request that can happen on the project. These checklists represent everything
we need to be able to review and respond quickly.

<!-- TOC depthFrom:2 -->

- [HashiCorp vs. Community Providers](#hashicorp-vs-community-providers)
- [Issues](#issues)
    - [Issue Reporting Checklists](#issue-reporting-checklists)
        - [Bug Reports](#bug-reports)
        - [Feature Requests](#feature-requests)
        - [Questions](#questions)
    - [Issue Lifecycle](#issue-lifecycle)
- [Pull Requests](#pull-requests)
    - [Pull Request Lifecycle](#pull-request-lifecycle)
    - [Checklists for Contribution](#checklists-for-contribution)
        - [Documentation Update](#documentation-update)
        - [Enhancement/Bugfix to a Resource](#enhancementbugfix-to-a-resource)
        - [New Resource](#new-resource)
        - [Terraform Schema and Code Idiosyncracies](#terraform-schema-and-code-idiosyncracies)
    - [Writing Acceptance Tests](#writing-acceptance-tests)
        - [Running an Acceptance Test](#running-an-acceptance-test)
        - [Writing an Acceptance Test](#writing-an-acceptance-test)

<!-- /TOC -->

## HashiCorp vs. Community Providers

We separate providers out into what we call "HashiCorp Providers" and
"Community Providers".

HashiCorp providers are providers that we'll dedicate full time resources to
improving, supporting the latest features, and fixing bugs. These are providers
we understand deeply and are confident we have the resources to manage
ourselves.

Community providers are providers where we depend on the community to
contribute fixes and enhancements to improve. HashiCorp will run automated
tests and ensure these providers continue to work, but will not dedicate full
time resources to add new features to these providers. These providers are
available in official Terraform releases, but the functionality is primarily
contributed.

Our testing standards are the same for both HashiCorp and Community providers,
and HashiCorp runs full acceptance test suites for every provider nightly to
ensure Terraform remains stable.

We make the distinction between these two types of providers to help
highlight the vast amounts of community effort that goes in to making Terraform
great, and to help contributors better understand the role HashiCorp employees
play in the various areas of the code base.

## Issues

### Issue Reporting Checklists

We welcome issues of all kinds including feature requests, bug reports, and
general questions. Below you'll find checklists with guidelines for well-formed
issues of each type.

#### Bug Reports

 - [ ] __Test against latest release__: Make sure you test against the latest
   released version. It is possible we already fixed the bug you're experiencing.

 - [ ] __Search for possible duplicate reports__: It's helpful to keep bug
   reports consolidated to one thread, so do a quick search on existing bug
   reports to check if anybody else has reported the same thing. You can scope
   searches by the label "bug" to help narrow things down.

 - [ ] __Include steps to reproduce__: Provide steps to reproduce the issue,
   along with your `.tf` files, with secrets removed, so we can try to
   reproduce it. Without this, it makes it much harder to fix the issue.

 - [ ] __For panics, include `crash.log`__: If you experienced a panic, please
   create a [gist](https://gist.github.com) of the *entire* generated crash log
   for us to look at. Double check no sensitive items were in the log.

#### Feature Requests

 - [ ] __Search for possible duplicate requests__: It's helpful to keep requests
   consolidated to one thread, so do a quick search on existing requests to
   check if anybody else has reported the same thing. You can scope searches by
   the label "enhancement" to help narrow things down.

 - [ ] __Include a use case description__: In addition to describing the
   behavior of the feature you'd like to see added, it's helpful to also lay
   out the reason why the feature would be important and how it would benefit
   Terraform users.

#### Questions

 - [ ] __Search for answers in Terraform documentation__: We're happy to answer
   questions in GitHub Issues, but it helps reduce issue churn and maintainer
   workload if you work to find answers to common questions in the
   documentation. Often times Question issues result in documentation updates
   to help future users, so if you don't find an answer, you can give us
   pointers for where you'd expect to see it in the docs.

### Issue Lifecycle

1. The issue is reported.

2. The issue is verified and categorized by a Terraform collaborator.
   Categorization is done via GitHub labels. We generally use a two-label
   system of (1) issue/PR type, and (2) section of the codebase. Type is
   usually "bug", "enhancement", "documentation", or "question", and section
   can be any of the providers or provisioners or "core".

3. Unless it is critical, the issue is left for a period of time (sometimes
   many weeks), giving outside contributors a chance to address the issue.

4. The issue is addressed in a pull request or commit. The issue will be
   referenced in the commit message so that the code that fixes it is clearly
   linked.

5. The issue is closed. Sometimes, valid issues will be closed to keep
   the issue tracker clean. The issue is still indexed and available for
   future viewers, or can be re-opened if necessary.

## Pull Requests

Thank you for contributing! Here you'll find information on what to include in
your Pull Request to ensure it is accepted quickly.

 * For pull requests that follow the guidelines, we expect to be able to review
   and merge very quickly.
 * Pull requests that don't follow the guidelines will be annotated with what
   they're missing. A community or core team member may be able to swing around
   and help finish up the work, but these PRs will generally hang out much
   longer until they can be completed and merged.

### Pull Request Lifecycle

1. You are welcome to submit your pull request for commentary or review before
   it is fully completed. Please prefix the title of your pull request with
   "[WIP]" to indicate this. It's also a good idea to include specific
   questions or items you'd like feedback on.

2. Once you believe your pull request is ready to be merged, you can remove any
   "[WIP]" prefix from the title and a core team member will review. Follow
   [the checklists below](#checklists-for-contribution) to help ensure that
   your contribution will be merged quickly.

3. One of Terraform's core team members will look over your contribution and
   either provide comments letting you know if there is anything left to do. We
   do our best to provide feedback in a timely manner, but it may take some
   time for us to respond.

4. Once all outstanding comments and checklist items have been addressed, your
   contribution will be merged! Merged PRs will be included in the next
   Terraform release. The core team takes care of updating the CHANGELOG as
   they merge.

5. In rare cases, we might decide that a PR should be closed. We'll make sure
   to provide clear reasoning when this happens.

### Checklists for Contribution

There are several different kinds of contribution, each of which has its own
standards for a speedy review. The following sections describe guidelines for
each type of contribution.

#### Documentation Update

Because [Terraform's website][website] is in the same repo as the code, it's
easy for anybody to help us improve our docs.

 - [ ] __Reasoning for docs update__: Including a quick explanation for why the
   update needed is helpful for reviewers.
 - [ ] __Relevant Terraform version__: Is this update worth deploying to the
   site immediately, or is it referencing an upcoming version of Terraform and
   should get pushed out with the next release?

#### Enhancement/Bugfix to a Resource

Working on existing resources is a great way to get started as a Terraform
contributor because you can work within existing code and tests to get a feel
for what to do.

 - [ ] __Acceptance test coverage of new behavior__: Existing resources each
   have a set of [acceptance tests][acctests] covering their functionality.
   These tests should exercise all the behavior of the resource. Whether you are
   adding something or fixing a bug, the idea is to have an acceptance test that
   fails if your code were to be removed. Sometimes it is sufficient to
   "enhance" an existing test by adding an assertion or tweaking the config
   that is used, but often a new test is better to add. You can copy/paste an
   existing test and follow the conventions you see there, modifying the test
   to exercise the behavior of your code.
 - [ ] __Documentation updates__: If your code makes any changes that need to
   be documented, you should include those doc updates in the same PR. The
   [Terraform website][website] source is in this repo and includes
   instructions for getting a local copy of the site up and running if you'd
   like to preview your changes.
 - [ ] __Well-formed Code__: Do your best to follow existing conventions you
   see in the codebase, and ensure your code is formatted with `go fmt`. (The
   Travis CI build will fail if `go fmt` has not been run on incoming code.)
   The PR reviewers can help out on this front, and may provide comments with
   suggestions on how to improve the code.
 - [ ] __Vendor additions__: Create a separate PR if you are updating the vendor
   folder. This is to avoid conflicts as the vendor versions tend to be fast
   moving targets.

#### New Resource

Implementing a new resource is a good way to learn more about how Terraform
interacts with upstream APIs. There are plenty of examples to draw from in the
existing resources, but you still get to implement something completely new.

 - [ ] __Minimal LOC__: It can be inefficient for both the reviewer
   and author to go through long feedback cycles on a big PR with many
   resources. We therefore encourage you to only submit **1 resource at a time**.
 - [ ] __Acceptance tests__: New resources should include acceptance tests
   covering their behavior. See [Writing Acceptance
   Tests](#writing-acceptance-tests) below for a detailed guide on how to
   approach these.
 - [ ] __Naming__: Resources should be named `auth0_<resource>` where
   `resource` is the corresponding Auth0 resource. Use `_` as a separator if the resource is described by more than one word. We encourage you to only submit **1 resource at a time**.
 - [ ] __Arguments_and_Attributes__: The HCL for arguments and attributes should mimic the types and structs presented by the Auth0 Management API.
 - [ ] __Documentation__: Each resource gets a page in the Terraform
   documentation. The [Terraform website][website] source is in this
   repo and includes instructions for getting a local copy of the site up and
   running if you'd like to preview your changes. For a resource, you'll want
   to add a new file in the appropriate place and add a link to the sidebar for
   that page.
 - [ ] __Well-formed Code__: Do your best to follow existing conventions you
   see in the codebase, and ensure your code is formatted with `go fmt`. (The
   Travis CI build will fail if `go fmt` has not been run on incoming code.)
   The PR reviewers can help out on this front, and may provide comments with
   suggestions on how to improve the code.
 - [ ] __Vendor updates__: Create a separate PR if you are adding to the vendor
   folder. This is to avoid conflicts as the vendor versions tend to be fast
   moving targets.

### Writing Acceptance Tests

Terraform includes an acceptance test harness that does most of the repetitive
work involved in testing a resource.

#### Running an Acceptance Test

Acceptance tests can be run using the `testacc` target in the Terraform
`Makefile`. The individual tests to run can be controlled using a regular
expression. Prior to running the tests provider configuration details such as
access keys must be made available as environment variables.

For example, to run an acceptance test against the Auth0
provider, the following environment variables must be set:

```sh
export AUTH0_DOMAIN=...
export AUTH0_CLIENT_ID=...
export AUTH0_CLIENT_SECRET=...
```

Tests can then be run by specifying the `test` or `testacc` make target.

```sh
$ make testacc
```

The optional `TESTS` variable can specify a particular test to run instead of the entire test suite.

```sh
$ make testacc TESTS=TestAccConnection
```

#### Writing an Acceptance Test

Terraform has a framework for writing acceptance tests which minimises the
amount of boilerplate code necessary to use common testing patterns. The entry
point to the framework is the `resource.ParallelTest()` function.

Tests are divided into `TestStep`s. Each `TestStep` proceeds by applying some
Terraform configuration using the provider under test, and then verifying that
results are as expected by making assertions using the provider API. It is
common for a single test function to exercise both the creation of and updates
to a single resource. Most tests follow a similar structure.

For more details on acceptance tests and how to write them, the [documentation](https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html) can be a valuable resource. The acceptance tests in this provider should also be able to offer reasonable guidance.

[website]: https://github.com/hashicorp/terraform/tree/master/website
[acctests]: https://github.com/hashicorp/terraform#acceptance-tests
[ml]: https://groups.google.com/group/terraform-tool
