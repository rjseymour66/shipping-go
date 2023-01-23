# Shipping Go

## Delivering value

Segmented delivery
: The pipeline is a series of pipelines that requires manual intervention between each.

An item that is stuck in the pipeline is called a _work in progress_ (WIP). WIP is wasted value and money, because it is work that the customer has not received. It is a loss of value until it is delivered.

To prevent WIP, you want your pipeline to fail fast. To fail fast, you can scatter quality gates that check the software is functioning correctly before you spend resources.

## Three parts of a delivery pipeline

A delivery pipeline is broken into three parts that result in a quality product.

1. Continuous Integration
   Continuous integration is the practice of writing and sharing code. This includes the code and resources to run, deploy, or test the product. Think of these items as the raw materials that enter an assembly line.
2. Continuous testing
   Continuous testing is where a company verifies that they are delivering value to the customer. _Testing_ is a general term that is used to see if a system under test (SUT) is functioning as expected. Automated testing frees up testers to do exploratory testing and potentially uncover more bugs.
3. Continuous delivery
   Continuous delivery is when you deliver value to the customer. Delivery is the act of shipping an _artifact_--any product of the build process that a person can use. The process of making an artifact run is called _deployment_.
   A _release candidate_ is a product that is almost ready for GA. The org performs a variety of tests (load, UI, smoke, manual) to make sure it is ready for public consumption.

## Makefiles

Makefiles help manage projects, including the build system and installing or upgrading dependencies.

### Dependency Makefile

The following Makefile installs and upgrades Go 1.19 on a Linux machine. This creates a uniform development environment:

```Makefile
GO_VERSION := 1.19  # Define global variables at the top

setup:              # This target calls the nested targets 
	install-go      
	init-go

install-go          # Installs go using go website instructions
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go             # Add go env vars to .bashrc
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

upgrade-go          # Deletes existing go installation and installs 1.19
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz
```
> You might need to use the `sudo` command to change system-level files in the `/usr` directory.

For example, to set up your environment, enter `sudo make setup`.

## GitHub Actions

GitHub Actions is an integration tool. It allows you to keep your code and integration platform in one place.

GitHub Action workflows are stored in the .github directory:

```shell
$ mkdir -p .github/workflows
$ touch .github/workflows/pipeline.yml
```
You can create multiple workflows for a single repository.

## Continuous integration

Continuous integration requires a central code repository that contains all the raw materials for the project.

Trunk-based development
: Frequent changes to the main product. You can increase throughput--decrease WIP--if you commit and test smaller batch sizes of code. Smaller, faster changes allows you to correct problems and stabilize a system faster than hidde bugs and issues that might occur if you have a large, long-running branch.

Gitflow
: When you have more than one branch, and you must keep them in sync. For example, you might have a separate branch for production, QA, and development.

Release
: A deployable product. You might not make it generally available--you might just mark it as a place on the trunk where the product has passed specific tests.

A Makefile and GitHub actions enable continuous integration.






## Continuous testing

| Test | Description |
|------|:------------|
| Unit | Independent portions of code (functions, etc.) |
| Integration| How portions of the application interact internally or externally with external systems. |
| Acceptance | Verifies that enduser agrees that the application functions as expected. |
| Fuzz | Send arbitrary inputs into a system to verify that there are no bugs or security vulnerabilities. |
| Smoke | Does the application run? |
| Load | How the application runs under high levels of input or users. |
| Regression | Have past bugs reoccurred within the system or whether a feature no longer works as intended. |