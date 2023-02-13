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

Generally, test code falls in two categories:
1. `Unit-level tests` are small contained tests that run portions of code in isolation. These are the foundation of the testing pyramid.
2. `System-level tests` require interactions between various code segments or systems. These are difficult to manage and expensive, so they sit at the top of the testing pyramid.

### Types of tests

| Test | Description |
|------|:------------|
| Unit | Independent portions of code (functions, etc.) |
| Integration| How portions of the application interact internally or externally with external systems. |
| Acceptance | Verifies that enduser agrees that the application functions as expected. |
| Fuzz | Send arbitrary inputs into a system to verify that there are no bugs or security vulnerabilities. |
| Smoke | Does the application run? |
| Load | How the application runs under high levels of input or users. |
| Regression | Have past bugs reoccurred within the system or whether a feature no longer works as intended. |

### Testing pyramid

If the _unit tests_ do not pass, you cannot advance up the testing pyramid. 

Further up the pyramid, you find _integration tests_--tests that verify the functionality between units of work, often including an integration with an external dependency like a database.

_Acceptance tests_ verify results from a customer prospective.

_End-to-end_ tests verify that the system is working as expected. These tests include load-testing, where you test a system with a large amount of input or users.

As you move up the pyramid, each layer becomes smaller because they become more expensive to run. They require dependencies or more resources, and they are not always consistent and might not produce _deterministic_ results.

### What to test

Decompose your work into testable units called _system under tests_ (SUT). SUTs have clear boundaries and should be treated as a black box. You should test inputs and assert that the outputs are correct. As you move up the testing pyramid, results are _stochastic_, or random.

### How to test

Test-Driven Development (TDD) motivates developers to meet the requirements set forth by the customer and verify them through the use of tests. This motivation is to help develop the bare minimum and move forward while preventing waste. A similar approach is "duct tape programming": take a basic test and interface, then move forward and make it work. You refine and revisit as needed.

#### 3 steps

You need a system to prove that your code works. The scientific method: question, test, and results. The following steps mimic this method:

1. _Arrange_ your test code so that everything is set up
2. _Act_ on the code you are testing
3. _Assert_ the results

### Code coverage

Go provides code coverage tests that highlight a percentage of the code that has been tested and highlight areas that you might have missed.

While achieving high code coverage is important, trying to reach 100% might result in poorly written tests that are difficult to maintain over time and delay the delivery of the software.

### REST 

## Continuous delivery 

_Delivery_ is providing a product for someone to use. _Deployment_ is the final step in a process where you are running and using your product as a service. Not all products are deployed but all should be delivered. For example, libraries are delivered but not deployed.

Typically, a release is paired with an artifact and a message about what has changed since the previous release. An artifact is any sort of released item, such as a binary file, zip file, etc. A release is the latest version of an application--something that a user can use.

### Deployment

Servers are virtualized and controlled by a set of unique API commands that could allow for the easy creation and destruction of server instances. This abstraction is called Infrastructure as a Service (IaaS)

X as a Service is an abstraction over deployment hardware. The higher the abstraction, the higher the overall cost.

Infrastructure as a Service (IaaS)
: OS and VM. Servers are virtualized and controlled by a set of unique API commands that could allow for the easy creation and destruction of server instances.
: _AWS EC2_, _Google Compute_

Containers as a Service (CaaS)
: VM, OS, Runtime.
: _AWS ECS_, _Google Cloud Run_

Platform as a Service (PaaS)
: VM, OS, Runtime, Application. Allows developers to quickly create and iterate on their platforms. You provide the source code, and the platform identifies, builds, and runs your application.
: _Heroku_, _Google App Engine_

Function as a Service (FaaS)
: VM, OS, Runtime, Application, Functions. Called _serverless_, because the developer doesn't have to know anything about the platform or runtime--just upload a function to run at an endpoint.
: _AWS Lambda_, _Google Cloud Functions_.

### Scale

You need a way to check that your application is running and healthy. You can use a _health check_ endpoint that tells the running platform that the service is working and ready.

# Code quality enforcement

Add rules on PRs to protect the source code and require contributors to explain the changes that they are making. In your repo, complete the following:
1. Go to Settings > Branches.
2. Under **Branch protection rule**, click **Add branch protection rule**.
3. Select **Require a pull request before merging**.

## PR template

