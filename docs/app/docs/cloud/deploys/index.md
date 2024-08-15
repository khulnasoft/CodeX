---
title: "Introduction"
sidebar_position: 1
---

Khulnasoft Deployments is an easy, Codex friendly way to deploy a stateless application to the cloud in a few minutes. Khulnasoft can build and run any Docker container, and provides easy tools for configuring your project's secrets, custom domains, and more. Khulnasoft connects to your projects Github repo to ensure that you always have the latest version of your application deployed. 

:::info
If you want to invite team members to your projects, you will need to add a payment option and upgrade your account to a Solo Plan or higher. For more details, see our [Plans and Pricing](https://www.khulnasoft/cloud/pricing).
:::

## Quickstart

This quickstart will walk through how to configure and deploy a project with Khulnasoft Cloud. We'll start by forking an example repo that is configured for Khulnasoft Deployments, and then demonstrate how to connect your Github repo and activate deploys for your account.

## Forking the Example Repo

To help you get started with Khulnasoft Cloud, we've created an [example Rails app](https://github.com/khulnasoft/khulnasoft-deploy-example) that's been configured to deploy with Khulnasoft Cloud. 

You can fork this repo from the Github UI to add it to your account, or clone and push the repo to your Github account.

## Connecting your Repo to Khulnasoft Cloud

First, you'll need to sign-in with Github and connect your project to Khulnasoft Cloud:

1. From the Create Project screen, select Continue with Github to sign in with Github
2. Select a Github Org to import your project from. If you are only a member of one org, it will be selected for you by default. 
   1. If this is your first time importing a project from your org, you will need to install the Codex Cloud app to provide access to your project
3. Select a Repository to import your repo. If your project is not in the root directory of your repository, you can specify a subdirectory for Khulnasoft to search for your project. 

Once your project is added to Khulnasoft Cloud, you can configure your secrets or deployments. 

## Deploying your Site

1. Select the Deploys tab in your project
2. Click the **Enable Deployments** button to turn on Khulnasoft Deployments for your project
3. Once activated, Khulnasoft will automatically attempt to deploy your repository. You can select the deployment to view its status and build logs

If your site fails to deploy, or if you want to update your deployment, push a commit to the default branch of your repo to trigger a new deploy. 

## Visiting your Site

When your site has finished deploying, Khulnasoft Cloud will display a preview URL that you can visit to test your application. 

Congratulations! You've now deployed your first site with Khulnasoft Cloud.

## Next Steps

* Learn more about [setting up your project](./setup.md)
* Set up a [custom domain](./custom_domains.md) for your application
* Learn how to setup databases, caches, and other [integrations](./integrations/index.md)
