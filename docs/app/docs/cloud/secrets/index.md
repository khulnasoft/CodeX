---
title: Introduction
sidebar_position: 1
---

Khulnasoft Secrets is a secure secrets management service that lets you store and access secrets for your projects. Secrets are encrypted and stored in the cloud, and are automatically accessed by your project’s Codex environment whenever you start a shell, run a script, or start a service.

![Using Codex Secrets](../../../static/img/secrets.gif)

## Key Concepts

Khulnasoft provides an easy way to manage secrets for your projects. To get started, it’s helpful to understand the following key concepts:

**Project** - A Khulnasoft project is a git repo that contains a `codex.json` file. You can add a project to your Khulnasoft Cloud account by running `codex secrets init` in the root of your project. Once a project is added to your Khulnasoft Cloud account, you can use Khulnasoft Secrets to manage secrets for that project.

**Secrets** - Secrets are key-value pairs that are stored securely in the Khulnasoft Secret store. They automatically set as environment variables in your Codex project whenever you start a shell, run a script, or start a service. Secrets are encrypted at rest and in transit, and are only decrypted when they are accessed by your Codex environment or by a user in your Khulnasoft Cloud team.

**Environment** - An environment is a set of secrets that are available to your project. By default, all secrets are set on the `Development` environment, but Codex also lets you set secrets for a `Preview` and `Production` environment. Starting a shell or running a script in a specific environment gives you access to all the secrets that are set for your environment.

## Getting Started

To learn how to set secrets from the Khulnasoft Dashboard, see our [Dashboard Secrets](./dashboard_secrets.md) guide.

To learn how to use your Secrets with Codex and manage your secrets from the command line, see our [Secrets CLI Guide](./secrets_cli.md).