You can create a PR template in `./github/PULL_REQUEST_TEMPLATE.md`.

## Constraints

Constraints--or bottlenecks--are the part of the application that determines throughput. Any optimizations that are not at a constraint are pointless.

The slowest part of the development pipeline is the developer--the thought process, and the actual developing.

## Static code analysis

_Static code analysis_ tools detect bad coding practices and antipatterns that might result in bugs or security vulnerabilities. Go provides go vet:
```shell
$ go vet ./...
```
The preceding command logs any bugs or errors to the console.

In addition, you can use `golangci-lint`. This tool lets you select from a variety of linting and static checking libraries to check for unused code, ineffective variable assignments, missing error checks, security risks, etc. To do this, you must create a file called `.golangci.yml` in the project root.

# Git hooks

A _hook_ runs either before or after a specific function executes. Development pipelines can use a _pre-commit hook_ that runs before the developer commits changes. This verifies that the code is working properly, and it prevents failing tests from running in the public repo.

# Dependency inversion principle

Always depend on abstractions, not implementation. Instead of using an implemented class or function directly, you should use an _interface_. You create a structure that fulfills the interface and then you inject it into the consuming struct.

> Think about an electrical plug: you can access electricity with any appliance that has a two-pronged plug. This prevents having to wire your appliance directly into the home or building electrical system. The outlet is the interface, and the plug follows the interface contract.

An interface is an abstraction that allows you to easily use something that is more complex, behind the scenes. In software development, an interface defines the functions for a given struct or class. A struct that has all requested functions statisfies the interface, and you can use it in place of another service that satisfies the interface. An interface is a protocol that defines the boundaries between systems and provides a way to communicate across those boundaries.

Interfaces simplify code testing. _Interface segregation_ is when you split interfaces into small chunks to make them more composable and reusable. Think of the `io.Reader` and `io.Writer` interface.

# Testing dependency injection

Dependency injection allows you to constrain and isolate various parts of the underlying code to help minimize the effects of any _independent variables_, which are the variables that we test.

For example, if you have a service that translates a word into a set amount of languages, do not test the specific languages but rather what would trigger responses. If you test for a specific language, the implementation and testing are tightly coupled. You might want to test that valid requests return a `200` code, and invalid requests return a `404`.

## Stubs

Dependency injection means that you can create a _stub_, a service specifically for testing. Use stubs to test any structure (services, repositories, utilities, etc.). Stubs are not complicated--they return hard-coded values.

Stubs help you focus on testing the results rather than pushing logic through the tests. You do not test dependencies, you test how the part of your application works with dependencies. So, if you are testing a service that translates a word into another language, you might have test cases for the following:
- The default language is returned correctly
- If a word is translated, return the new new word
- If a word is not translated, return the correct error response

For example:
```go
type stubbedService struct{}

// Implements the interface under test.
func (s *stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "bar"
	}
	return ""
}

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/translate/foo",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "bar",
		},
		...
```

## Mocking

Mocks are stubs with more detail. Mocks have methods that allow you to test error handling and edge cases.

## Fakes

> Fakes can test _external dependencies_, or systems outside your control. For a more robust solution, you can look at the Circuit Breakers pattern.

A _fake_ is an object, struct, or service with limited capabilites that stands in for an external service. You can use this to create an HTTP test server.

# Implementing interfaces

Use interfaces so you do not have to directly call other methods in your code. For example, instead of calling a function directly, you can do the following:
1. Create an interface for that function. For example, 
2. Create a type that has a field that implements that interface
3. Create a type that implements the interface.
4. Pass the type that implements the interface as the field to the other type.

Create an interface with the method that you want to implement:

```go
type Translator interface {
	Translate(word string, language string) string
}
```

Create a type that has a field that implements the interface. Create a factory method that returns an implementation:
```go
type TranslateHandler struct {
	service Translator
}

func NewTranslateHandler(service Translator) *TranslateHandler {...}
```

Create the type that implements the interface:
```go
type StaticService struct{}

func NewStaticService() *StaticService {...}

func (s *StaticService) Translate(word string, language string) string {...}
```

To use this, create a `StaticService` object, and pass it to a `TranslateHandler` object:

```go
translationService := translation.NewStaticService()
translateHandler := rest.NewTranslateHandler(translationService)
```

![Interface diagram](/images/interface.svg "Interface implementation")