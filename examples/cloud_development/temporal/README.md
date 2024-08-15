# Temporal

[![Built with Codex](https://www.khulnasoft/codex/img/shield_galaxy.svg)](https://www.khulnasoft/codex/docs/contributor-quickstart/)

Example codex for testing and developing Temporal workflows using Temporalite and the Python Temporal SDK.

For more details, check out:

* [Temporal.io](https://temporal.io/)
* [Temporalite](https://github.com/temporalio/temporalite)
* [Temporal Python SDK](https://github.com/temporalio/sdk-python)
* [Temporal Python Samples](https://github.com/temporalio/samples-python)

## Starting Temporal

```bash
codex run start-temporal
```

This will start the temporalite server for testing.

* You can view the WebUI at `localhost:8233`
* By default, Temporal will listen for activities/requests on port `7233`

## Starting a Codex Shell

```bash
codex shell
```

This will activate a virtual environment and install the Temporal Python SDK.

## Testing the Temporal Workflows

From inside your `codex shell`

```bash
cd temporal_example/hello
python run hello_activity.py
```

This should start the workflow using temporalite.
